# FTP2Paperless

FTP2Paperless ist ein Python-basiertes Tool, das entwickelt wurde, um Dokumente von einem FTP-Server herunterzuladen und diese über die API in Paperless-ngx zu importieren. Dieses Tool automatisiert den Prozess der Dokumentenübertragung und -speicherung, was die Verwaltung von Dokumenten in einer digitalen Umgebung erleichtert.

## Funktionen

- Herunterladen von Dokumenten von einem FTP-Server.
- Hochladen von Dokumenten zur Paperless-ngx API.
- Automatisierung dieser Prozesse durch Nutzung von Umgebungsvariablen und Docker.

## Voraussetzungen

- Python 3.9 oder höher
- Docker (für die Verwendung von Docker-Containern)
- Zugang zu einem FTP-Server
- Zugang zu einer Paperless-ngx Instanz mit API-Unterstützung

## Installation

1. **Projekt klonen**:

   ```bash
   git clone https://github.com/dein-benutzername/ftp2paperless.git
   cd ftp2paperless
   ```

2. **Virtuelle Umgebung erstellen und aktivieren** (optional aber empfohlen):

   ```bash
   python3 -m venv venv
   source venv/bin/activate
   ```

3. **Abhängigkeiten installieren**:

   ```bash
   pip install -r requirements.txt
   ```

## Konfiguration

FTP2Paperless nutzt Umgebungsvariablen für Konfigurationseinstellungen. Du kannst diese Variablen entweder in deiner Shell setzen oder eine `.env`-Datei verwenden (mit `python-dotenv`).

Erforderliche Umgebungsvariablen:

- `FTP_SERVER`: Adresse des FTP-Servers (z.B. `ftp.example.com`)
- `FTP_USER`: Benutzername für den FTP-Zugang
- `FTP_PASS`: Passwort für den FTP-Zugang
- `REMOTE_DIR`: Verzeichnis auf dem FTP-Server, von dem Dateien heruntergeladen werden sollen
- `LOCAL_DIR`: Lokales Verzeichnis, in dem die heruntergeladenen Dateien gespeichert werden
- `API_URL`: URL zur Paperless-ngx API (z.B. `http://example.com/api/documents/post_document/`)
- `API_TOKEN`: API-Token für die Authentifizierung bei Paperless-ngx

Beispiel `.env` Datei:

```env
FTP_SERVER=ftp.example.com
FTP_USER=username
FTP_PASS=password
REMOTE_DIR=/remote/path/
LOCAL_DIR=/tmp/ftp_downloads
API_URL=http://example.com/api/documents/post_document/
API_TOKEN=your_api_token_here
```

## Verwendung

### Direktes Ausführen des Python-Skripts

Stelle sicher, dass alle erforderlichen Umgebungsvariablen gesetzt sind und führe das Skript aus:

```bash
python -m ftp2paperless.core
```

### Verwendung mit Docker

1. **Docker-Image erstellen**:

   ```bash
   docker build -t ftp2paperless .
   ```

2. **Docker-Container ausführen**:

   ```bash
   docker run --rm \
       -e FTP_SERVER=ftp.example.com \
       -e FTP_USER=username \
       -e FTP_PASS=password \
       -e REMOTE_DIR=/remote/path/ \
       -e LOCAL_DIR=/tmp/ftp_downloads \
       -e API_URL=http://example.com/api/documents/post_document/ \
       -e API_TOKEN=your_api_token_here \
       -v /path/to/local/downloads:/tmp/ftp_downloads \
       ftp2paperless
   ```

Dieser Befehl startet den Container, setzt die erforderlichen Umgebungsvariablen und mountet ein lokales Verzeichnis, um heruntergeladene Dateien zu speichern.

## Fehlerbehandlung

- **Verbindungsprobleme zum FTP-Server**: Überprüfe die FTP-Serveradresse, den Benutzernamen und das Passwort.
- **Fehler beim Hochladen zu Paperless-ngx**: Stelle sicher, dass die API-URL und das API-Token korrekt sind und dass die API-Instanz erreichbar ist.

## Mitwirkende

- Dein Name (oder die Mitwirkenden)

## Lizenz

Dieses Projekt ist unter der MIT-Lizenz lizenziert – siehe die [LICENSE](LICENSE) Datei für Details.

## Kontakt

Falls du Fragen hast, melde dich gerne bei uns: [deine-email@example.com](mailto:deine-email@example.com)
