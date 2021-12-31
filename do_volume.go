package catacomb

type DoSpaceConfig struct {
}

type DoSpaceSession struct {
}

func doSpaceNewSession(config LinodeVolumeConfig) LinodeVolumeSession {

	/*
		key := os.Getenv("SPACES_KEY")
			secret := os.Getenv("SPACES_SECRET")

			s3Config := &aws.Config{
				Credentials: credentials.NewStaticCredentials(key, secret, ""),
				Endpoint:    aws.String("https://nyc3.digitaloceanspaces.com"),
				Region:      aws.String("us-east-1"),
			}

			newSession := session.New(s3Config)
			s3Client := s3.New(newSession)
	*/

	session := LinodeVolumeSession{}
	return session
}

func (session LinodeVolumeSession) doSpaceCloseSession() {

}

func (session DoSpaceSession) doSpaceListObjects() {
	/*
	   input := &s3.ListObjectsInput{
	     Bucket:  aws.String("example-space-name"),
	   }

	   objects, err := s3Client.ListObjects(input)
	   if err != nil {
	     fmt.Println(err.Error())
	   }

	   for _, obj := range objects.Contents {
	     fmt.Println(aws.StringValue(obj.Key))
	   }
	*/
}

func (session DoSpaceSession) doSpaceObjectUpload() {
	/*object := s3.PutObjectInput{
		Bucket: aws.String("example-space-name"),
		Key:    aws.String("file.ext"),
		Body:   strings.NewReader("The contents of the file."),
		ACL:    aws.String("private"),
		Metadata: map[string]*string{
								 "x-amz-meta-my-key": aws.String("your-value"), //required
						 },
	}
	_, err := s3Client.PutObject(&object)
	if err != nil {
		fmt.Println(err.Error())
	}*/
}

func (session DoSpaceSession) doSpaceObjectDownload() {
	/*
	   input := &s3.GetObjectInput{
	       Bucket: aws.String("example-space-name"),
	       Key:    aws.String("file.ext"),
	   }

	   result, err := s3Client.GetObject(input)
	   if err != nil {
	       fmt.Println(err.Error())
	   }

	   out, err := os.Create("/tmp/local-file.ext")
	   defer out.Close()

	   _, err = io.Copy(out, result.Body)
	   if err != nil {
	       fmt.Println(err.Error())
	   }
	*/
}

func (session DoSpaceSession) doSpaceObjectDelete() {
	/*
	   input := &s3.DeleteObjectInput{
	       Bucket: aws.String("example-space-name"),
	       Key:    aws.String("example-file-to-delete.ext"),
	   }

	   result, err := s3Client.DeleteObject(input)
	   if err != nil {
	       fmt.Println(err.Error())
	   }
	*/
}
