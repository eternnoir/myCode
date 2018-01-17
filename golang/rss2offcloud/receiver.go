package main

import (
	"fmt"
	"time"

	"github.com/ungerik/go-rss"
)

func NewReceiver(cfg SourceConfig, it, igd time.Duration, rc chan Feed) *Receiver {
	return &Receiver{Name: cfg.Name, Url: cfg.Url, Interval: it, IgnoreDuration: igd, resultChan: rc}
}

type Receiver struct {
	Name           string
	Url            string
	Interval       time.Duration
	IgnoreDuration time.Duration
	resultChan     chan Feed
}

func (r Receiver) Start() {
	for {
		feeds, err := r.Fetch()
		if err != nil {
			fmt.Println("Fetch fail.", err.Error())
		}

		for _, f := range feeds {
			r.resultChan <- f
		}
		fmt.Println("Done. Sleep ", r.Interval)
		time.Sleep(r.Interval)
	}
}

func (r Receiver) Fetch() ([]Feed, error) {
	channel, err := rss.Read(r.Url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(channel.Title)

	ret := make([]Feed, 0)

	for _, item := range channel.Item {
		pubDate, err := item.PubDate.Parse()
		if err != nil {
			fmt.Println("Parse PubDate fail.", err.Error())
			continue
		}
		if pubDate.Add(r.IgnoreDuration).Before(time.Now()) {
			continue
		}
		fmt.Println(r.Name, " fetched ", item.Title)
		for _, enrouse := range item.Enclosure {
			feed := Feed{Link: enrouse.URL, Title: item.Title}
			ret = append(ret, feed)
		}
	}
	return ret, nil
}