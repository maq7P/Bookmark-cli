package cli

import (
	"fmt"
	"strings"

	"bookmarkapp/services"
	colors "bookmarkapp/utils/colors"
	console "bookmarkapp/utils/console"
)

var (
	Green  = colors.Green
	Blue = colors.Blue
	White  = colors.White
	Red = colors.Red
	Magenta= colors.Magenta
	Default = colors.Default
	Sleep  = console.Sleep
	Print = console.Println
	PrintlnWithColor = console.PrintlnWithColor
	PrintlnEmptyLine = console.PrintlnEmptyLine
	Log = console.Log
)
func UserInputValue[V console.InputValue](text string, color colors.Color) V {
	return console.UserInputValue[V](text, color)
}

type BookmarkCli struct {
	bookmarkService *services.Bookmarks
}

func NewBookmarkCli(bookmarkService *services.Bookmarks) *BookmarkCli {
	return &BookmarkCli{
		bookmarkService: bookmarkService,
	}
}

func (c *BookmarkCli) Run() {
	iterationsCount := 0
	
	for {
		if iterationsCount != 0 {
			userInput := UserInputValue[string]("Show menu panel again? (y/n): ", White)
			
			if strings.ToLower(userInput) != "y" {
				break
			}
		}
		iterationsCount++

		PrintlnWithColor("What do you want to do?", Magenta)
		PrintlnWithColor("1. Show all bookmarks", Default)
		PrintlnWithColor("2. Add bookmark", Default)
		PrintlnWithColor("3. Get bookmark", Default)
		PrintlnWithColor("4. Delete bookmark", Default)
		PrintlnWithColor("5. Exit", Default)
		PrintlnEmptyLine()

		userInput := UserInputValue[int]("Enter your choice number: ", White)
		PrintlnEmptyLine()

		switch userInput {
			case 1:
				c.showAllBookmarks()
			case 2:
				c.addBookmark()
			case 3:
				c.getBookmark()
			case 4:
				c.deleteBookmark()
			case 5:
				PrintlnWithColor("Exiting...", Green)
				PrintlnEmptyLine()
				return
			default:
				PrintlnWithColor("Invalid choice. Please try again.", Green)
				PrintlnEmptyLine()
			}
	}
}

// showAllBookmarks выводит все закладки.
func (c *BookmarkCli) showAllBookmarks() {
	all, err := c.bookmarkService.GetAll()

	if err != nil {
			Log("Error getting bookmarks:", err)
	}

	if len(all) == 0 {
		PrintlnWithColor("No bookmarks yet", Red)
		PrintlnEmptyLine()

		return
	}

	Print("---------Your bookmarks------------")
	counter := 1
	for key, value := range all {
		PrintlnWithColor(fmt.Sprintf("%d) %s | %s", counter, key, value), Blue)
		counter++
	}
	Print("----------------------")
	PrintlnEmptyLine()
}

// addBookmark добавляет новую закладку.
func (c *BookmarkCli) addBookmark() {
	name := UserInputValue[string]("Enter short name of URL: ", Blue)
	url := UserInputValue[string]("Enter URL: ", Blue)
	PrintlnEmptyLine()

	exists, err := c.bookmarkService.Add(name, url)

	if err != nil {
			Log("Error getting bookmarks:", err)
	}

	if exists {
		PrintlnWithColor("Bookmark already exists.", Red)
		userInput := UserInputValue[string]("Do you want to replace it? (y/n): ", Default)

		if strings.ToLower(userInput) == "y" {
			c.bookmarkService.Update(name, url)
			PrintlnWithColor("Bookmark has been updated.", Green)
			PrintlnEmptyLine()
		}
	} else {
		PrintlnWithColor("Bookmark has been added.", Green)
		PrintlnEmptyLine()
	}
}

// getBookmark получает закладку по имени.
func (c *BookmarkCli) getBookmark() {
	name := UserInputValue[string]("Enter short name of URL: ", White)

	bookmark, err := c.bookmarkService.Get(name)

	if err != nil {
			Log("Error getting bookmarks:", err)
	}

	if bookmark == "" {
		PrintlnWithColor("Bookmark not found.", Red)
		return
	}

	PrintlnWithColor(fmt.Sprintf("Bookmark: %s", bookmark), Green)
}

// deleteBookmark удаляет закладку по имени.
func (c *BookmarkCli) deleteBookmark() {
	name := UserInputValue[string]("Enter short name of URL: ", White)

	successDeleted, err := c.bookmarkService.Delete(name)

	if err != nil {
			Log("Error getting bookmarks:", err)
	}

	if successDeleted {
		PrintlnWithColor("Bookmark successfully deleted.", Green)
		PrintlnEmptyLine()
	} else {
		PrintlnWithColor("Bookmark not found.", Red)
		PrintlnEmptyLine()
	}
}