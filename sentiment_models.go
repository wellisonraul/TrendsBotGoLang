package main

type Document struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type sentimentRequest struct {
	EncodingType string   `json:"encodingType"`
	Document     Document `json:"document"`
}

func NewSentimentRequest(content string) *sentimentRequest {
	return &sentimentRequest{
		EncodingType: "UTF8",
		Document: Document{
			Type:    "PLAIN_TEXT",
			Content: content,
		},
	}
}

type sentimentResponse struct {
	DocumentSentiment struct {
		Magnitude float64 `json:"magnitude"`
		Score     float64 `json:"score"`
	} `json:"documentSentiment"`
	Language  string `json:"language"`
	Sentences []struct {
		Text struct {
			Content     string `json:"content"`
			BeginOffset int    `json:"beginOffset"`
		} `json:"text"`
		Sentiment struct {
			Magnitude float64 `json:"magnitude"`
			Score     float64 `json:"score"`
		} `json:"sentiment"`
	} `json:"sentences"`
}
