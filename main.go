package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	oldVersion := "4.9.1"
	newVersion := "4.10.0"

	isNewVersion := NewVersionIsHigher(oldVersion, newVersion)

	fmt.Println("The latest version is higher: ", isNewVersion)
}

func splitVersion(versionString string) (int, int, int) {
	var major, minor, patch int

	minVersion := strings.Split(versionString, ".")

	if len(minVersion) > 2 {
		major, _ = strconv.Atoi(minVersion[0])
		minor, _ = strconv.Atoi(minVersion[1])
		patch, _ = strconv.Atoi(minVersion[2])
	}

	return major, minor, patch
}

func NewVersionIsHigher(oldVersion, newVersion string) bool {
	major, minor, patch := splitVersion(oldVersion)

	majorNew, minorNew, patchNew := splitVersion(newVersion)

	if majorNew > major {
		return true
	} else if majorNew == major {
		if minorNew > minor {
			return true
		} else if minorNew == minor {
			if patchNew > patch {
				return true
			} else if patchNew == patch {
				return false
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}
