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

    //reading os-release
    osRelease := read("/etc/os-release")
    appendArray(osRelease, mainMap, "=")

    //reading meminfo
    meminfo := read("/proc/meminfo")
    for i := range meminfo {
	meminfo[i] = strings.TrimSuffix(meminfo[i], " kB")
    }
    appendArray(meminfo, mainMap, ":")

    //reading hostname
    hostname, err := os.ReadFile("/etc/hostname")
    if err != nil { log.Fatal(err) }
    mainMap["HOSTNAME"] = string(hostname)

    //reading kernel version
    kernelVersion, err := os.ReadFile("/proc/sys/kernel/osrelease")
    if err != nil { log.Fatal(err) }
    mainMap["KERNEL_VERSION"] = string(kernelVersion)


    fmt.Println(mainMap)
}

func appendArray(array []string, destinationMap map[string]string, splitter string) {
    for i := range array {
	key, value, _ := strings.Cut(array[i], splitter)
	destinationMap[strings.ToUpper(strings.TrimSpace(key))] = strings.TrimSpace(value)
    }
}

func read(path string) []string {
    resultStr, err := os.ReadFile(path)
    if err != nil { log.Fatal(err) }
    result := strings.Split(string(resultStr), "\n")
    return result
}
