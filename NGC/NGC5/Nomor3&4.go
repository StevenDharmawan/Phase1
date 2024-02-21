package main

import (
	"fmt"
	"math/rand"
)

type Hero struct {
	Name                                             string
	BaseAttack, Defence, CriticalDamage, HealthPoint int
	weapon                                           Weapon
}

type Weapon struct {
	Attack int
}

func (hero Hero) CountDamage() int {
	return hero.BaseAttack + hero.weapon.Attack + hero.CriticalDamage*rand.Intn(2-0)
}

func (hero *Hero) IsAttackedBy(enemy Hero) {
	totalDamage := hero.Defence - enemy.CountDamage()
	if totalDamage < 0 {
		hero.HealthPoint = hero.HealthPoint + totalDamage
	}
}

func main() {
	hero1 := Hero{"Hero1", 50, 100, 50, 500, Weapon{20}}
	hero2 := Hero{"Hero2", 50, 100, 100, 500, Weapon{20}}
	index := 1
	for hero1.HealthPoint >= 0 && hero2.HealthPoint >= 0 {
		fmt.Println("BATTLE ", index)
		battle(hero2, &hero1)
		battle(hero1, &hero2)
		if hero1.HealthPoint <= 0 {
			fmt.Println("WINNER HERO 2\n")
		} else if hero2.HealthPoint <= 0 {
			fmt.Println("WINNER HERO 1\n")
		}
		index++
	}
}

func battle(hero1 Hero, hero2 *Hero) {
	hero2.IsAttackedBy(hero1)
	if hero2.HealthPoint < 0 {
		fmt.Printf("%s is attacked and HP now is 0\n", hero2.Name)
	} else {
		fmt.Printf("%s is attacked and HP now is %d\n", hero2.Name, hero2.HealthPoint)
	}
}
