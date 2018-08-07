package main

import (
	"errors"
	"time"
)

var (
	fDontAsk   bool
	fUuid      string
	fEnum      string
	fEname     string
	fStartDate string
	fEndDate   string
	fapiUrl    string
	StartDate  time.Time
	EndDate    time.Time
)

func setUuid() error {
	if fUuid == "" {
		return errors.New("missing uuid")
	}
	return nil
}

func setEnum() error {
	if fEnum == "" {
		return errors.New("missing em number")
	}
	return nil
}

func setEname() error {
	if fEname == "" {
		return errors.New("missing em name")
	}
	return nil
}

func setStartDate() error {
	if fStartDate == "" {
		return errors.New("missing start date")
	}
	var err error
	StartDate, err = time.Parse(datelayout, fStartDate)
	if err != nil {
		return err
	}
	return nil
}
func setEndDate() error {
	if fEndDate == "" {
		return errors.New("missing end date")
	}
	var err error
	EndDate, err = time.Parse(datelayout, fEndDate)
	if err != nil {
		return err
	}
	return nil
}
