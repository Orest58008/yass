# YASS is yet another sysfetch software written in Go

## Features
* environment variables, os-release, meminfo, hostname, kernel version support
* basic configuration with ANSI styling support
* no dependencies

## Configuration
### ANSI styling
* `<b>` for bold text
* `<u>` for underline
* `<r>` for reversed colors
### Colors
* `<color>` for color
  * supported colors: black, red, green, yellow, blue, magenta, cyan and white
* `<brcolor>` for bright version of color
* `<bgcolor>` for background color
* `<bgbrcolor>` for bright version of background color
### Using variables from envvars, os-release, meminfo etc.
* `|VARIABLE_NAME|` will be replaced with variable's value
  * the name should be in ALL CAPS
  * names are the same as in source(e.g. `MEMREE` for free memory); hostname is `HOSTNAME`, kernel version is `KERNEL_VERSION`

## TODO
* [ ] example configuration
* [ ] config in ~/.config/yass
* [ ] memory and swap represented in MB and GB
* [ ] uptime support
* [x] automatic distro color detection(e.g. cyan for Arch, red for Ubuntu etc.)
* [ ] left-side ASCII art support
  * [ ] distro logo for ASCII art
