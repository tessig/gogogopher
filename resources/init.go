package resources

import (
	"bytes"
	"image"
	"io"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/tessig/gogogopher/resources/fonts"
	"github.com/tessig/gogogopher/resources/images"
	"github.com/tessig/gogogopher/resources/music"
	"github.com/tessig/gogogopher/resources/sounds"
)

const (
	fontSize      = 32
	smallFontSize = fontSize / 2

	MusicGreenHills = iota
	MusicFunnyChase
	MusicSpringThing
	MusicBirthdayCake
	MusicRetroNoHope
	MusicVictory
	MusicThemeSongFull
	MusicUnderTheSun
	MusicPlatform
	MusicALittleJourney
	MusicProperSummer
)

var (
	audioContext   = audio.NewContext(44100)
	GopherSprite   *ebiten.Image
	GopherEmojis   *ebiten.Image
	PlainsTiles    *ebiten.Image
	ForestTiles    *ebiten.Image
	Elephpant      *ebiten.Image
	AlienElephpant *ebiten.Image
	Python         *ebiten.Image
	AlienPython    *ebiten.Image
	Duke           *ebiten.Image
	AlienDuke      *ebiten.Image
	Ferris         *ebiten.Image
	AlienFerris    *ebiten.Image
	Alien          *ebiten.Image
	Coin           *ebiten.Image
	Chest          *ebiten.Image
	CoinPlayer     *audio.Player
	JumpPlayer     *audio.Player
	HurtPlayer     *audio.Player
	LifePlayer     *audio.Player

	BGPlayer map[int]*audio.Player

	ArcadeFont      font.Face
	SmallArcadeFont font.Face
)

// init images
func init() {
	img, _, err := image.Decode(bytes.NewReader(images.GopherSprite))
	if err != nil {
		panic(err)
	}
	GopherSprite = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(images.GopherEmojis))
	if err != nil {
		panic(err)
	}
	GopherEmojis = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(images.PlainsTiles))
	if err != nil {
		panic(err)
	}
	PlainsTiles = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(images.ForestTiles))
	if err != nil {
		panic(err)
	}
	ForestTiles = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(images.Elephpant))
	if err != nil {
		panic(err)
	}
	Elephpant = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(images.AlienElephpant))
	if err != nil {
		panic(err)
	}
	AlienElephpant = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(images.Python))
	if err != nil {
		panic(err)
	}
	Python = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(images.AlienPython))
	if err != nil {
		panic(err)
	}
	AlienPython = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(images.Duke))
	if err != nil {
		panic(err)
	}
	Duke = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(images.AlienDuke))
	if err != nil {
		panic(err)
	}
	AlienDuke = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(images.Ferris))
	if err != nil {
		panic(err)
	}
	Ferris = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(images.AlienFerris))
	if err != nil {
		panic(err)
	}
	AlienFerris = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(images.Alien))
	if err != nil {
		panic(err)
	}
	Alien = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(images.Coin))
	if err != nil {
		panic(err)
	}
	Coin = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(images.Chest))
	if err != nil {
		panic(err)
	}
	Chest = ebiten.NewImageFromImage(img)
}

// init sounds
func init() {
	oggSound, err := vorbis.Decode(audioContext, bytes.NewReader(sounds.CoinSound))
	if err != nil {
		panic(err)
	}
	CoinPlayer, err = audio.NewPlayer(audioContext, oggSound)
	if err != nil {
		panic(err)
	}
	CoinPlayer.SetVolume(.05)
	oggSound, err = vorbis.Decode(audioContext, bytes.NewReader(sounds.JumpSound))
	if err != nil {
		panic(err)
	}
	JumpPlayer, err = audio.NewPlayer(audioContext, oggSound)
	if err != nil {
		panic(err)
	}
	JumpPlayer.SetVolume(.5)
	oggSound, err = vorbis.Decode(audioContext, bytes.NewReader(sounds.HurtSound))
	if err != nil {
		panic(err)
	}
	HurtPlayer, err = audio.NewPlayer(audioContext, oggSound)
	if err != nil {
		panic(err)
	}
	HurtPlayer.SetVolume(.5)
	oggSound, err = vorbis.Decode(audioContext, bytes.NewReader(sounds.LifeSound))
	if err != nil {
		panic(err)
	}
	LifePlayer, err = audio.NewPlayer(audioContext, oggSound)
	if err != nil {
		panic(err)
	}
	LifePlayer.SetVolume(.1)
}

// init music
func init() {
	var oggs = map[int]struct {
		loop   bool
		data   []byte
		volume float64
	}{
		MusicGreenHills:     {loop: true, data: music.GreenHillsMusic, volume: 1},
		MusicFunnyChase:     {loop: true, data: music.FunnyChase, volume: 1},
		MusicSpringThing:    {loop: true, data: music.SpringThing, volume: 1},
		MusicBirthdayCake:   {loop: true, data: music.BirthdayCake, volume: 1},
		MusicRetroNoHope:    {loop: false, data: music.RetroNoHope, volume: .2},
		MusicVictory:        {loop: false, data: music.Victory, volume: .1},
		MusicUnderTheSun:    {loop: true, data: music.UnderTheSun, volume: 1},
		MusicPlatform:       {loop: true, data: music.Platform, volume: .2},
		MusicALittleJourney: {loop: true, data: music.ALittleJourney, volume: .1},
		MusicProperSummer:   {loop: true, data: music.ProperSummer, volume: 1},
	}

	BGPlayer = make(map[int]*audio.Player, len(oggs)+1)

	// special handling for theme song with intro
	length := int64(len(music.ThemeSongFull) * audioContext.SampleRate())
	intro := int64(59 * audioContext.SampleRate())
	oggSound, err := vorbis.Decode(audioContext, bytes.NewReader(music.ThemeSongFull))
	if err != nil {
		panic(err)
	}
	BGPlayer[MusicThemeSongFull], err = audio.NewPlayer(
		audioContext,
		audio.NewInfiniteLoopWithIntro(oggSound, intro, length-intro),
	)
	BGPlayer[MusicThemeSongFull].SetVolume(.1)
	if err != nil {
		panic(err)
	}

	// normal music
	for key, ogg := range oggs {
		var oggSound io.ReadSeeker
		oggSound, err = vorbis.Decode(audioContext, bytes.NewReader(ogg.data))
		if err != nil {
			panic(err)
		}
		if ogg.loop {
			length := int64(len(ogg.data) * audioContext.SampleRate())
			oggSound = audio.NewInfiniteLoop(oggSound, length)
		}
		BGPlayer[key], err = audio.NewPlayer(audioContext, oggSound)
		BGPlayer[key].SetVolume(ogg.volume)
		if err != nil {
			panic(err)
		}
	}
}

// init fonts
func init() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		panic(err)
	}
	const dpi = 72
	ArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}
	SmallArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    smallFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}
}
