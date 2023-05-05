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
    var config = []string{" <distrocolor><d><b><u>|USER|<>@<distrocolor><d><b><u>|HOSTNAME|",
	      " <distrocolor><b>os<>     |PRETTY_NAME|",
	      " <distrocolor><b>kernel<> |KERNEL_VERSION|",
	      " <distrocolor><b>memory<> |MEMFREE_MB| / |MEMTOTAL_MB| MiB",
	      " <distrocolor><b>swap<>   |SWAPFREE_MB| / |SWAPTOTAL_MB| MiB",
	      " <distrocolor><b>uptime<> |UPTIME_HRS|:|UPTIME_MINS|:|UPTIME_SECS|", ""}
    var art = []string{"<b>\\   ", "<b>\\\\  ", "<b> \\\\ ", "<b>  \\\\", "<b>  //", "<b> // ", "<b>//  ", "<b>/   ", ""}

    if _, err := os.Stat(mainMap["HOME"] + "/.config/yass/"); !os.IsNotExist(err) {
        config = getConfig(mainMap["HOME"] + "/.config/yass/config")
	art = getConfig(mainMap["HOME"] + "/.config/yass/art")
    }
    config = parseConfig(config, false)
    art = parseConfig(art, true)

    //printing results
    artSize := len(art) - 1
    configSize := len(config) - 1
    sizeDifference := math.Abs(float64(artSize - configSize) / 2)
    if artSize > configSize {
	for i := 0; i < artSize; i++ {
	    if i >= int(math.Floor(sizeDifference)) && i <= configSize + int(math.Floor(sizeDifference)) {
                fmt.Println(art[i], config[i - int(math.Floor(sizeDifference))])
	    } else {
		fmt.Println(art[i])
	    }
	}
    } else {
	for i := 0; i < configSize; i++ {
	    if i >= int(math.Floor(sizeDifference)) && i <= artSize + int(math.Floor(sizeDifference)) {
                fmt.Println(art[i - int(math.Floor(sizeDifference))], config[i])
	    } else {
		fmt.Println(config[i])
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

func getConfig(pathToConfig string) []string {
    configStr, err := os.ReadFile(pathToConfig)
    if err != nil { log.Fatal(err) }
    config := strings.Split(string(configStr), "\n")

    return config
}


func parseConfig(config []string, prependDistrocolor bool) []string {
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
	/*
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
	*/

	styleSheet := map[string]string{
	    //clear all styling and coloring
	    "<c>": "\x1B[0m",
	    "<>":  "\x1B[0m",
	    //text styling
	    "<b>":  "\x1B[1m",
	    "<d>":  "\x1B[2m",
	    "<i>":  "\x1B[3m",
	    "<u>":  "\x1B[4m",
	    "<uu>": "\x1B[21m",
	    "<r>":  "\x1B[7m",
	    "<s>":  "\x1B[9m",
	    //coloring
	    "<black>":   "\x1B[30m",
	    "<red>":     "\x1B[31m",
	    "<green>":   "\x1B[32m",
	    "<yellow>":  "\x1B[33m",
	    "<blue>":    "\x1B[34m",
	    "<magenta>": "\x1B[35m",
	    "<cyan>":    "\x1B[36m",
	    "<white>":   "\x1B[37m",
	    //bright coloring
	    "<brblack>":   "\x1B[30m;1m",
	    "<brred>":     "\x1B[31m;1m",
	    "<brgreen>":   "\x1B[32m;1m",
	    "<bryellow>":  "\x1B[33m;1m",
	    "<brblue>":    "\x1B[34m;1m",
	    "<brmagenta>": "\x1B[35m;1m",
	    "<brcyan>":    "\x1B[36m;1m",
	    "<brwhite>":   "\x1B[37m;1m",
	    //background coloring
	    "<bgblack>":   "\x1B[40m",
	    "<bgred>":     "\x1B[41m",
	    "<bggreen>":   "\x1B[42m",
	    "<bgyellow>":  "\x1B[43m",
	    "<bgblue>":    "\x1B[44m",
	    "<bgmagenta>": "\x1B[45m",
	    "<bgcyan>":    "\x1B[46m",
	    "<bgwhite>":   "\x1B[47m",
	    //bright background coloring
	    "<bgbrblack>":   "\x1B[40m;1m",
	    "<bgbrred>":     "\x1B[41m;1m",
	    "<bgbrgreen>":   "\x1B[42m;1m",
	    "<bgbryellow>":  "\x1B[43m;1m",
	    "<bgbrblue>":    "\x1B[44m;1m",
	    "<bgbrmagenta>": "\x1B[45m;1m",
	    "<bgbrcyan>":    "\x1B[46m;1m",
	    "<bgbrwhite>":   "\x1B[47m;1m",
	}
	
	for key, code := range styleSheet {
	    config[i] = strings.ReplaceAll(config[i], key, code)
	}

	//distro color
	idLike := strings.Split(mainMap["ID_LIKE"], " ")

	if distroColor(mainMap["ID"]) != "" {
	    config[i] = strings.ReplaceAll(config[i], "<distrocolor>", distroColor(mainMap["ID"]))
	} else {
	    config[i] = strings.ReplaceAll(config[i], "<distrocolor>", idLike[0])
	}
    }
    
    return config
}

func distroColor(distroID string) string {
    switch mainMap["ID"] {
	case "centos", "debian", "rhel", "ubuntu":
	    return "\x1B[31m"
	case "linux-mint", "manjaro", "nixos", "opensuse-leap", "opensuse-tumbleweed", "void":
	    return "\x1B[32m"
	case "alpine", "fedora", "kali", "slackware", "scientific":
	    return "\x1B[34m"
	case "endeavouros", "gentoo":
	    return "\x1B[35m"
	case "arch", "clearlinux", "mageia":
	    return "\x1B[36m"
	default:
	    return ""
    }
}
