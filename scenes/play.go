package scenes

import (
	"errors"
	"fmt"
	"image/color"
	"os"
	"sync"
	"time"

	"github.com/SolarLune/resolv"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/tessig/gogogopher/levels"
	"github.com/tessig/gogogopher/mechanics"
	"github.com/tessig/gogogopher/music"
	"github.com/tessig/gogogopher/objects/objects"
	"github.com/tessig/gogogopher/resources"
)

type (
	Symbol struct {
		Img   *ebiten.Image
		Scale float64
	}

	Play struct {
		level         *mechanics.Level
		width, height int
		char          *mechanics.Character
		lives         int
		livesSymbol   *Symbol
		lifeSound     *audio.Player
		coins         int
		coinSymbol    *Symbol
		coinSound     *audio.Player

		winSound *audio.Player

		collisionSpace resolv.Space
		objects        resolv.Space

		isLevelFinished bool
		isGameOver      bool

		// Camera
		cameraX int32
		cameraY int32

		once           sync.Once
		objectProvider mechanics.ObjectProvider
	}
)

var _ mechanics.Scene = new(Play)

func NewLevelScene(w, h int, character *mechanics.Character, livesSymbol, coinSymbol *Symbol, coinSound, lifeSound, winSound *audio.Player, objectProvider mechanics.ObjectProvider) *Play {
	l := &Play{
		width:          w,
		height:         h,
		char:           character,
		livesSymbol:    livesSymbol,
		coinSymbol:     coinSymbol,
		objectProvider: objectProvider,
		lives:          3,
		lifeSound:      lifeSound,
		coinSound:      coinSound,
		winSound:       winSound,
	}
	l.init()
	l.LoadLevel(levels.L1)
	return l
}

func (l *Play) Init() {
	music.SetTrack(resources.BGPlayer[resources.MusicALittleJourney])
	music.Play()
}

func (l *Play) init() {
	l.once = sync.Once{}
	l.isLevelFinished = false
	l.char.IsDead = false
	l.char.SetXY(90, 0)
	l.cameraX = 100
	l.cameraY = 0
}

func (l *Play) Reset() {
	l.isLevelFinished = false
	l.isGameOver = false
	l.lives = 3
	l.coins = 0
}

func (l *Play) ReloadAfter(secs int) {
	go func() {
		time.Sleep(time.Duration(secs) * time.Second)
		music.Play()
		l.init()
		l.LoadLevel(l.level)
	}()
}

func (l *Play) LoadLevel(level *mechanics.Level) {
	l.level = level
	l.level.SetObjectProvider(l.objectProvider)
	l.collisionSpace = level.CollisionSpace()
	l.collisionSpace.Move(-l.cameraX, -l.cameraY)
	l.objects = level.Objects()
	l.objects.Move(-l.cameraX, -l.cameraY)
}

func (l *Play) Layout(int, int) (int, int) {
	return l.width, l.height
}

func (l *Play) Update() (mechanics.SceneType, error) {
	l.checkLoseConditions()
	if l.isGameOver {
		l.Reset()
		l.ReloadAfter(1)
		return mechanics.SceneTypeGameOver, nil
	}
	l.checkWinConditions()

	sceneType := l.checkControls()

	l.checkObjects()

	l.checkCollisions()

	return sceneType, nil
}

func (l *Play) checkLoseConditions() {
	if l.char.IsDead {
		l.once.Do(func() {
			l.lives--
			if l.lives < 0 {
				l.lives = 0
				go func() {
					time.Sleep(2 * time.Second)
					l.isGameOver = true
				}()

				return
			}
			l.ReloadAfter(2)
		})
	}
}

func (l *Play) checkWinConditions() {
	winZones := l.collisionSpace.FilterByTags(mechanics.CollisionGroupWin)
	if l.char.IsColliding(winZones) {
		if !l.isLevelFinished {
			music.Pause()
			_ = l.winSound.Rewind()
			l.winSound.Play()
			zone := winZones.GetCollidingShapes(l.char).Get(0)
			l.char.SetXY(zone.GetXY())
			l.ReloadAfter(4)
		}
		l.isLevelFinished = true
		if l.char.Vy == 0 {
			l.char.Vy -= 8
		}
	}
}

