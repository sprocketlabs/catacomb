package catacomb

import (
	//"github.com/sprocketlabs/gocom"
	"github.com/pkg/sftp"
)

type SftpRemoteConfig struct {
	sftpPort    string
	sftpAddress string
}

type SftpRemoteSession struct {
	session *sftp.Client
}

/*
func sftpNewSession(config SftpRemoteConfig) SftpRemoteSession {
	user := ""
	pass := ""

	// get host public key
	hostKey := nil //getHostKey(config.sftpAddress)

	sshClientConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		// HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	// connect
	sftpHost := config.sftpAddress + ":" + config.sftpPort
	conn, err := ssh.Dial("tcp", sftpHost, sshClientConfig)
	if err != nil {
		log.Fatal(err)
	}

	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Close()

	session := SftpRemoteSession{
		session: client,
	}
	return session
}

func (session SftpRemoteSession) sftpCloseSession() {
	session.session.Close()
}

func (session SftpRemoteSession) sftpListObjects(remoteDir string) (theFiles []FileData, err error) {

	files, err := session.session.ReadDir(remoteDir)
	if err != nil {
		return theFiles, fmt.Errorf("Unable to list remote dir: %v", err)
	}

	for _, f := range files {
		var name, modTime, size string

		name = f.Name()
		modTime = f.ModTime().Format("2006-01-02 15:04:05")
		size = fmt.Sprintf("%12d", f.Size())

		if f.IsDir() {
			name = name + "/"
			modTime = ""
			size = "PRE"
		}

		theFiles = append(theFiles, FileData{
			Name:    name,
			Size:    size,
			ModTime: modTime,
		})
	}

	return theFiles, nil
}

func (session SftpRemoteSession) sftpUploadObject(localPath string, remotePath string) {
	// create destination file
	dstFile, err := session.session.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	// create source file
	srcFile, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	// copy source file to destination file
	bytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes copied\n", bytes)
}

func (session SftpRemoteSession) sftpDownloadObject() {
	// create destination file
	dstFile, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	// open source file
	srcFile, err := session.session.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	// copy source file to destination file
	bytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes copied\n", bytes)

	// flush in-memory copy
	err = dstFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
*/
