// +build !windows

package main

import (
	"fmt"
	"os"
	"syscall"
)

const DISK_PATH = "/"

func getFilesystemStats() (output filesystemStats) {

	// this only works on linux / mac
	syscallResult := syscall.Statfs_t{}
	err := syscall.Statfs(DISK_PATH, &syscallResult)
	if err != nil {
		fmt.Println("Unable to get Filesystem data", err)
		syscall.Exit(1)
	}
	output.Total = float64(syscallResult.Blocks * uint64(syscallResult.Bsize) / GB)
	output.Free = float64(syscallResult.Bavail * uint64(syscallResult.Bsize) / GB)
	output.Percent = output.Free / output.Total * 100
	output.Filesystem = "" //convertToString(syscallResult.Fstypename) // mac can do this
	output.Hostname, _ = os.Hostname()

	return
}
