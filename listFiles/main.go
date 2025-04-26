package main

import (
	"errors"
	"fmt"
	"os"
)

func DirList(folder string) ([]string, error) {
	var list []string

	files, err := os.ReadDir(folder)

	if err != nil {
		return nil, errors.New("invalid repo")
	}

	for _, file := range files {
		list = append(list, file.Name())
	}

	return list, nil
}

func main() {
	data, _ := DirList("dirdata")

	fmt.Println(data)
}
