# FAV/E
[![License](https://img.shields.io/badge/license-MIT-_red.svg)](https://opensource.org/licenses/MIT)

<img src="img/fave.png">

FAV/E (Find A Vulnerability/Exposure) utilizes the NIST CVE database search API to search for vulnerabilities and exposures while filtering based on age, keywords, and other parameters.

## Usage:

Use the `-h` flag to display available flags
```
$ fave -h
```
| Flag | Description | Example |
|------|-------------|---------|
| `-cwe` | Search for CVEs based on a CWE number. | `fave -cwe 79` |
| `-exact` | Return only items matching the exact keyword(s) specified with -k | `fave -k un4gi -exact` |
| `-fd` | Number of days to filter results (prior to today; maximum 120) | `fave -fd 5` |
| `-k` | Search for CVEs based on a keyword (or words) | `fave -k "Microsoft Windows 10" -exact` |
| `-s` | Search for CVEs based on the CVSS V3 severity. | `fave -cvss CRITICAL` |

Example usage:
```
$ fave -k "Windows 10" -exact -cvss CRITICAL -fd 2
```

## Installation
To install, use:
```
go get -u github.com/un4gi/fave
```