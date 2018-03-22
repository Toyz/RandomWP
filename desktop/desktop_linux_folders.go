// +build linux

package desktop

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

const ROOT_FILE string = "/etc/xdg/user-dirs.conf"

var USER_FILE string = expand("~/.config/user-dirs.dirs")

const DEFAULTS_FILE string = "/etc/xdg/user-dirs.defaults"

type DesktopFolders struct {
	root     map[string]string
	rootLast int

	user     map[string]string
	userLast int

	defaults     map[string]string
	defaultsLast int
}

// user application data folder
func getAppDataFolder() string {
	return path("XDG_CONFIG_HOME", "", "$HOME/.config")
}

// user home "/home/user"
func getHomeFolder() string {
	return expand("$HOME")
}

// user my documents "~/Documents"
func getDocumentsFolder() string {
	return path("XDG_DOCUMENTS_DIR", "DOCUMENTS", "$HOME/Documents")
}

// user downloads "~/Downloads"
func getDownloadsFolder() string {
	return path("XDG_DOWNLOAD_DIR", "DOWNLOAD", "$HOME/Downloads")
}

// user desktop "~/Desktop"
func getDesktopFolder() string {
	return path("XDG_DESKTOP_DIR", "DESKTOP", "$HOME/Desktop")
}

func trim(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "\"")
	s = strings.TrimSpace(s)
	return s
}

func expand(s string) string {
	s = strings.Replace(s, "~", "$HOME", -1)
	s = os.ExpandEnv(s)

	return s
}

func getini(s string) map[string]string {
	m := make(map[string]string)

	f, err := os.Open(s)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	p := regexp.MustCompile("([^=]*)=([^\n]*)")

	for scanner.Scan() {
		s := scanner.Text()
		s = trim(s)
		if s == "" {
			continue
		}
		if strings.HasPrefix(s, "#") {
			continue
		}
		if !p.MatchString(s) {
			continue
		}
		groups := p.FindStringSubmatch(s)
		key := trim(groups[1])
		value := trim(groups[2])
		m[key] = value
	}
	return m
}

func exists(s string) bool {
	if _, err := os.Stat(s); err == nil {
		return true
	}

	return false
}

// key
//            - key like "XDG_DOCUMENTS_DIR"
// xdgDefaultKey
//            - "DOCUMENTS", from /etc/xdg/user-dirs.defaults
// xdgDefaultPath
//            - default fallback value, "$HOME/Documents"

func path(k string, dk string, dp string) string {
	// 1) we have to check /etc/xdg/user-dirs.conf
	// if it is diabled, file is not here. rollback to xdgDefault parameter
	//
	// 2) ~/.config/user-dirs.dirs
	// try to locate user dirs file. and look for a value in there. if here
	// is no value or no file, switch to step 3
	//
	// 3) /etc/xdg/user-dirs.defaults
	// try to get system defaults for specified value. is here is no default
	// switch to xdgDefault parameter

	enabled := false
	if exists(ROOT_FILE) {
		root := getini(ROOT_FILE)
		e := strings.ToLower(root["enabled"])
		enabled = (e == "yes") || (e == "true")
	}

	if !enabled {
		return expand(dp)
	}

	if exists(USER_FILE) {
		user := getini(USER_FILE)
		if val, ok := user[k]; ok {
			return expand(val)
		}
	}

	if dk != "" {
		if exists(DEFAULTS_FILE) {
			defaults := getini(DEFAULTS_FILE)
			if val, ok := defaults[dk]; ok {
				return expand("$HOME/" + val)
			}
		}
	}

	return expand(dp)
}
