package main

import (
	"flag"

	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	blue   = color.New(color.FgBlue).SprintFunc()
)

func start() {
	confirm()
	hc := NewHhchecker(fapiUrl, fUuid, fEname, fEnum)
	if err := hc.StartCheck(StartDate, EndDate); err != nil {
		panic(err)
	}
}

func confirm() {
	fmt.Printf("Use UUID:%s, Ename:%s, Enum:%s \n Checkin from %s to %s (y/n)\n", fUuid, fEname, fEnum, StartDate.Format("2006/01/02"), EndDate.Format("2006/01/02"))
	if !fDontAsk {
		getInput()
	}
}

func main() {
	flag.StringVar(&fUuid, "u", "", "")
	flag.StringVar(&fEname, "ei", "", "")
	flag.StringVar(&fEnum, "en", "", "")
	flag.StringVar(&fStartDate, "sd", "", "")
	flag.StringVar(&fEndDate, "ed", "", "")
	flag.StringVar(&fapiUrl, "api", "", "")
	flag.BoolVar(&fDontAsk, "y", false, "Donnot ask")
	flag.Parse()
	checkErr(setUuid())
	checkErr(setEname())
	checkErr(setEnum())
	checkErr(setStartDate())
	checkErr(setEndDate())
	start()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}
