package main

import "fmt"

type Player struct {
	Alive  bool
	Health int
	Name   string
}

func newPlayer(name string) *Player {
	return &Player{
		Alive:  true,
		Health: 100,
		Name:   name,
	}
}

func (p *Player) takeDamage(damage int) {
	p.Health -= damage

	if p.Health <= 0 {
		p.Alive = false
	}
}

func main() {
	playerA := newPlayer("Agnes")

	fmt.Printf("Defined player: %+v\n", playerA)

	playerA.takeDamage(50)

	fmt.Printf("Defined player: %+v\n", playerA)

	playerA.takeDamage(50)

	fmt.Printf("Defined player: %+v\n", playerA)
}
