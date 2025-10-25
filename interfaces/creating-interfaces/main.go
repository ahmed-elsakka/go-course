package main

import "fmt"

type Playable interface {
	Play()
}

type MP3 struct {
	Name string
}

type WAV struct {
	Name string
}

func (m MP3) Play() {
	fmt.Printf("Playing MP3 file: %s \n", m.Name)
}

func (w WAV) Play() {
	fmt.Printf("Playing WAV file: %s \n", w.Name)
}

func main() {

}
