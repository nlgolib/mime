package mime

import (
	"path/filepath"
	"strings"
)

// ExtToMime converts a file extension to a MIME type.
// The format of the extension should be like ".txt", ".html", etc.
func ExtToMime(ext string) string {
	ext = strings.TrimPrefix(ext, ".")
	ext = strings.ToLower(ext)
	ext = "." + ext
	mime, ok := mapExtToMime[ext]
	if !ok {
		return ""
	}
	return mime
}

// FilenameToMime converts a filename to a MIME type.
// The format of the filename should be like "example.txt", "example.html", etc.
func FilenameToMime(filename string) string {
	ext := filepath.Ext(filename)
	return ExtToMime(ext)
}

// MimeToExt converts a MIME type to a file extension.
// The format of the MIME type should be like "text/plain", "text/html", etc.
func MimeToExt(mime string) string {
	ext, ok := mapMimeToExt[mime]
	if !ok {
		return ""
	}
	return ext
}
