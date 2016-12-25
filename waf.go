package main

import (
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

var url string

func init() {
	url = os.Getenv("WAF_URL")
	if url == "" {
		panic("WAF_URL variable must be defined.")
	}
}

func main() {
	runtime.LockOSThread()
	gallium.Loop(os.Args, render)
}
