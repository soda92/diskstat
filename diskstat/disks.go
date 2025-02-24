package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/windows"
)

var KB = float64(1024)
var GB = KB * KB * KB

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func get_disks() []string {
	x := []string{}
	for char := 'A'; char <= 'Z'; char++ {
		path := fmt.Sprintf("%c%s", char, ":\\")
		if pathExists(path) {
			x = append(x, path)
		}
	}
	return x
}

func get_disk_name(v string) string {
	volumePath := v // Replace with the desired volume path
	var volumeNameBuffer [windows.MAX_PATH + 1]uint16
	var fileSystemNameBuffer [windows.MAX_PATH + 1]uint16
	var volumeSerialNumber uint32
	var maxComponentLength uint32
	var fileSystemFlags uint32

	err := windows.GetVolumeInformation(
		windows.StringToUTF16Ptr(volumePath),
		&volumeNameBuffer[0],
		uint32(len(volumeNameBuffer)),
		&volumeSerialNumber,
		&maxComponentLength,
		&fileSystemFlags,
		&fileSystemNameBuffer[0],
		uint32(len(fileSystemNameBuffer)),
	)

	if err != nil {
		log.Fatalf("Error getting volume information: %v", err)
	}

	volumeName := windows.UTF16ToString(volumeNameBuffer[:])
	return fmt.Sprintf("%s (%c:)", volumeName, v[0])
}
