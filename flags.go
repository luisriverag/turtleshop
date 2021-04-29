package main

import "flag"

func flags() {
	flag.BoolVar(&wantsLicense, "license", false, "Show the license")
	flag.Parse()
}
