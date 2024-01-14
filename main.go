package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"syscall"
)

const VERSION = "v1.0.3 2024-01-13"
const GB = 1024 * MB
const MB = 1024 * 1024
const TERMINAL_COLOR_RED = "\033[41m"
const TERMINAL_COLOR_GREEN = "\033[42m"
const TERMINAL_COLOR_RESET = "\033[0m"

var pathToCheckFreespace = ""
var JSON_OUTPUT = false
var ERROR_IF_UNDER_PERCENT = 0.0
var ERROR_IF_UNDER_GB = 0.0
var NO_COLOR = false

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
	if !fs.Passed {
		syscall.Exit(2)
	}
}

func parseCommandLineFlags() {
	showVersion := flag.Bool("version", false, VERSION)
	flag.BoolVar(&JSON_OUTPUT, "json", false, "switch to json output")
	flag.Float64Var(&ERROR_IF_UNDER_PERCENT, "percent", 10, "a number like 2.5 that will trigger an alert if free space is under 2.5 percent")
	flag.Float64Var(&ERROR_IF_UNDER_GB, "gb", 0, "a number like 1.2 that will trigger an alert if free space is less than 1.2GB")
	flag.StringVar(&pathToCheckFreespace, "path", "", "optionally sets the path to check for free space (ng /mnt/usbdrive or D: )")
	flag.BoolVar(&NO_COLOR, "nocolor", false, "do not apply terminal color")
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

		colorStart := ""
		colorEnd := ""
		if !NO_COLOR {
			if fs.Passed {
				colorStart = TERMINAL_COLOR_GREEN
			} else {
				colorStart = TERMINAL_COLOR_RED
			}
			colorEnd = TERMINAL_COLOR_RESET
		}
		fmt.Println(colorStart, fs.Status, colorEnd)
	}
}

func roundOneDecimal(input float64) float64 {
	return math.Round(input*10) / 10
}
