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
	mp3 := MP3{Name: "mp3_song.mp3"}
	wav := WAV{Name: "wav_song.wav"}

	PlayAudio(mp3)
	PlayAudio(wav)
}

func PlayAudio(p Playable) {
	fmt.Println("Starting playback...")
	p.Play()
	fmt.Println("Playback finished")
}
