package main

import (
	"fmt"
	"unicode"
)

func main() {
	var pass string
	pass = "aaaabbbbccccddeeddeeddeedd"
	pass = "!"
	pass = "aaaaabbbbbbccccccddeeddeeddeedd"

	fmt.Print(pass, "\n")
	//fmt.Println(repDelTest(pass))
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

//func repDelTest(p string) repDel {
//
//	a := repDel{
//		replaseMust: len(p),
//		deleteMust:  len(p) / 3,
//	}
//	return a
//}

func clusterAnalyzer(p string) clusterStatus {
	//var a []repDel
	//
	//a = append(a, repDel{
	//	replaseMust: len(p),
	//	deleteMust:  len(p) / 3,
	//})
	//
	//b := clusterStatus{
	//	repit:    false,
	//	cluster:  a,
	//	upLowDig: 1,
	//	tail:     2,
	//}
	//return b

	clusterCounter := 1
	var char int32
	l := len(p)

	var up, low, dig int = 1, 1, 1

	var cluster []repDel
	var clusterito clusterStatus

	for i, v := range p {
		if char == v {
			clusterCounter++
		}

		if clusterCounter >= 3 && (char != v || i+1 == l) {

			cluster = append(cluster, repDel{
				replaseMust: clusterCounter / 3,
				deleteMust:  clusterCounter - 2,
			})

			clusterito.repit = true
			clusterito.cluster = cluster
			clusterCounter = 1
			char = 0
		}

		if char != v {
			clusterCounter = 1
		}

		if clusterCounter < 3 {
			char = v
		}

		if up != 0 {
			if unicode.IsUpper(v) {
				up = 0
			}
		}

		if low != 0 {
			if unicode.IsLower(v) {
				low = 0
			}
		}

		if dig != 0 {
			if unicode.IsDigit(v) {
				dig = 0
			}
		}

	}
	clusterito.upLowDig = up + low + dig

	return clusterito
}
