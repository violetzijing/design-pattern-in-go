package main

import "fmt"

type MediaPlayer interface {
	Play(string, string)
}

type AdvancedPlayer interface {
	PlayVlc(string)
	PlayMP4(string)
}

type VlcPlayer struct{}
type MP4Player struct{}

func (*VlcPlayer) PlayVlc(filename string) {
	fmt.Println("play vlc: ", filename)
}

func (*VlcPlayer) PlayMP4(filename string) {
}

func (*MP4Player) PlayVlc(filename string) {
}

func (*MP4Player) PlayMP4(filename string) {
	fmt.Println("play mp4: ", filename)
}

type MediaAdapter struct {
	vlcPlayer *VlcPlayer
	mp4Player *MP4Player
}

func NewMediaAdaper() *MediaAdapter {
	return &MediaAdapter{
		vlcPlayer: &VlcPlayer{},
		mp4Player: &MP4Player{},
	}
}

func (m *MediaAdapter) Play(audioType string, filename string) {
	if audioType == "vlc" {
		m.vlcPlayer.PlayVlc(filename)
	}
	if audioType == "mp4" {
		m.mp4Player.PlayMP4(filename)
	}
}

func main() {
	var mediaPlayer MediaPlayer
	mediaPlayer = NewMediaAdaper()
	mediaPlayer.Play("vlc", "vlc filename")
	mediaPlayer.Play("mp4", "mp4 filename")
}
