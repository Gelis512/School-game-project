package playableMap

import (
	"main/room"
	"math/rand/v2"
)

func GenerateMap(width int, height int) [][]*room.Room {
	roomNames := [5]string{"X", "O", "O", "O", "O"}
	mapLayout := make([][]*room.Room, width)
	for i := 0; i < width; i++ {
		mapLayout[i] = make([]*room.Room, height)
	}

	/*
		fmt.Println(roomNames, rand.IntN(2))
		mapLayout[0][0] = NewRoom("test", 5, 5)
		fmt.Println(mapLayout)
		mapLayout[0][1] = NewRoom("test", 5, 5)
		fmt.Println(mapLayout)
		mapLayout[0][2] = NewRoom("test", 5, 5)
		fmt.Println(mapLayout)
		mapLayout[1][0] = NewRoom("test", 5, 5)
		fmt.Println(mapLayout)
		mapLayout[1][1] = NewRoom("test", 5, 5)
		fmt.Println(mapLayout)
		mapLayout[1][2] = NewRoom("test", 5, 5)
		fmt.Println(mapLayout)
		mapLayout[2][0] = NewRoom("test", 5, 5)
		fmt.Println(mapLayout)
		mapLayout[2][1] = NewRoom("test", 5, 5)
		fmt.Println(mapLayout)
		mapLayout[2][2] = NewRoom("test", 5, 5)
		fmt.Println(mapLayout)
	*/

	for i0 := 0; i0 < width; i0++ {
		for i1 := 0; i1 < height; i1++ {
			mapLayout[i0][i1] = room.NewRoom(roomNames[rand.IntN(5)], i0, i1)
		}
	}
	//mapDisplay(width, height, mapLayout)
	return mapLayout
}

/*
func mapDisplay(len int, wid int, mapSlice [][]*room.Room) {
	for i0 := 0; i0 < len; i0++ {
		for i1 := 0; i1 < wid; i1++ {
			fmt.Print(mapSlice[i0][i1].GetRoomType())
		}
		fmt.Println()
	}
} */
