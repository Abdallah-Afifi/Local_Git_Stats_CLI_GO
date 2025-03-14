package main

import (
	"flag"
)

// scan a path to crawl it and its sub-directories
// search for git repos
func scan(path string) {
	print("scan")
}

// stats to generate a graph of local git contributions
func stats(email string) {
	print("stats")
}

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "add a folder to scan")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}
	stats(email)

}
