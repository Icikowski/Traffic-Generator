# Traffic Generator

**Traffic Generator** is a simple tool that allows user to put traffic with specified properties on given target.

> **⚠️ Warning**
>
> This repository is hosted on _git.sr.ht_ and mirrored to GitHub.
> You should always refer to _git.sr.ht_ version as the primary instance.

## Command line options

| Option | Description | Required | Accepted values | Default value |
|-|-|-|-|-|
| `-config` | Configuration file in JSON or YAML format which should be loaded; if specified, it will be used instead of default/flag provided values | no | string with filename or `--` for piped input | _N/A_ |
| `-name` | Traffic name (for logging purposes) | no | string | `some-traffic` |
| `-target` | Traffic target | **yes** | URL, eg. `http://example.com` | _N/A_ |
| `-success` | Desired success ratio in percents | no | float from range `1` - `100` | `90.0` |
| `-requests` | Number of simultaneous requests to be sent in given interval | no | positive integer | `30` |
| `-interval` | Requests interval | no | duration | `2s` |
| `-timeout` | Requests timeout | no | duration (≤ interval) | `1s` |
| `-insecure` | Insecure mode (SSL certificates of the target will not be verified) | no | just the presence of the flag which equals `true` | `false` |
| `-record` | Enable recording of success ratio in CSV file | no | just the presence of the flag which equals `true` | `false` |
| `-verbose` | Verbose logging console output flag | no | just the presence of the flag which equals `true` | `false` |
| `-help`, `-h` | Shows help message with command line options' descriptions | no | none | _N/A_ |
| `-version` | Show version information | no | none | _N/A_ |

## Configuration file fields

**Important:** All fields are required by default!

| Field | Description | Accepted values | Example |
|-|-|-|-|
| `name` | Traffic name (for logging purposes) | string | `"some-traffic"` |
| `target` | Traffic target | string with URL | `"http://example.com"` |
| `success_ratio` | Desired success ratio in percents | float from range `1` - `100` | `90.0` |
| `simultaneous_requests` | Number of simultaneous requests to be sent in given interval | positive integer | `30` |
| `interval` | Requests interval | Number of milliseconds or string representing duration | `2000`, `"2s"` |
| `timeout` | Requests timeout | Number of milliseconds or string representing duration | `1000`, `"1s"` |
| `insecure` | Insecure mode (SSL certificates of the target will not be verified) | boolean | `false`, `true` |
| `record` | Enable recording of success ratio in CSV file | boolean | `false`, `true` |

<details>
<summary>Example JSON configuration file</summary>

```json
{
    "name": "some-traffic",
    "target": "http://example.com",
    "success_ratio": 90.0,
    "simultaneous_requests": 30,
    "interval": "2s",
    "timeout": "1s",
    "insecure": true,
    "record": false
}
```
</details>

<details>
<summary>Example YAML configuration file</summary>

```yaml
name: some-traffic
target: http://example.com
success_ratio: 90.0
simultaneous_requests: 30
interval: 2s
timeout: 1s
insecure: true
record: false
```
</details>
