package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/GrosseBen/FTP2Paperless/internal/ftpclient"
	"github.com/GrosseBen/FTP2Paperless/internal/paperless"
)

func main() {
	// Umgebungsvariablen lesen oder Standardwerte verwenden
	ftpServer := getEnv("FTP_SERVER", "ftp.example.com")
	ftpUser := getEnv("FTP_USER", "consume")
	ftpPass := getEnv("FTP_PASS", "consume")
	remoteDir := getEnv("REMOTE_DIR", "/")
	localDir := getEnv("LOCAL_DIR", "./downloads")
	paperlessAPIURL := getEnv("PAPERLESS_API_URL", "http://localhost:8000/api/documents/post_document/")
	paperlessAPIKey := getEnv("PAPERLESS_API_KEY", "")
	ignoreSSLErrors := getEnv("IGNORE_SSL_ERRORS", "false")

	// Überprüfen, ob alle notwendigen Umgebungsvariablen gesetzt sind
	if ftpServer == "" || ftpUser == "" || ftpPass == "" || remoteDir == "" || localDir == "" || paperlessAPIURL == "" || paperlessAPIKey == "" {
		log.Fatal("Missing required environment variables")
	}

	// Erstelle das lokale Verzeichnis, falls es nicht existiert
	if _, err := os.Stat(localDir); os.IsNotExist(err) {
		err := os.MkdirAll(localDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating local directory: %v", err)
		}
	}

	// FTP-Client erstellen und verbinden
	client := ftpclient.NewFTPClient(ftpServer, ftpUser, ftpPass)
	err := client.Connect()
	if err != nil {
		log.Fatalf("Error connecting to FTP server: %v", err)
	}
	defer client.Disconnect()

	// HTTP-Client konfigurieren
	httpClient := &http.Client{}
	if ignoreSSLErrors == "true" {
		// SSL-Zertifikatsfehler ignorieren
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient = &http.Client{Transport: tr}
	}

	// Dateien aus Remote-Verzeichnis abrufen
	err = client.DownloadFiles(remoteDir, localDir)
	if err != nil {
		log.Fatalf("Error downloading files: %v", err)
	}

	// Dateien zu Paperless hochladen und dann löschen
	err = filepath.Walk(localDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// Versuche die Datei hochzuladen
			err := paperless.UploadToPaperless(path, paperlessAPIURL, paperlessAPIKey, httpClient)
			if err != nil {
				log.Printf("Fehler beim Hochladen der Datei %s zu Paperless: %v", path, err)
			} else {
				fmt.Printf("Datei erfolgreich zu Paperless hochgeladen: %s\n", path)
				// Lösche die Datei vom FTP-Server nach erfolgreichem Upload
				remoteFilePath := filepath.Join(remoteDir, info.Name())
				err = client.DeleteFile(remoteFilePath)
				if err != nil {
					log.Printf("Fehler beim Löschen der Datei %s vom FTP-Server: %v", remoteFilePath, err)
				} else {
					fmt.Printf("Datei %s erfolgreich vom FTP-Server gelöscht.\n", remoteFilePath)
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error processing local directory: %v", err)
	}

	fmt.Println("Alle Dateien erfolgreich heruntergeladen, hochgeladen und gelöscht.")
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
