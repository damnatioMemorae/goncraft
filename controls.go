package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	MAX_CAM_TILT = -math.Pi * 0.5
	MIN_CAM_TILT = -math.Pi * 1.5
)



func (c *Game) HandleMouseMove() {
	xt, yt := ebiten.CursorPosition()
	x := float64(xt)
	y := float64(yt)

	dt := float64(1) / float64(60)

	c.PlayerRotation.X += (((c.Cursor.X - x) * dt) * math.Pi)
	c.PlayerRotation.Y += (((c.Cursor.Y - y) * dt) * math.Pi)

	if c.PlayerRotation.Y > MAX_CAM_TILT {
		c.PlayerRotation.Y = MAX_CAM_TILT
	} else if c.PlayerRotation.Y < MIN_CAM_TILT {
		c.PlayerRotation.Y = MIN_CAM_TILT
	}

	c.Cursor.X = x
	c.Cursor.Y = y
}

func (c *Camera) HandleMovement() {
	dt := float64(1) / float64(60)

	var dv *Vector3

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		dv = &Vector3{-1 * dt, 0, 0}
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		dv = &Vector3{1 * dt, 0, 0}
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		dv = &Vector3{0, 0, 1 * dt}
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		dv = &Vector3{0, 0, -1 * dt}
	}

	if dv == nil {
		return
	}

	c. = c.PlayerPosition.Add(dv)
}

func (g *Game) HandleJump() {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		dt := float64(1) / float64(60)
		g.PlayerPosition.Y += 1 * dt
		return
	}
}

func (g *Game) HandleDescent() {
	if ebiten.IsKeyPressed(ebiten.KeyShiftLeft) {
		dt := float64(1) / float64(60)
		g.PlayerPosition.Y += -1 * dt
		return
	}
}
