package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func RunBucketTest() {
	_, err := client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String("test"),
	})
	eval("CreateBucket", err)

	listBucketsOutput, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if listBucketsOutput != nil {
		assert("ListBuckets", len(listBucketsOutput.Buckets) == 1, err)
	}

	_, err = client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: aws.String("test"),
	})
	eval("HeadBucket", err)

	_, err = client.GetBucketVersioning(context.TODO(), &s3.GetBucketVersioningInput{
		Bucket: aws.String("test"),
	})
	eval("GetBucketVersioning", err)

	_, err = client.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
		Bucket: aws.String("test"),
	})
	eval("DeleteBucket", err)
}
