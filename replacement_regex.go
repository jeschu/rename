package main

import (
	"fmt"
	"regexp"
)

type RegexReplacement struct {
	Regex *regexp.Regexp
	To    string
}

func (r RegexReplacement) Apply(path string) string {
	return r.Regex.ReplaceAllString(path, r.To)
}

func (r RegexReplacement) String() string {
	return fmt.Sprintf("{RegexReplacement Regex:%s To:%s}", r.Regex, r.To)
}
