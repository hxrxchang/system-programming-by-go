package main

import "fmt"

type Talker interface {
	Talk()
}

type Greeter struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Printf("Hello my name is %s\n", g.name)
}

var talker Talker

func main() {
	talker = &Greeter{"wozozo"}
	talker.Talk()
}
