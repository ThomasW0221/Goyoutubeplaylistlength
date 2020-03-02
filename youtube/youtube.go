package youtube

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`([0-9]+):([0-9]{2})`)

func getPlaylistLength(playlistId string, strCh chan string, errCh chan error) {
	url := "https://www.youtube.com/playlist?list=" + playlistId
	resp, err := http.Get(url)
	if err != nil {
		errCh <- fmt.Errorf("%v,%v", playlistId, err)
		return
	}
	defer resp.Body.Close()

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		errCh <- fmt.Errorf("%v,%v", playlistId, err)
		return
	}

	titleString := strings.TrimSpace(document.Find(".pl-header-title").Text())

	lengthSelector := document.Find(".timestamp")
	totalLength := 0
	for _, node := range lengthSelector.Nodes {
		timeString := node.FirstChild.FirstChild.Data
		minSecond := re.FindStringSubmatch(timeString)
		minuteValue, minuteErr := strconv.Atoi(minSecond[1])
		if minuteErr != nil {
			errCh <- fmt.Errorf("%v,%v", titleString, err)
			return
		}

		secondValue, secondErr := strconv.Atoi(minSecond[2])
		if secondErr != nil {
			errCh <- fmt.Errorf("%v,%v", titleString, err)
			return
		}

		totalLength = totalLength + (60 * minuteValue) + secondValue
	}

	totalHours := totalLength / 3600
	totalMinutes := (totalLength % 3600) / 60
	totalSeconds := totalLength % 60
	strCh <- fmt.Sprintf("%v,%v:%v:%02v", titleString, totalHours, totalMinutes, totalSeconds)
}

func GetLengthOfMultiplePlaylists(playlistIds []string) []PlaylistResult{
	strCh := make(chan string)
	errCh := make(chan error)
	results := make([]PlaylistResult, 0)

	for _, playlistId := range playlistIds {
		go getPlaylistLength(playlistId, strCh, errCh)
	}

	i := 0
	for {
		select {
			case s := <- strCh:
				result := strings.Split(s, ",")
				results = append(results, PlaylistResult{Id:result[0], Result:result[1]})
				i++
				break
			case e := <- errCh:
				result := strings.Split(e.Error(), ",")
				results = append(results, PlaylistResult{Id:result[0], Result:result[1]})
				break
		}
		if i >= len(playlistIds){
			break
		}
	}
	fmt.Println(results)
	return results
}
