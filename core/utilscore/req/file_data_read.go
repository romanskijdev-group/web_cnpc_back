package reqresutils

import (
	"cnpc_backend/core/typescore"
	"errors"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

func GetBaseFileInfoFromFromData(r *http.Request) (multipart.File, string, *typescore.WEvent) {
	log.Println("GetBaseFileInfoFromFromData")
	// Чтение файла из запроса
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Println("file not found: ", err)
		return nil, "", nil
	}
	defer func() {
		if file != nil {
			file.Close()
		}
	}()
	// Ограничение размера файла
	const maxFileSize = 5 << 20 // 5 MB
	if r.ContentLength > maxFileSize {
		return nil, "", &typescore.WEvent{
			Err:  errors.New("file_too_large"),
			Text: "file_too_large",
		}
	}
	if file != nil && header != nil {
		// Получаем расширение файла из оригинального имени файла
		ext := strings.ToLower(filepath.Ext(header.Filename))
		// Список разрешенных форматов файлов
		allowedExtensions := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}

		if _, allowed := allowedExtensions[ext]; !allowed {
			return nil, "", &typescore.WEvent{
				Err:  errors.New("invalid_file_format"),
				Text: "invalid_file_format",
			}
		}
		return file, ext, nil
	}
	return nil, "", nil
}
