package ftpclient

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jlaffaye/ftp"
)

type FTPClient struct {
	server   string
	username string
	password string
	conn     *ftp.ServerConn
}

func NewFTPClient(server, username, password string) *FTPClient {
	return &FTPClient{
		server:   server,
		username: username,
		password: password,
	}
}

func (c *FTPClient) Connect() error {
	conn, err := ftp.Dial(fmt.Sprintf("%s:21", c.server))
	if err != nil {
		return err
	}
	err = conn.Login(c.username, c.password)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *FTPClient) Disconnect() {
	if c.conn != nil {
		c.conn.Quit()
	}
}

func (c *FTPClient) DownloadFiles(remoteDir, localDir string) error {
	c.conn.ChangeDir(remoteDir)
	entries, err := c.conn.List("")
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.Type == ftp.EntryTypeFile {
			fmt.Printf("Downloading %s/%s\n", localDir, entry.Name)

			r, err := c.conn.Retr(entry.Name)
			if err != nil {
				return err
			}
			defer r.Close()

			localFilePath := filepath.Join(localDir, entry.Name)
			file, err := os.Create(localFilePath)
			if err != nil {
				return err
			}

			_, err = file.ReadFrom(r)
			if err != nil {
				return err
			}

			file.Close()
			r.Close()
		}
	}
	return nil
}

// DeleteFile löscht eine Datei vom FTP-Server
func (c *FTPClient) DeleteFile(remoteFilePath string) error {
	err := c.conn.Delete(remoteFilePath)
	if err != nil {
		return fmt.Errorf("Fehler beim Löschen der Datei %s: %v", remoteFilePath, err)
	}
	return nil
}
