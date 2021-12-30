package catacomb

//import (
//	"github.com/Azure/azure-sdk-for-go"
//)

type AzureStorageConfig struct {
}

type AzureStorageSession struct {
}

func azureStorageNewSession(config AzureStorageConfig) AzureStorageSession {
	session := AzureStorageSession{}
	return session
}

func (session AzureStorageSession) azureStorageCloseSession() {

}

func (session AzureStorageSession) azureStorageListObjects() {

}

func (session AzureStorageSession) azureStorageObjectUpload() {

}

func (session AzureStorageSession) azureStorageObjectDownload() {

}

func (session AzureStorageSession) azureStorageObjectDelete() {

}
