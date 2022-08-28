package main

import (
	"fmt"
)

func main() {
	var pass string
	pass = "aaaabbbbccccddeeddeeddeedd"
	pass = "aaaaaabbbbbbccccccddeeddeeddeedd"

	fmt.Print(pass, "\n")
	fmt.Println(repDelTest(pass))
	fmt.Println(clusterAnalyzer(pass))
}

type repDel struct {
	replaseMust int
	deleteMust  int
}

type clusterStatus struct {
	repit    bool
	cluster  []repDel
	upLowDig int
	tail     int
}

const (
	min = 6
	max = 20
)

func repDelTest(p string) repDel {

	a := repDel{
		replaseMust: len(p),
		deleteMust:  len(p) / 3,
	}
	return a
}

func clusterAnalyzer(p string) clusterStatus {
	var a []repDel

	a = append(a, repDel{
		replaseMust: len(p),
		deleteMust:  len(p) / 3,
	})

	b := clusterStatus{
		repit:    false,
		cluster:  a,
		upLowDig: 1,
		tail:     2,
	}
	return b
}
