package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"time"
)

func main() {
	logPtr := flag.String("log", "openshift_install.log", "openshift_install.log file location")
	flag.Parse()

	logContent, err := ioutil.ReadFile(*logPtr)
	if err != nil {
		log.Fatal(err)
	}

	reLiteral := [7]string{
		`(?m)time="(.*?)"\s+level=info\s+msg="(Creating\sinfrastructure\sresources...)"`,
		`(?m)time="(.*?)"\s+level=info\s+msg="(Waiting\sup\sto\s\d+m\d+s\sfor\sthe\sKubernetes\sAPI\sat.*?)"`,
		`(?m)time="(.*?)"\s+level=info\s+msg="(Waiting\sup\sto\s\d+m\d+s\sfor\sbootstrapping.*?)"`,
		`(?m)time="(.*?)"\s+level=info\s+msg="(Destroying\sthe\sbootstrap.*?)"`,
		`(?m)time="(.*?)"\s+level=info\s+msg="(Waiting\sup\sto\s\d+m\d+s\sfor\sthe\scluster.*?initialize...)"`,
		`(?m)time="(.*?)"\s+level=info\s+msg="(Waiting\sup\sto\s\d+m\d+s\sfor\sthe\sopenshift-console.*?created...)"`,
		`(?m)time="(.*?)"\s+level=info\s+msg="(Install\scomplete!)"`,
	}

	description := [6]string{
		"Infrastructure created",
		"API server available",
		"Bootstrap completed",
		"Bootstrap destroyed",
		"Cluster initialized",
		"Console route created",
	}

	results := [][]string{}
	for i := range reLiteral {
		re := regexp.MustCompile(reLiteral[i])
		result := re.FindStringSubmatch(string(logContent))
		results = append(results, result)
	}

	timestamps := []time.Time{}
	for i, result := range results {
		if len(result) > 0 {
			t2, err := time.Parse(time.RFC3339, result[1])
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
			timestamps = append(timestamps, t2)
		} else {
			fmt.Printf("Result %d: NO REGEX MATCH FOUND\n", i+1)
		}
	}

	if len(timestamps) < 2 {
		fmt.Println("ERROR: Not enoungh regex matches to generate diffs")
		return
	}

	fmt.Printf("Step, Time (s)\n")
	for i := 0; i < len(timestamps)-1; i++ {
		diff := timestamps[i+1].Sub(timestamps[i])
		fmt.Printf("%s, %v\n", description[i], diff.Seconds())
	}
}
