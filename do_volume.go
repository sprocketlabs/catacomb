package catacomb

type DoStorageConfig struct {
}

type DoStorageSession struct {
}

func digitaloceanNewSession(config LinodeVolumeConfig) LinodeVolumeSession {

	session := LinodeVolumeSession{}
	return session
}

func (session LinodeVolumeSession) digitaloceanCloseSession() {

}

func (session DoStorageSession) digitaloceanListObjects() {

}

func (session DoStorageSession) digitaloceanObjectUpload() {

}

func (session DoStorageSession) digitaloceanObjectDownload() {

}

func (session DoStorageSession) digitaloceanObjectDelete() {

}
