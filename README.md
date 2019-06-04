# ocp-install-scrape

Simple OpenShift log file scraper to measure the duration of the installation steps.

## Usage
```
[user@host ocp-install-scrape]$ ./ocp-install-scrape -h
Usage of ./ocp-install-scrape:
  -log string
       openshift_install.log file location (default "openshift_install.log")
```

## Output
```
[user@host ocp-install-scrape]$ ./ocp-install-scrape -log openshift_install-new.log
Step, Time (s)
Infrastructure created, 409
API server available, 435
Bootstrap completed, 269
Bootstrap destroyed, 48
Cluster initialized, 608
Console route created, 0
```
