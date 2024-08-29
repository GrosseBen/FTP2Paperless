# Verwende ein minimalistisches Go-Basisimage, um das Programm zu bauen
FROM golang:1.23-alpine AS builder

# Setze das Arbeitsverzeichnis im Container
WORKDIR /app

# Kopiere go.mod und go.sum ins Arbeitsverzeichnis
COPY go.mod go.sum ./

# Installiere die Abhängigkeiten
RUN go mod download

# Kopiere den Rest des Codes ins Arbeitsverzeichnis
COPY . .

# Baue das Go-Binary
RUN go build -o ftp2paperless ./cmd/ftp2paperless

# Verwende ein minimalistisches Basisimage für die endgültige Ausführung
FROM alpine:latest

# Setze das Arbeitsverzeichnis im Container
WORKDIR /app

# Kopiere das Go-Binary aus dem Build-Container
COPY --from=builder /app/ftp2paperless .

# Setze die Standard-Umgebungsvariablen (optional)
#ENV FTP_SERVER="ftp.example.com"
#ENV FTP_USER="consume"
#ENV FTP_PASS="consume"
#ENV REMOTE_DIR="/"
#ENV LOCAL_DIR="./downloads"
#ENV PAPERLESS_API_URL="http://localhost:8000/api/documents/post_document/"
#ENV PAPERLESS_API_KEY="your_api_key_here"
#ENV IGNORE_SSL_ERRORS="false"

# Definiere den Einstiegspunkt
ENTRYPOINT ["./ftp2paperless"]
