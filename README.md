# ocp-install-scrape

Simple OpenShift log file scraper to measure the duration of the installation steps.

## Usage
```
[user@host ocp-install-scrape]$ go run main.go -h
Usage of /tmp/go-build829195785/b001/exe/main:
  -log string
    	openshift_install.log file location (default "openshift_install.log")
```

## Output
```
[user@host ocp-install-scrape]$ go run main.go -log openshift_install-new.log
Step, Time (s)
Infrastructure created, 409
API server available, 435
Bootstrap completed, 269
Bootstrap destroyed, 48
Cluster initialized, 608
Console route created, 0
```
