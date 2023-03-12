package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

    //reading config and outputting final result
    result := parseConfig("./config")
    for i := 0; i < len(result) - 1; i++ {
        fmt.Println(result[i])
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

func parseConfig(pathToConfig string) []string {
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
	config[i] = strings.Join(line, "")
    }

    //parsing basic styling
    for i := range config {
	//clear styles
	config[i] = strings.ReplaceAll(config[i], "<c>", "\u001b[0m")
	config[i] = strings.ReplaceAll(config[i], "<>", "\u001b[0m")
        //bold, underlined and reversed
	config[i] = strings.ReplaceAll(config[i], "<b>", "\u001b[1m")
	config[i] = strings.ReplaceAll(config[i], "<u>", "\u001b[4m")
	config[i] = strings.ReplaceAll(config[i], "<r>", "\u001b[7m")
        //colors
	config[i] = strings.ReplaceAll(config[i], "<black>", "\u001b[30m")
	config[i] = strings.ReplaceAll(config[i], "<red>", "\u001b[31m")
	config[i] = strings.ReplaceAll(config[i], "<green>", "\u001b[32m")
	config[i] = strings.ReplaceAll(config[i], "<yellow>", "\u001b[33m")
	config[i] = strings.ReplaceAll(config[i], "<blue>", "\u001b[34m")
	config[i] = strings.ReplaceAll(config[i], "<magenta>", "\u001b[35m")
	config[i] = strings.ReplaceAll(config[i], "<cyan>", "\u001b[36m")
	config[i] = strings.ReplaceAll(config[i], "<white>", "\u001b[37m")
	//bright colors
	config[i] = strings.ReplaceAll(config[i], "<brblack>", "\u001b[30m;1m")
	config[i] = strings.ReplaceAll(config[i], "<brred>", "\u001b[31m;1m")
	config[i] = strings.ReplaceAll(config[i], "<brgreen>", "\u001b[32m;1m")
	config[i] = strings.ReplaceAll(config[i], "<bryellow>", "\u001b[33m;1m")
	config[i] = strings.ReplaceAll(config[i], "<brblue>", "\u001b[34m;1m")
	config[i] = strings.ReplaceAll(config[i], "<brmagenta>", "\u001b[35m;1m")
	config[i] = strings.ReplaceAll(config[i], "<brcyan>", "\u001b[36m;1m")
	config[i] = strings.ReplaceAll(config[i], "<brwhite>", "\u001b[37m;1m")
	//background colors
	config[i] = strings.ReplaceAll(config[i], "<bgblack>", "\u001b[40m")
	config[i] = strings.ReplaceAll(config[i], "<bgred>", "\u001b[41m")
	config[i] = strings.ReplaceAll(config[i], "<bggreen>", "\u001b[42m")
	config[i] = strings.ReplaceAll(config[i], "<bgyellow>", "\u001b[43m")
	config[i] = strings.ReplaceAll(config[i], "<bgblue>", "\u001b[44m")
	config[i] = strings.ReplaceAll(config[i], "<bgmagenta>", "\u001b[45m")
	config[i] = strings.ReplaceAll(config[i], "<bgcyan>", "\u001b[46m")
	config[i] = strings.ReplaceAll(config[i], "<bgwhite>", "\u001b[47m")
	//background bright colors
	config[i] = strings.ReplaceAll(config[i], "<bgbrblack>", "\u001b[40m;1m")
	config[i] = strings.ReplaceAll(config[i], "<bgbrred>", "\u001b[41m;1m")
	config[i] = strings.ReplaceAll(config[i], "<bgbrgreen>", "\u001b[42m;1m")
	config[i] = strings.ReplaceAll(config[i], "<bgbryellow>", "\u001b[43m;1m")
	config[i] = strings.ReplaceAll(config[i], "<bgbrblue>", "\u001b[44m;1m")
	config[i] = strings.ReplaceAll(config[i], "<bgbrmagenta>", "\u001b[45m;1m")
	config[i] = strings.ReplaceAll(config[i], "<bgbrcyan>", "\u001b[46m;1m")
	config[i] = strings.ReplaceAll(config[i], "<bgbrwhite>", "\u001b[47m;1m")
	//distro color
	switch mainMap["NAME"] {
	case "Ubuntu", "Debian":
	    config[i] = strings.ReplaceAll(config[i], "<distrocolor>", "\u001b[31m") //red
	case "OpenSUSE", "Linux Mint":
	    config[i] = strings.ReplaceAll(config[i], "<distrocolor>", "\u001b[32m") //green
	case "Fedora Linux", "Slackware":
	    config[i] = strings.ReplaceAll(config[i], "<distrocolor>", "\u001b[34m") //blue
	case "Gentoo":
	    config[i] = strings.ReplaceAll(config[i], "<distrocolor>", "\u001b[35m") //magenta
	case "Arch Linux":
	    config[i] = strings.ReplaceAll(config[i], "<distrocolor>", "\u001b[36m") //cyan
	}
    }
    
    return config
}
