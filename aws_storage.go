package catacomb

import (
	"fmt"
	"os"
	"github.com/aws/aws-sdk-go-v2"
)

type AwsS3Config struct {
}

type AwsS3Session struct {
	session
}

func (config AwsS3Config) awsNewSession() AwsS3Session {

	session := AwsS3Session{
		session = session.Must(session.NewSession()),
	}

	return session
}

func (session AwsS3Session) awsCloseSession() {

}

func awsStorageObjectUpload() {
	sess := session.Must(session.NewSession())

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", filename, err)
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(myString),
		Body:   f,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(result.Location))
}

func awsStorageObjectDownload() {
	// The session the S3 Downloader will use
	sess := session.Must(session.NewSession())

	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(sess)

	// Create a file to write the S3 Object contents to.
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %q, %v", filename, err)
	}

	// Write the contents of S3 Object to the file
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(myString),
	})
	if err != nil {
		return fmt.Errorf("failed to download file, %v", err)
	}
	fmt.Printf("file downloaded, %d bytes\n", n)
}

func awsStorageObjectDelete() {

}
