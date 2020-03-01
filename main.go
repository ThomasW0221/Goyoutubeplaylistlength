package main

import "fmt"

func main() {
	playlistLength, err := getPlaylistLength("https://www.youtube.com/playlist?list=PL0qTfdf9DoTgQDG61aOO90_bMUK0XOXMS")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(playlistLength)
}
