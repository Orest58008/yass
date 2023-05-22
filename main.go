package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	distros "codeberg.org/Orest58008/yass/distros"
)

//map to store values later outputted
var mainMap map[string]string

func main() {
    mainMap = make(map[string]string)
    configPath := os.Getenv("HOME") + "/.config/yass/config"
    artPath := ""
    ascii := ""
    version := "1.4.0"
    dumpmap := false

    //read arguments
    for i := range os.Args {
	switch os.Args[i] {
	case "-c", "--config":
	    i++
	    configPath = os.Args[i]
	case "-a", "--ascii":
	    i++
	    ascii = os.Args[i]
	case "-ap", "--asciipath":
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
	    fmt.Println("  \tSpecify path to config file.")
	    fmt.Println("  \tDefault: $HOME/.config/yass/config")
	    fmt.Println("  \x1B[1m-a, --ascii\x1B[0m logo-name:")
	    fmt.Println("  \tSpecify ascii logo.")
	    fmt.Println("  \tDefault: your distro's ID from /etc/os-release")
	    fmt.Println("  \x1B[1m-ap, --asciipath\x1B[0m /path/to/ascii:")
	    fmt.Println("  \tSpecify path to custom ascii logo.")
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

    //updating ascii
    if ascii == "" { ascii = mainMap["ID"] }
    if _, ok := distros.Distros[ascii]; !ok { ascii = "linux" }
    if ascii == "opensuse-tumbleweed" || ascii == "opensuse-leap" { ascii = "suse" }

    //reading configs
    var config = []string{"<distrocolor><d><b><u>$USER$<clear>@<distrocolor><d><b><u>$HOSTNAME$",
	      "<distrocolor><b>os<clear>     $PRETTY_NAME$",
	      "<distrocolor><b>kernel<clear> $KERNEL_VERSION$",
	      "<distrocolor><b>memory<clear> $MEMFREE_MB$ / $MEMTOTAL_MB$ MiB",
	      "<distrocolor><b>swap<clear>   $SWAPFREE_MB$ / $SWAPTOTAL_MB$ MiB",
	      "<distrocolor><b>uptime<clear> $UPTIME_HRS$:$UPTIME_MINS$:$UPTIME_SECS$",""}
    var art = distros.Distros[ascii]
    var artPurged = purgeConfig(art)
    var artWidth = len(artPurged[len(artPurged) - 1])
    var artSpacer = strings.Repeat(" ", artWidth)

    if _, err := os.Stat(configPath); !os.IsNotExist(err) {
        config = read(configPath)
    }

    if _, err := os.Stat(artPath); !os.IsNotExist(err) {
        art = read(artPath)
    }

    config = parseConfig(config, art[0])
    art = parseConfig(art, art[0])


    //printing results
    if len(config) > len(art) {
	for i := 0; i < len(config) + 2; i++ {
	    if i < len(art) && i > 1 {
		fmt.Println(" \x1B[1m" + art[i] + " " + config[i - 2])
	    } else if i > 1 {
		fmt.Println(artSpacer + "  " + config[i - 2])
	    } else if i < len(art) {
		fmt.Println(" \x1B[1m" + art[i])
	    }
	}
    } else {
	for i := range art {
	    if i < len(config) + 1 && i > 1 {
		fmt.Println(" \x1B[1m" + art[i] + " " + config[i - 2])
	    } else {
		fmt.Println(" \x1B[1m" + art[i])
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

func parseConfig(config []string, distroColor string) []string {
    //result holder to not mess up the original config
    result := make([]string, len(config))
    copy(result, config)

    //parsing values
    for i := range result {
	line := strings.Split(result[i], "$")
	for j := range line {
	    val, ok := mainMap[line[j]]
	    if ok { line[j] = val }
	}
	line = append(line, "<clear>")
	result[i] = strings.Join(line, "")
    }

    //creating style maps styling
    styleSheet := map[string]string{
	//clear all styling and coloring
	"<clear>": "\x1B[0m",
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

    //parsing styling
    for i := range result {
	result[i] = strings.ReplaceAll(result[i], "<distrocolor>", distroColor)

	for key, code := range styleSheet {
	    result[i] = strings.ReplaceAll(result[i], key, code)
	}
    }
    
    return result
}

func purgeConfig(config []string) []string {
    //result holder to not mess up the original config
    result := make([]string, len(config))
    copy(result, config)

    //creating style maps styling
    styleSheet := map[string]string{
	//clear all styling and coloring
	"<clear>": "\x1B[0m",
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

    //parsing styling
    for i := range result {
	result[i] = strings.ReplaceAll(result[i], "<distrocolor>", "")

	for key := range styleSheet {
	    result[i] = strings.ReplaceAll(result[i], key, "")
	}
    }
    
    return result
}
