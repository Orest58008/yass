package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"
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

    //reading model
    model := ""
    modelName, err := os.ReadFile("/sys/devices/virtual/dmi/id/product_name")
    model += strings.TrimSpace(string(modelName))
    modelVersion, err := os.ReadFile("/sys/devices/virtual/dmi/id/product_version")
    model += " " + strings.TrimSpace(string(modelVersion))
    modelModel, err := os.ReadFile("/sys/firmware/devicetree/base/model")
    model += " " + strings.TrimSpace(string(modelModel))
    mainMap["MODEL"] = model

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

    //reading package counts
    packageManagers := map[string][]string{
	//listing packages with command
	"apk":		{"apk", "info"},
	"bonsai":	{"bonsai", "list"},
	"brew":		{"brew", "list"},
	"crux":		{"pkginfo", "-i"},
	"dpkg":		{"dpkg", "-l"},
	"guix":		{"quix", "package", "--list-installed"},
	"opkg":		{"opkg", "list-installed"},
	"rpm":		{"rpm", "-qa"},
	"xbps-query":	{"xbps-query", "-l"},
	//directories containing packages
	"cpt-list":	{"ls", "/var/db/cpt/installed/"},
	"emerge":	{"ls", "/var/db/pkg/*/"},
	"eopkg":	{"ls", "/var/lib/eopkg/package/"},
	"kiss":		{"ls", "/var/db/kiss/installed/"},
	"nix":		{"ls", "/nix/store/"},
	"pacman":	{"ls", "/var/lib/pacman/local/"},
	"pkgtool":	{"ls", "/var/log/packages/"},
    }

    maxCount := 0
    for pm, pmargs := range packageManagers {
	path, err := exec.LookPath(pm)
	if err == nil {
	    out, _ := exec.Command(pmargs[0], pmargs[1:]...).Output()
	    count := strings.Count(string(out), "\n")
	    mainMap[strings.ToUpper(pm) + "_COUNT"] = strconv.Itoa(count)
	    mainMap[strings.ToUpper(pm) + "_PATH"] = path
	    if count > maxCount {
		maxCount = count
		mainMap["PM_COUNT"] = strconv.Itoa(count)
		mainMap["PM_NAME"] = pm
		mainMap["PM_PATH"] = path
	    }
	}
    }

    //palette
    mainMap["PALETTE"] = "\x1B[40m  \x1B[41m  \x1B[42m  \x1B[43m  \x1B[44m  \x1B[45m  \x1B[46m  \x1B[47m  \x1B[0m"
    mainMap["BRPALETTE"] = "\x1B[100m  \x1B[101m  \x1B[102m  \x1B[103m  \x1B[104m  \x1B[105m  \x1B[106m  \x1B[107m  \x1B[0m"

    //environment variables
    //this MUST be after all other reads
    appendArray(os.Environ(), mainMap, "=")

    //dumpmap behaviour
    if dumpmap { dumpMap() }

    //updating ascii
    if ascii == "" { ascii = mainMap["ID"] }
    if _, ok := distros.Distros[ascii]; !ok { ascii = "linux" }

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
	    if i < len(art) && i > 0 {
		fmt.Println(" \x1B[1m" + art[i] + " " + config[i - 1])
	    } else if i > 0 {
		fmt.Println(artSpacer + "  " + config[i - 1])
	    } else if i < len(art) {
		fmt.Println(" \x1B[1m" + art[i])
	    }
	}
    } else {
	for i := range art {
	    if i < len(config) + 1 && i > 0 {
		fmt.Println(" \x1B[1m" + art[i] + " " + config[i - 1])
	    } else {
		fmt.Println(" \x1B[1m" + art[i])
	    }
	}
    }
}



