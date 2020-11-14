package mechanics

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SceneTypeTitle SceneType = iota
	SceneTypeLevel
	SceneTypeGameOver
	SceneTypeCredits
	SceneTypeQuit
)

type (
	Scene interface {
		Update() (SceneType, error)
		Draw(screen *ebiten.Image)
		Layout(outsideWidth int, outsideHeight int) (int, int)
		Init()
	}

	SceneType int

	Game struct {
		scenes      map[SceneType]Scene
		activeScene Scene
	}
)

var _ ebiten.Game = new(Game)

func NewGame(title, level, gameover, credits, quit Scene) *Game {
	sc := &Game{
		scenes: map[SceneType]Scene{
			SceneTypeTitle:    title,
			SceneTypeLevel:    level,
			SceneTypeGameOver: gameover,
			SceneTypeCredits:  credits,
			SceneTypeQuit:     quit,
		},
	}

	sc.init()

	return sc
}

func (m *Game) init() {
	m.activeScene = m.scenes[SceneTypeTitle]
	m.activeScene.Init()
}

func (m *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return m.activeScene.Layout(outsideWidth, outsideHeight)
}

func (m *Game) Update() error {
	sceneType, err := m.activeScene.Update()
	if err != nil {
		return err
	}

	if m.activeScene != m.scenes[sceneType] {
		m.activeScene = m.scenes[sceneType]
		m.activeScene.Init()
	}

	return nil
}

func (m *Game) Draw(screen *ebiten.Image) {
	m.activeScene.Draw(screen)
}
