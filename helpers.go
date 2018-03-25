package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"reflect"
	"unsafe"
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

func createFolder(folder string) {
	os.MkdirAll(folder, 0644)
}

func execFolder() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := path.Dir(ex)

	return exPath
}

func bytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func fileExist(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
