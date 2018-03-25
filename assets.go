package main

import (
	"io/ioutil"
	"path"
)

var (
	assetPath string
)

func getAsset(name ...string) string {
	args := make([]string, len(name)+1)
	args = append(args, assetPath)
	for _, n := range name {
		args = append(args, n)
	}

	return path.Join(args...)
}

func setAsset(data []byte, name ...string) string {
	asset := getAsset(name...)

	ioutil.WriteFile(asset, data, 0644)

	return asset
}

func saveAssets() {
	if !fileExist(path.Join(assetPath, "icon.png")) {
		data, _ := Asset("assets/icon.png")
		setAsset(data, "icon.png")
	}

	if !fileExist(path.Join(assetPath, "information.png")) {
		data, _ := Asset("assets/information.png")
		setAsset(data, "information.png")
	}
}
