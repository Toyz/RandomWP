package main

import (
	"fmt"
	"math/rand"
	"os"
)

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
