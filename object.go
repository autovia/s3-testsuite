package main

import (
	"context"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func RunObjectTest() {
	setup()

	listObjectsV2Output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("test"),
	})
	assert("ListObjectsV2", len(listObjectsV2Output.Contents) == 0, err)

	content := "testfile\n123"
	r := strings.NewReader(content)

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("test"),
		Key:    aws.String("test.txt"),
		Body:   r,
	})
	eval("PutObject", err)

	GetObjectOutput, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String("test"),
		Key:    aws.String("test.txt"),
	})
	eval("GetObject", err)
	buf, err := io.ReadAll(GetObjectOutput.Body)
	assert("GetObject body", string(buf) == content, err)

	_, err = client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String("test"),
		Key:    aws.String("test.txt"),
	})
	eval("HeadObject", err)

	_, err = client.ListObjectVersions(context.TODO(), &s3.ListObjectVersionsInput{
		Bucket: aws.String("test"),
		Prefix: aws.String("test.txt"),
	})
	eval("ListObjectVersions", err)

	_, err = client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String("test"),
		Key:    aws.String("test.txt"),
	})
	eval("DeleteObject", err)

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
