package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

const DISK_PATH = "C:"

// Windows doesn't have the syscall that linux and mac have, so we will patch in the kernel32.dll
func getFilesystemStats(path string) (output filesystemStats) {
	if path == "" {
		path = DISK_PATH
	}

	kernel32, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		fmt.Println("Unable to load kernel32.dll:", err)
		syscall.Exit(1)
	}
	defer syscall.FreeLibrary(kernel32)
	GetDiskFreeSpaceEx, err := syscall.GetProcAddress(syscall.Handle(kernel32), "GetDiskFreeSpaceExW")
	if err != nil {
		fmt.Println("Unable to get Filesystem data", err)
		syscall.Exit(1)
	}

	freeBytesAvailable := int64(0)
	totalNumberOfBytes := int64(0)
	totalNumberOfFreeBytes := int64(0)

	syscall.Syscall6(uintptr(GetDiskFreeSpaceEx), 4,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(path))),
		uintptr(unsafe.Pointer(&freeBytesAvailable)),
		uintptr(unsafe.Pointer(&totalNumberOfBytes)),
		uintptr(unsafe.Pointer(&totalNumberOfFreeBytes)), 0, 0)

	output.Total = float64(totalNumberOfBytes / GB)
	output.Free = float64(totalNumberOfFreeBytes / GB)
	output.Percent = output.Free / output.Total * 100
	output.Filesystem = ""
	output.Hostname, _ = os.Hostname()

	return
}
