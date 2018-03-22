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
	options = make([]wallhaven.Option, 0)
	// cats specifies the enabled wallpaper categories.
	cats wallhaven.Categories
	// purity specifies the enabled purity modes.
	purity wallhaven.Purity
	// res specifies the enabled screen resolutions.
	res wallhaven.Resolutions
	// ratios specifies the enabled aspect rations.
	ratios wallhaven.Ratios
	// Run Once
	runOnce bool
	// Delay for loop between changes
	delay int64
	// Send notifcations (default false)
	notify bool
)

func main() {
	flag.Var(&cats, "cats", "Wallpaper categories (general,anime,people)")
	flag.Var(&purity, "purity", "Purity modes (SFW,sketchy)")
	flag.Var(&res, "res", "Screen resolutions (1024x768,1280x800,1366x768,1280x960,1440x900,1600x900,1280x1024,1600x1200,1680x1050,1920x1080,1920x1200,2560x1440,2560x1600,3840x1080,5760x1080,3840x2160)")
	flag.Var(&ratios, "ratios", "Aspect ratios (4x3,5x4,16x9,16x10,21x9,32x9,48x9)")
	flag.Int64Var(&delay, "delay", 30, "Delay between background changes (in seconds)")
	flag.BoolVar(&runOnce, "once", false, "Only run the program once")
	flag.BoolVar(&notify, "notify", false, "Show notification when wallpaper changes")
	flag.Parse()

	if cats != 0 {
		options = append(options, cats)
	}
	if purity != 0 {
		options = append(options, purity)
	}
	if res != 0 {
		options = append(options, res)
	}
	if ratios != 0 {
		options = append(options, ratios)
	}

	rand.Seed(time.Now().Unix())
	if runOnce {
		changeWallpaper()
		return
	}

	for {
		changeWallpaper()
		time.Sleep(time.Duration(delay) * time.Second)
	}
}

func changeWallpaper() {
	var page wallhaven.Page
	page.Set(strconv.Itoa(random(1, 3))) // between 1 or 2...
	options = append(options, page)
	options = append(options, wallhaven.SortRandom)

	havenIDs, _ := wallhaven.Search("", options...)

	background, err := wallpaper.Get()

	if isError(err) {
		return
	}

	fmt.Printf("Current wallpaper: %s\n", background)
	file, _ := havenIDs[rand.Intn(len(havenIDs))].Download(os.TempDir())
	fmt.Printf("New Wallpaper: %s\n", file)
	wallpaper.SetFromFile(file)

	if notify {
		err = beeep.Notify("Changed Wallpaper", "Your wallpaper has been changed", "assets/information.png")
		isError(err)
	}

	deleteFile(file)
}

func deleteFile(path string) {
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("==> deleted temp wallpaper")
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
