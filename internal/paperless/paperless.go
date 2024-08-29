package paperless

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// UploadToPaperless lädt eine Datei zu Paperless NGX hoch
func UploadToPaperless(filePath, apiURL, apiKey string, client *http.Client) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Fehler beim Öffnen der Datei: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("document", filepath.Base(file.Name()))
	if err != nil {
		return fmt.Errorf("Fehler beim Erstellen des Formulars: %v", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("Fehler beim Kopieren der Datei: %v", err)
	}

	writer.Close()

	req, err := http.NewRequest("POST", apiURL, body)
	if err != nil {
		return fmt.Errorf("Fehler beim Erstellen der Anfrage: %v", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", apiKey))

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Fehler bei der Anfrage an Paperless: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("Paperless returned unexpected status code %d", resp.StatusCode)
	}

	return nil
}
