package main

import (
	"fmt"
	"unicode"
)

func main() {
	var pass string
	pass = "!"
	pass = "!!!!!!!!!.........!!!!!!.!.!.!.!.!"
	pass = "aaaaaabbbbbbccccccddeeddee" // OK
	pass = "aaaaabbbbbcccccddeeddeedde" // OK
	pass = "aaaabbbbccccddeeddeeddeedd" // OK
	pass = "aaabbbcccdeeddeeddeeddeded" // OK

	pass = "aaaaaabbbbbbccccccddeeddee"

	fmt.Print(pass, "\n")
	//fmt.Println(repDelTest(pass))
	fmt.Println(clusterAnalyzer(pass))
	fmt.Println(strongPasswordChecker(pass))

}

type repDel struct {
	replaceMust int
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
//		replaceMust: len(p),
//		deleteMust:  len(p) / 3,
//	}
//	return a
//}

func strongPasswordChecker(password string) int {
	var sheat int
	//
	c := clusterAnalyzer(password)
	if c.repit {
		for i, _ := range c.cluster {

			// replace except last one in cluster
			if c.upLowDig > 0 && (c.cluster[i].deleteMust+2)%3 != 0 {
				r := c.cluster[i].replaceMust
				c.cluster[i].deleteMust = (c.cluster[i].deleteMust+2)/3 - r
				c.cluster[i].replaceMust -= r
				c.upLowDig -= r
				sheat += r
			}

			// replace last for AAA, AAAAAA,
			if c.upLowDig > 0 && c.cluster[i].replaceMust > 1 && (c.cluster[i].deleteMust+2)%3 == 0 {
				var r int
				if c.upLowDig >= c.cluster[i].replaceMust {
					r = c.cluster[i].replaceMust - 1
				} else {
					r = c.upLowDig
				}

				c.cluster[i].deleteMust = (c.cluster[i].deleteMust+2)/3 - r
				c.cluster[i].replaceMust -= r
				c.upLowDig -= r
				sheat += r
			}

			//if c.upLowDig > 0 && c.cluster[i].replaceMust >= 2 {
			//	r := c.cluster[i].replaceMust - 1
			//	c.cluster[i].deleteMust = (c.cluster[i].deleteMust+2)/3 - r
			//	c.cluster[i].replaceMust -= r
			//	c.upLowDig -= r
			//	sheat += r
			//}

			// delete last one
			if c.upLowDig > 0 && c.cluster[i].deleteMust == 1 {
				sheat++
				c.cluster[i].replaceMust = 0
				c.cluster[i].deleteMust = 0
				//c.upLowDig--
				c.tail--
			}

			// delete last two for aaaa [{1 2}]
			if c.upLowDig > 0 && c.cluster[i].deleteMust == 2 {
				c.cluster[i].replaceMust = 0
				c.cluster[i].deleteMust = 0
				sheat += 2
				c.tail -= 2
			}

			// replace all except last one
			// todo: not working for AAA, AAAAAA,...
			//	don't work at all
			//if c.upLowDig == 0 && c.cluster[i].replaceMust >= 2 {
			//	r := c.cluster[i].replaceMust - 1
			//	c.cluster[i].deleteMust = (c.cluster[i].deleteMust+2)/3 - r
			//	c.cluster[i].replaceMust -= r
			//	sheat += r
			//}

			// delete last
			//if c.upLowDig == 0 && c.cluster[i].replaceMust == 1 && c.cluster[i].deleteMust <= c.tail {
			if c.upLowDig == 0 && c.cluster[i].deleteMust != 0 && c.cluster[i].deleteMust <= c.tail {
				sheat += c.cluster[i].deleteMust
				c.tail -= c.cluster[i].deleteMust
				c.cluster[i].deleteMust = 0
				c.cluster[i].replaceMust = 0
			}

		}

		// if not replaced during clusters iteration
		if c.upLowDig != 0 {
			sheat += c.upLowDig
			c.upLowDig = 0
		}

		//	delete tail
		if c.tail > 0 {
			sheat += c.tail
			c.tail = 0
		}
	}
	fmt.Println(c)
	return sheat
}

func clusterAnalyzer(p string) clusterStatus {
	//var a []repDel
	//
	//a = append(a, repDel{
	//	replaceMust: len(p),
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
				replaceMust: clusterCounter / 3,
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

	clusterito.tail = l - max

	return clusterito
}
