package room

import (
	"fmt"
	"main/coordinates"
)

type Room struct {
	position *coordinates.Coordinate
	roomType string
	symbol   string
}

func NewRoom(t string, xCord int, yCord int) *Room {
	var r *Room = new(Room)
	r.SetRoomType(t)
	r.SetPosition(coordinates.NewCoordinate(xCord, yCord))
	return r
}

func (room *Room) GetPosition() *coordinates.Coordinate {
	return room.position
}

func (room *Room) GetRoomType() string {
	return room.roomType
}

func (room *Room) GetSymbol() string {
	return room.symbol
}

func (room *Room) SetPosition(position *coordinates.Coordinate) *Room {
	room.position = position
	return room
}

func (room *Room) SetRoomType(roomType string) *Room {
	room.roomType = roomType
	return room
}

func (room *Room) SetSymbol(symbol string) *Room {
	room.symbol = symbol
	return room
}

func (room *Room) RoomInteraction(gameCurrentScore int, gameMaxScore int) bool { //Wenn der Spieler bei einem X ist, wird ein Auswahl angebeten (allgemein unbenutzt)
	for loop := true; loop; {
		switch room.GetRoomType() { //Dieses switch case erlaubt alternative Räume hinzuzufügen
		case "X":
			fmt.Println("You are in a room. Type the number of the action you want:")
			fmt.Println("1. Test action 1")
			fmt.Println("2. Test action 2") //Testaktionen, könnten für etwas in der Zukunft verwendet werden. Zurzeit machen das gleiche: einen Punkt an den Spieler geben
			var inputStr string
			fmt.Scan(&inputStr)
			switch inputStr {
			case "1":
				fmt.Println("You selected action 1, and got a point! Please press any button to continue")
				fmt.Println("Current score: ", gameCurrentScore+1, "/", gameMaxScore) //Der Punktenanzahl muss mit +1 ausgegeben werden, da es nur später inkrementiert wird
				loop = false
				return true //return ob der Spieler einen Punkt bekommen hat oder nicht (der Plan war eine Option zu geben, denn X nicht zu sammeln)
			case "2":
				fmt.Println("You selected action 2, and got a point! Please press any button to continue")
				fmt.Println("Current score: ", gameCurrentScore+1, "/", gameMaxScore)
				loop = false
				return true
			default:
				fmt.Println("Wrong option, try again")
			}
		default:
			loop = false
		}
	}
	return false
}
