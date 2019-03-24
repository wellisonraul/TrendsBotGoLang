package main

import (
	"net/http"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// TwitterTrendsSvc accesses twitter
type TwitterTrendsSvc interface {
	Trends(woeid int) ([]twitter.Trend, error)
	Close()
	TweetsFor(string) ([]string, error)
	TwitterOfAProfile(string) ([]string, error)
}

type twitterTrends struct {
	client *twitter.Client
}

func (tt *twitterTrends) Close() {}

func (tt *twitterTrends) Trends(woeid int) ([]twitter.Trend, error) {
	ts, res, err := tt.client.Trends.Place(int64(woeid), nil)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, err
	}
	var trends []twitter.Trend
	for _, xyz := range ts {
		for _, trend := range xyz.Trends {
			trends = append(trends, trend)
		}
	}
	return trends, nil
}

func (tt *twitterTrends) TweetsFor(query string) ([]string, error) {
	var tweets []string
	search, _, err := tt.client.Search.Tweets(&twitter.SearchTweetParams{
		Query: query,
		Count: 100,
		Lang:  "pt",
	})
	if err != nil {
		return nil, err
	}
	for _, tweet := range search.Statuses {
		tweets = append(tweets, tweet.Text)
	}
	return tweets, nil
}


func (tt *twitterTrends) TwitterOfAProfile(query string) ([]string, error) {
	var tweets []string
	// Home Timeline
	search, resp, err := tt.client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: query,
		Count: 100,
	})
	if err != nil {
		return nil, err
	}
	for _, tweet := range search {
		tweets = append(tweets, tweet.Text)
	}
	return tweets, nil
	print(resp, err)

	return tweets, nil
}

// NewTwitterTrendsSvc creates a new TwitterTrendsSvc
func NewTwitterTrendsSvc() TwitterTrendsSvc {
	return &twitterTrends{
		client: newClient(),
	}
}

func newClient() *twitter.Client {
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_SECRET")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	client := twitter.NewClient(config.Client(oauth1.NoContext, token))

	// Tests connection and crashes app if there is bad auth data
	_, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})
	if err != nil {
		panic(err.Error())
	}

	return client
}
