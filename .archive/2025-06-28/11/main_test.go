package main

import (
	"reflect"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	t.Log("Testing newPlayer")

	expected := &Player{
		Alive:  true,
		Health: 100,
		Name:   "Tester",
	}

	result := newPlayer("Tester")

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestTakeDamage(t *testing.T) {
	t.Log("Testing takeDamage")

	tests := []struct {
		name           string
		initialHealth  int
		damage         int
		expectedHealth int
		expectedAlive  bool
	}{
		{
			name:           "normal damage",
			initialHealth:  100,
			damage:         30,
			expectedHealth: 70,
			expectedAlive:  true,
		},
		{
			name:           "damage that kills",
			initialHealth:  50,
			damage:         60,
			expectedHealth: -10,
			expectedAlive:  false,
		},
		{
			name:           "exact lethal damage",
			initialHealth:  50,
			damage:         50,
			expectedHealth: 0,
			expectedAlive:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player := &Player{
				Alive:  true,
				Health: tt.initialHealth,
				Name:   "TestPlayer",
			}

			player.takeDamage(tt.damage)

			if player.Health != tt.expectedHealth {
				t.Errorf("Expected health %d, got %d", tt.expectedHealth, player.Health)
			}

			if player.Alive != tt.expectedAlive {
				t.Errorf("Expected alive status %t, got %t", tt.expectedAlive, player.Alive)
			}
		})
	}
}
