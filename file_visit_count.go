 /*cat log.txt
learngoprogramming.com 10
learngoprogramming.com 10
golang.org 4
golang.org 96
blog.golang.org 20
blog.golang.org 10

*/
 //go run string.go<log.txt 
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	//wd1 := bufio.NewScanner(os.Stdin)
	//wd1.Split(bufio.ScanWords)
	var (
		//lst      = []string{}
		lst          []string
		total, lines int
		dom_dict     = map[string]int{}
	)
	//sum:=0
	scan := bufio.NewScanner(os.Stdin)
	if err := scan.Err(); err != nil {
		fmt.Println("scan file  err", err)
	}
	for scan.Scan() {
		lines++
		slc := strings.Fields(scan.Text())

		domain := slc[0]
		visits, err := strconv.Atoi(slc[1])
		if err != nil {
			fmt.Printf("error %s at line %d\n", err, lines)
			return // terminate the program
		}
		//fmt.Println(dom_dict[domain], visits)
		if _, ok := dom_dict[domain]; !ok {
			lst = append(lst, domain)
		}

		dom_dict[domain] += visits

	}
	sort.Strings(lst)
	fmt.Println(lst, dom_dict)

	for _, k := range lst {
		fmt.Printf("%10s %10d\n", k, dom_dict[k])
		total += dom_dict[k]

	}
	fmt.Printf("%-10s %10d\n", "total", total)

}