func (l *Play) checkControls() mechanics.SceneType {
	if !l.isPlayable() {
		return mechanics.SceneTypeLevel
	}

	// Controls
	acc := 4.0
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		acc = 8
	}
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft):
		l.char.Vx = -acc
	case ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight):
		l.char.Vx = acc
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		l.char.Jump()
	}
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return mechanics.SceneTypeTitle
	}
	return mechanics.SceneTypeLevel
}

func (l *Play) checkObjects() {
	for _, object := range l.objects {
		o := object.(mechanics.Object)
		err := o.Update(l.char, &l.collisionSpace)
		if errors.Is(err, mechanics.ErrDestroy) {
			l.objects.Remove(o)
			switch ob := o.(type) {
			case *mechanics.Collectable:
				switch ob.Type() {
				case objects.Coin:
					l.addCoin()
				case objects.Life:
					l.addLife()
				}
			}
		}
	}
}

func (l *Play) addCoin() {
	_ = l.coinSound.Rewind()
	l.coinSound.Play()
	l.coins++
	if l.coins >= 100 {
		l.addLife()
		l.coins -= 100
	}
}

func (l *Play) addLife() {
	_ = l.lifeSound.Rewind()
	l.lifeSound.Play()
	l.lives++
}

func (l *Play) checkCollisions() {
	collisions := &resolv.Space{}
	if !l.char.IsDead {
		collisions = l.objects.FilterByTags(mechanics.ObjectTypeEnemy)
	}
	collisions.Add(l.collisionSpace...)
	dx, dy := l.char.Update(collisions)
	x, y := l.char.GetXY()
	if dx > 0 && l.width-int(x) < l.width/2 ||
		dx < 0 && int(x) < l.width/2 {
		l.MoveCamera(dx, 0)
	}
	if dy > 0 && l.height-int(y) < l.height/2 ||
		dy < 0 && int(y) < l.height/2 {
		l.MoveCamera(0, dy)
	}

	// too low -> die
	if y > int32(l.height+l.level.TileSize) {
		l.char.Die()
	}
}

func (l *Play) isPlayable() bool {
	return !(l.char.IsDead || l.isLevelFinished || l.isGameOver)
}

func (l *Play) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0x80, G: 0xa0, B: 0xc0, A: 0xff})
	l.level.SetCamera(l.cameraX, l.cameraY)
	l.level.Draw(screen, l.char.Draw, l.objects)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(5, 5)
	screen.DrawImage(symbolWithText(l.livesSymbol, fmt.Sprintf("x%d", l.lives)), op)
	op.GeoM.Reset()
	coins := symbolWithText(l.coinSymbol, fmt.Sprintf("x%d", l.coins))
	w, _ := coins.Size()
	op.GeoM.Translate(float64(l.width-w-5), 5)
	screen.DrawImage(coins, op)

	if _, debug := os.LookupEnv("DEBUG"); debug {
		x, y := l.char.GetXY()
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f, TPS: %0.2f, X: %d, Y:%d, VX: %0.2f, VY: %0.2f", ebiten.CurrentFPS(), ebiten.CurrentTPS(), x, y, l.char.Vx, l.char.Vy))
	}
}

func symbolWithText(s *Symbol, txt string) *ebiten.Image {
	height := resources.SmallArcadeFont.Metrics().Height.Floor()
	w, h := s.Img.Size()
	w = int(float64(w) * s.Scale)
	h = int(float64(h) * s.Scale)
	livesImg := ebiten.NewImage(len(txt)*height+w+3, h)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(s.Scale, s.Scale)
	livesImg.DrawImage(s.Img, op)
	text.Draw(livesImg, txt, resources.SmallArcadeFont, w+3, h-height/4, color.White)

	return livesImg
}

func (l *Play) MoveCamera(dx, dy int32) {
	if !l.isPlayable() {
		return
	}
	for l.cameraX+dx < 0 {
		dx += -dx - l.cameraX
	}
	endX := int32(l.level.TileSize*len(l.level.CollisionMap.Map[0]) - l.width)
	for l.cameraX+dx > endX {
		dx += endX - dx - l.cameraX
	}
	endY := int32(l.level.TileSize*len(l.level.CollisionMap.Map) - l.height)
	for l.cameraY+dy > endY {
		dy += endY - dy - l.cameraY
	}
	for l.cameraY+dy < 0 {
		dy += -dy - l.cameraY
	}
	l.cameraX += dx
	l.cameraY += dy

	l.collisionSpace.Move(-dx, -dy)
	l.char.Move(-dx, -dy)
	l.objects.Move(-dx, -dy)
}
