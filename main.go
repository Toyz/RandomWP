package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/Toyz/RandomWP/wallhaven"
	"github.com/Toyz/RandomWP/wallpaper"
	"github.com/gen2brain/beeep"
)

var (
	runOnce = flag.Bool("once", false, "Only run the program once")
)

func main() {
	flag.Parse()

	rand.Seed(time.Now().Unix())

	if *runOnce {
		changeWallpaper()
		return
	}

	for {
		changeWallpaper()
		time.Sleep(30 * time.Second)
	}
}

func changeWallpaper() {
	var page wallhaven.Page
	page.Set(strconv.Itoa(random(1, 3))) // between 1 or 2...

	havenIDs, _ := wallhaven.Search("", wallhaven.CatAnime, wallhaven.Ratio16x9, wallhaven.SortRandom, wallhaven.PuritySketchy, page)

	background, err := wallpaper.Get()

	if isError(err) {
		return
	}

	fmt.Printf("Current wallpaper: %s\n", background)
	file, _ := havenIDs[rand.Intn(len(havenIDs))].Download(os.TempDir())
	fmt.Printf("New Wallpaper: %s\n", file)
	wallpaper.SetFromFile(file)

	err = beeep.Notify("Changed Wallpaper", "Your wallpaper has been changed", "assets/information.png")
	isError(err)

	deleteFile(file)
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

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
