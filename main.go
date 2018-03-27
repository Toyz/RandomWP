package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/Toyz/GoHaven"
	"github.com/Toyz/RandomWP/desktop"

	"github.com/Toyz/RandomWP/settings"
	"github.com/Toyz/RandomWP/wallpaper"
	"github.com/gen2brain/beeep"
	"github.com/gen2brain/dlgs"
	"github.com/marcsauter/single"
)

var (
	options = make([]GoHaven.Option, 0)

	conf *settings.Config
	// Run Once
	runOnce bool

	running bool

	CurrentVersion      string
	CurrentVersionShort string

	WallHaven *GoHaven.WallHaven
)

func main() {
	s := single.New("RandomWP")
	if err := s.CheckLock(); err != nil && err == single.ErrAlreadyRunning {
		log.Println("another instance of the app is already running, exiting")
		dlgs.Error("Already Running!", "Looks like another instance is already running!")
		os.Exit(0)
	} else if err != nil {
		// Another error occurred, might be worth handling it as well
		log.Fatalf("failed to acquire exclusive app lock: %v", err)
	}
	defer s.TryUnlock()

	WallHaven = GoHaven.New()

	ver, _ := Asset("assets/version.txt")
	runes := []rune(string(ver))
	CurrentVersionShort = string(runes[0:7])
	CurrentVersion = string(ver)

	confFolder := path.Join(desktop.GetDocumentsFolder(), "RandomWP")
	assetPath = path.Join(confFolder, "assets")
	createFolder(confFolder)
	createFolder(assetPath)

	saveAssets()

	conf = settings.New(path.Join(confFolder, "config.json"))

	createFolder(conf.SaveFolder)

	handleArgs()
	createOptions()

	rand.Seed(time.Now().Unix())

	if runOnce {
		changeWallpaper()
		os.Exit(0)
	}

	if conf.AutoStart {
		go startEndlessLoop()
	}
	setupTrayIcon()
}

func startEndlessLoop() {
	running = true
	for running {
		changeWallpaper()
		time.Sleep(time.Duration(conf.Delay) * time.Second)
	}
}

func changeWallpaper() {
	createOptions()
	var page GoHaven.Page
	page = 1
	options = append(options, page)
	options = append(options, GoHaven.SortRandom)

	havenIDs, _ := WallHaven.Search("", options...)
	background, err := wallpaper.Get()

	if isError(err) {
		return
	}

	if len(havenIDs) <= 0 {
		time.Sleep(30 * time.Second)
		changeWallpaper()
		return
	}
	currID := havenIDs[rand.Intn(len(havenIDs))]
	if conf.LastImageID != currID {
		conf.LastImageID = currID
	} else {
		for currID != conf.LastImageID {
			conf.LastImageID = havenIDs[rand.Intn(len(havenIDs))]
		}
	}

	conf.Save()

	fmt.Printf("Current wallpaper: %s\n", background)
	detail, _ := conf.LastImageID.Details()
	fmt.Println(detail)

	file, err := detail.Download(conf.SaveFolder)
	isError(err)

	fmt.Printf("New Wallpaper: %s\n", file)
	wallpaper.SetFromFile(file)

	if conf.Notify {
		err = beeep.Notify("Changed Wallpaper", "Your wallpaper has been changed", getAsset("information.png"))
		isError(err)
	}

	if conf.AutoDelete {
		deleteFile(file)
	}
}

func createOptions() {
	options = make([]GoHaven.Option, 0)
	options = append(options, conf.Category)
	options = append(options, conf.Purity)
	options = append(options, conf.Ratio)
	//options = append(options, res) // TODO: Fixme
}

func handleArgs() {
	flag.Var(&conf.Category, "cats", "Wallpaper categories (general,anime,people)")
	flag.Var(&conf.Purity, "purity", "Purity modes (sfw,sketchy)")
	//flag.Var(&conf, "res", "Screen resolutions (1024x768,1280x800,1366x768,1280x960,1440x900,1600x900,1280x1024,1600x1200,1680x1050,1920x1080,1920x1200,2560x1440,2560x1600,3840x1080,5760x1080,3840x2160)")
	flag.Var(&conf.Ratio, "ratios", "Aspect ratios (4x3,5x4,16x9,16x10,21x9,32x9,48x9)")
	flag.Int64Var(&conf.Delay, "delay", 3600, "Delay between background changes (in seconds)") // defaults to 1 hour
	flag.BoolVar(&runOnce, "once", false, "Only run the program once")
	flag.BoolVar(&conf.Notify, "notify", false, "Show notification when wallpaper changes")
	flag.Parse()
}
