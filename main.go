package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"math"
)

//map to store values later outputted
var mainMap map[string]string
    
func main() {
    mainMap = make(map[string]string)

    //environment variables
    envVars := os.Environ()
    appendArray(envVars, mainMap, "=")

    //reading os-release
    osrelease := read("/etc/os-release")
    for i := range osrelease {
	osrelease[i] = strings.ReplaceAll(osrelease[i], "\"", "")
    }
    appendArray(osrelease, mainMap, "=")

    //reading meminfo
    meminfo := read("/proc/meminfo")
    appendArray(meminfo, mainMap, ":")

    //reading hostname
    hostname, err := os.ReadFile("/etc/hostname")
    if err != nil { log.Fatal(err) }
    mainMap["HOSTNAME"] = strings.TrimSpace(string(hostname)) 

    //reading kernel version
    kernelVersion, err := os.ReadFile("/proc/sys/kernel/osrelease")
    if err != nil { log.Fatal(err) }
    mainMap["KERNEL_VERSION"] = strings.TrimSpace(string(kernelVersion))

    //reading uptime
    uptimeStr, err := os.ReadFile("/proc/uptime")
    if err != nil { log.Fatal(err) }
    uptimeArr := strings.Split(string(uptimeStr), " ")
    uptimeSecsFloat, err := strconv.ParseFloat(uptimeArr[0], 64)
    if err != nil { log.Fatal(err) }
    uptimeSecs := int(uptimeSecsFloat)
    uptimeHrs := uptimeSecs / 3600
    uptimeSecs -= uptimeHrs * 3600
    uptimeMins := uptimeSecs / 60
    uptimeSecs -= uptimeMins * 60
    mainMap["UPTIME_HRS"] = strconv.Itoa(uptimeHrs)
    mainMap["UPTIME_MINS"] = strconv.Itoa(uptimeMins)
    mainMap["UPTIME_SECS"] = strconv.Itoa(uptimeSecs)

    //reading configs
    var result []string 
    var art []string 
    if _, err := os.Stat(mainMap["HOME"] + "/.config/yass/"); !os.IsNotExist(err) {
	result = parseConfig(mainMap["HOME"] + "/.config/yass/config", false)
	art = parseConfig(mainMap["HOME"] + "/.config/yass/art", true)
    } else {
	result = parseConfig("./config/config", false)
	art = parseConfig("./config/art", true)
    }
    
    //printing results
    artSize := len(art) - 1
    resultSize := len(result) - 1
    sizeDifference := math.Abs(float64(artSize - resultSize) / 2)
    if artSize > resultSize {
	for i := 0; i < artSize; i++ {
	    if i >= int(math.Floor(sizeDifference)) && i <= resultSize + int(math.Floor(sizeDifference)) {
                fmt.Println(art[i], result[i - int(math.Floor(sizeDifference))])
	    } else {
		fmt.Println(art[i])
	    }
	}
    } else {
	for i := 0; i < resultSize; i++ {
	    if i >= int(math.Floor(sizeDifference)) && i <= artSize + int(math.Floor(sizeDifference)) {
                fmt.Println(art[i - int(math.Floor(sizeDifference))], result[i])
	    } else {
		fmt.Println(result[i])
	    }
	}
    }
}

func appendArray(array []string, destinationMap map[string]string, splitter string) {
    for i := range array {
	key, value, _ := strings.Cut(array[i], splitter)
	value = strings.TrimSpace(value)
	if strings.Contains(value, "kB") {
            value = strings.ReplaceAll(value, " kB", "")
	    valueNum, err := strconv.Atoi(value)
	    if err != nil { log.Fatal(err) }
	    valueMb := valueNum / 1024
	    valueGb := valueMb / 1024
	    destinationMap[strings.ToUpper(strings.TrimSpace(key))+"_MB"] = strconv.Itoa(valueMb)
	    destinationMap[strings.ToUpper(strings.TrimSpace(key))+"_GB"] = strconv.Itoa(valueGb)
	}
	destinationMap[strings.ToUpper(strings.TrimSpace(key))] = value
    }
}

func read(path string) []string {
    resultStr, err := os.ReadFile(path)
    if err != nil { log.Fatal(err) }
    result := strings.Split(string(resultStr), "\n")
    return result
}

