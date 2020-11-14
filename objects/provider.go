package objects

import (
	"fmt"

	"github.com/tessig/gogogopher/mechanics"
	"github.com/tessig/gogogopher/objects/enemies"
	"github.com/tessig/gogogopher/objects/objects"
)

type (
	Provider struct{}
)

var _ mechanics.ObjectProvider = new(Provider)

func (p *Provider) CreateEnemy(typ int) *mechanics.Enemy {
	var enemy *mechanics.Enemy
	switch typ {
	default:
		panic(fmt.Sprintf("enemy type %d unknown", typ))
	case enemies.Elephpant:
		enemy = enemies.ElephpantProvider()
	case enemies.AlienElephpant:
		enemy = enemies.AlienElephpantProvider()
	case enemies.Python:
		enemy = enemies.PythonProvider()
	case enemies.AlienPython:
		enemy = enemies.AlienPythonProvider()
	case enemies.Duke:
		enemy = enemies.DukeProvider()
	case enemies.AlienDuke:
		enemy = enemies.AlienDukeProvider()
	case enemies.Ferris:
		enemy = enemies.FerrisProvider()
	case enemies.AlienFerris:
		enemy = enemies.AlienFerrisProvider()
	case enemies.Alien:
		enemy = enemies.AlienProvider()
	}

	return enemy
}

func (p *Provider) CreateObject(typ int) mechanics.Object {
	switch typ {
	default:
		panic(fmt.Sprintf("object type %d unknown", typ))
	case objects.Chest:
		return objects.ChestProvider()
	case objects.Life:
		return objects.LifeProvider()
	case objects.Coin:
		return objects.CoinProvider()
	}
}
