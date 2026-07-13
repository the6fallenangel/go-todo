package main

import "strings"

type tagList []string

func (t *tagList) String() string {
	return strings.Join(*t, ",")
}

func (t *tagList) Set(value string) error {
	*t = append(*t, value)
	return nil
}
