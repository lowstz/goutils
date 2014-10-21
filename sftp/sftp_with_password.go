package sftp

import (
	"bufio"
	"code.google.com/p/go.crypto/ssh"
	"github.com/pkg/sftp"
	"io"
	"os"
)

func GetSFTPClient(domain string, user string, password string) (*sftp.Client, error) {
	clientConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(password)},
	}
	ssh_client, err := ssh.Dial("tcp", domain, clientConfig)
	if err != nil {
		return nil, err
	}

	client, err := sftp.NewClient(ssh_client)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetSFTPAllFilesName(path string, domain string, user string, password string) ([]string, error) {
	client, err := GetSFTPClient(domain, user, password)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	var results []string
	walker := client.Walk(path)
	for walker.Step() {
		if err := walker.Err(); err != nil {
			continue
		}
		results = append(results, walker.Path())
	}
	return results, nil
}
func ReadSFTPFile(path string, domain string, user string, password string) (*sftp.File, error) {
	client, err := GetSFTPClient(domain, user, password)
	if err != nil {
		return nil, err
	}

	sftpFile, err := client.Open(path)
	if err != nil {
		return nil, err
	}

	return sftpFile, err

}

func SaveSFTPFilesToLocal(sftpPath string, localPath string, domain string, user string, password string) (bool, error) {
	sftpFile, err := ReadSFTPFile(sftpPath, domain, user, password)
	if err != nil {
		return false, err
	}
	defer sftpFile.Close()

	localFile, err := os.Create(localPath)
	if err != nil {
		return false, err
	}

	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := sftpFile.Read(buf)
		if err != nil && err != io.EOF {
			return false, err
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := localFile.Write(buf[:n]); err != nil {
			return false, err
		}
	}
	defer localFile.Close()
	return true, nil
}

func ReadSFTPFileToString(path string, domain string, user string, password string) ([]string, error) {
	sftpFile, err := ReadSFTPFile(path, domain, user, password)
	if err != nil {
		return nil, err
	}
	var lines []string
	scanner := bufio.NewScanner(sftpFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	defer sftpFile.Close()
	return lines, scanner.Err()
}

func LeaveSFTPAFile(path string, content string, domain string, user string, password string) error {
	client, err := GetSFTPClient(domain, user, password)
	if err != nil {
		return err
	}
	defer client.Close()

	f, err := client.Create(path)
	if err != nil {
		return err
	}
	if _, err := f.Write([]byte(content)); err != nil {
		return err
	}
	defer f.Close()
	return nil
}
func LeaveSFTPAFileWithByte(path string, content []byte, domain string, user string, password string) error {
	client, err := GetSFTPClient(domain, user, password)
	if err != nil {
		return err
	}
	defer client.Close()

	f, err := client.Create(path)
	if err != nil {
		return err
	}
	if _, err := f.Write(content); err != nil {
		return err
	}
	defer f.Close()
	return nil
}
