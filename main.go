package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	W, H               = 800, 600
	TITLE              = "Pong Game"
	PADDLE_W, PADDLE_H = 20, 100
	BALL_R             = 10
	PADDLE_SPEED       = 6
	BALL_SPEED         = 8
	SPEED_BUFF         = 2
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
	x     int32
	y     int32
	dir   bool
	score int32
}

func newPaddle(playerNum int8) *Paddle {
	if playerNum == 1 {
		return &Paddle{x: PADDLE_W, y: (H / 2) - (PADDLE_H / 2), dir: false, score: 0}
	}
	return &Paddle{x: W - (PADDLE_W * 2), y: (H / 2) - (PADDLE_H / 2), dir: false, score: 0}
}

func (p *Paddle) moveUp() {
	if p.y < 0 {
		p.y = 0
	}
	if p.y == 0 {
		return
	}
	p.y -= PADDLE_SPEED
	p.dir = true
}

func (p *Paddle) moveDown() {
	if p.y > H-PADDLE_H {
		p.y = H - PADDLE_H
	}
	if p.y == H-PADDLE_H {
		return
	}
	p.y += PADDLE_SPEED
	p.dir = false
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

func (b *Ball) moveBall(p1, p2 *Paddle) {
	if b.x <= 0 {
		p1.score++
		ball = newBall()
	} else if b.x >= W {
		p2.score++
		ball = newBall()
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
	if (b.x-br <= p1.x+PADDLE_W) && (b.x-br >= p1.x) && (b.y-br >= p1.y) && (b.y+br <= p1.y+PADDLE_H) {
		if p1.dir == true {
			b.speedx = int32(float32(b.speedx)*-0.9) + SPEED_BUFF
			b.speedy = int32(float32(b.speedy)*1.1) + SPEED_BUFF
		} else {
			b.speedx = int32(float32(b.speedx)*-1.1) + SPEED_BUFF
			b.speedy = int32(float32(b.speedy)*0.9) + SPEED_BUFF
		}
		b.x = p1.x + PADDLE_W + 5
	} else if (b.x+br >= p2.x) && (b.x+br <= p2.x+PADDLE_W) && (b.y-br >= p2.y) && (b.y+br <= p2.y+PADDLE_H) {
		if p2.dir == true {
			b.speedx = int32(float32(b.speedx)*-0.9) + SPEED_BUFF
			b.speedy = int32(float32(b.speedy)*1.1) + SPEED_BUFF
		} else {
			b.speedx = int32(float32(b.speedx)*-1.1) + SPEED_BUFF
			b.speedy = int32(float32(b.speedy)*0.9) + SPEED_BUFF
		}
		b.x = p2.x - 5
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
	ball.moveBall(playerOne, playerTwo)
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

func initialize() {
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
	fmt.Println("pong")
	initialize()
	for !rl.WindowShouldClose() {
		input()
		update()
		draw()
	}
	kill()
}
