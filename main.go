package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"syscall"
)

const VERSION = "v1.0.1 2022-03-18"
const GB = 1024 * MB
const MB = 1024 * 1024

var pathToCheckFreespace = ""
var JSON_OUTPUT = false
var ERROR_IF_UNDER_PERCENT = 0.0
var ERROR_IF_UNDER_GB = 0.0

type filesystemStats struct {
	Filesystem string  `json:"Filesystem"`
	Total      float64 `json:"TotalGb"`
	Free       float64 `json:"FreeGb"`
	Percent    float64 `json:"FreePercentage"`
	Hostname   string  `json:"Hostname"`
	Status     string  `json:"Status"`
	Passed     bool    `json:"Passed"`
	Path       string  `json:"Path"`
}

func main() {
	parseCommandLineFlags()
	fs := getFilesystemStats(pathToCheckFreespace)

	if ERROR_IF_UNDER_GB > 0 && fs.Free < ERROR_IF_UNDER_GB {
		fs.Passed = false
		fs.Status = fmt.Sprintf("[FAIL] Free disk space under %.1fGB", ERROR_IF_UNDER_GB)
	} else if ERROR_IF_UNDER_GB == 0 && fs.Percent < ERROR_IF_UNDER_PERCENT {
		fs.Passed = false
		fs.Status = fmt.Sprintf("[FAIL] Free disk space under %.1f%%", ERROR_IF_UNDER_PERCENT)
	} else {
		fs.Passed = true
		fs.Status = "[PASS] Disk OK"
	}

	displayOutput(fs)
	if fs.Passed == false {
		syscall.Exit(2)
	}
}

func parseCommandLineFlags() {
	showVersion := flag.Bool("version", false, VERSION)
	flag.BoolVar(&JSON_OUTPUT, "json", false, "switch to json output")
	flag.Float64Var(&ERROR_IF_UNDER_PERCENT, "percent", 10, "a number like 2.5 that will trigger an alert if free space is under 2.5 percent")
	flag.Float64Var(&ERROR_IF_UNDER_GB, "gb", 0, "a number like 1.2 that will trigger an alert if free space is less than 1.2GB")
	flag.StringVar(&pathToCheckFreespace, "path", "", "optionally sets the path to check for free space (ng /mnt/usbdrive or D: )")
	flag.Parse()
	if *showVersion {
		fmt.Println(VERSION)
		syscall.Exit(0)
	}
}

func displayOutput(fs filesystemStats) {
	if JSON_OUTPUT {
		jsonBytes, err := json.Marshal(fs)
		if err != nil {
			fmt.Printf("{'error':'%s'}\n", err)
		}
		fmt.Println(string(jsonBytes))
	} else {
		fmt.Println(fs.Path, "on", fs.Hostname)
		fmt.Println(" Free:   ", fs.Free, "GB /", fs.Total, "GB")
		fmt.Printf(" Percent: %.2f%%\n", fs.Percent)
		fmt.Println(fs.Status)
	}
}

func convertToString(input [16]int8) (output string) {
	for i := range input {
		if input[i] == 0 {
			return
		}
		output += string(byte(input[i]))
	}
	return
}

func roundOneDecimal(input float64) float64 {
	return math.Round(input*10) / 10
}
