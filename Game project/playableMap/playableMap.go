package playableMap

import (
	"fmt"
	"main/character"
	"main/room"
	"math/rand/v2"
)

var maxScore int

func GenerateMap(width int, height int) [][]*room.Room { //Eine Karte wird in einem zweidimensionalen Slice zum ersten Mal erstellt
	maxScore = 0
	roomNames := [5]string{"X", ".", ".", ".", "."} //Dabei werden zufällig die X-Räume verteilt
	mapLayout := make([][]*room.Room, width)
	for i := 0; i < width; i++ { //Die Karte wird erstellt und ausgegeben
		mapLayout[i] = make([]*room.Room, height) //ein zweidimensionaler Slice wird erstellt
	}
	for i0 := 0; i0 < width; i0++ {
		for i1 := 0; i1 < height; i1++ {
			var nameInt int = rand.IntN(5)
			mapLayout[i0][i1] = room.NewRoom(roomNames[nameInt], i0, i1)
			if mapLayout[i0][i1].GetRoomType() == "X" { //Berechnung der maximalen zu erreichenden Menge an Punkten
				maxScore++
			}
		}
	}
	fmt.Println("The map has ", maxScore, " points to collect")
	return mapLayout //return den Karten-Slice
}

func MapRender(char *character.Character, mapSlice [][]*room.Room, currentScore int, mapWidth int, mapHeight int) { //Ausgabe der Karte und den Informationen
	fmt.Println("You are the P symbol")
	fmt.Println("Current score: ", currentScore, "/", ReturnMaxScore())
	fmt.Println("Please use the arrow keys to move, press ESC to exit the program: ")
	for i0 := 0; i0 < mapWidth; i0++ {
		for i1 := 0; i1 < mapHeight; i1++ {
			if i0 == char.GetPosX() && i1 == char.GetPosY() { //Der Kartenslice wird vollständig ausgegeben, dabei wird der Typ des
				fmt.Print("P")                    //Räumes bei dem Spieler auf . gesetzt, um das Sammeln von X zu erlauben.
				mapSlice[i0][i1].SetRoomType(".") //Außerdem wird der Spieler bei einer Übereinstimmung mit einer Koordinate als P ausgegeben
			} else {
				fmt.Print(mapSlice[i0][i1].GetRoomType())
			}
		}
		fmt.Println()
	}
}

func ReturnMaxScore() int {
	return maxScore
}
