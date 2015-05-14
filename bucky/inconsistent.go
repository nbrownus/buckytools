package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

// import "github.com/jjneely/buckytools/hashing"

func init() {
	usage := "[options]"
	short := "Find metrics not in correct locations."
	long := `Metrics are located according to the hash ring, this finds metrics that
are not in the correct locations.

Search the entire cluster unless -s is used and check that the consistent hash
ring matches the location where the metric is found.  Print to STDOUT
server:metric for each metric that is in the wrong location.  The server
is the server the metric is presently found on.

Use bucky rebalance to correct.`

	c := NewCommand(inconsistentCommand, "inconsistent", usage, short, long)
	SetupHostname(c)
	SetupSingle(c)
	SetupJSON(c)

	c.Flag.BoolVar(&listForce, "f", false,
		"Force the remote daemons to rebuild their cache.")
}

func InconsistentMetrics(hostports []string) (map[string][]string, error) {
	var list map[string][]string
	var err error
	list, err = ListAllMetrics(hostports, listForce)
	if err != nil {
		log.Printf("Error retrieving metric lists: %s", err)
		return nil, err
	}

	results := make(map[string][]string)
	ring := buildHashRing(GetRings())
	log.Printf("Hashing...")
	t := time.Now().Unix()
	for server, metrics := range list {
		for _, m := range metrics {
			if strings.HasPrefix(m, "carbon.agents.") {
				// These metrics are inserted into the stream after hashing
				// is done.  They will never be consistent and shouldn't be.
				continue
			}
			if ring.GetNode(m).Server != server {
				results[server] = append(results[server], m)
			}
		}
	}
	log.Printf("Hashing time was: %ds", time.Now().Unix()-t)

	// sort for sanity
	for server, metrics := range results {
		log.Printf("%d inconsistent metrics found on %s", len(metrics), server)
		sort.Strings(metrics)
	}

	return results, nil
}

// inconsistentCommand runs this subcommand.
func inconsistentCommand(c Command) int {
	servers := GetAllBuckyd()
	if servers == nil {
		return 1
	}

	var err error
	results, err := InconsistentMetrics(servers)
	if JSONOutput {
		blob, err := json.Marshal(results)
		if err != nil {
			log.Printf("%s", err)
		} else {
			os.Stdout.Write(blob)
			os.Stdout.Write([]byte("\n"))
		}
	} else {
		for server, metrics := range results {
			for _, m := range metrics {
				fmt.Printf("%s: %s\n", server, m)
			}
		}
	}

	if err != nil {
		return 1
	}
	return 0
}