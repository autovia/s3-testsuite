package main

import (
	"context"
	"io"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func RunObjectTest() {
	setup()

	listObjectsV2Output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("test"),
	})
	if listObjectsV2Output != nil {
		assert("ListObjectsV2", len(listObjectsV2Output.Contents) == 0, err)
	}

	content := "testfile\n123"
	r := strings.NewReader(content)

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("test"),
		Key:    aws.String("one/two/test.txt"),
		Body:   r,
	})
	eval("PutObject", err)

	GetObjectOutput, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String("test"),
		Key:    aws.String("one/two/test.txt"),
	})
	eval("GetObject", err)
	if err == nil {
		GetObjectBuf, err := io.ReadAll(GetObjectOutput.Body)
		assert("GetObject body", string(GetObjectBuf) == content, err)
	}

	_, err = client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String("test"),
		Key:    aws.String("one/two/test.txt"),
	})
	eval("HeadObject", err)

	_, err = client.CopyObject(context.TODO(), &s3.CopyObjectInput{
		Bucket:     aws.String("test"),
		Key:        aws.String("test2.txt"),
		CopySource: aws.String("/test/one/two/test.txt"),
	})
	eval("CopyObject", err)

	CopiedObject, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String("test"),
		Key:    aws.String("test2.txt"),
	})
	eval("CopiedObject", err)
	if err == nil {
		CopiedObjectBuf, err := io.ReadAll(CopiedObject.Body)
		assert("CopiedObject body", string(CopiedObjectBuf) == content, err)
	}

	CreateMultipartUploadOutput, err := client.CreateMultipartUpload(context.TODO(), &s3.CreateMultipartUploadInput{
		Bucket: aws.String("test"),
		Key:    aws.String("one/two/test22.txt"),
	})
	eval("CreateMultipartUpload", err)
	log.Print(CreateMultipartUploadOutput)

	_, err = client.ListObjectVersions(context.TODO(), &s3.ListObjectVersionsInput{
		Bucket: aws.String("test"),
		Prefix: aws.String("test2.txt"),
	})
	eval("ListObjectVersions", err)

	_, err = client.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
		Bucket: aws.String("test"),
		Delete: &types.Delete{
			Objects: []types.ObjectIdentifier{
				{
					Key: aws.String("one/two/test.txt"),
				},
				{
					Key: aws.String("test2.txt"),
				},
				{
					Key: aws.String("one/two/test22.txt"),
				},
			},
		},
	})
	eval("DeleteObjects", err)

	teardown()
}

func setup() {
	_, err := client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String("test"),
	})
	eval("CreateBucket", err)
}

func teardown() {
	_, err := client.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
		Bucket: aws.String("test"),
	})
	eval("DeleteBucket", err)
}
