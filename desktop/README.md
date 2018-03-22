# Desktop

Java desktop functions. Have you tried to find user default Download folder using java? If so, you would find this
library very helpful.

## Example Desktop Folders
    
```go
package main

import (
  "github.com/axet/desktop/go"
)

func main() {
  fmt.Println("Home:", desktop.GetHomeFolder())
  fmt.Println("Documents:" desktop.GetDocumentsFolder())
  fmt.Println("AppFolder:" desktop.GetAppDataFolder())
  fmt.Println("Desktop:" desktop.GetDesktopFolder())
  fmt.Println("Downloads:" desktop.GetDownloadsFolder())
}
```

## Example Sys Tray Icon
(aka Notification Area Icons or Status Bar icons)

Full example:

  * [systraytest.go](test/systraytest.go)

```go
package main

import (
	"fmt"
	"image"
	"os"

	"github.com/axet/desktop/go"
)

type SysTest struct {
	s *desktop.DesktopSysTray
}

func (m *SysTest) Click(mn *desktop.Menu) {
	fmt.Println("m", mn.Name)
}

func (m *SysTest) ClickBox(mn *desktop.Menu) {
	fmt.Println(mn.Name)
	mn.State = !mn.State
	m.s.Update()
}

func main() {
	m := SysTest{desktop.DesktopSysTrayNew()}

	file, err := os.Open("icon.png")
	if err != nil {
		panic(err)
	}
	icon, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	menu := []desktop.Menu{
		desktop.Menu{Icon: icon, Type: desktop.MenuItem, Enabled: true, Name: "test1", Action: m.Click},
		desktop.Menu{Type: desktop.MenuSeparator},
		desktop.Menu{Icon: icon, Type: desktop.MenuItem, Enabled: true, Name: "test2", Menu: []desktop.Menu{
			desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "test21", Action: m.Click},
			desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "test22", Action: m.Click},
		}},
		desktop.Menu{Type: desktop.MenuItem, Enabled: false, Name: "test3", Action: m.Click},
		desktop.Menu{Type: desktop.MenuCheckBox, Enabled: true, Name: "test4", State: true, Action: m.ClickBox},
		desktop.Menu{Type: desktop.MenuSeparator},
		desktop.Menu{Icon: icon, Type: desktop.MenuItem, Enabled: true, Name: "test5", Action: m.Click},
	}

	m.s.SetIcon(icon)
	m.s.SetTitle("go menu hoho!")
	m.s.SetMenu(menu)
	m.s.Show()

	desktop.Main()
}
```