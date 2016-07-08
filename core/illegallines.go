package core

import (
	"fmt"
	"github.com/hpcloud/tail"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	waitGrp sync.WaitGroup
	caches  []string
	keys    = make(map[string]int)
	ctrl    = make(chan struct{})
)

func IllegalLines() {

	caches = []string{"carbon-cache-a", "carbon-cache-b"}

	c := make(chan string)
	// stop the program after 1 minute
	go func() {
		defer close(ctrl)
		select {
		case <-time.After(1 * time.Minute):
			go saveMap()
		}
	}()

	waitGrp.Add(1)
	go updateMap(c)

	// define regex
	r, _ := regexp.Compile("invalid line \\(([^)]+)\\)")

	for _, cache := range caches {
		waitGrp.Add(1)
		go parser(fmt.Sprintf("/var/log/carbon/%s/listener.log", cache), c, r)
	}
	waitGrp.Wait()
	log.Println("Stopped.")
}

func saveMap() {
	f, err := os.Create("/tmp/output.txt")
	if err != nil {
		log.Fatal("Could not create result file.")
	}
	defer f.Close()
	for k, _ := range keys {
		_, err := f.WriteString(k + "\n")
		if err != nil {
			log.Fatal("Writing to file failed.")
		}
	}
	f.Sync()
}

func updateMap(c <-chan string) {
loop:
	for {
		select {
		case found := <-c:
			log.Println(found)
			keys[found]++
		case <-ctrl:
			log.Println("quit")
			break loop
		}
	}
	waitGrp.Done()
}

func parser(file string, c chan<- string, re *regexp.Regexp) {
	log.Printf("Starting parser with file: %s", file)
	t, err := tail.TailFile(file, tail.Config{Follow: true})
	if err != nil {
		log.Fatalf("listen to %s failed", file)
	}
loop:
	for line := range t.Lines {
		select {
		case <-ctrl:
			log.Printf("closing parser %s.", file)
			break loop
		default:
			if res := re.FindAllStringSubmatch(line.Text, -1); res != nil {
				split := strings.Split(res[0][1], " ")
				c <- split[0]
			}
		}
	}
	waitGrp.Done()
}
