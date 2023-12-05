package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	var endpoint, region string
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: %s [options...] <file> <s3url>
  <file> string
	Filepath to upload
  <s3url> string
	URL of upload destination (s3://<bucket>/<key>)
`, os.Args[0])
		flag.PrintDefaults()
	}
	flag.StringVar(&endpoint, "endpoint", "", "Endpoint URL (use if you are using S3-compatible storage)")
	flag.StringVar(&region, "region", "us-east-1", "Region")
	flag.Parse()

	creds := credentials.NewEnvCredentials()
	config := &aws.Config{
		Endpoint:    &endpoint,
		Region:      &region,
		Credentials: creds,
	}
	ss, err := session.NewSession(config)
	if err != nil {
		log.Fatalf("failed to create AWS session: %+v", err)
	}

	src := flag.Arg(0)
	f, err := os.Open(src)
	if err != nil {
		log.Fatalf("failed to open input file('%s'): %+v", src, err)
	}
	defer f.Close()

	// e.g: s3://<bucket>/<key>
	dst := flag.Arg(1)
	dstUrl, err := url.Parse(dst)
	if err != nil {
		log.Fatalf("failed to parse S3 URL('%s'): %+v", dst, err)
	}
	bucket := dstUrl.Host
	key := dstUrl.Path
	if dstUrl.Scheme != "s3" || bucket == "" {
		log.Fatalf("destination URL must be 's3://<bucket>/<key>' format")
	}
	res, err := s3.New(ss).PutObject(&s3.PutObjectInput{
		Bucket: &bucket,
		Body:   f,
		Key:    &key,
	})
	if err != nil {
		log.Fatalf("failed to put object: %+v", err)
	}
	log.Printf("uploaded: %s", res)
}
