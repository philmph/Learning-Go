package main

import "fmt"

type Player struct {
	HP int
}

func takeDamage(p *Player, d int) {
	p.HP -= d
}

func (p *Player) takeDamageReceiver(d int) {
	p.HP -= d
}

func main() {
	p := &Player{
		HP: 100,
	}

	fmt.Println("HP: ", p.HP)

	takeDamage(p, 10)

	fmt.Println("HP: ", p.HP)

	p.takeDamageReceiver(10)

	fmt.Println("HP: ", p.HP)
}
