package main

/*
	This factory doesn't actually create new types, it's
	just a skeleton for the version that will return
	the new types
*/

import (
	"fmt"
	"errors"
	"math/rand"
)

type Player interface {

	getHp() int
	attack() int
	defend(dmg int)
	scream()
}

type Character struct {
	hp int
	damage int
	defense int
	phrase string
}

func (c *Character) scream() {
	fmt.Println(c.phrase)
}

func (c *Character) getHp() int {
	return c.hp
}

func (c *Character) attack() int {
	return c.damage
}

func (c *Character) defend(dmg int) {
	if c.defense <= dmg {
		c.hp -= (dmg - c.defense)
	}
}

const (
	WEAK = iota
	AVG
	STRONG
)

func enemyFactory(t int) (Character, error) {
	switch t {
		// Character{hp: , damage: , defense: , phrase : }

		case WEAK:
			return Character{hp: 5, damage: 1, defense: 0, phrase : "weak"}, nil
		case AVG:
			return Character{hp: 8, damage: 4, defense: 3, phrase : "avg"}, nil
			
		case STRONG:
			return Character{hp: 12, damage: 7, defense: 5, phrase : "strong"}, nil
			
		default:
			return Character{}, errors.New("Invalid enemy type")
	}
}

func main() {
	player := Character{hp: 20, damage: 7, defense: 5, phrase: "I'm the new player"}
	myEnemy, _ := enemyFactory(rand.Intn(3))
	myEnemy.scream()

	for player.getHp() > 0 {
		myEnemy.defend(player.attack())

		if myEnemy.getHp() <= 0 {
			fmt.Println("Enemy is dead !")

			myEnemy, _ = enemyFactory(rand.Intn(3))
			myEnemy.scream()

			continue
		}
		
		player.defend(myEnemy.attack())
	}

	fmt.Println("Game Over")
}
