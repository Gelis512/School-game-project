package character

import (
	"fmt"
	"main/coordinates"
	"main/entity"
	"main/story"
	"main/weapon"
)

type Character struct { //Die Player Klasse
	entity.Entity
	equippedWeapon *weapon.Weapon          //Adresse zur verwendeten Waffe (unbenutzt)
	position       *coordinates.Coordinate //Position des Spielers
}

func (character *Character) GetEquippedWeapon() *weapon.Weapon {
	return character.equippedWeapon
}

func (character *Character) SetEquippedWeapon(equippedWeapon *weapon.Weapon) {
	character.equippedWeapon = equippedWeapon
}

func NewCharacter(PosX int, PosY int) *Character { //Erstellt eine neue Spieler-Instanz
	var c *Character = new(Character)
	var inputString string
	c.SetAc(10)
	c.SetHp(20) //placeholder oder einfach Grundwerte
	c.SetBaseWeapon("Fists")
	c.SetBasedmg(3)
	c.position = coordinates.NewCoordinate(PosX, PosY)

	if story.GetPlayerStoryName() != "" { //Wenn der Spieler die Story überspringt, wird nach eine Name gefragt
		c.SetName(story.GetPlayerStoryName())
	} else {
		fmt.Println("Please input your character's name:")
		fmt.Scan(&inputString)
		c.SetName(inputString)
	}
	return c
}

func (character *Character) AssignNewShortsword() { //Erstellt einen neuen Shortsword für den Spieler (unbenutzt)
	character.SetEquippedWeapon(weapon.NewShortsword())
}

func (character *Character) GetEffectiveDamage() int { //Die Summe der Basis-Schaden, und den DMG-Bonus der Waffe
	if character.equippedWeapon == nil {
		return character.GetBasedmg()
	} else {
		return character.GetEquippedWeapon().GetWdmg()
	}

}

func (character *Character) GetPos() (int, int) {
	var x int = character.position.GetX()
	var y int = character.position.GetY()
	return x, y
}

func (character *Character) GetPosY() int {
	var y int = character.position.GetY()
	return y
}

func (character *Character) GetPosX() int {
	var x int = character.position.GetX()
	return x
}

func (character *Character) SetPosX(positionX int) {
	character.position.SetX(positionX)
}

func (character *Character) SetPosY(positionY int) {
	character.position.SetY(positionY)
}
