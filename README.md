![Go version](https://img.shields.io/github/go-mod/go-version/Icikowski/Traffic-Generator?filename=application%2Fgo.mod&style=for-the-badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/Icikowski/Traffic-Generator?style=for-the-badge)](https://goreportcard.com/report/github.com/Icikowski/Traffic-Generator)
![License](https://img.shields.io/github/license/Icikowski/Traffic-Generator?style=for-the-badge)

# Traffic Generator

**Traffic Generator** is a simple tool that allows user to put traffic with specified properties on given target.

## Command line options

| Option | Description | Required | Accepted values | Default value |
|-|-|-|-|-|
| `-name` | Traffic name (for logging purposes) | no | string | `some-traffic` |
| `-target` | Traffic target | **yes** | URL, eg. `http://example.com` | _N/A_ |
| `-success` | Desired success ratio in percents | no | integer from range `1` - `100` | `90` |
| `-requests` | Number of simultaneous requests to be sent in given interval | no | positive integer | `30` |
| `-interval` | Requests interval | no | duration | `2s` |
| `-timeout` | Requests timeout | no | duration (≤ interval) | `1s` |
| `-help`, `-h` | Shows help message with command line options' descriptions | no | none | _N/A_ |
