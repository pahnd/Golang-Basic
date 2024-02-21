package main

import "fmt"

// #1: Create struct type
type song struct {
	title, artist string
}

// #5: Structs can be contain other structs
type playlist struct {
	genre string

	// #6: Inculde a slice of song structs
	songs []song
}

func main() {
	song1 := song{title: "wonderwall", artist: "oasis"}
	song2 := song{title: "super sonic", artist: "oasis"}

	/// #4: Structs are copied
	// song1 = song2
	// #3: Structs can be compared
	if song1 == song2 {
		// #2: Struct comparison works like this
		// if song1.title == song2.title &&
		//		song1.artist == song2.artist {
		fmt.Println("song are equal.")
	} else {
		fmt.Println("song are not equal")
	}

	songs := []song{
		{title: "wonderfall", artist: "oasis"},
		{title: "supersonic", artist: "oasis"},
	}
	// #7 a struct can be include another struct
	rock := playlist{
		genre: "indie rock",
		songs: songs,
	}
	// #9: you can't compare struct values that contains incomparable fields
	// you need to compare them manually

	// clone := rock
	// func comparePlaylists(p1, p2 playlist) bool {
	// 	if p1.genre != p2.genre {
	// 		return false
	// 	}

	// 	if len(p1.songs) != len(p2.songs) {
	// 		return false
	// 	}

	// 	for i := range p1.songs {
	// 		if p1.songs[i] != p2.songs[i] {
	// 			return false
	// 		}
	// 	}

	// 	return true
	// }

	// #11: song is a clone, it cannot change the original struct value
	song := rock.songs[0]
	song.title = "live forever"
	fmt.Printf("\n%+v\n%+v\n", song, rock.songs[0])
	// #11c: directly set the original one
	rock.songs[0].title = "live forever"
	// #11b
	fmt.Printf("\n%+v\n%+v\n", song, rock.songs[0])

	// #10: printing
	fmt.Printf("\n%-20s %20s\n", "TITLE", "ARTIST")

	for _, s := range rock.songs {
		// s := rock.songs[i]

		// #12b: s is a copy inside scope because struct values are copied
		s.title = "destroy"
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}

	// #12
	fmt.Printf("\n%-20s %20s\n", "TITLE", "ARTIST")
	for _, s := range rock.songs {
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}
}
