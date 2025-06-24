package main

import (
	"io"
	"log"
	"os"
	"path"
	"time"

	"github.com/jlaffaye/ftp"
)

func DownloadFTPDirectory(host, username, password, remoteDir, localDir string) error {
	conn, err := ftp.Dial(host, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return err
	}
	defer conn.Quit()

	err = conn.Login(username, password)
	if err != nil {
		return err
	}
	return downloadDir(conn, remoteDir, localDir)
}

func downloadDir(conn *ftp.ServerConn, remoteDir, localDir string) error {
	entries, err := conn.List(remoteDir)
	if err != nil {
		return err
	}

	err = os.MkdirAll(localDir, 0755)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		remotePath := path.Join(remoteDir, entry.Name)
		localPath := path.Join(localDir, entry.Name)

		log.Printf("Downloading %s to %s", remotePath, localPath)
		if entry.Type == ftp.EntryTypeFile {
			err = downloadFile(conn, remotePath, localPath)
			if err != nil {
				log.Printf("Error downloading file %s: %v", remotePath, err)
			}
		} else if entry.Type == ftp.EntryTypeFolder {
			err = downloadDir(conn, remotePath, localPath)
			if err != nil {
				log.Printf("Error downloading directory %s: %v", remotePath, err)
			}
		}
		
	}

	return nil
}

func downloadFile(conn *ftp.ServerConn, remotePath, localPath string) error {
	response, err := conn.Retr(remotePath)
	if err != nil {
		return err
	}
	defer response.Close()

	outFile, err := os.Create(localPath)
	if err != nil {
		return err
	}

	_, err = io.Copy(outFile, response)
	return err
}