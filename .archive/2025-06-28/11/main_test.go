package main

import (
	"reflect"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	t.Log("Testing newPlayer")

	expected := &Player{
		Name: "Tester",
	}

	result := newPlayer("Tester")

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestDamage(t *testing.T) {
	t.Log("Testing takeDamage")

	testPlayers := []*Player{
		{
			Alive:  true,
			Health: 100,
			Name:   "A",
		},
		{
			Alive:  true,
			Health: 10,
			Name:   "B",
		},
	}

	for _, p := range testPlayers {
		t.Run("taking damage", func(t *testing.T) {
			expected := p.Health - 50
			p.takeDamage(50)

			if expected != p.Health {
				t.Errorf("Expected health %d after damage, got %d", expected, p.Health)
			}
		})

		t.Run("dying", func(t *testing.T) {
			expected := p.Health >= 50
			p.takeDamage(50)

			if expected != p.Alive {
				t.Errorf("Expected alive status %t after damage, got %t", expected, p.Alive)
			}
		})
	}
}
