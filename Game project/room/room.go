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

func (room *Room) RoomInteraction() {

	for loop := true; loop; {
		switch room.GetRoomType() {
		case "X":
			fmt.Println(`You are in a room. Type the number of the action you want:
1. Test action 1
2. Test action 2`)
			var inputStr string
			fmt.Scan(&inputStr)
			switch inputStr {
			case "1":
				fmt.Println("You selected action 1, please press any button to continue")
				loop = false
			case "2":
				fmt.Println("You selected action 2, please press any button to continue")
				loop = false
			default:
				fmt.Println("Wrong option, try again")
			}
		default:
			loop = false
		}
	}

}
