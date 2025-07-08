package main

import "fmt"

type WeaponType int

func (w WeaponType) String() string {
	var res string

	switch w {
	case Axe:
		res = "Axe"
	case Sword:
		res = "Sword"
	}

	return res
}

const (
	_ WeaponType = iota
	Axe
	Sword
)

func getDamage(wt WeaponType) int {
	var res int

	switch wt {
	case Axe:
		res = 50
	case Sword:
		res = 80
	}

	return res
}

func main() {
	fmt.Printf("Damage of weapon: (%s) (%d):\n", Axe, getDamage(Axe))
}
