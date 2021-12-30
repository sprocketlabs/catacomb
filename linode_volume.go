package catacomb

type LinodeVolumeConfig struct {
}

type LinodeVolumeSession struct {
}

func linodeNewSession(config LinodeVolumeConfig) LinodeVolumeSession {

	session := LinodeVolumeSession{}
	return session
}

func (session LinodeVolumeSession) linodeCloseSession() {

}

func (session LinodeVolumeSession) linodeVolumeListObjects() {

}

func (session LinodeVolumeSession) linodeVolumeObjectUpload() {

}

func (session LinodeVolumeSession) linodeVolumeObjectDownload() {

}

func (session LinodeVolumeSession) linodeVolumeObjectDelete() {

}
