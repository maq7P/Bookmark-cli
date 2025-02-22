package utils_colors

type Color string

const (
	White   Color = "white"
	Green   Color = "green"
	Blue  Color = "blue"
	Red     Color = "red"
	Magenta Color = "magenta"
	Default Color = "default"
)

var COLOR_CODES = map[Color]string{
	White:   "\033[36;47m",
	Green:   "\033[37;42m",
	Blue:  "\033[33;44m",
	Red: "\033[37;41m",
	Magenta: "\033[39;45m",
	Default: "",
}

// IsValidColor проверяет, является ли цвет допустимым.
func IsValidColor(c Color) bool {
	_, exists := COLOR_CODES[c]
	return exists
}

// GetColorKeys возвращает список всех допустимых цветов.
func GetColorKeys() []string {
	keys := make([]string, 0, len(COLOR_CODES))
	for key := range COLOR_CODES {
		keys = append(keys, string(key))
	}
	return keys
}

// GetColorCode возвращает ANSI-код для указанного цвета.
func GetColorCode(color Color) string {
	return COLOR_CODES[color]
}