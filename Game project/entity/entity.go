package entity

type Entity struct {
	hp         int
	basedmg    int
	ac         int
	name       string
	baseWeapon string
}

func (entity *Entity) GetHp() int {
	return entity.hp
}

func (entity *Entity) GetBasedmg() int {
	return entity.basedmg
}

func (entity *Entity) GetAc() int {
	return entity.ac
}

func (entity *Entity) GetName() string {
	return entity.name
}

func (entity *Entity) GetBaseWeapon() string {
	return entity.baseWeapon
}

func (entity *Entity) SetHp(hp int) *Entity {
	entity.hp = hp
	return entity
}

func (entity *Entity) SetBasedmg(basedmg int) *Entity {
	entity.basedmg = basedmg
	return entity
}

func (entity *Entity) SetAc(ac int) *Entity {
	entity.ac = ac
	return entity
}

func (entity *Entity) SetName(name string) *Entity {
	entity.name = name
	return entity
}

func (entity *Entity) SetBaseWeapon(baseWeapon string) *Entity {
	entity.baseWeapon = baseWeapon
	return entity
}
