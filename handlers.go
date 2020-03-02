package main

import (
	"encoding/json"
	"github.com/ThomasW0221/Goyoutubeplaylistlength/youtube"
	"net/http"
	"strings"
)

func handleGetPlaylistLengths(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// following call returns the list as one big string which requires further splitting
		playlistIds, ok := r.URL.Query()["playlistIds"]
		if !ok || len(playlistIds) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// if request param contains , (i.e. more than 1 playlist is sent in)
		// split it into pieces, otherwise it can be worked with as is
		if strings.Contains(playlistIds[0], ",") {
			playlistIds = strings.Split(playlistIds[0], ",")
		}

		results := youtube.GetLengthOfMultiplePlaylists(playlistIds)
		resultsJson, err := json.Marshal(results)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resultsJson)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	return
}