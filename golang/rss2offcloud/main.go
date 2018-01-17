package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const (
	DefaultIgnoreDuration = "96h"
	DefualtInterval       = "30m"
)

var (
	aConfig = flag.String("c", "config.json", "Config path.")
)

type Config struct {
	IgnoreDuration string         `json:"ignoreDuration"`
	Interval       string         `json:"interval"`
	Soruces        []SourceConfig `json:"sources"`
	Targets        []TargetConfig `json:"targets"`
}

func (c *Config) GetIgnoreDuration() (time.Duration, error) {
	if c.IgnoreDuration == "" {
		c.IgnoreDuration = DefaultIgnoreDuration
	}
	return time.ParseDuration(c.IgnoreDuration)
}

func (c *Config) GetInterval() (time.Duration, error) {
	if c.Interval == "" {
		c.Interval = DefualtInterval
	}
	return time.ParseDuration(c.Interval)
}

type SourceConfig struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type TargetConfig struct {
	Name   string            `json:"name"`
	Url    string            `json:"url"`
	Params map[string]string `json:"params"`
}

func main() {
	flag.Parse()
	cfg := readConfig()

	feedChan := make(chan Feed)
	rs := buildSources(cfg, feedChan)
	ts := buildTarget(cfg, feedChan)
	StartDaemon(rs, ts, feedChan)
	stop := make(chan struct{})
	<-stop
}

func StartDaemon(rs []*Receiver, ts []*Target, c chan Feed) {
	go func() {
		for _, r := range rs {
			go r.Start()
		}
	}()

	for {
		feed := <-c
		fmt.Println("Receive feed ", feed.Title)
		for _, t := range ts {
			go func(target *Target) {
				if err := target.Fire(feed); err != nil {
					fmt.Println("[Error] ", err.Error())
				}
			}(t)
		}
	}
}

func buildSources(cfg *Config, result chan Feed) []*Receiver {
	ret := make([]*Receiver, len(cfg.Soruces))
	it, err := cfg.GetInterval()
	if err != nil {
		panic(err.Error())
	}
	igd, err := cfg.GetIgnoreDuration()
	if err != nil {
		panic(err.Error())
	}
	for i, sc := range cfg.Soruces {
		ret[i] = NewReceiver(sc, it, igd, result)
	}
	return ret
}

func buildTarget(cfg *Config, result chan Feed) []*Target {
	ret := make([]*Target, len(cfg.Targets))
	for i, t := range cfg.Targets {
		ret[i] = &Target{Url: t.Url, Name: t.Name, Params: t.Params}
	}
	return ret
}

func readConfig() *Config {
	raw, err := ioutil.ReadFile(*aConfig)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var config Config
	if err := json.Unmarshal(raw, &config); err != nil {
		panic(err.Error())
	}
	return &config
}
