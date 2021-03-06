package updater

import (
	"encoding/json"
	"path"
	"strings"

	"github.com/parnurzeal/gorequest"
)

var GitUpdater = &Updater{}

func (git *Updater) LoadStatus() {
	_, body, err := gorequest.New().Get("https://api.github.com/repos/Toyz/RandomWP/statuses/master").End()

	if err != nil {
		git = nil
		return
	}

	json.Unmarshal([]byte(body), &git.Statues)
}

func (git *Updater) GetSHASum() string {
	if git == nil || len(git.Statues) <= 0 {
		return ""
	}

	data := git.Statues[0]

	return strings.TrimSuffix(path.Base(data.URL), "\n")
}

func (git *Updater) IsStable() bool {
	if git == nil || len(git.Statues) <= 0 {
		return false
	}

	data := git.Statues[0]

	return strings.EqualFold(data.State, "success")
}
