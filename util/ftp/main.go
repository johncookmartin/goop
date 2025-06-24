package main

import (
	"flag"
	"log"
	"os"
)

func main() {

	username := flag.String("username", "", "FTP username")
	password := flag.String("password", "", "FTP password")
	remoteDir := flag.String("remoteDir", "./", "Remote directory to download")
	localDir := flag.String("localDir", ".", "Local directory to save files")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatalf("Usage: %s <host:port> [--username] [--password] [--remoteDir] [--localDir]", os.Args[0])
	}

	host := args[0]

	err := DownloadFTPDirectory(host, *username, *password, *remoteDir, *localDir)
	if err != nil {
		log.Fatalf("Error downloading FTP directory: %v", err)
	}

	log.Println("FTP directory downloaded successfully")
}