package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`(?P<minutes>[0-9]+):(?P<seconds>[0-9]{2})`)

func getPlaylistLength(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	lengthSelector := document.Find(".timestamp")
	totalLength := 0
	for _, node := range lengthSelector.Nodes {
		timeString := node.FirstChild.FirstChild.Data
		minSecond := re.FindStringSubmatch(timeString)
		minuteValue, minuteErr := strconv.Atoi(minSecond[1])
		if minuteErr != nil {
			return "", minuteErr
		}

		secondValue, secondErr := strconv.Atoi(minSecond[2])
		if secondErr != nil {
			return "", secondErr
		}

		totalLength = totalLength + (60 * minuteValue) + secondValue
		fmt.Println(totalLength)
	}
	return "hello", nil
}