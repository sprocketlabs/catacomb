package catacomb

type DoStorageConfig struct {
}

type DoStorageSession struct {
}

func doVolumeNewSession(config LinodeVolumeConfig) LinodeVolumeSession {

	session := LinodeVolumeSession{}
	return session
}

func (session LinodeVolumeSession) doVolumeCloseSession() {

}

func (session DoStorageSession) doVolumeListObjects() {

}

func (session DoStorageSession) doVolumeObjectUpload() {

}

func (session DoStorageSession) doVolumeObjectDownload() {

}

func (session DoStorageSession) doVolumeObjectDelete() {

}
