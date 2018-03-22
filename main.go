package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Toyz/RandomWP/desktop"
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

	runningMenuOption *desktop.Menu

	sys *SysTest
)

type SysTest struct {
	S *desktop.DesktopSysTray
}

func (m *SysTest) ChangeBackground(mn *desktop.Menu) {
	changeWallpaper()
}

func (m *SysTest) ChageCategory(mn *desktop.Menu) {
	var c wallhaven.Categories
	c.Set(strings.ToLower(mn.Name))
	cats = c

	for i := 0; i < len(m.S.Menu[3].Menu); i++ {
		m.S.Menu[3].Menu[i].State = false
	}

	mn.State = true
	m.S.Update()

	createOptions()
}

func (m *SysTest) ChageSafety(mn *desktop.Menu) {
	var c wallhaven.Purity
	c.Set(strings.ToLower(mn.Name))
	purity = c

	for i := 0; i < len(m.S.Menu[4].Menu); i++ {
		m.S.Menu[4].Menu[i].State = false
	}

	mn.State = true
	m.S.Update()

	createOptions()
}

func (m *SysTest) ChangeRatio(mn *desktop.Menu) {
	var c wallhaven.Ratios
	c.Set(strings.ToLower(mn.Name))
	ratios = c

	for i := 0; i < len(m.S.Menu[5].Menu); i++ {
		m.S.Menu[5].Menu[i].State = false
	}

	mn.State = true
	m.S.Update()

	createOptions()
}

func (m *SysTest) QuitProgram(mn *desktop.Menu) {
	os.Exit(0)
}

func (m *SysTest) StopForeverRunning(mn *desktop.Menu) {
	mn.State = !mn.State
	m.S.Update()

	if mn.State {
		go DoForeverLoop()
	} else {
		running = false
	}
}

func main() {
	s := single.New("RandomWP")
	if err := s.CheckLock(); err != nil && err == single.ErrAlreadyRunning {
		log.Fatal("another instance of the app is already running, exiting")
		os.Exit(0)
	} else if err != nil {
		// Another error occurred, might be worth handling it as well
		log.Fatalf("failed to acquire exclusive app lock: %v", err)
	}
	defer s.TryUnlock()

	flag.Var(&cats, "cats", "Wallpaper categories (general,anime,people)")
	flag.Var(&purity, "purity", "Purity modes (SFW,sketchy)")
	flag.Var(&res, "res", "Screen resolutions (1024x768,1280x800,1366x768,1280x960,1440x900,1600x900,1280x1024,1600x1200,1680x1050,1920x1080,1920x1200,2560x1440,2560x1600,3840x1080,5760x1080,3840x2160)")
	flag.Var(&ratios, "ratios", "Aspect ratios (4x3,5x4,16x9,16x10,21x9,32x9,48x9)")
	flag.Int64Var(&delay, "delay", 30, "Delay between background changes (in seconds)")
	flag.BoolVar(&runOnce, "once", false, "Only run the program once")
	flag.BoolVar(&autoQuit, "quit", false, "Auto quit after task finishes (Only works with once)")
	flag.BoolVar(&notify, "notify", false, "Show notification when wallpaper changes")
	flag.Parse()
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
		go DoForeverLoop()
		setupTrayIcon(true)
	}
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

func DoForeverLoop() {
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

func setupTrayIcon(forever bool) {
	sys = &SysTest{desktop.DesktopSysTrayNew()}

	file, err := os.Open("assets/icon.png")
	if err != nil {
		panic(err)
	}
	icon, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	menu := []desktop.Menu{
		desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "Change Background", Action: sys.ChangeBackground},
		desktop.Menu{Type: desktop.MenuCheckBox, State: forever, Enabled: true, Name: "Run Forever", Action: sys.StopForeverRunning},
		desktop.Menu{Type: desktop.MenuSeparator},
		desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "Change Category", Menu: []desktop.Menu{
			desktop.Menu{Type: desktop.MenuCheckBox, State: cats == wallhaven.CatAnime, Enabled: true, Name: "Anime", Action: sys.ChageCategory},
			desktop.Menu{Type: desktop.MenuCheckBox, State: cats == wallhaven.CatPeople, Enabled: true, Name: "People", Action: sys.ChageCategory},
			desktop.Menu{Type: desktop.MenuCheckBox, State: cats == wallhaven.CatGeneral, Enabled: true, Name: "General", Action: sys.ChageCategory},
		}},
		desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "Change Purity", Menu: []desktop.Menu{
			desktop.Menu{Type: desktop.MenuCheckBox, State: purity == wallhaven.PuritySFW, Enabled: true, Name: "SFW", Action: sys.ChageSafety},
			desktop.Menu{Type: desktop.MenuCheckBox, State: purity == wallhaven.PuritySketchy, Enabled: true, Name: "Sketchy", Action: sys.ChageSafety},
		}},
		// SFW,sketchy
		desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "Change Display Raio", Menu: []desktop.Menu{
			desktop.Menu{Type: desktop.MenuCheckBox, State: ratios == wallhaven.Ratio4x3, Enabled: true, Name: "4x3", Action: sys.ChangeRatio},
			desktop.Menu{Type: desktop.MenuCheckBox, State: ratios == wallhaven.Ratio5x4, Enabled: true, Name: "5x4", Action: sys.ChangeRatio},
			desktop.Menu{Type: desktop.MenuCheckBox, State: ratios == wallhaven.Ratio16x9, Enabled: true, Name: "16x9", Action: sys.ChangeRatio},
			desktop.Menu{Type: desktop.MenuCheckBox, State: ratios == wallhaven.Ratio16x10, Enabled: true, Name: "16x10", Action: sys.ChangeRatio},
			desktop.Menu{Type: desktop.MenuCheckBox, State: ratios == wallhaven.Ratio21x9, Enabled: true, Name: "21x9", Action: sys.ChangeRatio},
			desktop.Menu{Type: desktop.MenuCheckBox, State: ratios == wallhaven.Ratio32x9, Enabled: true, Name: "32x9", Action: sys.ChangeRatio},
			desktop.Menu{Type: desktop.MenuCheckBox, State: ratios == wallhaven.Ratio48x9, Enabled: true, Name: "49x9", Action: sys.ChangeRatio},
		}},
		//4x3,5x4,16x9,16x10,21x9,32x9,48x9
		desktop.Menu{Type: desktop.MenuSeparator},
		desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "Quit", Action: sys.QuitProgram},
	}

	sys.S.SetIcon(icon)
	sys.S.SetTitle("RandomWP")
	sys.S.SetMenu(menu)
	sys.S.Show()

	desktop.Main()
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
