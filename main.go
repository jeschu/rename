package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	cmdline := ParseCommandLine()

	replacements := make(map[string]string)
	for _, path := range cmdline.Files {
		target := path
		for _, replacement := range cmdline.Replacements {
			target = replacement.Apply(target)
		}
		if path != target {
			replacements[path] = target
		}
	}

	if cmdline.Execute {
		for path, target := range replacements {
			fmt.Println(path, "->", filepath.Base(target))
			err := os.Rename(path, target)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "unable to rename '%s': %v", path, err)
			}
		}
		fmt.Printf("renamed %d files.\n", len(replacements))
	} else {
		for path, target := range replacements {
			fmt.Println("mv", path)
			fmt.Println("  ", target)
		}
	}

}
