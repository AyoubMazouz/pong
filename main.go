package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	W, H               = 800, 600
	TITLE              = "Pong Game"
	PADDLE_W, PADDLE_H = 20, 100
	BALL_R             = 10
	PADDLE_SPEED       = 3
	BALL_SPEED         = 5
)

var (
	bgColor     = rl.Black
	paddleColor = rl.White
	ballColor   = rl.White

	playerOne *Paddle
	playerTwo *Paddle
	ball      *Ball
)

// Paddle Start.
type Paddle struct {
	x int32
	y int32
}

func newPaddle(playerNum int8) *Paddle {
	if playerNum == 1 {
		return &Paddle{x: PADDLE_W, y: (H / 2) - (PADDLE_H / 2)}
	}
	return &Paddle{x: W - (PADDLE_W * 2), y: (H / 2) - (PADDLE_H / 2)}
}
func (p *Paddle) moveUp() {
	if p.y < 0 {
		p.y = 0
	}
	if p.y == 0 {
		return
	}
	p.y -= PADDLE_SPEED
}
func (p *Paddle) moveDown() {
	if p.y > H-PADDLE_H {
		p.y = H - PADDLE_H
	}
	if p.y == H-PADDLE_H {
		return
	}
	p.y += PADDLE_SPEED
}
func (p *Paddle) drawPaddle() {
	rl.DrawRectangle(p.x, p.y, PADDLE_W, PADDLE_H, paddleColor)
}

// Paddle ENDs
// Ball Start
type Ball struct {
	x      int32
	y      int32
	speedy int32
	speedx int32
}

func newBall() *Ball {
	speedx := rl.GetRandomValue(-1, 1)
	for speedx == 0 {
		speedx = rl.GetRandomValue(-1, 1)
	}
	speedy := rl.GetRandomValue(-1, 1)
	for speedy == 0 {
		speedy = rl.GetRandomValue(-1, 1)
	}
	return &Ball{x: W / 2, y: H / 2, speedx: BALL_SPEED * 1, speedy: BALL_SPEED * -1}
}
func (b *Ball) moveBall() {
	if b.x <= 0 || b.x >= W {
		b.speedx *= -1
	} else if b.y <= 0 || b.y >= H {
		b.speedy *= -1
	}
	b.x += b.speedx
	b.y += b.speedy
}
func (b *Ball) drawBall() {
	rl.DrawCircle(b.x, b.y, BALL_R, ballColor)
}

// BALL ENDs.
// Collision.
func collision(p1, p2 *Paddle, b *Ball) {
	var br int32 = BALL_R / 2
	if (b.x-br <= p1.x+PADDLE_W) && (b.x-br >= p1.x) && (b.y+br >= p1.y) && (b.y+br <= p1.y+PADDLE_H) {
		b.speedx *= -1
	}
	if (b.x+br >= p2.x) && (b.x+br <= p1.x+PADDLE_W) && (b.y+br >= p2.y) && (b.y+br <= p2.y+PADDLE_H) {
		b.speedx *= -1
	}
}
func input() {
	switch true {
	// Player One.
	case rl.IsKeyDown(rl.KeyW):
		playerOne.moveUp()
		break
	case rl.IsKeyDown(rl.KeyS):
		playerOne.moveDown()
		break

		// Player Two.
	case rl.IsKeyDown(rl.KeyUp):
		playerTwo.moveUp()
		break
	case rl.IsKeyDown(rl.KeyDown):
		playerTwo.moveDown()
		break
	}
}
func update() {
	collision(playerOne, playerTwo, ball)
	ball.moveBall()
}
func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(bgColor)
	// Start.

	playerOne.drawPaddle()
	playerTwo.drawPaddle()
	ball.drawBall()

	// End.
	rl.EndDrawing()
}

func init() {
	rl.InitWindow(W, H, TITLE)
	rl.SetTargetFPS(60)

	playerOne = newPaddle(1)
	playerTwo = newPaddle(2)
	ball = newBall()
}
func kill() {
	rl.CloseWindow()
}

func main() {
	for !rl.WindowShouldClose() {
		input()
		update()
		draw()
	}
}
