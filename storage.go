package catacomb

const CloudStorageProvider (
	AWS_S3 iota
	AZURE_STORAGE
	DO_VOLUMES
	LINODE_STORAGE
)

type CatacombStorage struct {
	provider string
	accountid string
	password string
	address string
}

/*
type CatacombStorageConfig struct {
	provider int
	accountid string
	password string
	address string
}
*/

func init_storage() *CatacombStorage {
	/*
	switch config.provider {
	case AWS_S3:
		fmt.Printf("Unsupported")
	case AZURE_STORAGE:
		fmt.Printf("Unsupported")
	case DO_VOLUMES:
		fmt.Printf("Unsupported")
	case LINODE_STORAGE:
		fmt.Printf("Unsupported")
	default:
		fmt.Printf("Unsupported")
	}
	*/
}

func storageObjectUpload() {
	switch {

	}
}

func storageObjectDownload() {
	switch {

	}
}

func storageObjectDelete() {
	switch {

	}
}


/*
func verify_storage_config() {

}


func upload(&LaputaStorage storage, localpath string, remotepath string) {
	switch config.provider {
	case AWS_S3:
		sess := session.Must(session.NewSession())
		// Create an uploader with the session and default options
		uploader := s3manager.NewUploader(sess)
		f, err  := os.Open(filename)
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
	case AZURE_STORAGE:
		fmt.Printf("Unsupported")
	case DO_VOLUMES:
		fmt.Printf("Unsupported")
	case LINODE_STORAGE:
		fmt.Printf("Unsupported")
	default:
		fmt.Printf("Unsupported")
	}
}

func download(&LaputaStorage config, localpath string, remotepath string) {
	switch config.provider {
	case AWS_S3:
		sess := session.Must(session.NewSession())
		// Create a downloader with the session and default options
		downloader := s3manager.NewDownloader(sess)

		// Create a file to write the S3 Object contents to.
		f, err := os.Create(localpath)
		if err != nil {
			return fmt.Errorf("failed to create file %q, %v", filename, err)
		}
		// Write the contents of S3 Object to the file
		n, err := downloader.Download(f, &s3.GetObjectInput{
			Bucket: aws.String(myBucket),
			Key:    aws.String(remotepath),
		})
		if err != nil {
			return fmt.Errorf("failed to download file, %v", err)
		}
	case AZURE_STORAGE:
		fmt.Printf("Unsupported")
	case DO_VOLUMES:
		fmt.Printf("Unsupported")
	case LINODE_STORAGE:
		fmt.Printf("Unsupported")
	default:
		fmt.Printf("Unsupported")
	}
}*/