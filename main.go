package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var client *s3.Client
var count int

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	RunBucketTest()
	RunObjectTest()
}

func assert(test string, b bool, err error) {
	if err != nil {
		log.Println(test, err)
	} else {
		if b {
			log.Println(test, "OK")
		} else {
			log.Println(test, "ERR")
		}
	}
	count += 1
}

func eval(test string, err error) {
	if err != nil {
		log.Println(test, err)
	} else {
		log.Println(test, "OK")
	}
	count += 1
}

func stats() {
	log.Println("testcases: ", count)
}
