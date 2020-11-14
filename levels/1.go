package levels

import (
	"github.com/tessig/gogogopher/levels/level1"
	"github.com/tessig/gogogopher/mechanics"
	"github.com/tessig/gogogopher/resources"
)

var (
	L1 *mechanics.Level
)

func init() {
	L1 = &mechanics.Level{
		TileSize: 64,
		Back: []*mechanics.Layer{
			{
				Map:         level1.Background,
				TilesSprite: resources.PlainsTiles,
				TilesPerRow: 10,
			},
		},
		Front: []*mechanics.Layer{
			{
				Map:         level1.Foreground,
				TilesSprite: resources.PlainsTiles,
				TilesPerRow: 10,
			},
			{
				Map:         level1.Environment,
				TilesSprite: resources.ForestTiles,
				TilesPerRow: 10,
			},
		},
		CollisionMap: &mechanics.CollisionMap{
			Map: level1.Collision,
		},
		EnemiesMap: level1.Enemies,
		ObjectsMap: level1.Objects,
	}
}
