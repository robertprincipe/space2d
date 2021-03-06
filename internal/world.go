package internal

import "github.com/faiface/pixel"

type World struct {
	width  float64
	height float64
	bg     pixel.Picture
}

func NewWorld(w, h float64) *World {
	return &World{
		width:  w,
		height: h,
	}
}

func (w *World) AddBackground(path string) error {
	bg, err := loadPicture(path)

	if err != nil {
		return nil
	}

	w.bg = bg

	return nil
}

func (w World) Draw(t pixel.Target) {
	spriteBg := pixel.NewSprite(w.bg, w.bg.Bounds())
	bgBatch := pixel.NewBatch(&pixel.TrianglesData{}, w.bg)

	for x := 0.0; x <= w.width; x += spriteBg.Frame().W() {
		for y := 0.0; y <= w.height; y += spriteBg.Frame().H() {
			pos := pixel.V(x, y)
			spriteBg.Draw(bgBatch, pixel.IM.Moved(pos))
		}
	}

	bgBatch.Draw(t)
}
