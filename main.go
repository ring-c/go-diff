package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"go-diff/pkg/safetensors"
)

func main() {
	spew.Config.Indent = "\t"

	println("\n\n\n\n\n\n\n")

	var err = safetensors.Read("/mnt/files/sd/models/ponyRealism_v22MainVAE.safetensors")
	if err != nil {
		fmt.Printf("\n\n\nERROR: %s\n\n\n", err.Error())
		return
	}
}
