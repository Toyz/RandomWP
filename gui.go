package main

import (
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/Toyz/GoHaven"
	"github.com/Toyz/RandomWP/desktop"
	"github.com/Toyz/RandomWP/updater"
	"github.com/gen2brain/dlgs"
	"github.com/skratchdot/open-golang/open"
)

var (
	sys *SysTest
)

type SysTest struct {
	S *desktop.DesktopSysTray
}

func setupTrayIcon() {
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
		desktop.Menu{Type: desktop.MenuCheckBox, State: conf.AutoStart, Enabled: true, Name: "Auto start on load", Action: sys.StopForeverRunning},
		desktop.Menu{Type: desktop.MenuSeparator},
		desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "Change Category", Menu: []desktop.Menu{
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Category == GoHaven.CatAnime, Enabled: true, Name: "Anime", Action: sys.ChageCategory},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Category == GoHaven.CatPeople, Enabled: true, Name: "People", Action: sys.ChageCategory},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Category == GoHaven.CatGeneral, Enabled: true, Name: "General", Action: sys.ChageCategory},
		}},
		desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "Change Purity", Menu: []desktop.Menu{
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Purity == GoHaven.PuritySFW, Enabled: true, Name: "SFW", Action: sys.ChageSafety},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Purity == GoHaven.PuritySketchy, Enabled: true, Name: "Sketchy", Action: sys.ChageSafety},
		}},
		// SFW,sketchy
		desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "Change Display Raio", Menu: []desktop.Menu{
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Ratio == GoHaven.Ratio4x3, Enabled: true, Name: "4x3", Action: sys.ChangeRatio},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Ratio == GoHaven.Ratio5x4, Enabled: true, Name: "5x4", Action: sys.ChangeRatio},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Ratio == GoHaven.Ratio16x9, Enabled: true, Name: "16x9", Action: sys.ChangeRatio},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Ratio == GoHaven.Ratio16x10, Enabled: true, Name: "16x10", Action: sys.ChangeRatio},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Ratio == GoHaven.Ratio21x9, Enabled: true, Name: "21x9", Action: sys.ChangeRatio},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Ratio == GoHaven.Ratio32x9, Enabled: true, Name: "32x9", Action: sys.ChangeRatio},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Ratio == GoHaven.Ratio48x9, Enabled: true, Name: "49x9", Action: sys.ChangeRatio},
		}},
		//4x3,5x4,16x9,16x10,21x9,32x9,48x9
		desktop.Menu{Type: desktop.MenuSeparator},
		desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "Save Current Image", Action: sys.SaveCurrentImage},
		desktop.Menu{Type: desktop.MenuSeparator},
		desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "Setting", Menu: []desktop.Menu{
			desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "Change current image folder", Action: sys.ChangeCurrentImageSaveFolder},
			desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "Change cache folder", Action: sys.ChangeImageCacheFolder},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.Notify, Enabled: true, Name: "Notify on Change", Action: sys.SendNotification},
			desktop.Menu{Type: desktop.MenuCheckBox, State: conf.AutoDelete, Enabled: true, Name: "Auto Delete Image", Action: sys.AutoDeleteImage},
			desktop.Menu{Type: desktop.MenuSeparator},
			desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "Check for Update", Action: sys.CheckForUpdate},
		}},
		desktop.Menu{Type: desktop.MenuSeparator},
		desktop.Menu{Type: desktop.MenuItem, Enabled: false, Name: fmt.Sprintf("Version: %s", CurrentVersionShort)},
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

func (m *SysTest) ChangeImageCacheFolder(mn *desktop.Menu) {
	folder, ok, err := dlgs.File("Select cache folder", "", true)
	if isError(err) {
		dlgs.Error("Something Happened", err.Error())
		return
	}

	if !ok {
		return
	}

	conf.SaveFolder = folder
	conf.Save()
	dlgs.Info("Changed Cache Folder", fmt.Sprintf("Cache folder has been set to:\n\n%s", folder))
}

func (m *SysTest) ChangeCurrentImageSaveFolder(mn *desktop.Menu) {
	folder, ok, err := dlgs.File("Select save current folder", "", true)
	if isError(err) {
		dlgs.Error("Something Happened", err.Error())
		return
	}

	if !ok {
		return
	}

	conf.SaveCurrentImageFolder = folder
	conf.Save()
	dlgs.Info("Changed Folder for saving", fmt.Sprintf("Folder for saving the current image is now at:\n\n%s", folder))
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
	var c GoHaven.Categories
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
	var c GoHaven.Purity
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
	var c GoHaven.Ratios
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
	yes, _ := dlgs.Question("Are you sure?", "Are you sure you wish to quit?", true)

	if yes {
		conf.Save()
		os.Exit(0)
	}

}

func (m *SysTest) StopForeverRunning(mn *desktop.Menu) {
	mn.State = !mn.State
	m.S.Update()

	if mn.State {
		conf.AutoStart = true
		go startEndlessLoop()
	} else {
		conf.AutoStart = false
		running = false
	}

	conf.Save()
}

func (m *SysTest) SaveCurrentImage(mn *desktop.Menu) {
	if conf.LastImageID <= 0 {
		dlgs.Error("Saved Image Failed", "Last image ID was less then zero\nThis happens when you first run!")
		return
	}
	go func() {
		details, _ := conf.LastImageID.Details()

		p, _ := details.Download(conf.SaveCurrentImageFolder)
		dlgs.Info("Saved Image", fmt.Sprintf("Saved image to:\n\n%s", p))
	}()
}

func (m *SysTest) CheckForUpdate(mn *desktop.Menu) {
	mn.Enabled = false
	go func(t *SysTest, m *desktop.Menu) {
		updater.GitUpdater.LoadStatus()

		if !updater.GitUpdater.IsStable() {
			m.Enabled = false
			t.S.Update()

			dlgs.Info("No New Version", "There is currently no new version!")
			return
		}

		remoteVer := updater.GitUpdater.GetSHASum()
		if !strings.EqualFold(remoteVer, strings.TrimSpace(CurrentVersion)) {
			ok, _ := dlgs.Question("New Version available", "Do you wish to goto the current build page?", true)
			if ok {
				open.Run("https://gitlab.com/Toyz/RandomWP/-/jobs/artifacts/master/browse?job=build")
			}
		} else {
			dlgs.Info("No New Version", "There is currently no new version!")
		}
		m.Enabled = false
		t.S.Update()
	}(m, mn)
	m.S.Update()
}
