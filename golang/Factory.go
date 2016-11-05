package main

import (
	"fmt"
	"errors"
	"math/rand"
)

type Character interface {

	getHp() int
	attack() int
	defend(dmg int)
	scream()
	init()
}

type Player struct {
	hp int
	damage int
	defense int
	phrase string
	healingPotions int
}

func (p *Player) scream() {
	fmt.Println(p.phrase)
}

func (p *Player) getHp() int {
	return p.hp
}

func (p *Player) attack() int {
	return p.damage
}

func (p *Player) defend(dmg int) {
	if p.defense <= dmg {
		p.hp -= (dmg - p.defense)
	}

	if p.hp <= 10 && p.healingPotions > 0 {
		p.hp += 7
		p.healingPotions -= 1
	}
}

// Weak enemy
type WeakEnemy struct {
	hp int
	damage int
	defence int
	phrase string
}

func (we *WeakEnemy)init() {
	we.hp = 5
	we.damage = 2
	we.defence = 0
	we.phrase = "Have mercy !!"

	we.scream()
}

func (we *WeakEnemy)getHp() int {
	return we.hp
}

func (we *WeakEnemy)attack() int {
	return we.damage
}

func (we *WeakEnemy)defend(damage int) {
	if damage > we.defence {
		we.hp -= (damage - we.defence)
	}
}

func (we *WeakEnemy)scream() {
	fmt.Println(we.phrase)
}

// AVG enemy
type AvgEnemy struct {
	hp int
	damage int
	defence int
	phrase string
}

func (avg *AvgEnemy)init() {
	avg.hp = 10
	avg.damage = 4
	avg.defence = 3
	avg.phrase = "Ha Ha Ha ... I'm not as weak as you think ! Die !"

	avg.scream()
}

func (avg *AvgEnemy)getHp() int {
	return avg.hp
}

func (avg *AvgEnemy)attack() int {
	return avg.damage
}

func (avg *AvgEnemy)defend(damage int) {
	if damage > avg.defence {
		avg.hp -= (damage - avg.defence)
	}

	if avg.hp <= 3 {
		fmt.Println("Alright ... I give up !!")
		avg.hp = 0
	}
}

func (avg *AvgEnemy)scream() {
	fmt.Println(avg.phrase)
}

// Strong enemy
type StrongEnemy struct {
	hp int
	damage int
	defence int
	phrase string
	frenzy bool
}

func (s *StrongEnemy)init() {
	s.hp = 15
	s.damage = 7
	s.defence = 5
	s.phrase = "You are dead ..."
	s.frenzy = false

	s.scream()
}

func (s *StrongEnemy)getHp() int {
	return s.hp
}

func (s *StrongEnemy)attack() int {
	return s.damage
}

func (s *StrongEnemy)defend(damage int) {
	if damage > s.defence {
		s.hp -= (damage - s.defence)
	}

	if s.hp <= 7 && s.frenzy == false {
		s.damage += 3
		s.frenzy = true

		fmt.Println("YOU ARE MAKING ME ANGRY !! GRRRwaaa !!")
	}
}

func (s *StrongEnemy)scream() {
	fmt.Println(s.phrase)
}

/*
	This is the factory function. Note that this factory returns the interface
	rather than the struct type. This is handy when dealing with multiple
	implementations of the same Interface.
*/

const (
	WEAK = iota
	AVG
	STRONG
)

func enemyFactory(t int) (Character, error) {
	switch t {

		case WEAK:
			return new(WeakEnemy), nil
		case AVG:
			return new(AvgEnemy), nil
		case STRONG:
			return new(StrongEnemy), nil
		default:
			return nil, errors.New("Invalid enemy type")
	}
}

func main() {
	player := Player{hp: 20, damage: 7, defense: 5, healingPotions: 3, phrase: "I'm the new player"}
	myEnemy, _ := enemyFactory(rand.Intn(3))
	myEnemy.init()

	for player.getHp() > 0 {
		myEnemy.defend(player.attack())

		if myEnemy.getHp() <= 0 {
			fmt.Println("Enemy is dead !")

			myEnemy, _ = enemyFactory(rand.Intn(3))
			myEnemy.init()

			continue
		}
		player.defend(myEnemy.attack())
	}

	fmt.Println("Game Over")
}
