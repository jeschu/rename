package main

type Replacement interface {
	Apply(path string) string
}
