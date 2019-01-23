package main

import (
	"errors"
	"github.com/dghubble/sling"
	"log"
	"net/http"
)

type ModelApi struct {
	sling *sling.Sling
}

func NewModelApi(baseUrl string, client *http.Client) *ModelApi {
	return &ModelApi{
		sling: sling.New().Client(client).Base(baseUrl),
	}
}

func (m *ModelApi) ScoreImage(url string) (*GetScoreResponse, error) {
	req := &GetScoreRequest{Url: url}
	scoreResponse := &GetScoreResponse{}

	res, err := m.sling.New().Get("/score-image").QueryStruct(req).ReceiveSuccess(scoreResponse)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		log.Print("Error: Status Code: ", res.StatusCode)
		return nil, errors.New("Error in call to model")
	}

	return scoreResponse, nil
}

type GetScoreRequest struct {
	Url string `json:"url"`
}

type GetScoreResponse struct {
	Brand       string  `json:"brand"`
	Probability float32 `json:"probability"`
}
