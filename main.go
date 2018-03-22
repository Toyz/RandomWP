package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Toyz/RandomWP/wallhaven"
	"github.com/Toyz/RandomWP/wallpaper"
)

func main() {
	for {
		havenIDs, _ := wallhaven.Search("anime", wallhaven.CatAnime, wallhaven.Ratio16x9, wallhaven.SortRandom, wallhaven.PuritySketchy)

		background, err := wallpaper.Get()

		if isError(err) {
			continue
		}

		fmt.Printf("Current wallpaper: %s\n", background)
		file, _ := havenIDs[0].Download(os.TempDir())
		fmt.Printf("New Wallpaper: %s\n", file)
		wallpaper.SetFromFile(file)
		deleteFile(file) // will get set at loop start
		time.Sleep(30 * time.Second)
	}
}

func deleteFile(path string) {
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("==> done deleting file")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