func parseConfig(pathToConfig string, prependDistrocolor bool) []string {
    configStr, err := os.ReadFile(pathToConfig)
    if err != nil { log.Fatal(err) }
    config := strings.Split(string(configStr), "\n")

    //parsing values
    for i := range config {
	line := strings.Split(config[i], "|")
	for j := range line {
	    val, ok := mainMap[line[j]]
	    if ok { line[j] = val }
	}
	line = append(line, "<>")
	if prependDistrocolor { line = append([]string{"<distrocolor>"}, line...) }
	config[i] = strings.Join(line, "")
    }

    //parsing basic styling
    for i := range config {
	//clear styles
	config[i] = strings.ReplaceAll(config[i], "<c>", "\x1B[0m")
	config[i] = strings.ReplaceAll(config[i], "<>", "\x1B[0m")
        //decorations
	config[i] = strings.ReplaceAll(config[i], "<b>", "\x1B[1m")
	config[i] = strings.ReplaceAll(config[i], "<d>", "\x1B[2m")
	config[i] = strings.ReplaceAll(config[i], "<i>", "\x1B[3m")
	config[i] = strings.ReplaceAll(config[i], "<u>", "\x1B[4m")
	config[i] = strings.ReplaceAll(config[i], "<uu>", "\x1B[21m")
	config[i] = strings.ReplaceAll(config[i], "<r>", "\x1B[7m")
	config[i] = strings.ReplaceAll(config[i], "<s>", "\x1B[9m")
        //colors
	config[i] = strings.ReplaceAll(config[i], "<black>", "\x1B[30m")
	config[i] = strings.ReplaceAll(config[i], "<red>", "\x1B[31m")
	config[i] = strings.ReplaceAll(config[i], "<green>", "\x1B[32m")
	config[i] = strings.ReplaceAll(config[i], "<yellow>", "\x1B[33m")
	config[i] = strings.ReplaceAll(config[i], "<blue>", "\x1B[34m")
	config[i] = strings.ReplaceAll(config[i], "<magenta>", "\x1B[35m")
	config[i] = strings.ReplaceAll(config[i], "<cyan>", "\x1B[36m")
	config[i] = strings.ReplaceAll(config[i], "<white>", "\x1B[37m")
	//bright colors
	config[i] = strings.ReplaceAll(config[i], "<brblack>", "\x1B[30m;1m")
	config[i] = strings.ReplaceAll(config[i], "<brred>", "\x1B[31m;1m")
	config[i] = strings.ReplaceAll(config[i], "<brgreen>", "\x1B[32m;1m")
	config[i] = strings.ReplaceAll(config[i], "<bryellow>", "\x1B[33m;1m")
	config[i] = strings.ReplaceAll(config[i], "<brblue>", "\x1B[34m;1m")
	config[i] = strings.ReplaceAll(config[i], "<brmagenta>", "\x1B[35m;1m")
	config[i] = strings.ReplaceAll(config[i], "<brcyan>", "\x1B[36m;1m")
	config[i] = strings.ReplaceAll(config[i], "<brwhite>", "\x1B[37m;1m")
	//background colors
	config[i] = strings.ReplaceAll(config[i], "<bgblack>", "\x1B[40m")
	config[i] = strings.ReplaceAll(config[i], "<bgred>", "\x1B[41m")
	config[i] = strings.ReplaceAll(config[i], "<bggreen>", "\x1B[42m")
	config[i] = strings.ReplaceAll(config[i], "<bgyellow>", "\x1B[43m")
	config[i] = strings.ReplaceAll(config[i], "<bgblue>", "\x1B[44m")
	config[i] = strings.ReplaceAll(config[i], "<bgmagenta>", "\x1B[45m")
	config[i] = strings.ReplaceAll(config[i], "<bgcyan>", "\x1B[46m")
	config[i] = strings.ReplaceAll(config[i], "<bgwhite>", "\x1B[47m")
	//background bright colors
	config[i] = strings.ReplaceAll(config[i], "<bgbrblack>", "\x1B[40m;1m")
	config[i] = strings.ReplaceAll(config[i], "<bgbrred>", "\x1B[41m;1m")
	config[i] = strings.ReplaceAll(config[i], "<bgbrgreen>", "\x1B[42m;1m")
	config[i] = strings.ReplaceAll(config[i], "<bgbryellow>", "\x1B[43m;1m")
	config[i] = strings.ReplaceAll(config[i], "<bgbrblue>", "\x1B[44m;1m")
	config[i] = strings.ReplaceAll(config[i], "<bgbrmagenta>", "\x1B[45m;1m")
	config[i] = strings.ReplaceAll(config[i], "<bgbrcyan>", "\x1B[46m;1m")
	config[i] = strings.ReplaceAll(config[i], "<bgbrwhite>", "\x1B[47m;1m")
	//distro color
	switch mainMap["NAME"] {
	case "Debian", "Ubuntu":
	    config[i] = strings.ReplaceAll(config[i], "<distrocolor>", "\x1B[31m") //red
	case "Linux Mint", "OpenSUSE", "Void":
	    config[i] = strings.ReplaceAll(config[i], "<distrocolor>", "\x1B[32m") //green
	case "Fedora Linux", "Slackware":
	    config[i] = strings.ReplaceAll(config[i], "<distrocolor>", "\x1B[34m") //blue
	case "Gentoo":
	    config[i] = strings.ReplaceAll(config[i], "<distrocolor>", "\x1B[35m") //magenta
	case "Arch Linux":
	    config[i] = strings.ReplaceAll(config[i], "<distrocolor>", "\x1B[36m") //cyan
	default:
	    config[i] = strings.ReplaceAll(config[i], "<distrocolor>", "\x1B[37m") //white
	}
    }
    
    return config
}
