package main

import (
	"fmt"

	"github.com/Masterminds/glide/msg"

	"github.com/heewa/glide-brew/brew"
)

func main() {
	lock, err := brew.LoadLockFile()
	if err != nil {
		msg.Die(err.Error())
	}

	resources, err := brew.ConvertLock(lock)
	if err != nil {
		msg.Die(err.Error())
	}

	if len(resources) == 0 {
		fmt.Println("No Go dependencies found to convert")
	}

	for _, res := range resources {
		fmt.Printf("%s\n\n", res)
	}
}
