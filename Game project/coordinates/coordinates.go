package coordinates

type Coordinate struct { //Jede Position auf der Karte ist ein Objekt, der einen Koordinatenwert sowie Eigenschaften besitzt
	x int
	y int
}

func (coordinates *Coordinate) GetX() int {
	return coordinates.x
}

func (coordinates *Coordinate) GetY() int {
	return coordinates.y
}

func (coordinates *Coordinate) SetX(x int) *Coordinate {
	coordinates.x = x
	return coordinates
}

func (coordinates *Coordinate) SetY(y int) *Coordinate {
	coordinates.y = y
	return coordinates
}

func NewCoordinate(xCord int, yCord int) *Coordinate {
	var c *Coordinate = new(Coordinate)
	c.SetX(xCord)
	c.SetY(yCord)
	return c
}
