//Story des Spiel
//Lillian Hummel

package story

import (
	"fmt"
	"math/rand/v2"
)

var y string //Namensvariable für den Spielernamen, da Charakterpacket noch nicht fertig und importiert

func GetPlayerStoryName() string {
	return y
}

func StoryStart1() { //Erster Tag der Einführung von 2

	var x string //VArible eingabe zum Text weiterkommen

	var z int //VAriable für die Endscheidung

	fmt.Println(" Der Dungen von Lurenz ") //fmt.Println --> Zeigt den Text in Klammern an

	fmt.Println(" ")

	fmt.Println(" ")

	fmt.Println("-Um zum weiteren Text zukommen, drücke bitte Enter-")

	fmt.Println(" ")

	fmt.Println("------------------------------------------------------------------------------------------------------------------------ ")

	fmt.Println(" ")

	fmt.Scanln(&x) //fmt.Scanln --> Leist die eingabe und peichert die als x

	fmt.Println("Die Sonne verschwindet langsam hinterm Horizont. Der Himmel färbt sich in verschiedenen Tönen. ")

	fmt.Println("Langsam schließen die Läden am Straßenrand.")

	fmt.Scanln(&x)

	fmt.Println("Müde läufst du die Hauptstraße entlang. Nach drei Tagen bist du endlich in Lurenz angekommen. ")

	fmt.Println("Es war ein beschwerlicher Weg.")

	fmt.Scanln(&x)

	fmt.Println("Nun bist du auf der Suche nach einem Schlafplatz.")

	fmt.Println("Auf der linken Straßenseite siehst du eine kleine Gaststätte “der Krug” und hoffst das diese noch ein Bett für die Nacht hat.")

	fmt.Scanln(&x)

	fmt.Println("Die Gaststätte ist voll, als du Sie betrittst.")

	fmt.Println("Menschen Essen, sie trinken und schwatzen.")

	fmt.Scanln(&x)

	fmt.Println("Durch die Menschenmenge bahnst du dir einen Weg zu den Tressen, hinter dem der Wirt steht. ")

	fmt.Scanln(&x)

	fmt.Println("“Abend”, grüßt der schon etwas in die Jahre gekommene Mann.")

	fmt.Scanln(&x)

	fmt.Println("Spieler: “Guten Abend. Haben Sie noch ein Bett frei für heute Nacht?”")

	fmt.Scanln(&x)

	fmt.Println("Wirt: “Ja. Das macht mit Essen heute Abend und morgen 20 Bien”")

	fmt.Scanln(&x)

	fmt.Println("Spieler: “In Ordnung” ")

	fmt.Scanln(&x)

	fmt.Println("-Du bezahlst die 20 Bien.-")

	fmt.Scanln(&x)

	fmt.Println("Wirt: “Auf welchen Namen soll ich das Zimmer buchen?”")

	fmt.Scanln(&x)

	fmt.Println("-Bitte gebe deinen Namen ein-")

	fmt.Println(" ")

	//Fest legung des Namens
	for y == "" {
		fmt.Scanln(&y)
		if y == "" {
			fmt.Println("Bitte gebe einen passenden Namen ein")
		}
	}

	fmt.Println(" ")

	fmt.Println("Wirt: “Okay", y, ",also Herzlich Wilkommen. Möchtest du in der Gaststube Essen oder auf deinem Zimmer?“") //

	fmt.Println(" ") //Schleife von Herr Herker geklaut um zu endscheiden ob auf dem Zimmer essen oder in der Stube

	for z != 1 && z != 2 { //Nur wenn 1 oder 2 Eingegeben wird geht es weiter

		fmt.Print("Bitte wähle 1 wenn du in der Stube Essen möchtest und 2 wenn du auf deinem Zimmer Essen möchtest.")

		fmt.Println(" ")

		fmt.Scanln(&z)
		if z == 1 { //Option 1 => Gaststube

			fmt.Println(" ")

			fmt.Println(" Du gehst zurück in Gaststube. Dort suchst du dir einen Platz und beobachtest die Leute.  ")

			fmt.Scanln(&x)

			fmt.Println("In der einen Ecke setzen ein paar alte Damen und schwatzen. In der anderen spielen ein paar Herren Skard. ")

			fmt.Println("In der Mitte des Raumes sitzen Paare, Freunde und einzelne Leute und genießen ihr Essen. ")

			fmt.Scanln(&x)

			fmt.Println("Es dauert nicht lange bis dein Essen dir gebracht wird. Der Duft eines Eintopfes steigt dir in die Nase.")

			fmt.Scanln(&x)

			fmt.Println("Bedienung: „Bitte schön. Rindereintopf mit frischem Brot und ein Glass Wasser. Guten Appetit.“  ")

			fmt.Scanln(&x)

			fmt.Println(y, ": “Vielen Dank” ")

			fmt.Scanln(&x)

			fmt.Println("Du fängst an zu essen.")

			fmt.Println("Es schmeckt gut. Der Eintopf schmeckt hervorragend und das Brot ist frisch gebacken, warm und weich. ")

			fmt.Scanln(&x)

			fmt.Println("Nach dem du deinem Essen fertig bist, schlurfst du zu deinem Bett. Es war ein langer Tag ")

			fmt.Println("Sobald dein Kopf die Kissen berührt, schläfst du ein.")

			break

		} else if z == 2 { //Option 2 =>Zimmer

			fmt.Println(" ")

			fmt.Println("Du machst dich auf den Weg zu deinem Zimmer. Es ist nicht groß und es ist nicht klein.  ")

			fmt.Scanln(&x)

			fmt.Println("In dem Zimmer befindet sich ein Bett, eine Kommode, ein Tisch und ein Stuhl. Eine Tür führt in ein kleines Bad.")

			fmt.Scanln(&x)

			fmt.Println("Ein Klopfen an der Tür kündigt dein Essen an. Du öffnest die Tür und nimmst das Tablet der Bedienung entgegen.")

			fmt.Scanln(&x)

			fmt.Println("Bedienung: „Bitte schön. Rindereintopf mit frischem Brot und ein Glass Wasser. Guten Appetit.“  ")

			fmt.Scanln(&x)

			fmt.Println(y, ": “Vielen Dank” ")

			fmt.Scanln(&x)

			fmt.Println("Du fängst an zu essen")

			fmt.Println("Es schmeckt gut. Der Eintopf schmeckt hervorragend und das Brot ist frisch gebacken, warm und weich")

			fmt.Scanln(&x)

			fmt.Println("Nach dem du dein Essen aufgegessen hast machst du dich Bett fertig. Ein paar Minuten später liegst du im Bett und schläfst ein. ")

			break
		}

	}

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func StoryStart2() { //Zweiter Tag der Einführung von 2

	var x string //Variable zum weiter kommen Text

	fmt.Println(" ")

	fmt.Println(" ")

	fmt.Println("---------------------------------------------------------------------------------------- ")

	fmt.Println(" ")

	fmt.Println(" Nach einer ruhigen Nacht, erwachst du gut ausgeruht in deinem Zimmer.")

	fmt.Println(" Du stehst auf sammelst alles was du für den neuen Tag brauchst zusammen und machst dich fertig. ")

	fmt.Scanln(&x)

	fmt.Println(" Nach dem du deine Tasche und deine Sachen zusammen hast, verlässt du dein Zimmer, und gehst in die Gaststube. ")

	fmt.Scanln(&x)

	Frühstück() //Frühstück Funktion funktionier

	fmt.Scanln(&x)

	fmt.Println(y, ":“Danke für’s Frühstück. Gibt es eine Möglichkeit hier Geld zuverdienen?“ ")

	fmt.Scanln(&x)

	fmt.Println(" Wirt: “Mhm du könntest es bei Jörg probieren.“ ")

	fmt.Scanln(&x)

	fmt.Println(y, ": „Jörg?“  ")

	fmt.Scanln(&x)

	fmt.Println("Wirt: „Ja Jörg er ist der Dungen Verwalter. Bei ihm kannst du dir eine Zulassung hohlen. ")

	fmt.Scanln(&x)

	fmt.Println(y, ": „Und was mache ich mit einer Zulassung?“ ")

	fmt.Scanln(&x)

	fmt.Println("Wirt: „Mit der Zulassung darfst du in den Dungen um Monster zu bekämpfen. Den Rest kann dir Jörg besser erklären. ")

	fmt.Scanln(&x)

	fmt.Println(y, ":„Hört sich gut an. Wo finde ich Jörg?“ ")

	fmt.Scanln(&x)

	fmt.Println("Wirt: „Geh die Straße runter, nach Norden. Ganz am Ende findest du Margrets Krämmerladen. Geh hinein und frage nach Jörg.“")

	fmt.Scanln(&x)

	fmt.Println(y, ": „Okay. Vielen Dank.“ ")

	fmt.Println("Du machst dich auf den Weg. ")

	fmt.Scanln(&x)

	fmt.Println("Die Hauptstraße von Lurenz ist in den Morgenstunden leer. ")

	fmt.Println("Einzelne Läden öffnen grade ihre Türen. ")

	fmt.Println("Nachbarn grüßen einander.")

	fmt.Scanln(&x)

	fmt.Println("Du läufst Richtung Norden und folgst der Straßen. Nach 5 Minuten siehst du ein Gebäude am Ende der Straße. ")

	fmt.Println("In großen Buchstaben steht MARGRETS KRÄMMERLADEN über der Tür. ")

	fmt.Scanln(&x)

	fmt.Println("Als du die Tür öffnest, ertönt eine kleine Glocke. ")

	fmt.Println("Frau: “Klein Moment. Ich bin gleich für Sie da”, ertönt aus einem Hinterzimmer. ")

	fmt.Scanln(&x)

	fmt.Println("Eine etwas ältere Frau mit Kleid und Schürze taucht im Türrahmen hinter der Kasse auf. ")

	fmt.Scanln(&x)

	fmt.Println("Frau: “Guten Tag.  Willkommen in Margrets Krämmerladen.  ")

	fmt.Println("Ich bin Margret und bei mir können sie Rubis in Bien umtauschen, Waffen, Tränke und Haushaltsutensilien kaufen.” ")

	fmt.Scanln(&x)

	fmt.Println("Spieler: “Hallo mein Name ist", y, ". Der Wirt aus „dem Krug“  hat mich hierher geschickt.  ")

	fmt.Println("Er meinte, wenn ich Geld verdienen will, kann ich es bei einem Jörg probieren?”")

	fmt.Scanln(&x)

	fmt.Println("Margret: “Ah ja, eine Sekunde bitte. JÖRG!!!")

	fmt.Scanln(&x)

	fmt.Println("Ein Mann mit grauen Harre kommt aus dem Hinterzimmer.")

	fmt.Println("Jörg: “Was ist den Margret?” ")

	fmt.Scanln(&x)

	fmt.Println("Margret: “Hier ist jemand für den Dungen.” ")

	fmt.Println("Jörg mustert dich von oben bis unten. ")

	fmt.Scanln(&x)

	fmt.Println("Jörg: “Du willst in den Dungen?” ")

	fmt.Scanln(&x)

	fmt.Println(y, ": “Äh wenn ich damit Geld verdienen kann… Ja.”")

	fmt.Scanln(&x)

	fmt.Println("Jörg: „Hast du Kampferfahrung?“")

	fmt.Scanln(&x)

	fmt.Println(y, ": „Ja“")

	fmt.Scanln(&x)

	fmt.Println(" Jörg: “ Gut komm mit.”")

	fmt.Println("Du folgst Jörg, als dieser den Laden verlässt.  ")

	fmt.Scanln(&x)

	fmt.Println(" Ein schmaler Pfad verläuft an der Hausseite. Der Pfad fürt in den Wald hinter Lurenz. ")

	fmt.Println("Der Pfad ist schmal und kaum genutz. Der Wald dicht und lebendig. Vögel singen und im Unterholz raschelt es.")

	fmt.Scanln(&x)

	fmt.Println("Nach einer guten halben Stunde erreichen Jörg und du eine Höhle, welche mit einer Eisen Tür verschlossen ist. ")

	fmt.Scanln(&x)

	fmt.Println("Jörg: „Willkommen im Dungen von Lurenz.“ ")

	fmt.Println("„Lass es mich erklären. Der Dungeon ist seit seiner Entdeckung im Besitz meiner Familie.  ")

	fmt.Println("Mein Uropa entwickelte das Genehmigungskonzept. Das heißt, nur Leute mit Erlaubnis von mir, komm hierein.“ ")

	fmt.Scanln(&x)

	fmt.Println("Jörg deutet auf die Tür. ")

	fmt.Scanln(&x)

	fmt.Println("Jörg: „Wenn du den Dungeon betrifft, landest du in einem Raum, in welchem sich Monster befinden. Diese werden auf dieser Karte angezeigt.“")

	fmt.Println("-Jörg gibt dir eine Karte-   ")

	fmt.Scanln(&x)

	fmt.Println("Jörg: „Diese Monster greifen erst an, wenn du auf sie zu gehst ")

	fmt.Println("Solltest du den Dungen verlassen wollen, gehst du einfach wieder zum Ausgang. Solltest du in diesem Dungen sterben, dann ähm… war´s das für dich.  ")

	fmt.Println("Jedes Monster lässt, nach dem du es besiegt hast, einen Ru fallen. Die Ru´s kannst du bei Margret eintauschen. 1 Ru zu 10 Bien. ")

	fmt.Println("Noch Fragen?“ ")

	fmt.Scanln(&x)

	fmt.Println(y, ": „Nein.“  ")

	fmt.Scanln(&x)

	fmt.Println("Jörg: „Gut, hier ist der Schlüssel für die Tür.  ")

	fmt.Println("Schließe ab, wenn du den Dungen verlässt. Den Schlüssel gibst du Margret ab und kannst ihn dir am nächsten Tag wieder abholen. ")

	fmt.Println("Viel Glück.“ ")

	fmt.Println("Du betrittst den Dungen von Lurenz. ")

	fmt.Scanln(&x)

	//Funktion für den Dungen//////////////////////////					//

	fmt.Scanln(&x)

	fmt.Println("Mit deiner Ausbeute verlässt du den Dungen. Nach dem Verschließen der Tür folgst du dem Pfad zurück nach Florenz. ")

	fmt.Scanln(&x)

	fmt.Println("An Margrets Krämerladen angekommen betritst du den Laden. ")

	fmt.Println("Margret: „Du bist wieder da. Wie war´s?“ ")

	fmt.Scanln(&x)

	fmt.Println(y, ": Aufregend. Ich freue mich auf morgen.“ ")

	fmt.Scanln(&x)

	fmt.Println("Margret: „Das ist gut. Zeig her deine Ausbeute.“ ")

	fmt.Scanln(&x)

	fmt.Println("-Du gibst Margret die Ru´s.- ")

	fmt.Scanln(&x)

	//Umtauschfunktion 													//Funktion zum Tauschen von Ru zu Bien -> exestiert noch nicht

	fmt.Println("Margret: „ Hier bitte deine Bien. Sonst noch was? ")

	fmt.Scanln(&x)

	//Handelmöglich														//Funktion zum Handeln -> exestiert noch nicht

	fmt.Println(y, ": „Nein Danke. Bis morgen ") //Satz da Funktion Handel nicht exsestiert

	fmt.Println("Margret: „Bis Morgen“")

	fmt.Println("Du verlässt den Laden und siehst, wie Ladenbesitzer ihre Geschäfte schließen.")

	fmt.Scanln(&x)

	fmt.Println("Du machst dich auf den Weg zurück „zum Krug“")

	fmt.Println("Als du diesen betritts, merkst du es ist genau so voll wie zu vor.")

	fmt.Println("Du gehst zur Theke.")

	fmt.Scanln(&x)

	fmt.Println("Wirt: „Abend. Selbe wie gestern?“  ")

	fmt.Scanln(&x)

	fmt.Println("Spieler: „ Ja, Bitte.“  ")

	fmt.Scanln(&x)

	fmt.Println("Wirt: „20 Bien Bitte.“ ")

	fmt.Scanln(&x)

	fmt.Println("-Du gibst die 20 Bien dem Wirt- ")

	fmt.Scanln(&x)

	Abendessen()

	fmt.Scanln(&x)

	fmt.Println("Du hast dein Abendessen aufgegessen und schleppst dich auf dein Zimmer.")

	fmt.Scanln(&x)

	fmt.Println("Nachdem du dich Umgezogen und Bett fertig gemacht hast, fällst du wie ein Stein ins Bett. ")

	fmt.Println("Morgen ist ein neuer Tag. ")

	fmt.Println("------------------------------------------------------------------------------------------------ ")

}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func zufall(max int) int { //Zuffallfunktion für die Wettervariable, von HErrn HErker geklaut
	return rand.IntN(max + 1)

}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Wetter() { //Funktion Wetter zur Induvidialliesierung des Tages, keine Auswirkung auf den Tag

	var x int = zufall(10) //Zufalls variable

	// || Zeichen für oder

	for {

		fmt.Println("Am nächsten Tag")

		if x == 1 { //Sonniges Wetter

			fmt.Println("Das Singen der Vögel, weckt dich aus deinem Schlaf. Die Sonne strahlt.")

			fmt.Println("Es sieht so aus, also ob heute ein schöner Tag sein wird.")

			break

		} else if x == 2 {

			fmt.Println("Das Singen der Vögel, weckt dich aus deinem Schlaf. Die Sonne strahlt.")

			fmt.Println("Es sieht so aus, also ob heute ein schöner Tag sein wird.")

			break

		} else if x == 3 {

			fmt.Println("Das Singen der Vögel, weckt dich aus deinem Schlaf. Die Sonne strahlt.")

			fmt.Println("Es sieht so aus, also ob heute ein schöner Tag sein wird.")

			break

		} else if x == 4 { //Regenwetter

			fmt.Println("Das Trommeln von Regentropfen an der Fensterscheibe, begrüßt dich als du aufwachst.")

			fmt.Println("Zum Glück regnet es nicht im Dungen")

			break

		} else if x == 5 {

			fmt.Println("Das Trommeln von Regentropfen an der Fensterscheibe, begrüßt dich als du aufwachst.")

			fmt.Println("Zum Glück regnet es nicht im Dungen")

			break

		} else if x == 6 {

			fmt.Println("Das Trommeln von Regentropfen an der Fensterscheibe, begrüßt dich als du aufwachst.")

			fmt.Println("Zum Glück regnet es nicht im Dungen")

			break

		} else if x == 7 { //Bewölktes Wetter

			fmt.Println("Langsam erwachst du aus deinem Schlaf. Durchs Fenster erhascht du einen Blick auf den Hellgrauen Himmel.")

			fmt.Println("Immerhin wirst du Wetter technisch nichts verpassen.")

			break

		} else if x == 0 {

			fmt.Println("Langsam erwachst du aus deinem Schlaf. Durchs Fenster erhascht du einen Blick auf den Hellgrauen Himmel.")

			fmt.Println("Immerhin wirst du Wetter technisch nichts verpassen.")

			break

		} else if x == 8 {

			fmt.Println("Langsam erwachst du aus deinem Schlaf. Durchs Fenster erhascht du einen Blick auf den Hellgrauen Himmel.")

			fmt.Println("Immerhin wirst du Wetter technisch nichts verpassen.")

			break

		} else if x == 9 {

			fmt.Println("Langsam erwachst du aus deinem Schlaf. Durchs Fenster erhascht du einen Blick auf den Hellgrauen Himmel.")

			fmt.Println("Immerhin wirst du Wetter technisch nichts verpassen.")

			break

		} else if x == 10 { //Sturm

			fmt.Println("Donnergrollen reißt dich aus dem Schlaf. Blitze zucken über den Himmel.")

			fmt.Println("Heute ist es aber so richtig ungemütlich.")

			break
		}

	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Frühstück() { //Funktion zur Frühstück auswahl für Zweiter Tag und Tag

	var z int

	var x int //Schleifenvariable

	fmt.Println("Wirt: „ Guten Morgen. Es gibt Frühstück. Was möchtest du essen?“")

	fmt.Scanln(&x)

	fmt.Println(y, " : „Was gibt es den zum Essen?“")

	fmt.Scanln(&x)

	fmt.Println("Wirt: “ Brötchen mit Wurst, Brötchen mit Käse, Brötchen mit Mamelade. Such dir was aus. ")

	fmt.Scanln(&x)

	for {
		fmt.Print("Bitte wähle 1 wenn du Brötchen mit Wurst, 2 wenn du Brötchen mit Käse, 3 wenn du Brötchen mit Mamelade.")

		fmt.Println(" ")

		fmt.Scanln(&z)
		if z == 1 { //Option 1 => Wurst

			fmt.Println(y, ": „Bitte das Brötchen mit Wurst“")

			fmt.Println("Wirt: „Bitte schön deine Brötchen mit Wurst lass es dir schmecken.“  ")

			fmt.Println("Du suchst dir einen Platz in der Stube und setzt dich hin.   ")

			fmt.Scanln(&x)

			fmt.Println(" Das Brötchen ist noch warm und knackig. Die Wurst ist lecker. ")

			fmt.Println("Schnell ist der Teller leer und du bringst ihn dem Wirt zurück. ")

			break

		} else if z == 2 { //Option 2 => Käse

			fmt.Println(y, ": „Bitte das Brötchen mit Käse“")

			fmt.Println("Wirt: „Bitte schön deine Brötchen mit Käse lass es dir schmecken.“  ")

			fmt.Println("Du suchst dir einen Platz in der Stube und setzt dich hin.  ")

			fmt.Scanln(&x)

			fmt.Println("Das Brötchen ist noch warm und knackig. Der Käse ist lecker.")

			fmt.Println("Schnell ist der Teller leer und du bringst ihn dem Wirt zurück. ")

			break

		} else if z == 3 { //Option 3 =>Mamelade

			fmt.Println(y, ": „Bitte das Brötchen mit Mamelade“")

			fmt.Println(" Wirt: „Bitte schön deine Brötchen mit Käse lass es dir schmecken.“  ")

			fmt.Println("Du suchst dir einen Platz in der Stube und setzt dich hin.   ")

			fmt.Scanln(&x)

			fmt.Println("Das Brötchen ist noch warm und knackig. Die Marmelade ist fruchtig. ")

			fmt.Println("Schnell ist der Teller leer und du bringst ihn dem Wirt zurück. ")

			break
		}
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Abendessen() { //Funktion zum Abendessen auswahl für Zweiter Tag und Tag

	var z int //Schleifenvariable

	var x int

	fmt.Println("Wirt: „ Danke. Das Abendessen ist fertig. Möchtest du was essen?“")

	fmt.Scanln(&x)

	fmt.Println(y, " : „Gerne was gibt es denn?“")

	fmt.Scanln(&x)

	fmt.Println("Wirt: “Es gibt Nudeln mit Bolognessesoße und Käse, Katoffelbrei mit Hähnchenbrust und Buttergemüse, sowie Rinderrulade mit Rotkohl, Kartoffeln und Bratensoße.“ ")

	fmt.Println("Wirt: “Deine Endscheidung.“ ")

	fmt.Scanln(&x)

	for {
		fmt.Print("Bitte wähle 1 wenn du Nudeln mit Bolognessesoße und Käse, 2 wenn du Katoffelbrei mit Hähnchenbrust und Buttergemüse, 3 wenn du Rinderrulade mit Rotkohl, Kartoffeln und Bratensoße.")

		fmt.Println(" ")

		fmt.Scanln(&z)
		if z == 1 { //Option 1 => Nudeln

			fmt.Println(y, ": „Heute habe ich Lust auf die Nudeln. Danke“")

			fmt.Println("Wirt: „Bitte schön, eine Portion Nudeln mit Tomatensoße und Käse. Lass es dir schmecken.“  ")

			fmt.Println("Du suchst dir einen Platz in der Stube und setzt dich hin.   ")

			fmt.Scanln(&x)

			fmt.Println("Hungrieg machst du dich über die Nudeln her. Die Soße ist lecker, der Käse ist genau richtig und die Nudeln sind durch. ")

			fmt.Println("Schnell ist der Teller leer und du bringst ihn dem Wirt zurück. ")

			break

		} else if z == 2 { //Option 2 => Hähnchen

			fmt.Println(y, ": „Bitte den Kartoffelbrei mit Hähnchen.“")

			fmt.Println("Wirt: „Bitte schön Kartoffelbrei, Hähnchenbrust und Buttergemüsse, wie bestellt. Lass es dir schmecken.“  ")

			fmt.Println("Du suchst dir einen Platz in der Stube und setzt dich hin.  ")

			fmt.Scanln(&x)

			fmt.Println("Das Hänchen ist zart und der Kartoffelbrei ist kremig. Das Buttergemüse ist lecker.")

			fmt.Println("Schnell ist der Teller leer und du bringst ihn dem Wirt zurück. ")

			break

		} else if z == 3 { //Option 3 =>Rulade

			fmt.Println(y, ": „Ich neheme heute die Rinderrulade, bitte“")

			fmt.Println(" Wirt: „Bitte schön, deine Rinderrulade mit Rotkohl und Kartoffeln. Lass es dir schmecken.“  ")

			fmt.Println("Du suchst dir einen Platz in der Stube und setzt dich hin.   ")

			fmt.Scanln(&x)

			fmt.Println("Die Rulade ist genau Richtig gefüllt. Die Kartoffeln sind durch und der Rotkohl ist hervorragen. ")

			fmt.Println("Schnell ist der Teller leer und du bringst ihn dem Wirt zurück. ")

			break
		}
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Tag() { //Noch nicht ganz zum Schliessen--> Startbildschirm

	var x string //VAriable zum weiterkommen Text

	var j int //Variable für einen Tag
	var t int //Variable fürs weiter Spielen (Schleife)

	j = 1

	for { //Schleife Nach 365 durchgängen ist schluss
		j++
		if j > 364 { //Nach einem Jahr ist das Spiel vorbei
			fmt.Println("Ein Jahr lang bist du jetzt schon in Lurenz. Langsam wird es Zeit weiter zuziehen.")

			fmt.Println("Du geniesst dein leztes Essen und gibst den Schlüssel beim Margert ab und machst dich uf den Weg.")

			break
		} else {
			if j < 364 { //Ein normaler Tag in Lurenz

				for {

					fmt.Println("Wenn du weiter Spielen möchtest dann drücke 1. Hast du aber genug, dann drücke zum Beenden des Spiel 2")

					fmt.Scanln(&t)
					if t == 2 { //Das Spiel wird Beendet

						fmt.Println("Du hast dich für das Beenden des Spiels endschieden.")

						fmt.Println(" ")

						fmt.Println("Der Wirt, Marget, Jörg und die Endwickler wünschen dir noch einen schönen Tag")

						break

					} else if t == 1 { //Das Spiel geht weiter. Spieler spielt einen Tag

						Wetter()

						fmt.Scanln(&x)

						fmt.Println("Du machst dich fertig für den Tag.")

						fmt.Println("Zähneputzen, waschen,  Haare kämmen und packen. ")

						fmt.Scanln(&x)

						fmt.Println("Nach dem du damit fertig bist, verlässt du dein Zimmer und gehst in die Gaststube.  ")

						fmt.Scanln(&x)

						fmt.Println(y, ": „Guten Morgen“")

						fmt.Scanln(&x)

						Frühstück()

						fmt.Println(y, ": „Danke für das Essen. Haben Sie einen schönen Tag“ ")

						fmt.Scanln(&x)

						fmt.Println("Wirt: „Du auch,", y, "“ ")

						fmt.Println("Du machst dich auf den Weg zu Margrets Krämerladen. Die Bewohner von Lurenz wachen langsam auf und starten ihren Tag.")

						fmt.Scanln(&x)

						fmt.Println("Die Läden öffnen und Nachbarn grüßen einander. ")

						fmt.Println("Du betrittst Margrets Laden. ")

						fmt.Scanln(&x)

						fmt.Println(y, ": „Guten Morgen, Margret.“ ")

						fmt.Scanln(&x)

						fmt.Println("Margret: „Guten Morgen ", y, ". Hier ist der Schlüssel. Viel Erfolg im Dungen.“ ")

						fmt.Scanln(&x)

						fmt.Println(y, ": „Danke, Bis dann.“ ")

						fmt.Scanln(&x)

						fmt.Println("Du verlässt den Laden und folgst den dir bekannten Weg durch den Wald.  ")

						fmt.Println("Auch der Wald scheint langsam zu erwachen. ")

						fmt.Scanln(&x)

						fmt.Println("Am Dungen angekommen schleichst du die eisen Tür auf und trittst hinein. ")

						fmt.Println("Hinter dir schließt du die Tür wieder ab. ")

						fmt.Scanln(&x)

						fmt.Println("-Du bist im Dungen von Lurenz.- ")

						fmt.Scanln(&x)

						//Dungen Funktion

						fmt.Println("Mit deiner Ausbeute verlässt du den Dungen. Nach dem Verschließen der Tür folgst du dem Pfad zurück nach Florenz. ")

						fmt.Println("An Margrets Krämerladen angekommen betritst du den Laden. ")

						fmt.Println("Margret: „Du bist wieder da. Wie war´s?“ ")

						fmt.Scanln(&x)

						fmt.Println(y, ": Gut. Keine großen Probleme.“ ")

						fmt.Scanln(&x)

						fmt.Println("Margret: „Das ist gut. Zeig her deine Ausbeute.“ ")

						fmt.Scanln(&x)

						fmt.Println("-Du gibst Margret die Ru´s.- ")

						fmt.Scanln(&x)

						//Umtauschfunktion 													//Funktion zum TAuschen von Ru zu Bien -> exestiert noch nicht

						fmt.Println("Margret: „ Hier bitte deine Bien. Sonst noch was? ")

						fmt.Scanln(&x)

						//Handelmöglich														//Funktion zum Handeln -> exestiert noch nicht

						fmt.Println(y, ": „Nein Danke. Bis morgen ") //Satz da Funktion Handel nicht exsestiert

						fmt.Scanln(&x)

						fmt.Println("Margret: „Bis Morgen“")

						fmt.Println("Du verlässt den Laden und siehst, wie Ladenbesitzer ihre Geschäfte schließen.")

						fmt.Scanln(&x)

						fmt.Println("Du machst dich auf den Weg zurück „zum Krug“")

						fmt.Println("Als du diesen betritts, merkst du es ist genau so voll wie sonst auch.")

						fmt.Println("Du gehst zur Theke.")

						fmt.Scanln(&x)

						fmt.Println("Wirt: „Abend. Selbe wie sonst auch?“  ")

						fmt.Scanln(&x)

						fmt.Println("Spieler: „ Ja, Bitte.“  ")

						fmt.Scanln(&x)

						fmt.Println("Wirt: „20 Bien Bitte.“ ")

						fmt.Scanln(&x)

						fmt.Println("-Du gibst die 20 Bien dem Wirt- ")

						fmt.Scanln(&x)

						Abendessen()

						fmt.Scanln(&x)

						fmt.Println("Du hast dein Abendessen aufgegessen und schleppst dich auf dein Zimmer.")

						fmt.Scanln(&x)

						fmt.Println("Nachdem du dich Umgezogen und Bett fertig gemacht hast, fällst du wie ein Stein ins Bett. ")

						fmt.Println("Morgen ist ein neuer Tag. ")

						fmt.Println("------------------------------------------------------------------------------------------------ ")

					}
				}
			}

		}

	}
}
