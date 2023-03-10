package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
    //map to store values later outputted
    mainMap := map[string]string{}

    //environment variables
    envVars := os.Environ()
    appendArray(envVars, mainMap, "=")

    //reading hostname

    //reading os-release
    osReleaseStr, err := os.ReadFile("/etc/os-release")
    if err != nil { log.Fatal(err) }
    osRelease := strings.Split(string(osReleaseStr), "\n")
    appendArray(osRelease, mainMap, "=")

    //reading meminfo
    meminfoStr, err := os.ReadFile("/proc/meminfo")
    if err != nil { log.Fatal(err) }
    meminfo := strings.Split(string(meminfoStr), "\n")
    for i := range meminfo {
	meminfo[i] = strings.TrimSuffix(meminfo[i], " kB")
    }
    appendArray(meminfo, mainMap, ":")

    fmt.Println(mainMap)
}

func appendArray(array []string, destinationMap map[string]string, splitter string) {
    for i := range array {
	key, value, _ := strings.Cut(array[i], splitter)
	destinationMap[strings.ToUpper(strings.TrimSpace(key))] = strings.TrimSpace(value)
    }
}
