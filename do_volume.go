package catacomb

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Additional imports needed for examples below

type DoSpaceConfig struct {
}

type DoSpaceSession struct {
	bucket string
	client *s3.S3
}

func DoSpaceNewSession(config DoSpaceConfig) DoSpaceSession {
	key := os.Getenv("SPACES_KEY")
	secret := os.Getenv("SPACES_SECRET")
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:    aws.String("https://nyc3.digitaloceanspaces.com"),
		Region:      aws.String("us-east-1"),
	}

	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)

	session := DoSpaceSession{
		client: s3Client,
	}
	return session
}

func (session DoSpaceSession) doSpaceCloseSession() {

}

func (session DoSpaceSession) doSpaceListObjects() {
	input := &s3.ListObjectsInput{
		Bucket: aws.String("example-space-name"),
	}

	objects, err := session.client.ListObjects(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, obj := range objects.Contents {
		fmt.Println(aws.StringValue(obj.Key))
	}
}

func (session DoSpaceSession) doSpaceUpload() {
	object := s3.PutObjectInput{
		Bucket: aws.String(session.bucket),
		Key:    aws.String("file.ext"),
		Body:   strings.NewReader("The contents of the file."),
		ACL:    aws.String("private"),
		Metadata: map[string]*string{
			"x-amz-meta-my-key": aws.String("your-value"), //required
		},
	}
	_, err := session.client.PutObject(&object)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (session DoSpaceSession) doSpaceDownload(objectKey string) {

	input := &s3.GetObjectInput{
		Bucket: aws.String(session.bucket),
		Key:    aws.String(objectKey),
	}

	result, err := session.client.GetObject(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	out, err := os.Create("/tmp/local-file.ext")
	defer out.Close()

	_, err = io.Copy(out, result.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func (session DoSpaceSession) doSpaceDelete(objectKey string) {

	input := &s3.DeleteObjectInput{
		Bucket: aws.String(session.bucket),
		Key:    aws.String(objectKey),
	}

	_, err := session.client.DeleteObject(input)
	if err != nil {
		fmt.Println(err.Error())
	}

}
