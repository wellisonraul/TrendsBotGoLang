package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dghubble/sling"
)

const (
	baseURL = "https://language.googleapis.com/v1/documents:analyzeSentiment"
)

type SentSvc interface {
	Score(corpus []string) (float64, float64)
}

type sentSvc struct {
	sling      *sling.Sling
	httpClient *http.Client
}

func (s *sentSvc) Score(corpus []string) (float64, float64) {
	content := strings.Join(corpus, ".")
	body := NewSentimentRequest(content)

	req, err := s.sling.BodyJSON(body).Request()
	if err != nil {
		panic(err.Error())
	}

	res, err := s.httpClient.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		panic(err.Error())
	}
	defer res.Body.Close()

	sentRes := sentimentResponse{}
	json.NewDecoder(res.Body).Decode(&sentRes)

	return sentRes.DocumentSentiment.Score,
		sentRes.DocumentSentiment.Magnitude
}

type params struct {
	Key string `url:"key,omitempty"`
}

func NewSentSvc() SentSvc {
	auth, ok := os.LookupEnv("NLP_API_KEY")
	if !ok {
		panic(errors.New("bad auth data for cloud nlp api"))
	}
	sling := sling.
		New().
		Set("Content-Type", "application/json; charset=utf-8").
		Post(baseURL).
		QueryStruct(&params{Key: auth})

	httpClient := http.Client{Timeout: time.Second * 10}
	return &sentSvc{
		sling:      sling,
		httpClient: &httpClient,
	}
}
