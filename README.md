# Bookmark App

Это простое CLI для управления закладками. Вы можете добавлять, просматривать, обновлять и удалять закладки. Все закладки сохраняются в файл, что позволяет сохранять их между запусками приложения.

---

## Оглавление

1. [Установка](#установка)
2. [Использование](#использование)

---

## Установка

1. Поддерживает Go версии 1.16 или выше
2. Сборка приложения:
```go build -o bookmarkapp```
3. Добавление env переменной для кастомного названия файла (только json)
```export BOOKMARKS_FILE="my_bookmarks.json"```
4. Запуск
```./bookmarkapp```