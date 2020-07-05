package main

import (
	"flag"
	"fmt"

	"moul.io/dockerself"
)

func main() {
	dockerize := flag.Bool("dockerize", false, "dockerize")
	flag.Parse()
	if *dockerize {
		dockerself.Dockerize("ubuntu")
		return
	}
	if dockerself.WithinDocker() {
		fmt.Println("I'm inside the matrix (docker)")
	} else {
		fmt.Println("I'm in the real world :(")
	}
}
