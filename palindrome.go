package main

import "strings"

type Palindrome struct {
}

func (p *Palindrome) reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func (p *Palindrome) Check(s string) bool {
	s = strings.ToLower(s)
	if s == p.reverse(s) {
		return true
	}
	return false
}
