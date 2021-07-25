package setup

import (
	"bytes"
	"embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed icons
	icons embed.FS
)

func platformSetup() {
	imgs := make([]image.Image, 3)
	for i, file := range []string{"icon48.png", "icon32.png", "icon16.png"} {
		content, err := icons.ReadFile("icons/" + file)
		if err != nil {
			panic(err)
		}
		imgs[i], _, err = image.Decode(bytes.NewReader(content))
		if err != nil {
			panic(err)
		}
	}

	ebiten.SetWindowIcon(imgs)
}
