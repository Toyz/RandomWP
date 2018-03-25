package main

import (
	"image"
	"os"
	"strings"

	"github.com/meskio/dialog"

	"github.com/Toyz/RandomWP/desktop"
	"github.com/Toyz/RandomWP/wallhaven"
)

var (
	sys *SysTest
)

type SysTest struct {
	S *desktop.DesktopSysTray
}

func setupTrayIcon(forever bool) {
	sys = &SysTest{desktop.DesktopSysTrayNew()}

	file, err := os.Open(getAsset("icon.png"))
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
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Category == wallhaven.CatAnime, Enabled: true, Name: "Anime", Action: sys.ChageCategory},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Category == wallhaven.CatPeople, Enabled: true, Name: "People", Action: sys.ChageCategory},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Category == wallhaven.CatGeneral, Enabled: true, Name: "General", Action: sys.ChageCategory},
		}},
		desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "Change Purity", Menu: []desktop.Menu{
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Purity == wallhaven.PuritySFW, Enabled: true, Name: "SFW", Action: sys.ChageSafety},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Purity == wallhaven.PuritySketchy, Enabled: true, Name: "Sketchy", Action: sys.ChageSafety},
		}},
		// SFW,sketchy
		desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "Change Display Raio", Menu: []desktop.Menu{
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Ratio == wallhaven.Ratio4x3, Enabled: true, Name: "4x3", Action: sys.ChangeRatio},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Ratio == wallhaven.Ratio5x4, Enabled: true, Name: "5x4", Action: sys.ChangeRatio},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Ratio == wallhaven.Ratio16x9, Enabled: true, Name: "16x9", Action: sys.ChangeRatio},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Ratio == wallhaven.Ratio16x10, Enabled: true, Name: "16x10", Action: sys.ChangeRatio},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Ratio == wallhaven.Ratio21x9, Enabled: true, Name: "21x9", Action: sys.ChangeRatio},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Ratio == wallhaven.Ratio32x9, Enabled: true, Name: "32x9", Action: sys.ChangeRatio},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Ratio == wallhaven.Ratio48x9, Enabled: true, Name: "49x9", Action: sys.ChangeRatio},
		}},
		//4x3,5x4,16x9,16x10,21x9,32x9,48x9
		desktop.Menu{Type: desktop.MenuSeparator},
		desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "Save Current Image", Action: sys.SaveCurrentImage},
		desktop.Menu{Type: desktop.MenuSeparator},
		desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Notify, Enabled: true, Name: "Notify on Change", Action: sys.SendNotification},
		desktop.Menu{Type: desktop.MenuCheckBox, State: conf.AutoDelete, Enabled: true, Name: "Auto Delete Image", Action: sys.AutoDeleteImage},
		desktop.Menu{Type: desktop.MenuSeparator},
		desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "Quit", Action: sys.QuitProgram},
	}

	sys.S.SetIcon(icon)
	sys.S.SetTitle("RandomWP")
	sys.S.SetMenu(menu)
	sys.S.Show()

	desktop.Main()
}

func (m *SysTest) ChangeBackground(mn *desktop.Menu) {
	go changeWallpaper()
}

func (m *SysTest) SendNotification(mn *desktop.Menu) {
	mn.State = !mn.State
	conf.Notify = mn.State

	m.S.Update()
	conf.Save()
}

func (m *SysTest) AutoDeleteImage(mn *desktop.Menu) {
	mn.State = !mn.State
	conf.AutoDelete = mn.State

	m.S.Update()
	conf.Save()
}

func (m *SysTest) ChageCategory(mn *desktop.Menu) {
	var c wallhaven.Categories
	c.Set(strings.ToLower(mn.Name))
	conf.Category = c

	for i := 0; i < len(m.S.Menu[3].Menu); i++ {
		m.S.Menu[3].Menu[i].State = false
	}

	mn.State = true
	m.S.Update()

	conf.Save()
	createOptions()
}

func (m *SysTest) ChageSafety(mn *desktop.Menu) {
	var c wallhaven.Purity
	c.Set(strings.ToLower(mn.Name))
	conf.Purity = c

	for i := 0; i < len(m.S.Menu[4].Menu); i++ {
		m.S.Menu[4].Menu[i].State = false
	}

	mn.State = true
	m.S.Update()

	conf.Save()
	createOptions()
}

func (m *SysTest) ChangeRatio(mn *desktop.Menu) {
	var c wallhaven.Ratios
	c.Set(strings.ToLower(mn.Name))
	conf.Ratio = c

	for i := 0; i < len(m.S.Menu[5].Menu); i++ {
		m.S.Menu[5].Menu[i].State = false
	}

	mn.State = true
	m.S.Update()

	conf.Save()
	createOptions()
}

func (m *SysTest) QuitProgram(mn *desktop.Menu) {
	ok := dialog.Message("%s", "Are you sure you wish to quit?").Title("Are you sure?").YesNo()
	if ok {
		os.Exit(0)
	}
}

func (m *SysTest) StopForeverRunning(mn *desktop.Menu) {
	mn.State = !mn.State
	m.S.Update()

	if mn.State {
		go startEndlessLoop()
	} else {
		running = false
	}
}

func (m *SysTest) SaveCurrentImage(mn *desktop.Menu) {
	desktopFolder := desktop.GetDesktopFolder() // will be changed when settings are a thing

	go lastID.Download(desktopFolder)
}
