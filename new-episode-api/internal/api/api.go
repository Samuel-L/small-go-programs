package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Samuel-L/new-episode-api/internal/helpers"
)

type episodes []struct {
	Name    string `json:"name"`
	Season  int    `json:"season"`
	Episode int    `json:"number"`
	AirDate string `json:"airdate"`
}

type request struct {
	ID int `json:"id"`
}

func NewEpisode(w http.ResponseWriter, r *http.Request) {
	var req request
	json.NewDecoder(r.Body).Decode(&req)

	url := helpers.ParseTvMazeUrl(req.ID)
	fetchedEpisodes, err := fetchEpisodes(url)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	var newReleases episodes
	for _, episode := range fetchedEpisodes {
		yesterday, err := helpers.IsYesterday(episode.AirDate)
		if err != nil {
			continue
		}
		if yesterday {
			newReleases = append(newReleases, episode)
		}
	}
	json.NewEncoder(w).Encode(newReleases)
}

func fetchEpisodes(url string) (episodes, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	request, err := http.NewRequest(
		"GET",
		fmt.Sprintf(url),
		nil,
	)
	if err != nil {
		return episodes{}, err
	}
	request.Header.Set("User-Agent", "newEpisodeChecker v1.0")

	response, err := client.Do(request)
	if err != nil {
		return episodes{}, err
	}
	defer response.Body.Close()

	body := episodes{}
	json.NewDecoder(response.Body).Decode(&body)

	return body, nil
}
