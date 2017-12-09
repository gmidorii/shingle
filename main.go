package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := []string{"a", "b", "c", "d", "e"}
	min := 2
	max := 4

	res := initialShingle(min, max, arr)
	fmt.Println(res)
	res = shingle2(min, max, arr)
	fmt.Println(res)
}

func initialShingle(min, max int, arr []string) []string {
	res := []string{}
	if min > max {
		return res
	}

	for i := 0; i < len(arr); i++ {
		cmin := min
		for j := i; j+cmin < len(arr)+1; j++ {
			if len(arr[i:j+cmin]) > max {
				break
			}
			var s string
			for _, v := range arr[i : j+cmin] {
				s += v
			}
			res = append(res, s)
		}
	}
	return res
}

func shingle2(min, max int, arr []string) []string {
	if min > max {
		return []string{}
	}

	// for alloc array
	maxLen := len(arr)
	resLen := 0
	for maxLen > 0 {
		resLen += maxLen
		maxLen--
	}
	res := make([]string, resLen)

	idx := 0
	for i := 0; i < len(arr); i++ {
		cmin := min
		for j := i; j+cmin < len(arr)+1; j++ {
			if len(arr[i:j+cmin]) > max {
				break
			}
			var s string
			for _, v := range arr[i : j+cmin] {
				s += v
			}
			res[idx] = s
			idx++
		}
	}

	return res[0:idx]
}

func shingle3(min, max int, arr []string) []string {
	res := []string{}
	if min > max {
		return res
	}

	for i := 0; i < len(arr); i++ {
		cmin := min
		for j := i; j+cmin < len(arr)+1; j++ {
			if len(arr[i:j+cmin]) > max {
				break
			}
			res = append(res, strings.Join(arr[i:j+cmin], ""))
		}
	}
	return res
}
