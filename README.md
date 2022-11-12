# FAV/E

[![License](https://img.shields.io/badge/license-MIT-_red.svg)](https://opensource.org/licenses/MIT) [![Twitter Follow](https://img.shields.io/twitter/follow/un4gi_io?label=%40un4gi_io&style=social)](https://twitter.com/un4gi_io)

<img src="img/fave.png">

FAV/E (Find A Vulnerability/Exposure) utilizes the NIST CVE database search API to search for vulnerabilities and exposures while filtering based on age, keywords, and other parameters.

## Usage

FAV/E currently features two commands: `search` and `describe`.

The `search` command allows you to search the CVE database while filtering on a number of different things, such as:

- keyword/words
- age (maximum of 120 days)
- CWE ID
- CVSS Severity

Use the `-h` flag to display available flags

```bash
fave search -h
```

| Flag | Description | Example |
|------|-------------|---------|
| `-c, --cwe int` | Search for CVEs based on a CWE number. | `fave -c 79` |
| `-d, --days int` | the number of days prior to today to filter (maximum of 120 days). | `fave --days 5` |
| `--exact` | Return only items matching the exact keyword(s) specified with -k | `fave -k un4gi --exact` |
| `-h, --help` | help for search | `fave -h` |
| `-k, --key string` | a word or phrase to search - this is required. | `fave -k "Microsoft Windows 10"` |
| `-s, --severity string` | filters based on the CVSS V3 severity rating (CRITICAL, HIGH, MEDIUM, or LOW). | `fave -s CRITICAL` |

Example usage:

```bash
fave search -k "Windows 10" --exact -s CRITICAL --days 2
```

Once an interesting CVE is found, you can pass the CVE-ID to the `describe` command to gather a more detailed description with references:

`fave describe CVE-XXXX-XXXXX`

## Installation

To install, use:

```bash
go install github.com/un4gi/fave@latest
```

Alternatively, you can download the latest release for your OS.