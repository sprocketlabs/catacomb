package catacomb

import "fmt"

type CloudStorageProvider int

const (
	AWS_S3         CloudStorageProvider = 0
	AZURE_STORAGE                       = 1
	DO_VOLUMES                          = 2
	LINODE_STORAGE                      = 3
	SFTP                                = 4
)

type CascadeOperation int

const (
	CASCADE_NONE            CascadeOperation = 0
	CASCADE_MIRROR_SERVICE                   = 1
	CASCADE_MIRROR_REGIONAL                  = 2
)

type CatacombStorageConfigs struct {
	provider string
	config   interface{}
}

type CatacombStorageSessions struct {
}

func init_storage(config CatacombStorageConfigs) *CatacombStorageSessions {

	switch config.provider {
	case "AWS_S3":
		fmt.Printf("Unsupported")
	case "AZURE_STORAGE":
		fmt.Printf("Unsupported")
	case "DO_VOLUMES":
		fmt.Printf("Unsupported")
	case "LINODE_STORAGE":
		fmt.Printf("Unsupported")
	default:
		fmt.Printf("Unsupported")
	}

	return nil
}

func verify_storage_config() {

}

func (sessions CatacombStorageSessions) storageObjectsList() {

}

func (sessions CatacombStorageSessions) storageObjectUpload(cascade CascadeOperation) {
	switch {

	}
}

func (sessions CatacombStorageSessions) storageObjectDownload() {
	switch {

	}
}

func (sessions CatacombStorageSessions) storageObjectDelete(cascade CascadeOperation) {
	switch {

	}
}
