package resources

import (
	"bytes"
	"embed"
	"image"
	"io"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
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
var (
	//go:embed embed
	resources embed.FS

	resourceTypes = map[string]fs.ReadFileFS{
		"images": nil, "sounds": nil, "music": nil, "fonts": nil,
	}
)

func mustReadFile(fs fs.ReadFileFS, name string) []byte {
	content, err := fs.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return content
}

// init fs
func init() {
	for key := range resourceTypes {
		f, err := fs.Sub(resources, "embed/"+key)
		if err != nil {
			panic(err)
		}
		resourceTypes[key] = f.(fs.ReadFileFS)
	}
}

// init images
func init() {
	images := resourceTypes["images"]
	img, _, err := image.Decode(bytes.NewReader(mustReadFile(images, "gopher-sprite.png")))
	if err != nil {
		panic(err)
	}
	GopherSprite = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(mustReadFile(images, "gopher-emojis.png")))
	if err != nil {
		panic(err)
	}
	GopherEmojis = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(mustReadFile(images, "plains.png")))
	if err != nil {
		panic(err)
	}
	PlainsTiles = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(mustReadFile(images, "forest.png")))
	if err != nil {
		panic(err)
	}
	ForestTiles = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(mustReadFile(images, "elephpant.png")))
	if err != nil {
		panic(err)
	}
	Elephpant = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(mustReadFile(images, "alien_elephpant.png")))
	if err != nil {
		panic(err)
	}
	AlienElephpant = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(mustReadFile(images, "python.png")))
	if err != nil {
		panic(err)
	}
	Python = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(mustReadFile(images, "alien_python.png")))
	if err != nil {
		panic(err)
	}
	AlienPython = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(mustReadFile(images, "duke.png")))
	if err != nil {
		panic(err)
	}
	Duke = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(mustReadFile(images, "alien_duke.png")))
	if err != nil {
		panic(err)
	}
	AlienDuke = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(mustReadFile(images, "ferris.png")))
	if err != nil {
		panic(err)
	}
	Ferris = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(mustReadFile(images, "alien_ferris.png")))
	if err != nil {
		panic(err)
	}
	AlienFerris = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(mustReadFile(images, "alien.png")))
	if err != nil {
		panic(err)
	}
	Alien = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(mustReadFile(images, "coin.png")))
	if err != nil {
		panic(err)
	}
	Coin = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(mustReadFile(images, "chest.png")))
	if err != nil {
		panic(err)
	}
	Chest = ebiten.NewImageFromImage(img)
}

// init sounds
func init() {
	sounds := resourceTypes["sounds"]
	oggSound, err := vorbis.DecodeWithSampleRate(audioContext.SampleRate(), bytes.NewReader(mustReadFile(sounds, "coin9.ogg")))
	if err != nil {
		panic(err)
	}
	CoinPlayer, err = audio.NewPlayer(audioContext, oggSound)
	if err != nil {
		panic(err)
	}
	CoinPlayer.SetVolume(.05)
	oggSound, err = vorbis.DecodeWithSampleRate(audioContext.SampleRate(), bytes.NewReader(mustReadFile(sounds, "jump22.ogg")))
	if err != nil {
		panic(err)
	}
	JumpPlayer, err = audio.NewPlayer(audioContext, oggSound)
	if err != nil {
		panic(err)
	}
	JumpPlayer.SetVolume(.5)
	oggSound, err = vorbis.DecodeWithSampleRate(audioContext.SampleRate(), bytes.NewReader(mustReadFile(sounds, "hurt.ogg")))
	if err != nil {
		panic(err)
	}
	HurtPlayer, err = audio.NewPlayer(audioContext, oggSound)
	if err != nil {
		panic(err)
	}
	HurtPlayer.SetVolume(.5)
	oggSound, err = vorbis.DecodeWithSampleRate(audioContext.SampleRate(), bytes.NewReader(mustReadFile(sounds, "powerup02.ogg")))
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
	music := resourceTypes["music"]
	var oggs = map[int]struct {
		loop   bool
		data   []byte
		volume float64
	}{
		MusicGreenHills:     {loop: true, data: mustReadFile(music, "GreenHills.ogg"), volume: 1},
		MusicFunnyChase:     {loop: true, data: mustReadFile(music, "Funny_Chase.ogg"), volume: 1},
		MusicSpringThing:    {loop: true, data: mustReadFile(music, "spring_thing.ogg"), volume: 1},
		MusicBirthdayCake:   {loop: true, data: mustReadFile(music, "Birthday_Cake.ogg"), volume: 1},
		MusicRetroNoHope:    {loop: false, data: mustReadFile(music, "Retro_No_hope.ogg"), volume: .2},
		MusicVictory:        {loop: false, data: mustReadFile(music, "Victory.ogg"), volume: .1},
		MusicUnderTheSun:    {loop: true, data: mustReadFile(music, "under_the_sun.ogg"), volume: 1},
		MusicPlatform:       {loop: true, data: mustReadFile(music, "Platform.ogg"), volume: .2},
		MusicALittleJourney: {loop: true, data: mustReadFile(music, "a_little_journey.ogg"), volume: .1},
		MusicProperSummer:   {loop: true, data: mustReadFile(music, "proper_summer.ogg"), volume: 1},
	}

	BGPlayer = make(map[int]*audio.Player, len(oggs)+1)
	themeSongFull := mustReadFile(music, "Theme_Song_full.ogg")
	// special handling for theme song with intro
	length := int64(len(themeSongFull) * audioContext.SampleRate())
	intro := int64(59 * audioContext.SampleRate())
	oggSound, err := vorbis.DecodeWithSampleRate(audioContext.SampleRate(), bytes.NewReader(themeSongFull))
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
		oggSound, err = vorbis.DecodeWithSampleRate(audioContext.SampleRate(), bytes.NewReader(ogg.data))
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
	fonts := resourceTypes["fonts"]
	tt, err := opentype.Parse(mustReadFile(fonts, "pressstart2p.ttf"))
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
