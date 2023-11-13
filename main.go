package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var client *s3.Client
var cases int
var errmap map[string]error

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	errmap = map[string]error{}

	RunBucketTest()
	RunObjectTest()

	stats()
}

func assert(test string, b bool, err error) {
	if err != nil {
		fmt.Print("E")
		errmap[test] = err
	} else {
		if b {
			fmt.Print(".")
		} else {
			fmt.Print("E")
			errmap[test] = errors.New("assert not valid")
		}
	}
	cases += 1
}

func eval(test string, err error) {
	if err != nil {
		fmt.Print("E")
		errmap[test] = err
	} else {
		fmt.Print(".")
	}
	cases += 1
}

func stats() {
	var c int
	for k, v := range errmap {
		fmt.Println("\nError: ", k, v)
		c++
	}

	fmt.Printf("\nTestcases: %v, Errors: %v\n", cases, c)
}
