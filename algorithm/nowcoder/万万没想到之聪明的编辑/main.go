package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	for n > 0 {
		n--
		var str string
		fmt.Scanln(&str)
		b := []byte(str)
		j := 0
		for i := 0; i < len(b); i++ {
			b[j] = b[i]
			j++
			if j >= 3 && b[j-1] == b[j-2] && b[j-2] == b[j-3] {
				j--
			}
			if j >= 4  &&  b[j-1]  == b[j-2] &&  b[j-3] == b[j-4]  {
				j--
			}
		}
		fmt.Println(string(b[:j]))
	}
}
