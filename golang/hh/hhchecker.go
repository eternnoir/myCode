package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var (
	datelayout      = "2006-01-02"
	checkinLayout   = "2006/01/02 15:04:05"
	monkeyBreakTime = time.Duration(3 * time.Second)
)

func NewHhchecker(url, uuid, ename, enum string) *Hhchecker {
	return &Hhchecker{
		apiUrl: url,
		uuid:   uuid,
		ename:  ename,
		enum:   enum,
	}
}

type Hhchecker struct {
	apiUrl string
	uuid   string
	ename  string
	enum   string
}

func (h *Hhchecker) StartCheck(start, end time.Time) error {
	for d := start; d.Before(end.AddDate(0, 0, 1)); d = d.AddDate(0, 0, 1) {
		fmt.Printf("Start check %s\n", blue(d.Format(datelayout)))
		dateType, err := h.getWorkDayType(d)
		if err != nil {
			return err
		}
		fmt.Printf("%s datetype is %s\n", d.Format(datelayout), dateType.DateType)
		if dateType.DateType != "1" {
			fmt.Printf("Song!! %s dont need to work. Go home. Zzzzzz\n", d.Format(datelayout))
			continue
		}
		fmt.Printf("Sad %s need to work. Start to checkin\n", d.Format(datelayout))

		lateDiff := rand.Intn(55)
		overworkDiff := rand.Intn(44)
		checkInTime := d.Add(9 * time.Hour).Add(time.Duration(lateDiff) * time.Minute)
		fmt.Printf("Checkin at %s \n", checkInTime.Format(checkinLayout))
		fmt.Println("Monkey need break.")
		time.Sleep(monkeyBreakTime)

		if err := h.checkIn(checkInTime); err != nil {
			panic(err)
		}

		checkoutTime := checkInTime.Add(9 * time.Hour).Add(time.Duration(overworkDiff) * time.Minute)
		fmt.Printf("Checkout at %s \n", checkoutTime.Format(checkinLayout))
		if err := h.checkOut(checkoutTime); err != nil {
			panic(err)
		}

		fmt.Println("Monkey need break.")
		time.Sleep(monkeyBreakTime)

		fmt.Printf("\n\nCheck in %s, checkout %s. Monkey Good Job\n\n", yellow(checkInTime.Format(checkinLayout)),
			yellow(checkoutTime.Format(checkinLayout)))
	}
	return nil
}

func (h *Hhchecker) checkIn(time time.Time) error {
	return h.insertWorkLog(time, "1", "1", "1")
}

func (h *Hhchecker) checkOut(time time.Time) error {
	return h.insertWorkLog(time, "1", "4", "2")
}

func (h *Hhchecker) insertWorkLog(time time.Time, dt, wi, wt string) error {
	reqPayload := &WorkTimeLog{
		UserName: h.ename,
		WorkTimeLogData: []WorkTimeData{
			{
				DateType:   dt,
				LeaveHours: 0,
				Memo:       "",
				WorkItem:   wi,
				WorkTime:   time.Format(checkinLayout),
				WorkType:   wt,
			},
		},
	}
	jsonValue, _ := json.Marshal(reqPayload)
	client := &http.Client{}
	requestUrl := fmt.Sprintf("%s/worktime/InsertTimeLog", h.apiUrl)
	req, _ := http.NewRequest("POST", requestUrl, bytes.NewBuffer(jsonValue))
	req.Header.Set("X-UUID", h.uuid)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("request fail. %s", string(bodyBytes))
	}

	ret := &WorkTimeLogResp{}
	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		return err
	}
	if !ret.Status && strings.Contains(ret.ErrorMessage, "已有請假紀錄，不需打卡!") {
		fmt.Printf("Already taken a leave. %s\n", ret.ErrorMessage)
		return nil
	}
	if !ret.Status {
		return fmt.Errorf("checkin fail %s", ret)
	}
	return nil
}

func (h *Hhchecker) getWorkDayType(date time.Time) (*DateTypeResp, error) {
	client := &http.Client{}
	requestUrl := fmt.Sprintf("%s/worktime/GetDateType/?empID=%s&date=%s", h.apiUrl, h.enum, date.Format(datelayout))
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("X-UUID", h.uuid)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("request fail. %s", string(bodyBytes))
	}
	ret := &DateTypeResp{}
	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		return nil, err
	}
	if !ret.Status {
		return nil, errors.New(fmt.Sprintf("%s", ret))
	}
	return ret, nil

}

type DateTypeResp struct {
	DateType     string `json:"DateType"`
	Status       bool   `json:"Status"`
	ErrorMessage string `json:"ErrorMessage"`
}

type WorkTimeLog struct {
	UserName        string         `json:"UserName"`
	WorkTimeLogData []WorkTimeData `json:"WorkTimeLogData"`
}

type WorkTimeData struct {
	DateType   string `json:"DateType"`
	LeaveHours int    `json:"LeaveHours"`
	Memo       string `json:"Memo"`
	WorkItem   string `json:"WorkItem"`
	WorkTime   string `json:"WorkTime"`
	WorkType   string `json:"WorkType"`
}

type WorkTimeLogResp struct {
	HolidayList  interface{} `json:"HolidayList"`
	Status       bool        `json:"Status"`
	ErrorMessage string      `json:"ErrorMessage"`
}
