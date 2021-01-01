package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "os"
    "syscall"
)

const VERSION = "v0.1 2020-DEC-29"
const GB = 1024 * 1024 * 1024

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
}

func main() {
    parseCommandLineFlags()
    fs := getFilesystemStats(ERROR_IF_UNDER_PERCENT)
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
    flag.Parse()
    if *showVersion {
        fmt.Println(VERSION)
        syscall.Exit(0)
    }
}

func getFilesystemStats(errorIfUnderPercent float64) (output filesystemStats) {
    // this only works on linux / mac
    syscallResult := syscall.Statfs_t{}
    err := syscall.Statfs("/", &syscallResult)
    if err != nil {
        fmt.Println("Unable to get Filesystem data", err)
        syscall.Exit(1)
    }
    output.Total = float64(syscallResult.Blocks * uint64(syscallResult.Bsize) / GB)
    output.Free = float64(syscallResult.Bavail * uint64(syscallResult.Bsize) / GB)
    output.Percent = output.Free / output.Total * 100
    //output.Filesystem = convertToString(syscallResult.Fstypename)
    output.Hostname, _ = os.Hostname()

    if ERROR_IF_UNDER_GB > 0 && output.Free < ERROR_IF_UNDER_GB {
        output.Passed = false
		output.Status = fmt.Sprintf("[FAIL] Free disk space under %.1fGB", ERROR_IF_UNDER_GB)
	}else if ERROR_IF_UNDER_GB == 0 && output.Percent < errorIfUnderPercent {
        output.Passed = false
        output.Status = fmt.Sprintf("[FAIL] Free disk space under %.1f%%", errorIfUnderPercent)
    }else {
        output.Passed = true
        output.Status = "[PASS] Disk OK"
    }
    return
}

func displayOutput(fs filesystemStats) {
    if JSON_OUTPUT {
        jsonBytes, err := json.Marshal(fs)
        if err != nil {
            fmt.Printf("{'error':'%s'}\n", err)
        }
        fmt.Println(string(jsonBytes))
    } else {
        fmt.Println("/ on", fs.Hostname)
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