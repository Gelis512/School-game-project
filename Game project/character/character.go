package character

import (
	"main/coordinates"
	"main/entity"
	"main/weapon"
)

type Character struct {
	entity.Entity
	equippedWeapon *weapon.Weapon
	position       *coordinates.Coordinate
}

func (character *Character) GetEquippedWeapon() *weapon.Weapon {
	return character.equippedWeapon
}

func (character *Character) SetEquippedWeapon(equippedWeapon *weapon.Weapon) {
	character.equippedWeapon = equippedWeapon
}

func NewCharacter(PosX int, PosY int) *Character {
	var c *Character = new(Character)
	//var inputString string
	c.SetAc(10)
	c.SetHp(20)
	c.SetBaseWeapon("Fists")
	c.SetBasedmg(3)
	c.position = coordinates.NewCoordinate(PosX, PosY)
	//fmt.Println("Please input your character's name:")
	//fmt.Scan(&inputString)
	//c.SetName(inputString)
	c.SetName("placeholder")
	return c
}

func (character *Character) AssignNewShortsword() {
	character.SetEquippedWeapon(weapon.NewShortsword())
}

func (character *Character) GetEffectiveDamage() int {
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
