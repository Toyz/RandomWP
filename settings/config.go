package settings

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/Toyz/RandomWP/desktop"

	"github.com/Toyz/RandomWP/wallhaven"
)

func New(confFile string) *Config {
	conf := &Config{
		SaveFolder: path.Join(desktop.GetDocumentsFolder(), "RandomWP", "pics"),
		Category:   wallhaven.CatGeneral,
		Purity:     wallhaven.PuritySFW,
		Ratio:      wallhaven.Ratio16x9,
		Delay:      3600,
		Notify:     false,
		AutoDelete: true, // Default to true unless the user say's other wise in the UI
		confFile:   confFile,
	}

	if fileExist(confFile) {
		conf.Load()
	} else {
		conf.Save()
	}

	return conf
}

func (conf *Config) Load() {
	data, _ := ioutil.ReadFile(conf.confFile)
	json.Unmarshal(data, conf)
}

func (conf *Config) Save() {
	data, _ := json.MarshalIndent(conf, "", "\t")
	ioutil.WriteFile(conf.confFile, data, 0644)
}

func fileExist(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
