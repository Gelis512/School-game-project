package main

import (
	"fmt"
	"main/character"
	"main/playableMap"
	"main/room"
	"os"

	"github.com/eiannone/keyboard"
)

var mapHeight int = 200
var mapWidth int = 200
var mapHeightP *int = &mapHeight
var mapWidthP *int = &mapWidth

func mapRender(char *character.Character, mapSlice [][]*room.Room) {

	for i0 := 0; i0 < mapWidth; i0++ {
		for i1 := 0; i1 < mapHeight; i1++ {
			if i0 == char.GetPosX() && i1 == char.GetPosY() {
				fmt.Print("P")
				mapSlice[i0][i1].SetRoomType(".")
			} else {
				fmt.Print(mapSlice[i0][i1].GetRoomType())
			}
		}
		fmt.Println()
	}
}

func playerMovement(player *character.Character, mapSlice [][]*room.Room) {
	fmt.Println("You are the P symbol")
	fmt.Println("Please use the arrow keys to move, press ESC to exit the program: ")
	mapRender(player, mapSlice)
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	char, key, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}
	fmt.Println(char, key)
	switch key {
	case keyboard.KeyArrowUp:
		if player.GetPosX() != 0 {
			var charNewPosX = player.GetPosX() - 1
			player.SetPosX(charNewPosX)
		}
	case keyboard.KeyArrowLeft:
		if player.GetPosY() != 0 {
			var charNewPosY = player.GetPosY() - 1
			player.SetPosY(charNewPosY)
		}
	case keyboard.KeyArrowDown:
		if player.GetPosX() != mapWidth {
			var charNewPosX = player.GetPosX() + 1
			if charNewPosX >= mapWidth {
				charNewPosX = charNewPosX - 1
			}
			player.SetPosX(charNewPosX)
		}
	case keyboard.KeyArrowRight:
		if player.GetPosY() != mapHeight {
			var charNewPosY = player.GetPosY() + 1
			if charNewPosY >= mapHeight {
				charNewPosY = charNewPosY - 1
			}
			player.SetPosY(charNewPosY)
		}
	case keyboard.KeyEsc:
		os.Exit(3)
	}
	mapSlice[player.GetPosX()][player.GetPosY()].RoomInteraction()
}

func test() {
	for mapHeight > 100 || mapHeight < 0 {
		var inputInt int
		fmt.Println("Please input the map height(max 100): ")
		fmt.Scan(&inputInt)
		*mapHeightP = inputInt
	}
	for mapWidth > 100 || mapWidth < 0 {
		var inputInt int
		fmt.Println("Please input the map width (max 100): ")
		fmt.Scan(&inputInt)
		*mapWidthP = inputInt
	}
	var playerCharacter *character.Character = character.NewCharacter(mapWidth/2, mapHeight/2)
	fmt.Println(playerCharacter.GetPosX(), playerCharacter.GetPosY())
	fmt.Println("Map: ")
	var playableMap [][]*room.Room = playableMap.GenerateMap(mapWidth, mapHeight)
	mapRender(playerCharacter, playableMap)
	for {
		playerMovement(playerCharacter, playableMap)
	}
}

func main() {

	test()
}
