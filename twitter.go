package socfetch

import (
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"time"
)

type TwitterApi struct {
	AccessToken,
	AccessTokenSecret,
	ConsumerKey,
	ConsumerSecret string

	api *anaconda.TwitterApi
}

func NewTwitterApi(AccessToken, AccessTokenSecret, ConsumerKey,
	ConsumerSecret string) (*TwitterApi, error) {
	anaconda.SetConsumerKey(ConsumerKey)
	anaconda.SetConsumerSecret(ConsumerSecret)
	api := anaconda.NewTwitterApi(AccessToken, AccessTokenSecret)
	TwApi := &TwitterApi{api: api}
	return TwApi, nil
}

func (api *TwitterApi) Feed(id string) (rsp []Media) {
	twits, err := api.api.GetUserTimeline(url.Values{"user_id": []string{id}})
	if err != nil {
		return
	}
	for _, v := range twits {
		rsp = append(rsp, Media(TwitterMedia{Status: v}))
	}
	return
}

func (api *TwitterApi) Search(q string) (rsp []Media) {
	v := url.Values{}
	v.Set("count", "2")
	searchResult, err := api.api.GetSearch("смолгу", v)
	if err != nil {
		return
	}
	for _, tweet := range searchResult.Statuses {
		rsp = append(rsp, Media(TwitterMedia{Status: tweet}))
	}
	return
}

type TwitterStatus struct{}

type TwitterMedia struct {
	Status anaconda.Tweet
}

func (tw TwitterMedia) Text() string {
	return tw.Status.Text
}

func (tw TwitterMedia) Created() time.Time {
	t, _ := tw.Status.CreatedAtTime()
	return t
}

func (tw TwitterMedia) Type() string {
	return "status"
}
