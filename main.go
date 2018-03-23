package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/Toyz/RandomWP/wallhaven"
	"github.com/Toyz/RandomWP/wallpaper"
	"github.com/gen2brain/beeep"
	"github.com/marcsauter/single"
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

	autoQuit bool

	lastID wallhaven.ID

	running bool
)

func main() {
	s := single.New("RandomWP")
	if err := s.CheckLock(); err != nil && err == single.ErrAlreadyRunning {
		log.Println("another instance of the app is already running, exiting")
		os.Exit(0)
	} else if err != nil {
		// Another error occurred, might be worth handling it as well
		log.Fatalf("failed to acquire exclusive app lock: %v", err)
	}
	defer s.TryUnlock()

	handleArgs()
	createOptions()

	rand.Seed(time.Now().Unix())
	if runOnce {
		if !autoQuit {
			go changeWallpaper()
			setupTrayIcon(false)
		} else {
			changeWallpaper()
		}
	} else {
		go startEndlessLoop()
		setupTrayIcon(true)
	}
}

func startEndlessLoop() {
	running = true

	for running {
		changeWallpaper()
		time.Sleep(time.Duration(delay) * time.Second)
	}
}

func changeWallpaper() {
	createOptions()
	var page wallhaven.Page
	page.Set(strconv.Itoa(random(1, 3))) // between 1 or 2...
	options = append(options, page)
	options = append(options, wallhaven.SortRandom)

	havenIDs, _ := wallhaven.Search("", options...)

	background, err := wallpaper.Get()

	if isError(err) {
		return
	}

	currID := havenIDs[rand.Intn(len(havenIDs))]
	if lastID != currID {
		lastID = currID
	} else {
		for currID != lastID {
			lastID = havenIDs[rand.Intn(len(havenIDs))]
		}
	}
	fmt.Printf("Current wallpaper: %s\n", background)
	file, _ := lastID.Download(os.TempDir())
	fmt.Printf("New Wallpaper: %s\n", file)
	wallpaper.SetFromFile(file)

	if notify {
		err = beeep.Notify("Changed Wallpaper", "Your wallpaper has been changed", "assets/information.png")
		isError(err)
	}

	deleteFile(file)
}

func createOptions() {
	options = make([]wallhaven.Option, 0)

	if cats != 0 {
		options = append(options, cats)
	} else {
		options = append(options, wallhaven.CatGeneral)
		cats = wallhaven.CatGeneral
	}

	if purity != 0 {
		options = append(options, purity)
	} else {
		options = append(options, wallhaven.PuritySFW)
		purity = wallhaven.PuritySFW
	}

	if ratios != 0 {
		options = append(options, ratios)
	} else {
		options = append(options, wallhaven.Ratio16x9)
		ratios = wallhaven.Ratio16x9
	}

	if res != 0 {
		options = append(options, res)
	}

}

func handleArgs() {
	flag.Var(&cats, "cats", "Wallpaper categories (general,anime,people)")
	flag.Var(&purity, "purity", "Purity modes (SFW,sketchy)")
	flag.Var(&res, "res", "Screen resolutions (1024x768,1280x800,1366x768,1280x960,1440x900,1600x900,1280x1024,1600x1200,1680x1050,1920x1080,1920x1200,2560x1440,2560x1600,3840x1080,5760x1080,3840x2160)")
	flag.Var(&ratios, "ratios", "Aspect ratios (4x3,5x4,16x9,16x10,21x9,32x9,48x9)")
	flag.Int64Var(&delay, "delay", 3600, "Delay between background changes (in seconds)") // defaults to 1 hour
	flag.BoolVar(&runOnce, "once", false, "Only run the program once")
	flag.BoolVar(&autoQuit, "quit", false, "Auto quit after task finishes (Only works with once)")
	flag.BoolVar(&notify, "notify", false, "Show notification when wallpaper changes")
	flag.Parse()
}
