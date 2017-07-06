//
//
package main

import (
	"fmt"
	"os"
        "bytes"
        "strings"
        "github.com/ryanuber/columnize"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Lists all objects in a bucket using pagination
//
// Usage:
// listbackends <bucket>
func main() {
	if len(os.Args) < 2 {
		fmt.Println("you must specify an AWS bucket")
		return
	}

	sess := session.Must(session.NewSession())

	svc := s3.New(sess)
	
	bucket := &os.Args[1]
	
	var output []string
	output = append(output, "Terraform State | Status")
	output = append(output, "--------------- | ------")

	i := 0
	
	// Change to not doing paging. 
	err := svc.ListObjectsPages(&s3.ListObjectsInput{
		Bucket: &os.Args[1],
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		//fmt.Println("Page,", i)
		i++

		for _, obj := range p.Contents {
			
			status := Get_Terraform_State(*bucket, *obj.Key)
			output_string := " " + *obj.Key + " | " + status
			output = append(output, output_string) 
		}
		return true
	})
	if err != nil {
		fmt.Println("failed to list objects", err)
		return
	}
	 listing := columnize.SimpleFormat(output)
         fmt.Println(listing)
}

//
// Retrieve terraform state file from S3 bucket and determine id any resources specified in state
//
func Get_Terraform_State(bucket string, key string) string {
	svc := s3.New(session.New())
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	result, err := svc.GetObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return ""
	}

        buf := new(bytes.Buffer)
	buf.ReadFrom(result.Body)
	data := buf.String()
	result.Body.Close()
	i := strings.Index(data, "\"resources\": {},")
	if i > 0 { 
	  return "NO RESOURCES"
	} else {
	   return "ACTIVE WITH RESOURCES"
	}
}
