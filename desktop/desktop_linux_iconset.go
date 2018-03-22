// +build linux

package desktop

import (
	"bufio"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

var rand uint32
var randmu sync.Mutex

func reseed() uint32 {
	return uint32(time.Now().UnixNano() + int64(os.Getpid()))
}

func nextSuffix() string {
	randmu.Lock()
	r := rand
	if r == 0 {
		r = reseed()
	}
	r = r*1664525 + 1013904223 // constants from Numerical Recipes
	rand = r
	randmu.Unlock()
	return strconv.Itoa(int(1e9 + r%1e9))[1:]
}

func TempFile(dir, prefix string, suffix string) (f *os.File, err error) {
	if dir == "" {
		dir = os.TempDir()
	}

	nconflict := 0
	for i := 0; i < 10000; i++ {
		name := filepath.Join(dir, prefix+nextSuffix()+suffix)
		f, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
		if os.IsExist(err) {
			if nconflict++; nconflict > 10 {
				randmu.Lock()
				rand = reseed()
				randmu.Unlock()
			}
			continue
		}
		break
	}
	return
}

type GtkIconSet struct {
	Path  string
	icons map[image.Image]string
}

func GtkIconSetNew() *GtkIconSet {
	m := &GtkIconSet{}
	m.icons = make(map[image.Image]string)
	var err error
	m.Path, err = ioutil.TempDir("", "systray")
	if err != nil {
		panic(err)
	}
	return m
}

func (m *GtkIconSet) Add(icon image.Image) string {
	if p, ok := m.icons[icon]; ok {
		return p
	}

	format := "png"
	suffix := "." + format
	f, err := TempFile(m.Path, "systray", suffix)
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)
	defer w.Flush()
	path := f.Name()
	png.Encode(w, icon)
	name := filepath.Base(path)
	name = strings.TrimRight(name, suffix)
	m.icons[icon] = name
	return name
}

func (m *GtkIconSet) Close() {
	os.RemoveAll(m.Path)
}
