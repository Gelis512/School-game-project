package main

import (
	"fmt"
	"main/character"
	"main/playableMap"
	"main/room"
	"main/story"
	"os"
	"os/exec"

	"github.com/eiannone/keyboard"
)

var mapHeight int = 200 //ein Placeholder-Wert wird an die beiden Kartengröße Variablen gegeben
var mapWidth int = 200
var mapHeightP *int = &mapHeight //Pointer auf die Größenvariablen (eigentlich unnötig, da ich die Variablen außerhalb von der main Funktion erstellt habe)
var mapWidthP *int = &mapWidth
var currentScore int

func WinScreen(playerCharacter *character.Character) { //Bei der Erfüllung der Bedingungen wird das Spiel gewonnen bzw. abgebrochen
	cmd := exec.Command("cmd", "/c", "cls") //diese drei Zeilen Entfernen alles aus der Konsole - Die erste erstellt einen Pointer zu einem neuem Command Objekt mit den Argumenten in Klammern
	cmd.Stdout = os.Stdout                  //Dieses Teil verstehe ich nur teilweise. Das Feld Stdout von cmd wird zu os.Stdout gesetzt
	cmd.Run()                               //Das Command wird ausgeführt
	fmt.Println("Congratulations, ", playerCharacter.GetName(), " , you've collected all the points! YOU WIN!!!")
	os.Exit(0) //Der Programm wird abgebrochen mit dem Zustand 3
}

func playerMovement(player *character.Character, mapSlice [][]*room.Room, gameCurrentScore int, gameMaxScore int) { //Die Position des Spielers wird entsprechend den gedrückten Tasten geändert
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	playableMap.MapRender(player, mapSlice, currentScore, mapWidth, mapHeight) //Zuerst wird die Karte ausgegeben
	if err := keyboard.Open(); err != nil {                                    //Die Tastatur-aufnahme wird eingeschaltet, wenn es ein Error beim aufnehmen der gedruckten Taste gibt, wird ein Fehler ausgegeben und der Program abgebrochen
		panic(err)
	}
	defer func() { //Weiß ich leider nicht genau was defer macht
		_ = keyboard.Close()
	}()
	_, key, err := keyboard.GetKey() //Die unterschiedliche Werte des aufnehmens einer Taste werden gespeichert (_ löscht den Wert, da es nicht verwendet wird)
	if err != nil {                  //wenn es ein Error beim aufnehmen der gedruckten Taste gibt, wird ein Fehler ausgegeben und der Program abgebrochen
		panic(err)
	}
	switch key { //in Abhängigkeit von dem, welche Taste wird gedruckt, wird die Position entsprechend geändert
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
			if charNewPosX >= mapWidth { //Da hier die Position des Spielers nicht von den Rändern des Slices begrenzt ist, muss sie in der Abhängigkeit von der Kartengröße entsprechend begrenzt werden
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
	case keyboard.KeyEsc: //Esc bricht den Program ab
		os.Exit(0)
	}
	if mapSlice[player.GetPosX()][player.GetPosY()].RoomInteraction(gameCurrentScore, gameMaxScore) { //Wenn bei der Interaktion mit einem Raum ein Punkt gegeben wird, wird der currentScore inkrementiert
		currentScore++
	}
	if currentScore == playableMap.ReturnMaxScore() { //hier wird überprüft, ob alle Punkte gesammelt wurden
		WinScreen(player)
	}
}

func game() {
	for mapHeight > 8 || mapHeight < 0 { //Erstellung der Karte, die Größe wird begrenzt, da je größer die Karte, desto weniger spielbar wird es
		var inputInt int                                    //Wenn die Karte länger als die Größe der Zeile ist, bricht der Bewegungssystem visuell
		fmt.Println("Please input the map height(max 8): ") //Außerdem dauert es auch sehr lange, wenn es zu viele X gibt
		fmt.Scan(&inputInt)
		*mapHeightP = inputInt
	}
	for mapWidth > 8 || mapWidth < 0 {
		var inputInt int
		fmt.Println("Please input the map width (max 8): ")
		fmt.Scan(&inputInt)
		*mapWidthP = inputInt
	}
	var playerCharacter *character.Character = character.NewCharacter(mapWidth/2, mapHeight/2) //der spieler wird in die Mitte gestellt, da es um Integer handelt, muss man nicht aufrunden
	var gameMap [][]*room.Room = playableMap.GenerateMap(mapWidth, mapHeight)                  //Die Karte wird erstellt und als gameMap gespeichert
	var mapMaxScore int = playableMap.ReturnMaxScore()                                         //maxScore wird aus dem package playableMap returned, um es hier weiter zu verwenden
	playableMap.MapRender(playerCharacter, gameMap, currentScore, mapWidth, mapHeight)
	for {
		playerMovement(playerCharacter, gameMap, currentScore, mapMaxScore) //Die Bewegung des Spielers findet so lange statt, bis das Spiel beendet ist. Deshalb ist es in einem unendlichem for-Loop
	}
}

func main() {
	for { //Auswahl, ob die Story angezeigt werden soll, wenn nicht, dann wird die Name des Spielers am Anfang angefragt
		var inputStr string
		fmt.Println("Would you like to experience the story? y/n")
		fmt.Scan(&inputStr)
		if inputStr == "y" {
			story.StoryStart1()
			story.StoryStart2()
			//story.Tag() //Der Schleife den Tagen von Lillian, unbenutzt
			break
		} else if inputStr == "n" {
			fmt.Println("You've skipped the story.")
			break
		}

	}
	game() //Die allgemeine Spielfunktion wird ausgeführt
}
