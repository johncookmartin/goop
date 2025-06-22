package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 5 {
		log.Fatalf("Usage: %s <host:port> <username> <password> <remoteDir>", os.Args[0])
	}

	host := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]
	remoteDir := os.Args[4]
	localDir := "."

	err := DownloadFTPDirectory(host, username, password, remoteDir, localDir)
	if err != nil {
		log.Fatalf("Error downloading FTP directory: %v", err)
	}

	log.Println("FTP directory downloaded successfully")
}