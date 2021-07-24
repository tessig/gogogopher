package music

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
)

var active *audio.Player

func Play() {
	active.Play()
}

func Pause() {
	active.Pause()
}

func IsPlaying() bool {
	return active.IsPlaying()
}

func Rewind() {
	_ = active.Rewind()
}

func SetVolume(v float64) {
	active.SetVolume(v)
}

func SetTrack(player *audio.Player) {
	if active != nil {
		active.Pause()
	}
	old := active
	active = player
	if old != active {
		Rewind()
	}
}
