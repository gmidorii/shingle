package main

import (
	"fmt"
	"testing"
)

type in struct {
	min int
	max int
	arr []string
}

var cases = []struct {
	in  in
	out []string
}{
	{
		in:  in{min: 2, max: 1, arr: []string{"a", "b", "c", "d", "e"}},
		out: []string{},
	},
	{
		in:  in{min: 2, max: 2, arr: []string{"a", "b", "c", "d", "e"}},
		out: []string{"ab", "bc", "cd", "de"},
	},
	{
		in:  in{min: 2, max: 3, arr: []string{"a", "b", "c", "d", "e"}},
		out: []string{"ab", "abc", "bc", "bcd", "cd", "cde", "de"},
	},
	{
		in:  in{min: 2, max: 5, arr: []string{"a", "b", "c", "d", "e"}},
		out: []string{"ab", "abc", "abcd", "abcde", "bc", "bcd", "bcde", "cd", "cde", "de"},
	},
	{
		in:  in{min: 2, max: 100, arr: []string{"a", "b", "c", "d", "e"}},
		out: []string{"ab", "abc", "abcd", "abcde", "bc", "bcd", "bcde", "cd", "cde", "de"},
	},
}

func equal(t *testing.T, res, exp []string) {
	if len(res) != len(exp) {
		t.Errorf("failed diff len\n e: %d\n a: %d\n", len(exp), len(res))
		return
	}

	for _, v := range exp {
		if !contains(v, res) {
			t.Errorf("unexpected value\n e: %s\n a", v)
			return
		}
	}
}

func contains(str string, arr []string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}

func initial() []string {
	s := "aa"
	arr := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = s
	}
	return arr
}

func TestInitialShingle(t *testing.T) {
	for _, c := range cases {
		res := initialShingle(c.in.min, c.in.max, c.in.arr)
		equal(t, res, c.out)
	}
}

func BenchmarkInitialShingle(b *testing.B) {
	arr := initial()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		initialShingle(2, 100, arr)
	}
}

func TestShingle2(t *testing.T) {
	for _, c := range cases {
		res := shingle2(c.in.min, c.in.max, c.in.arr)
		fmt.Println(res)
		equal(t, res, c.out)
	}
}

func BenchmarkShingle2(b *testing.B) {
	arr := initial()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shingle2(2, 100, arr)
	}
}
