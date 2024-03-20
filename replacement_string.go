package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

type StringReplacement struct {
	From string
	To   string
}

func (r StringReplacement) Apply(path string) string {
	dir := filepath.Dir(path)
	base := filepath.Base(path)
	newBase := strings.ReplaceAll(base, r.From, r.To)
	return filepath.Join(dir, newBase)
}

func (r StringReplacement) String() string {
	return fmt.Sprintf("{StringReplacement From:%s To:%s}", r.From, r.To)
}
