package main

import "fmt"

func main() {
	playlists := []string{"PL0qTfdf9DoTgQDG61aOO90_bMUK0XOXMS",
		"PLjpsoptsN4KCUuLuQ56Uy923hQhjPYlG_",
		"PLMxPYcr2zEkWInMyvvxmN22gjrRR0x__u"}
	results := getLengthOfMultiplePlaylists(playlists)
	for k, v := range results {
		fmt.Printf("%v: %v\n", k, v)
	}

}
