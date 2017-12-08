package main

import "fmt"

func main() {
	arr := []string{"a", "b", "c", "d", "e"}
	min := 2
	max := 6
	res := []string{}
	for i := 0; i < len(arr); i++ {
		cmin := min
		for j := i; j+cmin < len(arr)+1; j++ {
			if len(arr[i:j+cmin]) > max {
				break
			}
			res = append(res, merge(arr[i:j+cmin]))
		}
	}

	fmt.Println(res)
}

func merge(arr []string) string {
	var s string
	for _, v := range arr {
		s += v
	}
	return s
}
