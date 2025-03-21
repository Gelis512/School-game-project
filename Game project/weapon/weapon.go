package weapon //Von Hanna erstellt (unbenutzt)

type Weapon struct {
	wname string
	wtype string
	wdmg  int
}

func NewShortsword() *Weapon {

	var w *Weapon = new(Weapon)
	w.wdmg = 10
	w.wname = "Shortsword"
	w.wtype = "Blade"

	return w
}

func (weapon *Weapon) GetWname() string {
	return weapon.wname
}

func (weapon *Weapon) GetWtype() string {
	return weapon.wtype
}

func (weapon *Weapon) GetWdmg() int {
	return weapon.wdmg
}

func (weapon *Weapon) SetWname(wname string) *Weapon {
	weapon.wname = wname
	return weapon
}

func (weapon *Weapon) SetWtype(wtype string) *Weapon {
	weapon.wtype = wtype
	return weapon
}

func (weapon *Weapon) SetWdmg(wdmg int) *Weapon {
	weapon.wdmg = wdmg
	return weapon
}
