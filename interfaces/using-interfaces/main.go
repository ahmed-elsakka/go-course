package main

import "fmt"

type Playable interface {
	Play()
}

type AdvancedPlayable interface {
	Playable
	Compress()
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

func (m MP3) Compress() {
	fmt.Printf("Compressing MP3 file: %s \n", m.Name)
}

func (w WAV) Play() {
	fmt.Printf("Playing WAV file: %s \n", w.Name)
}

func main() {
	mp3 := MP3{Name: "mp3_song.mp3"}
	wav := WAV{Name: "wav_song.wav"}

	compressAndPlayAdvancedPlayable(mp3)
	prepareAndPlayPlayable(mp3)
	prepareAndPlayPlayable(wav)
}

func prepareAndPlayPlayable(p Playable) {
	fmt.Println("Configuring system audio...")
	fmt.Println("Starting playback...")
	p.Play()
	fmt.Println("Playback finished")
}

func compressAndPlayAdvancedPlayable(ab AdvancedPlayable) {
	ab.Compress()
	ab.Play()
}
