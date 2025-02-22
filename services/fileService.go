package services

import (
	"encoding/json"
	"os"
	"sync"
)

// FileService отвечает за работу с файловой системой.
type FileService struct {
	data   interface{} // Данные, которые загружаются из файла
	loaded bool        // Флаг, указывающий, загружены ли данные
	mutex  sync.Mutex  // Мьютекс для потокобезопасности
	file   string      // Имя файла
}

// NewFileService создаёт новый экземпляр FileService.
func NewFileService(file string, data interface{}) *FileService {
	return &FileService{
		data: data,
		file: file,
	}
}

// EnsureLoaded загружает данные из файла, если они ещё не загружены.
func (fs *FileService) EnsureLoaded() error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	if fs.loaded {
		return nil
	}

	err := fs.ReadFile(fs.file, fs.data)
	if err != nil {
		return err
	}

	fs.loaded = true
	return nil
}

// ReadFile читает данные из файла и десериализует их в указанный объект.
func (fs *FileService) ReadFile(filename string, target interface{}) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Файл не существует, это нормально
		}
		return err
	}
	return json.Unmarshal(data, target)
}

// WriteFile сериализует данные и записывает их в файл.
func (fs *FileService) WriteFile(filename string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, jsonData, 0644)
}