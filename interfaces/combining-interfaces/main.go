package main

import "fmt"

type Playable interface {
	Play()
}

type AdvancedPlayable interface {
	Playable
	Pause()
	Stop()
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

func (m MP3) Stop() {
	fmt.Printf("Stopped MP3 file: %s \n", m.Name)
}

func (m MP3) Pause() {
	fmt.Printf("Paused MP3 file: %s \n", m.Name)
}

func (w WAV) Play() {
	fmt.Printf("Playing WAV file: %s \n", w.Name)
}

func main() {
	mp3 := MP3{Name: "mp3_song.mp3"}
	wav := WAV{Name: "wav_song.wav"}

	PlayAudio(mp3)
	PlayAudio(wav)

	PauseAudio(mp3)
	PlayAudio(mp3)
	StopAudio(mp3)

}

func PlayAudio(p Playable) {
	fmt.Println("Starting playback...")
	p.Play()
	fmt.Println("Playback finished")
}

func StopAudio(a AdvancedPlayable) {
	a.Stop()
}

func PauseAudio(a AdvancedPlayable) {
	a.Pause()
}
