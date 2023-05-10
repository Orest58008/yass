package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

//map to store values later outputted
var mainMap map[string]string
    
func main() {
    mainMap = make(map[string]string)
    configPath := os.Getenv("HOME") + "/.config/yass/config"
    artPath := os.Getenv("HOME") + "/.config/yass/art"
    version := "1.3.0"
    dumpmap := false

    //read arguments
    for i := range os.Args {
	switch os.Args[i] {
	case "-c", "--config":
	    i++
	    configPath = os.Args[i]
	case "-a", "--ascii":
	    i++
	    artPath = os.Args[i]
	case "-d", "--dumpmap":
	    dumpmap = true
	case "-V", "--version":
	    fmt.Println(version)
	    os.Exit(0)
	case "-h", "--help":
	    fmt.Println("YASS - Yet Another Sysfetch Software")
	    fmt.Println("")
	    fmt.Println("\x1B[1mOptions:\x1B[0m")
	    fmt.Println("  \x1B[1m-c, --config\x1B[0m /path/to/config:")
	    fmt.Println("  \tUse custom config file.")
	    fmt.Println("  \tDefault: $HOME/.config/yass/config")
	    fmt.Println("  \x1B[1m-a, --ascii\x1B[0m /path/to/ascii:")
	    fmt.Println("  \tUse custom ascii art file.")
	    fmt.Println("  \tDefault: $HOME/.config/yass/art")
	    fmt.Println("  \x1B[1m-d, --dumpmap\x1B[0m:")
	    fmt.Println("  \tPrint all the values yass has collected and can use.")
	    fmt.Println("  \x1B[1m-V, --version\x1B[0m:")
	    fmt.Println("  \tPrint yass version and exit.")
	    fmt.Println("  \x1B[1m-h, --help\x1B[0m:")
	    fmt.Println("  \tPrint this message and exit.")
	    os.Exit(0)
	}
    }

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
    //adding MEMUSED and SWAPUSED
    calculateUsedMemory("MEMFREE", "MEMTOTAL", "MEMUSED");
    calculateUsedMemory("MEMFREE_MB", "MEMTOTAL_MB", "MEMUSED_MB");
    calculateUsedMemory("MEMFREE_GB", "MEMTOTAL_GB", "MEMUSED_GB");
    calculateUsedMemory("SWAPFREE", "SWAPTOTAL", "SWAPUSED");
    calculateUsedMemory("SWAPFREE_MB", "SWAPTOTAL_MB", "SWAPUSED_MB");
    calculateUsedMemory("SWAPFREE_GB", "SWAPTOTAL_GB", "SWAPUSED_GB");

    //reading hostname
    hostname, err := os.ReadFile("/etc/hostname")
    if err != nil {
	fmt.Println("Error while reading /etc/hostname:")
	log.Fatal(err) 
    }
    mainMap["HOSTNAME"] = strings.TrimSpace(string(hostname)) 

    //reading kernel version
    kernelVersion, err := os.ReadFile("/proc/sys/kernel/osrelease")
    if err != nil {
	fmt.Println("Error while reading /proc/sys/kernel/osrelease:")
	log.Fatal(err)
    }
    mainMap["KERNEL_VERSION"] = strings.TrimSpace(string(kernelVersion))

    //reading uptime
    uptimeStr, err := os.ReadFile("/proc/uptime")
    if err != nil { 
	fmt.Println("Error while reading /proc/uptime:")
	log.Fatal(err)
    }
    uptimeArr := strings.Split(string(uptimeStr), " ")
    uptimeSecsFloat, err := strconv.ParseFloat(uptimeArr[0], 64)
    if err != nil {
	fmt.Println("Error parsing uptime:")
	log.Fatal(err)
    }
    uptimeSecs := int(uptimeSecsFloat)
    uptimeHrs := uptimeSecs / 3600
    uptimeSecs -= uptimeHrs * 3600
    uptimeMins := uptimeSecs / 60
    uptimeSecs -= uptimeMins * 60
    mainMap["UPTIME_HRS"] = strconv.Itoa(uptimeHrs)
    mainMap["UPTIME_MINS"] = strconv.Itoa(uptimeMins)
    mainMap["UPTIME_SECS"] = strconv.Itoa(uptimeSecs)
    
    //dumpmap behaviour
    if dumpmap { dumpMap() }

    //reading configs
    var config = []string{" <distrocolor><d><b><u>|USER|<>@<distrocolor><d><b><u>|HOSTNAME|",
	      " <distrocolor><b>os<>     |PRETTY_NAME|",
	      " <distrocolor><b>kernel<> |KERNEL_VERSION|",
	      " <distrocolor><b>memory<> |MEMFREE_MB| / |MEMTOTAL_MB| MiB",
	      " <distrocolor><b>swap<>   |SWAPFREE_MB| / |SWAPTOTAL_MB| MiB",
	      " <distrocolor><b>uptime<> |UPTIME_HRS|:|UPTIME_MINS|:|UPTIME_SECS|", ""}
    var art = []string{"<b>\\   ", "<b>\\\\  ", "<b> \\\\ ", "<b>  \\\\", "<b>  //", "<b> // ", "<b>//  ", "<b>/   ", ""}

    if _, err := os.Stat(configPath); !os.IsNotExist(err) {
        config = read(configPath)
    }

    if _, err := os.Stat(artPath); !os.IsNotExist(err) {
        art = read(artPath)
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

func dumpMap() {
    keys := make([]string, 0, len(mainMap))
    for k := range mainMap {
	keys = append(keys, k)
    }
    sort.Strings(keys)
    for _, k := range keys {
	fmt.Println(k, mainMap[k])
    }
    os.Exit(0)
}

func read(path string) []string {
    resultStr, err := os.ReadFile(path)
    if err != nil {
	fmt.Println("Error while reading " + path + ":")
	log.Fatal(err)
    }
    result := strings.Split(string(resultStr), "\n")

    return result
}

func calculateUsedMemory(freeKey string, totalKey string, resultKey string) {
    memfree, err := strconv.Atoi(mainMap[freeKey])
    if err != nil {
	fmt.Println("Error getting " + freeKey + " from map:")
	log.Fatal(err)
    }
    memtotal, err := strconv.Atoi(mainMap[totalKey])
    if err != nil {
	fmt.Println("Error getting " + totalKey + " from map:")
	log.Fatal(err)
    }
    memused := memtotal - memfree
    mainMap[resultKey] = strconv.Itoa(memused)
}

func appendArray(array []string, destinationMap map[string]string, splitter string) {
    for i := range array {
	key, value, _ := strings.Cut(array[i], splitter)
	value = strings.TrimSpace(value)
	if strings.Contains(value, "kB") {
            value = strings.ReplaceAll(value, " kB", "")
	    valueNum, err := strconv.Atoi(value)
	    if err != nil {
		fmt.Println("Error conveting " + key + " to int:")
		log.Fatal(err)
	    }
	    valueMb := valueNum / 1024
	    valueGb := valueMb / 1024
	    destinationMap[strings.ToUpper(strings.TrimSpace(key))+"_MB"] = strconv.Itoa(valueMb)
	    destinationMap[strings.ToUpper(strings.TrimSpace(key))+"_GB"] = strconv.Itoa(valueGb)
	}
	destinationMap[strings.ToUpper(strings.TrimSpace(key))] = value
    }
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

    //creating style maps styling
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

    distroColors := map[string]string{
	//red
	"centos": "\x1B[31m",
	"debian": "\x1B[31m",
	"rhel": "\x1B[31m",
	"ubuntu": "\x1B[31m",
	//green
	"linux-mint": "\x1B[32m",
	"manjaro": "\x1B[32m",
	"opensuse-leap": "\x1B[32m",
	"opensuse-tumbleweed": "\x1B[32m",
	"void": "\x1B[32m",
	//blue
	"alpine": "\x1B[34m",
	"fedora": "\x1B[34m",
	"kali": "\x1B[34m",
	"slackware": "\x1B[34m",
	"scientific": "\x1B[34m",
	//magenta
	"endeavouros": "\x1B[35m",
	"gentoo": "\x1B[35m",
	//cyan
	"arch": "\x1B[36m",
	"clearlinux": "\x1B[36m",
	"mageia": "\x1B[36m",
    }

    idLike := strings.Split(mainMap["ID_LIKE"], " ")[0]

    //parsing styling
    for i := range config {
	for key, code := range styleSheet {
	    config[i] = strings.ReplaceAll(config[i], key, code)
	}

	if color, ok := distroColors[mainMap["ID"]]; ok {
	    config[i] = strings.ReplaceAll(config[i], "<distrocolor>", color)
	} else if colorLike, ok := distroColors[idLike]; ok {
	    config[i] = strings.ReplaceAll(config[i], "<distrocolor>", colorLike)
	}

    }
    
    return config
}
