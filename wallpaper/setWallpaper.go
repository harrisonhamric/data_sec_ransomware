package main

import (
	"fmt"

	"github.com/reujab/wallpaper"
)

func main() {
	err := wallpaper.SetFromURL("https://i.imgur.com/pIwrYeM.jpg")
	fmt.Println(err)
}
