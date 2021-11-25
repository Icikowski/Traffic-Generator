![Go version](https://img.shields.io/github/go-mod/go-version/Icikowski/Traffic-Generator?filename=application%2Fgo.mod&style=for-the-badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/Icikowski/Traffic-Generator?style=for-the-badge)](https://goreportcard.com/report/github.com/Icikowski/Traffic-Generator)
![License](https://img.shields.io/github/license/Icikowski/Traffic-Generator?style=for-the-badge)

# Traffic Generator

**Traffic Generator** is a simple tool that allows user to put traffic with specified properties on given target.

## Command line options

| Option | Description | Required | Accepted values | Default value |
|-|-|-|-|-|
| `-config` | Configuration file in JSON or YAML format which should be loaded; if specified, it will be used instead of default/flag provided values | no | string with filename or `--` for piped input | _N/A_ |
| `-name` | Traffic name (for logging purposes) | no | string | `some-traffic` |
| `-target` | Traffic target | **yes** | URL, eg. `http://example.com` | _N/A_ |
| `-success` | Desired success ratio in percents | no | integer from range `1` - `100` | `90` |
| `-requests` | Number of simultaneous requests to be sent in given interval | no | positive integer | `30` |
| `-interval` | Requests interval | no | duration | `2s` |
| `-timeout` | Requests timeout | no | duration (≤ interval) | `1s` |
| `-insecure` | Insecure mode (SSL certificates of the target will not be verified) | no | just the presence of the flag which equals `true` | `false` |
| `-verbose` | Verbose logging console output flag | no | just the presence of the flag which equals `true` | `false` |
| `-help`, `-h` | Shows help message with command line options' descriptions | no | none | _N/A_ |
| `-version` | Show version information | no | none | _N/A_ |

## Configuration file fields

**Important:** All fields are required by default!

| Field | Description | Accepted values | Example |
|-|-|-|-|
| `name` | Traffic name (for logging purposes) | string | `"some-traffic"` |
| `target` | Traffic target | string with URL | `"http://example.com"` |
| `success_ratio` | Desired success ratio in percents | integer from range `1` - `100` | `90` |
| `simultaneous_requests` | Number of simultaneous requests to be sent in given interval | positive integer | `30` |
| `interval` | Requests interval | Number of milliseconds or string representing duration | `2000`, `"2s"` |
| `timeout` | Requests timeout | Number of milliseconds or string representing duration | `1000`, `"1s"` |
| `insecure` | Insecure mode (SSL certificates of the target will not be verified) | boolean | `false`, `true` |

<details>
<summary>Example JSON configuration file</summary>

```json
{
    "name": "some-traffic",
    "target": "http://example.com",
    "success_ratio": 90,
    "simultaneous_requests": 30,
    "interval": "2s",
    "timeout": "1s",
    "insecure": true
}
```
</details>

<details>
<summary>Example YAML configuration file</summary>

```yaml
name: some-traffic
target: http://example.com
success_ratio: 90
simultaneous_requests: 30
interval: 2s
timeout: 1s
insecure: true
```
</details>

## Demo

<details>
<summary>Default logging</summary>
  
**Console:**

```log
PS G:\Demo> .\traffic-generator.exe -target http://localhost:8000
2021-28-08 01:09:03 INF configuration is valid configuration={"interval":2000,"name":"some-traffic","simultaneous_requests":30,"success_ratio":90,"target":"http://localhost:8000","timeout":1000}
2021-28-08 01:09:03 INF started logging to file filename="some-traffic (2021-28-08 01.09.03).log"
2021-28-08 01:09:03 INF starting analysis
2021-28-08 01:09:18 WRN target is DOWN cycle=7
2021-28-08 01:09:28 INF target is UP again cycle=12 downtime=0m10s
2021-28-08 01:09:34 WRN target is DOWN cycle=15
2021-28-08 01:09:42 INF target is UP again cycle=19 downtime=0m08s
2021-28-08 01:09:45 INF shutdown requested by user
2021-28-08 01:09:45 INF end statistics total downtime=0m18s
2021-28-08 01:09:45 INF shutdown successful
PS G:\Demo>
```

**File:**

```json
{"level":"info","filename":"some-traffic (2021-28-08 01.09.03).log","time":"2021-08-28T01:09:03+02:00","message":"started logging to file"}
{"level":"info","time":"2021-08-28T01:09:03+02:00","message":"starting analysis"}
{"level":"debug","cycle":1,"time":"2021-08-28T01:09:03+02:00","message":"cycle started"}
{"level":"debug","cycle":1,"success ratio":100,"time":"2021-08-28T01:09:05+02:00","message":"cycle finished"}
{"level":"debug","cycle":2,"time":"2021-08-28T01:09:05+02:00","message":"cycle started"}
{"level":"debug","cycle":2,"success ratio":100,"time":"2021-08-28T01:09:07+02:00","message":"cycle finished"}
{"level":"debug","cycle":3,"time":"2021-08-28T01:09:07+02:00","message":"cycle started"}
{"level":"debug","cycle":3,"success ratio":100,"time":"2021-08-28T01:09:09+02:00","message":"cycle finished"}
{"level":"debug","cycle":4,"time":"2021-08-28T01:09:09+02:00","message":"cycle started"}
{"level":"debug","cycle":4,"success ratio":100,"time":"2021-08-28T01:09:11+02:00","message":"cycle finished"}
{"level":"debug","cycle":5,"time":"2021-08-28T01:09:11+02:00","message":"cycle started"}
{"level":"debug","cycle":5,"success ratio":100,"time":"2021-08-28T01:09:13+02:00","message":"cycle finished"}
{"level":"debug","cycle":6,"time":"2021-08-28T01:09:13+02:00","message":"cycle started"}
{"level":"debug","cycle":6,"success ratio":100,"time":"2021-08-28T01:09:16+02:00","message":"cycle finished"}
{"level":"debug","cycle":7,"time":"2021-08-28T01:09:16+02:00","message":"cycle started"}
{"level":"debug","cycle":7,"ratio requested":90,"ratio actual":0,"time":"2021-08-28T01:09:18+02:00","message":"success ratio is below requested value"}
{"level":"warn","cycle":7,"time":"2021-08-28T01:09:18+02:00","message":"target is DOWN"}
{"level":"debug","cycle":7,"success ratio":0,"time":"2021-08-28T01:09:18+02:00","message":"cycle finished"}
{"level":"debug","cycle":8,"time":"2021-08-28T01:09:18+02:00","message":"cycle started"}
{"level":"debug","cycle":8,"ratio requested":90,"ratio actual":0,"time":"2021-08-28T01:09:20+02:00","message":"success ratio is below requested value"}
{"level":"debug","cycle":8,"success ratio":0,"time":"2021-08-28T01:09:20+02:00","message":"cycle finished"}
{"level":"debug","cycle":9,"time":"2021-08-28T01:09:20+02:00","message":"cycle started"}
{"level":"debug","cycle":9,"ratio requested":90,"ratio actual":0,"time":"2021-08-28T01:09:22+02:00","message":"success ratio is below requested value"}
{"level":"debug","cycle":9,"success ratio":0,"time":"2021-08-28T01:09:22+02:00","message":"cycle finished"}
{"level":"debug","cycle":10,"time":"2021-08-28T01:09:22+02:00","message":"cycle started"}
{"level":"debug","cycle":10,"ratio requested":90,"ratio actual":0,"time":"2021-08-28T01:09:24+02:00","message":"success ratio is below requested value"}
{"level":"debug","cycle":10,"success ratio":0,"time":"2021-08-28T01:09:24+02:00","message":"cycle finished"}
{"level":"debug","cycle":11,"time":"2021-08-28T01:09:24+02:00","message":"cycle started"}
{"level":"debug","cycle":11,"ratio requested":90,"ratio actual":0,"time":"2021-08-28T01:09:26+02:00","message":"success ratio is below requested value"}
{"level":"debug","cycle":11,"success ratio":0,"time":"2021-08-28T01:09:26+02:00","message":"cycle finished"}
{"level":"debug","cycle":12,"time":"2021-08-28T01:09:26+02:00","message":"cycle started"}
{"level":"info","cycle":12,"downtime":"0m10s","time":"2021-08-28T01:09:28+02:00","message":"target is UP again"}
{"level":"debug","cycle":12,"success ratio":90,"time":"2021-08-28T01:09:28+02:00","message":"cycle finished"}
{"level":"debug","cycle":13,"time":"2021-08-28T01:09:28+02:00","message":"cycle started"}
{"level":"debug","cycle":13,"success ratio":100,"time":"2021-08-28T01:09:30+02:00","message":"cycle finished"}
{"level":"debug","cycle":14,"time":"2021-08-28T01:09:30+02:00","message":"cycle started"}
{"level":"debug","cycle":14,"success ratio":100,"time":"2021-08-28T01:09:32+02:00","message":"cycle finished"}
{"level":"debug","cycle":15,"time":"2021-08-28T01:09:32+02:00","message":"cycle started"}
{"level":"debug","cycle":15,"ratio requested":90,"ratio actual":0,"time":"2021-08-28T01:09:34+02:00","message":"success ratio is below requested value"}
{"level":"warn","cycle":15,"time":"2021-08-28T01:09:34+02:00","message":"target is DOWN"}
{"level":"debug","cycle":15,"success ratio":0,"time":"2021-08-28T01:09:34+02:00","message":"cycle finished"}
{"level":"debug","cycle":16,"time":"2021-08-28T01:09:34+02:00","message":"cycle started"}
{"level":"debug","cycle":16,"ratio requested":90,"ratio actual":0,"time":"2021-08-28T01:09:36+02:00","message":"success ratio is below requested value"}
{"level":"debug","cycle":16,"success ratio":0,"time":"2021-08-28T01:09:36+02:00","message":"cycle finished"}
{"level":"debug","cycle":17,"time":"2021-08-28T01:09:36+02:00","message":"cycle started"}
{"level":"debug","cycle":17,"ratio requested":90,"ratio actual":0,"time":"2021-08-28T01:09:38+02:00","message":"success ratio is below requested value"}
{"level":"debug","cycle":17,"success ratio":0,"time":"2021-08-28T01:09:38+02:00","message":"cycle finished"}
{"level":"debug","cycle":18,"time":"2021-08-28T01:09:38+02:00","message":"cycle started"}
{"level":"debug","cycle":18,"ratio requested":90,"ratio actual":0,"time":"2021-08-28T01:09:40+02:00","message":"success ratio is below requested value"}
{"level":"debug","cycle":18,"success ratio":0,"time":"2021-08-28T01:09:40+02:00","message":"cycle finished"}
{"level":"debug","cycle":19,"time":"2021-08-28T01:09:40+02:00","message":"cycle started"}
{"level":"info","cycle":19,"downtime":"0m08s","time":"2021-08-28T01:09:42+02:00","message":"target is UP again"}
{"level":"debug","cycle":19,"success ratio":96.666664,"time":"2021-08-28T01:09:42+02:00","message":"cycle finished"}
{"level":"debug","cycle":20,"time":"2021-08-28T01:09:42+02:00","message":"cycle started"}
{"level":"debug","cycle":20,"success ratio":100,"time":"2021-08-28T01:09:44+02:00","message":"cycle finished"}
{"level":"debug","cycle":21,"time":"2021-08-28T01:09:44+02:00","message":"cycle started"}
{"level":"info","time":"2021-08-28T01:09:45+02:00","message":"shutdown requested by user"}
{"level":"info","total downtime":"0m18s","time":"2021-08-28T01:09:45+02:00","message":"end statistics"}
{"level":"info","time":"2021-08-28T01:09:45+02:00","message":"shutdown successful"}
```
</details>

<details>
<summary>Verbose logging</summary>

**Console:**

```log
PS G:\Demo> .\traffic-generator.exe -target http://localhost:8000 -verbose
2021-28-08 01:19:45 INF configuration is valid configuration={"interval":2000,"name":"some-traffic","simultaneous_requests":30,"success_ratio":90,"target":"http://localhost:8000","timeout":1000}
2021-28-08 01:19:45 INF started logging to file filename="some-traffic (2021-28-08 01.19.45).log"
2021-28-08 01:19:45 INF starting analysis
2021-28-08 01:19:45 DBG cycle started cycle=1
2021-28-08 01:19:47 DBG cycle finished cycle=1 success ratio=100
2021-28-08 01:19:47 DBG cycle started cycle=2
2021-28-08 01:19:49 DBG cycle finished cycle=2 success ratio=100
2021-28-08 01:19:49 DBG cycle started cycle=3
2021-28-08 01:19:51 DBG cycle finished cycle=3 success ratio=100
2021-28-08 01:19:51 DBG cycle started cycle=4
2021-28-08 01:19:53 DBG success ratio is below requested value cycle=4 ratio actual=0 ratio requested=90
2021-28-08 01:19:53 WRN target is DOWN cycle=4
2021-28-08 01:19:53 DBG cycle finished cycle=4 success ratio=0
2021-28-08 01:19:53 DBG cycle started cycle=5
2021-28-08 01:19:55 DBG success ratio is below requested value cycle=5 ratio actual=0 ratio requested=90
2021-28-08 01:19:55 DBG cycle finished cycle=5 success ratio=0
2021-28-08 01:19:55 DBG cycle started cycle=6
2021-28-08 01:19:57 DBG success ratio is below requested value cycle=6 ratio actual=0 ratio requested=90
2021-28-08 01:19:57 DBG cycle finished cycle=6 success ratio=0
2021-28-08 01:19:57 DBG cycle started cycle=7
2021-28-08 01:19:59 DBG success ratio is below requested value cycle=7 ratio actual=0 ratio requested=90
2021-28-08 01:19:59 DBG cycle finished cycle=7 success ratio=0
2021-28-08 01:19:59 DBG cycle started cycle=8
2021-28-08 01:20:01 INF target is UP again cycle=8 downtime=0m08s
2021-28-08 01:20:01 DBG cycle finished cycle=8 success ratio=100
2021-28-08 01:20:01 DBG cycle started cycle=9
2021-28-08 01:20:03 DBG cycle finished cycle=9 success ratio=100
2021-28-08 01:20:03 DBG cycle started cycle=10
2021-28-08 01:20:05 DBG cycle finished cycle=10 success ratio=100
2021-28-08 01:20:05 DBG cycle started cycle=11
2021-28-08 01:20:07 DBG cycle finished cycle=11 success ratio=100
2021-28-08 01:20:07 DBG cycle started cycle=12
2021-28-08 01:20:09 DBG cycle finished cycle=12 success ratio=100
2021-28-08 01:20:09 DBG cycle started cycle=13
2021-28-08 01:20:11 DBG success ratio is below requested value cycle=13 ratio actual=0 ratio requested=90
2021-28-08 01:20:11 WRN target is DOWN cycle=13
2021-28-08 01:20:11 DBG cycle finished cycle=13 success ratio=0
2021-28-08 01:20:11 DBG cycle started cycle=14
2021-28-08 01:20:13 DBG success ratio is below requested value cycle=14 ratio actual=0 ratio requested=90
2021-28-08 01:20:13 DBG cycle finished cycle=14 success ratio=0
2021-28-08 01:20:13 DBG cycle started cycle=15
2021-28-08 01:20:15 DBG success ratio is below requested value cycle=15 ratio actual=0 ratio requested=90
2021-28-08 01:20:15 DBG cycle finished cycle=15 success ratio=0
2021-28-08 01:20:15 DBG cycle started cycle=16
2021-28-08 01:20:17 DBG success ratio is below requested value cycle=16 ratio actual=0 ratio requested=90
2021-28-08 01:20:17 DBG cycle finished cycle=16 success ratio=0
2021-28-08 01:20:17 DBG cycle started cycle=17
2021-28-08 01:20:19 DBG success ratio is below requested value cycle=17 ratio actual=0 ratio requested=90
2021-28-08 01:20:19 DBG cycle finished cycle=17 success ratio=0
2021-28-08 01:20:19 DBG cycle started cycle=18
2021-28-08 01:20:22 INF target is UP again cycle=18 downtime=0m10s
2021-28-08 01:20:22 DBG cycle finished cycle=18 success ratio=100
2021-28-08 01:20:22 DBG cycle started cycle=19
2021-28-08 01:20:24 DBG cycle finished cycle=19 success ratio=100
2021-28-08 01:20:24 DBG cycle started cycle=20
2021-28-08 01:20:26 DBG cycle finished cycle=20 success ratio=100
2021-28-08 01:20:26 DBG cycle started cycle=21
2021-28-08 01:20:28 DBG cycle finished cycle=21 success ratio=100
2021-28-08 01:20:28 DBG cycle started cycle=22
2021-28-08 01:20:29 INF shutdown requested by user
2021-28-08 01:20:29 INF end statistics total downtime=0m18s
2021-28-08 01:20:29 INF shutdown successful
PS G:\Demo>
```

**File:**

```json
{"level":"info","filename":"some-traffic (2021-28-08 01.19.45).log","time":"2021-08-28T01:19:45+02:00","message":"started logging to file"}
{"level":"info","time":"2021-08-28T01:19:45+02:00","message":"starting analysis"}
{"level":"debug","cycle":1,"time":"2021-08-28T01:19:45+02:00","message":"cycle started"}
{"level":"debug","cycle":1,"success ratio":100,"time":"2021-08-28T01:19:47+02:00","message":"cycle finished"}
{"level":"debug","cycle":2,"time":"2021-08-28T01:19:47+02:00","message":"cycle started"}
{"level":"debug","cycle":2,"success ratio":100,"time":"2021-08-28T01:19:49+02:00","message":"cycle finished"}
{"level":"debug","cycle":3,"time":"2021-08-28T01:19:49+02:00","message":"cycle started"}
{"level":"debug","cycle":3,"success ratio":100,"time":"2021-08-28T01:19:51+02:00","message":"cycle finished"}
{"level":"debug","cycle":4,"time":"2021-08-28T01:19:51+02:00","message":"cycle started"}
{"level":"debug","cycle":4,"ratio requested":90,"ratio actual":0,"time":"2021-08-28T01:19:53+02:00","message":"success ratio is below requested value"}
{"level":"warn","cycle":4,"time":"2021-08-28T01:19:53+02:00","message":"target is DOWN"}
{"level":"debug","cycle":4,"success ratio":0,"time":"2021-08-28T01:19:53+02:00","message":"cycle finished"}
{"level":"debug","cycle":5,"time":"2021-08-28T01:19:53+02:00","message":"cycle started"}
{"level":"debug","cycle":5,"ratio requested":90,"ratio actual":0,"time":"2021-08-28T01:19:55+02:00","message":"success ratio is below requested value"}
{"level":"debug","cycle":5,"success ratio":0,"time":"2021-08-28T01:19:55+02:00","message":"cycle finished"}
{"level":"debug","cycle":6,"time":"2021-08-28T01:19:55+02:00","message":"cycle started"}
{"level":"debug","cycle":6,"ratio requested":90,"ratio actual":0,"time":"2021-08-28T01:19:57+02:00","message":"success ratio is below requested value"}
{"level":"debug","cycle":6,"success ratio":0,"time":"2021-08-28T01:19:57+02:00","message":"cycle finished"}
{"level":"debug","cycle":7,"time":"2021-08-28T01:19:57+02:00","message":"cycle started"}
{"level":"debug","cycle":7,"ratio requested":90,"ratio actual":0,"time":"2021-08-28T01:19:59+02:00","message":"success ratio is below requested value"}
{"level":"debug","cycle":7,"success ratio":0,"time":"2021-08-28T01:19:59+02:00","message":"cycle finished"}
{"level":"debug","cycle":8,"time":"2021-08-28T01:19:59+02:00","message":"cycle started"}
{"level":"info","cycle":8,"downtime":"0m08s","time":"2021-08-28T01:20:01+02:00","message":"target is UP again"}
{"level":"debug","cycle":8,"success ratio":100,"time":"2021-08-28T01:20:01+02:00","message":"cycle finished"}
{"level":"debug","cycle":9,"time":"2021-08-28T01:20:01+02:00","message":"cycle started"}
{"level":"debug","cycle":9,"success ratio":100,"time":"2021-08-28T01:20:03+02:00","message":"cycle finished"}
{"level":"debug","cycle":10,"time":"2021-08-28T01:20:03+02:00","message":"cycle started"}
{"level":"debug","cycle":10,"success ratio":100,"time":"2021-08-28T01:20:05+02:00","message":"cycle finished"}
{"level":"debug","cycle":11,"time":"2021-08-28T01:20:05+02:00","message":"cycle started"}
{"level":"debug","cycle":11,"success ratio":100,"time":"2021-08-28T01:20:07+02:00","message":"cycle finished"}
{"level":"debug","cycle":12,"time":"2021-08-28T01:20:07+02:00","message":"cycle started"}
{"level":"debug","cycle":12,"success ratio":100,"time":"2021-08-28T01:20:09+02:00","message":"cycle finished"}
{"level":"debug","cycle":13,"time":"2021-08-28T01:20:09+02:00","message":"cycle started"}
{"level":"debug","cycle":13,"ratio requested":90,"ratio actual":0,"time":"2021-08-28T01:20:11+02:00","message":"success ratio is below requested value"}
{"level":"warn","cycle":13,"time":"2021-08-28T01:20:11+02:00","message":"target is DOWN"}
{"level":"debug","cycle":13,"success ratio":0,"time":"2021-08-28T01:20:11+02:00","message":"cycle finished"}
{"level":"debug","cycle":14,"time":"2021-08-28T01:20:11+02:00","message":"cycle started"}
{"level":"debug","cycle":14,"ratio requested":90,"ratio actual":0,"time":"2021-08-28T01:20:13+02:00","message":"success ratio is below requested value"}
{"level":"debug","cycle":14,"success ratio":0,"time":"2021-08-28T01:20:13+02:00","message":"cycle finished"}
{"level":"debug","cycle":15,"time":"2021-08-28T01:20:13+02:00","message":"cycle started"}
{"level":"debug","cycle":15,"ratio requested":90,"ratio actual":0,"time":"2021-08-28T01:20:15+02:00","message":"success ratio is below requested value"}
{"level":"debug","cycle":15,"success ratio":0,"time":"2021-08-28T01:20:15+02:00","message":"cycle finished"}
{"level":"debug","cycle":16,"time":"2021-08-28T01:20:15+02:00","message":"cycle started"}
{"level":"debug","cycle":16,"ratio requested":90,"ratio actual":0,"time":"2021-08-28T01:20:17+02:00","message":"success ratio is below requested value"}
{"level":"debug","cycle":16,"success ratio":0,"time":"2021-08-28T01:20:17+02:00","message":"cycle finished"}
{"level":"debug","cycle":17,"time":"2021-08-28T01:20:17+02:00","message":"cycle started"}
{"level":"debug","cycle":17,"ratio requested":90,"ratio actual":0,"time":"2021-08-28T01:20:19+02:00","message":"success ratio is below requested value"}
{"level":"debug","cycle":17,"success ratio":0,"time":"2021-08-28T01:20:19+02:00","message":"cycle finished"}
{"level":"debug","cycle":18,"time":"2021-08-28T01:20:19+02:00","message":"cycle started"}
{"level":"info","cycle":18,"downtime":"0m10s","time":"2021-08-28T01:20:22+02:00","message":"target is UP again"}
{"level":"debug","cycle":18,"success ratio":100,"time":"2021-08-28T01:20:22+02:00","message":"cycle finished"}
{"level":"debug","cycle":19,"time":"2021-08-28T01:20:22+02:00","message":"cycle started"}
{"level":"debug","cycle":19,"success ratio":100,"time":"2021-08-28T01:20:24+02:00","message":"cycle finished"}
{"level":"debug","cycle":20,"time":"2021-08-28T01:20:24+02:00","message":"cycle started"}
{"level":"debug","cycle":20,"success ratio":100,"time":"2021-08-28T01:20:26+02:00","message":"cycle finished"}
{"level":"debug","cycle":21,"time":"2021-08-28T01:20:26+02:00","message":"cycle started"}
{"level":"debug","cycle":21,"success ratio":100,"time":"2021-08-28T01:20:28+02:00","message":"cycle finished"}
{"level":"debug","cycle":22,"time":"2021-08-28T01:20:28+02:00","message":"cycle started"}
{"level":"info","time":"2021-08-28T01:20:29+02:00","message":"shutdown requested by user"}
{"level":"info","total downtime":"0m18s","time":"2021-08-28T01:20:29+02:00","message":"end statistics"}
{"level":"info","time":"2021-08-28T01:20:29+02:00","message":"shutdown successful"}
```

</details>