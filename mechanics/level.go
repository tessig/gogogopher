package mechanics

import (
	"fmt"
	"image"
	"strconv"

	"github.com/SolarLune/resolv"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	CollisionTypeTraversable = iota - 1
	CollisionTypeWin
	CollisionTypeSolid
	CollisionTypeSolidHalfTop
	CollisionTypeSolidHalfBottom
	CollisionTypeRampUp
	CollisionTypeRampDown
	CollisionTypeDie
	CollisionTypeDieHalfTop
	CollisionTypeDieHalfBottom
	CollisionTypeStop

	ObjectTypeEnemy       = "enemy"
	ObjectTypeCollectable = "collectable"

	CollisionGroupTraversable = "traversable"
	CollisionGroupSolid       = "solid"
	CollisionGroupRamp        = "ramp"
	CollisionGroupDie         = "die"
	CollisionGroupStopper     = "stopper"
	CollisionGroupWin         = "win"
)

var (
	CollisionGroups = map[int]string{
		CollisionTypeTraversable:     CollisionGroupTraversable,
		CollisionTypeSolid:           CollisionGroupSolid,
		CollisionTypeSolidHalfTop:    CollisionGroupSolid,
		CollisionTypeSolidHalfBottom: CollisionGroupSolid,
		CollisionTypeRampUp:          CollisionGroupRamp,
		CollisionTypeRampDown:        CollisionGroupRamp,
		CollisionTypeDie:             CollisionGroupDie,
		CollisionTypeDieHalfTop:      CollisionGroupDie,
		CollisionTypeDieHalfBottom:   CollisionGroupDie,
		CollisionTypeWin:             CollisionGroupWin,
		CollisionTypeStop:            CollisionGroupStopper,
	}
)

type (
	Layer struct {
		Map         IntMap
		TilesSprite *ebiten.Image
		TilesPerRow int
	}
	CollisionMap struct {
		Map IntMap
	}
	Level struct {
		TileSize       int
		Back           []*Layer
		Front          []*Layer
		CollisionMap   *CollisionMap
		EnemiesMap     IntMap
		ObjectsMap     IntMap
		objectProvider ObjectProvider
		camera         struct {
			x, y float64
		}
	}
	IntMap [][]int
)

func (l *Level) CollisionSpace() resolv.Space {
	space := resolv.Space{}
	size := int32(l.TileSize)
	for y, row := range l.CollisionMap.Map {
		for x, tile := range row {
			x0 := int32(x) * size
			y0 := int32(y) * size
			var shape resolv.Shape
			switch tile {
			default:
				continue
			case CollisionTypeSolid,
				CollisionTypeDie,
				CollisionTypeStop:
				shape = resolv.NewRectangle(x0, y0, size, size)
			case CollisionTypeWin:
				shape = resolv.NewRectangle(x0+size/3, y0, size/3, size)
			case CollisionTypeSolidHalfTop,
				CollisionTypeDieHalfTop:
				shape = resolv.NewRectangle(x0, y0, size, size/2)
			case CollisionTypeSolidHalfBottom,
				CollisionTypeDieHalfBottom:
				shape = resolv.NewRectangle(x0, y0+size/2, size, size/2)
			case CollisionTypeRampUp:
				s := resolv.NewSpace()
				s.Add(
					resolv.NewLine(x0, y0+size, x0+size, y0),
					resolv.NewLine(x0, y0+size, x0+size, y0+size),
					resolv.NewLine(x0+size, y0, x0+size, y0+size),
				)
				shape = s
			case CollisionTypeRampDown:
				s := resolv.NewSpace()
				s.Add(
					resolv.NewLine(x0, y0, x0+size, y0+size),
					resolv.NewLine(x0, y0+size, x0+size, y0+size),
					resolv.NewLine(x0, y0, x0, y0+size),
				)
				shape = s
			}
			shape.AddTags(
				fmt.Sprintf("coord(%d,%d)", x, y),
				strconv.Itoa(tile),
				CollisionGroups[tile],
			)
			space.Add(shape)
		}
	}

	return space
}

func (l *Level) Objects() resolv.Space {
	space := resolv.Space{}
	for y, row := range l.EnemiesMap {
		for x, tile := range row {
			if tile < 0 {
				continue
			}
			x0 := int32(x * l.TileSize)
			y0 := int32(y * l.TileSize)
			enemy := l.objectProvider.CreateEnemy(tile)
			enemy.SetXY(x0, y0)
			enemy.AddTags(
				fmt.Sprintf("coord(%d,%d)", x, y),
				strconv.Itoa(tile),
				ObjectTypeEnemy,
				CollisionGroupSolid,
			)
			space.Add(enemy)
		}
	}
	for y, row := range l.ObjectsMap {
		for x, tile := range row {
			if tile < 0 {
				continue
			}
			object := l.objectProvider.CreateObject(tile)
			object.SetXY(PositionInTiles(x, y, float64(l.TileSize), object.Width(), object.Height()))
			tags := []string{
				fmt.Sprintf("coord(%d,%d)", x, y),
				strconv.Itoa(tile),
			}
			switch object.(type) {
			case *Collectable:
				tags = append(tags,
					ObjectTypeCollectable,
					CollisionGroupTraversable,
				)
			case *Chest:
				tags = append(tags,
					CollisionGroupSolid,
				)
			}
			object.AddTags(tags...)
			space.Add(object)
		}
	}

	return space
}

func (l *Level) SetObjectProvider(provider ObjectProvider) {
	l.objectProvider = provider
}

func (l *Level) SetCamera(x, y int32) {
	l.camera.x = float64(x)
	l.camera.y = float64(y)
}

func (l *Level) Draw(screen *ebiten.Image, drawChar func(*ebiten.Image), objects resolv.Space) {
	l.drawLayers(screen, l.Back)
	drawChar(screen)
	for _, o := range objects {
		o.(Object).Draw(screen)
	}

	l.drawLayers(screen, l.Front)
}

func (l *Level) drawLayers(screen *ebiten.Image, layers []*Layer) {
	for _, layer := range layers {
		layer.Draw(screen, l.TileSize, l.camera.x, l.camera.y)
	}
}

func (l *Layer) Draw(screen *ebiten.Image, tileSize int, offsetX, offsetY float64) {
	for i, row := range l.Map {
		for j, tile := range row {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(j*tileSize), float64(i*tileSize))
			op.GeoM.Translate(-offsetX, -offsetY)
			if tile < 0 {
				continue
			}
			t := tile
			x0 := (t % l.TilesPerRow) * tileSize
			y0 := (t / l.TilesPerRow) * tileSize

			screen.DrawImage(l.TilesSprite.SubImage(image.Rect(x0, y0, x0+tileSize, y0+tileSize)).(*ebiten.Image), op)
		}
	}
}
