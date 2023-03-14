# YASS is yet another sysfetch software written in Go

## Features
* environment variables, os-release, meminfo, hostname, kernel version support
* basic configuration with ANSI styling support
* no dependencies

## Configuration
### ANSI styling
* `<b>` for bold text
* `<d>` for dimmed
* `<i>` for italics
* `<u>` for underline
* `<uu>` for double underline
* `<s>` for strikethrough
* `<r>` for reversed colors
* `<c>` or `<>` to clear all styling including color
### Colors
* `<color>` for color
  * supported colors: black, red, green, yellow, blue, magenta, cyan and white
* `<brcolor>` for bright version of color
* `<bgcolor>` for background color
* `<bgbrcolor>` for bright version of background color
* `<distrocolor>` for your Linux distro's color
### Using variables from envvars, os-release, meminfo etc.
* `|VARIABLE_NAME|` will be replaced with variable's value
  * names are the same as in source(e.g. `MEMREE` for free memory); hostname is `HOSTNAME`, kernel version is `KERNEL_VERSION`
  * the name should be in ALL CAPS
  * by default, meminfo values are in kB, for values in MiB and GiB use `VARIABLE_MB` and `VARIABLE_GB`

## TODO
* [x] example configuration
* [x] config in ~/.config/yass
* [x] memory and swap represented in MB and GB
* [ ] uptime support
* [x] automatic distro color detection(e.g. cyan for Arch, red for Ubuntu etc.)
* [x] left-side ASCII art support
  * [ ] distro logo for ASCII art
