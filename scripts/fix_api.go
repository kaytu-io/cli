package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	count := 0
	var re = regexp.MustCompile(`var res \[\]struct \{\s+([a-zA-Z]+)\s+\}`)
	var re2 = regexp.MustCompile(`struct \{\s+([a-zA-Z]+)\s+\}\s+\` + "`json")

	err := filepath.Walk("./pkg/api/kaytu", func(path string, info fs.FileInfo, err error) error {
		if info == nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}

		c, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		s := re.ReplaceAllString(string(c), `var res []$1`)
		s = re2.ReplaceAllString(s, `$1 `+"`"+`json`)
		err = os.WriteFile(path, []byte(s), os.ModePerm)
		if err != nil {
			return err
		}

		count++
		return nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("fixed %s files\n", count)
}
