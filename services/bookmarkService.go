package services

import (
	"encoding/json"
	"os"
)

type Bookmarks struct {
	List map[string]string
	fileService *FileService
	configService *ConfigService
	fileName string
}

var DEFAULT_FILE_NAME = "bookmarks.json"

// NewBookmarksService создаёт новый сервис закладок и загружает данные из файла (если он существует).
func NewBookmarksService(fileService *FileService, configService *ConfigService) Bookmarks {
	fileName := configService.GetEnvValue("BOOKMARKS_FILE", "bookmarks.json")

	bookmarks := Bookmarks{
		List: make(map[string]string),
		fileService: fileService,
		configService: configService,
		fileName: fileName,
	}
	bookmarks.LoadFromFile(fileName) // Загружаем данные из файла
	return bookmarks
}

// GetAll возвращает все закладки.
func (b *Bookmarks) GetAll() (map[string]string, error) {
	if err := b.fileService.EnsureLoaded(); err != nil {
		return nil, err
	}

	return b.List, nil
}

// Get возвращает закладку по имени.
func (b Bookmarks) Get(name string) (string, error) {
	if err := b.fileService.EnsureLoaded(); err != nil {
		return "", err
	}

	return b.List[name], nil
}

// Add добавляет новую закладку.
func (b *Bookmarks) Add(name, url string) (bool, error) {
	if err := b.fileService.EnsureLoaded(); err != nil {
		return false, err
	}

	if _, exists := b.List[name]; exists {
		return true, nil
	}

	b.List[name] = url
	b.SaveToFile(b.fileName) // Сохраняем данные в файл
	return false, nil
}

// Update обновляет существующую закладку.
func (b *Bookmarks) Update(name, url string) (bool, error) {
	if err := b.fileService.EnsureLoaded(); err != nil {
		return false, err
	}

	if _, exists := b.List[name]; !exists {
		return false, nil
	}
	b.List[name] = url
	b.SaveToFile(b.fileName) // Сохраняем данные в файл
	return true, nil
}

// Delete удаляет закладку по имени.
func (b *Bookmarks) Delete(name string) (bool, error) {
	if err := b.fileService.EnsureLoaded(); err != nil {
		return false, err
	}
	
	if _, exists := b.List[name]; !exists {
		return false, nil
	}
	delete(b.List, name)
	b.SaveToFile(b.fileName) // Сохраняем данные в файл
	return true, nil
}

// SaveToFile сохраняет закладки в файл.
func (b *Bookmarks) SaveToFile(filename string) error {
	data, err := json.Marshal(b.List) // Преобразуем map в JSON
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644) // Записываем JSON в файл
}

// LoadFromFile загружает закладки из файла.
func (b *Bookmarks) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &b.List)
}