package main

import (
	"./scan_manager"
	_ "./scanners"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	/* */
	fmt.Println("[+] Looking for interesting environment variables...")
	scan_manager.ScanEnvironmentVariables()

	/* */
	fmt.Println("[+] Looking for interesting files...")
	for _, path := range FindRoots() {
		scan_manager.ScanDirectory(path)
	}
}

func FindRoots() []string {
	var roots []string

	// TODO: figure out how to find all filesystem roots in a system-independent manner
	roots = append(roots, "/")

	return roots
}
