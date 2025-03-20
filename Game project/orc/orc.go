package orc

import (
	"main/entity"
	"math/rand/v2"
)

var w = [3]string{"Mace", "Shortsword", "Spear"}

const orcMaceDMG = 6
const orcShortswordDMG = 4
const orcSpearDMG = 5
const orcBaseAC = 14
const orcBaseHP = 30
const orcBaseName = "Orc"

type Orc struct {
	entity.Entity
}

func NewOrc() *Orc {

	var o *Orc = new(Orc)
	o.SetAc(orcBaseAC)
	o.SetHp(orcBaseHP)
	var r int = rand.IntN(3)
	o.SetBaseWeapon(w[r])
	switch o.GetBaseWeapon() {
	case "Mace":
		o.SetBasedmg(orcMaceDMG)
	case "Shortsword":
		o.SetBasedmg(orcShortswordDMG)
	case "Spear":
		o.SetBasedmg(orcSpearDMG)
	}

	o.SetName(orcBaseName)

	return o

}