func appendArray(array []string, target map[string]string, splitter string) {
    for i := range array {
	key, value, _ := strings.Cut(array[i], splitter)
	value = strings.TrimSpace(value)
	if strings.Contains(value, " kB") {
	    value = strings.ReplaceAll(value, " kB", "")
	    valueNum, err := strconv.Atoi(value)
	    if err != nil {
		fmt.Println("Error conveting " + key + " to int:")
		log.Fatal(err)
	    }
	    valueMb := valueNum / 1024
	    valueGb := valueMb / 1024
	    target[strings.ToUpper(strings.TrimSpace(key))+"_MB"] = strconv.Itoa(valueMb)
	    target[strings.ToUpper(strings.TrimSpace(key))+"_GB"] = strconv.Itoa(valueGb)
	}
	target[strings.ToUpper(strings.TrimSpace(key))] = value
    }
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

    //creating styling map
    styleSheet := map[string]string{
	//clear all styling and coloring
	"<clear>":	"\x1B[0m",
	//text styling
	"<b>":		"\x1B[1m",
	"<d>":		"\x1B[2m",
	"<i>":		"\x1B[3m",
	"<u>":		"\x1B[4m",
	"<uu>":		"\x1B[21m",
	"<r>":		"\x1B[7m",
	"<s>":		"\x1B[9m",
	//coloring
	"<black>":	"\x1B[30m",
	"<red>":	"\x1B[31m",
	"<green>":	"\x1B[32m",
	"<yellow>":	"\x1B[33m",
	"<blue>":	"\x1B[34m",
	"<magenta>":	"\x1B[35m",
	"<cyan>":	"\x1B[36m",
	"<white>":	"\x1B[37m",
	//bright coloring
	"<brblack>":	"\x1B[90m",
	"<brred>":	"\x1B[91m",
	"<brgreen>":	"\x1B[92m",
	"<bryellow>":	"\x1B[93m",
	"<brblue>":	"\x1B[94m",
	"<brmagenta>":	"\x1B[95m",
	"<brcyan>":	"\x1B[96m",
	"<brwhite>":	"\x1B[97m",
	//background coloring
	"<bgblack>":	"\x1B[40m",
	"<bgred>":	"\x1B[41m",
	"<bggreen>":	"\x1B[42m",
	"<bgyellow>":	"\x1B[43m",
	"<bgblue>":	"\x1B[44m",
	"<bgmagenta>":	"\x1B[45m",
	"<bgcyan>":	"\x1B[46m",
	"<bgwhite>":	"\x1B[47m",
	//bright background coloring
	"<bgbrblack>":	"\x1B[100m",
	"<bgbrred>":	"\x1B[101m",
	"<bgbrgreen>":	"\x1B[102m",
	"<bgbryellow>":	"\x1B[103m",
	"<bgbrblue>":	"\x1B[104m",
	"<bgbrmagenta>":"\x1B[105m",
	"<bgbrcyan>":	"\x1B[106m",
	"<bgbrwhite>":	"\x1B[107m",
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

    //creating styling map
    styleSheet := map[string]string{
	//clear all styling and coloring
	"<clear>":	"\x1B[0m",
	//text styling
	"<b>":		"\x1B[1m",
	"<d>":		"\x1B[2m",
	"<i>":		"\x1B[3m",
	"<u>":		"\x1B[4m",
	"<uu>":		"\x1B[21m",
	"<r>":		"\x1B[7m",
	"<s>":		"\x1B[9m",
	//coloring
	"<black>":	"\x1B[30m",
	"<red>":	"\x1B[31m",
	"<green>":	"\x1B[32m",
	"<yellow>":	"\x1B[33m",
	"<blue>":	"\x1B[34m",
	"<magenta>":	"\x1B[35m",
	"<cyan>":	"\x1B[36m",
	"<white>":	"\x1B[37m",
	//bright coloring
	"<brblack>":	"\x1B[30m",
	"<brred>":	"\x1B[31m",
	"<brgreen>":	"\x1B[32m",
	"<bryellow>":	"\x1B[33m",
	"<brblue>":	"\x1B[34m",
	"<brmagenta>":	"\x1B[35m",
	"<brcyan>":	"\x1B[36m",
	"<brwhite>":	"\x1B[37m",
	//background coloring
	"<bgblack>":	"\x1B[40m",
	"<bgred>":	"\x1B[41m",
	"<bggreen>":	"\x1B[42m",
	"<bgyellow>":	"\x1B[43m",
	"<bgblue>":	"\x1B[44m",
	"<bgmagenta>":	"\x1B[45m",
	"<bgcyan>":	"\x1B[46m",
	"<bgwhite>":	"\x1B[47m",
	//bright background coloring
	"<bgbrblack>":	"\x1B[40m",
	"<bgbrred>":	"\x1B[41m",
	"<bgbrgreen>":	"\x1B[42m",
	"<bgbryellow>":	"\x1B[43m",
	"<bgbrblue>":	"\x1B[44m",
	"<bgbrmagenta>":"\x1B[45m",
	"<bgbrcyan>":	"\x1B[46m",
	"<bgbrwhite>":	"\x1B[47m",
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

func read(path string) []string {
    resultStr, err := os.ReadFile(path)
    if err != nil {
	fmt.Println("Error while reading " + path + ":")
	log.Fatal(err)
    }
    result := strings.Split(string(resultStr), "\n")

    return result
}
