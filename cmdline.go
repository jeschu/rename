package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type Cmdline struct {
	Recursive    bool
	Execute      bool
	Files        []string
	Replacements []Replacement
}

func ParseCommandLine() Cmdline {
	var (
		recursive    bool
		hidden       bool
		execute      bool
		replacements = make([]Replacement, 0)
		files        = make([]string, 0)
	)
	args := os.Args
	idx := 0
	for idx < len(args)-1 {
		idx++
		arg := args[idx]
		switch arg {
		case "-R":
			recursive = true
			break
		case "-e":
			execute = true
			break
		case "-h":
			printHelp(os.Stdout)
			break
		case "-hidden":
			hidden = true
			break
		case "-s":
			idx++
			from, to, err := parseReplacement(args[idx])
			if err != nil {
				fmt.Println(err)
				printHelp(os.Stdout)
				os.Exit(1)
			}
			replacements = append(replacements, StringReplacement{From: from, To: to})
			break
		case "-r":
			idx++
			from, to, err := parseReplacement(args[idx])
			if err != nil {
				fmt.Println(err)
				printHelp(os.Stdout)
				os.Exit(1)
			}
			regex, err := regexp.Compile(from)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			replacements = append(replacements, RegexReplacement{Regex: regex, To: to})
			break
		default:
			for ; idx < len(args); idx++ {
				files = append(files, args[idx])
			}
		}
	}

	files = parseFiles(recursive, hidden, files)

	return Cmdline{
		Recursive:    recursive,
		Execute:      execute,
		Files:        files,
		Replacements: replacements,
	}
}

func parseReplacement(arg string) (string, string, error) {
	parts := strings.SplitN(arg, "=", 2)
	if len(parts) != 2 {
		return "", "", errors.New(fmt.Sprintf("unable to parse replacement '%s'", arg))
	}
	return parts[0], parts[1], nil
}

func parseFiles(recursive bool, hidden bool, roots []string) []string {
	files := make([]string, 0)
	for _, root := range roots {
		err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() && (!recursive || (!hidden && d.Name()[0] == '.')) {
				return filepath.SkipDir
			}
			if !d.IsDir() && (hidden || d.Name()[0] != '.') {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "unable to scan '%s': %v", root, err)
		}
	}
	sort.Strings(files)
	return files
}

func printHelp(out io.Writer) {

}
