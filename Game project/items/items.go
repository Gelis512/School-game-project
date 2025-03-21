package classes

type Item struct { //Von Hanna erstellte Klasse, unbenutzt
	name     string
	itemType string
	hpregen  int
	dmgboost int
}

func (item *Item) GetName() string {
	return item.name
}

func (item *Item) GetType() string {
	return item.itemType
}

func (item *Item) GetHpregen() int {
	return item.hpregen
}

func (item *Item) GetDmgboost() int {
	return item.dmgboost
}

func (item *Item) SetName(name string) {
	item.name = name
}
func (item *Item) SetType(itemType string) {
	item.itemType = itemType
}
func (item *Item) SetHpregen(hpregen int) {
	item.hpregen = hpregen
}
func (item *Item) SetDmgboost(dmgboost int) {
	item.dmgboost = dmgboost
}
