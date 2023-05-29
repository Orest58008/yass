# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [to be added]
* Main package manager detection
* NO\_COLOR support
* Current time support
* Plugin system
* Refactor README

## [1.5.1]
* Added main package manager detection
* Moved environment variable detection to the bottom so you can replace arbitrary values at the runtime
* Updated examples

## [1.5.0]
* Added package count support
  * Currently supported package managers: apk, dnf, dpkg, equery(emerge), guix, pacman, pkginfo(pkgutils), rpm, xbps-query(xbps), yum
* Fixed vertical spacing issues

### [1.4.2]
* Fixed strconv issue, sometimes occuring in kitty due to incorrect memory variable detection

### [1.4.1]
* Fixed spacing issues
* Fixed `<>` leftovers

## [1.4.0]
* Added pfetch-like distro logo detection
* Changed variable denoters from `|` to `$`
* Removed ID\_LIKE support, instead defaulting to Linux logo

## [1.3.0]
* Added command line options, including --help, --version and others
* Added MEMUSED and SWAPUSED support

###  [1.2.1] - 2023-04-29
* Added ID\_LIKE support

##  [1.2.0] - 2023-04-21
* Added fallback config
* Moved `config` folder to `examples`
