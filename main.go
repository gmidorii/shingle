package main

import "fmt"

func main() {
	arr := []string{"a", "b", "c", "d", "e"}
	min := 2
	max := 6

	res := initialShingle(min, max, arr)
	fmt.Println(res)
	res = shingle2(min, max, arr)
	fmt.Println(res)
}

func merge(arr []string) string {
	var s string
	for _, v := range arr {
		s += v
	}
	return s
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
			res = append(res, merge(arr[i:j+cmin]))
		}
	}
	return res
}

func shingle2(min, max int, arr []string) []string {
	if min > max {
		return []string{}
	}

	// for alloc array
	if max > len(arr) {
		max = len(arr)
	}
	maxLen := max - min + 1
	resLen := 0
	for maxLen > 0 {
		resLen += maxLen
		maxLen--
	}
	fmt.Println(resLen)
	res := make([]string, resLen)

	idx := 0
	for i := 0; i < len(arr); i++ {
		cmin := min
		for j := i; j+cmin < len(arr)+1; j++ {
			if len(arr[i:j+cmin]) > max {
				break
			}
			res[idx] = merge(arr[i : j+cmin])
			idx++
		}
	}

	return res
}
