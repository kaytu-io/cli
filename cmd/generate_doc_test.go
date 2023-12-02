package cmd

import (
	"github.com/kaytu-io/cli-program/cmd/genv2"
	"github.com/spf13/cobra/doc"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestGenerateDoc(t *testing.T) {
	genv2.KaytuCmd.Execute()

	os.Mkdir("./docs", os.ModePerm)

	err := doc.GenMarkdownTree(genv2.KaytuCmd, "./docs/")
	if err != nil {
		log.Fatal(err)
	}

	kaytuDoc, err := os.ReadFile("./docs/kaytu.md")
	if err != nil {
		log.Fatal(err)
	}

	kaytuDoc = []byte(strings.ReplaceAll(string(kaytuDoc), "## kaytu", "# Kaytu CLI Commands\n\n## kaytu"))

	err = os.WriteFile("./docs/index.md", kaytuDoc, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove("./docs/kaytu.md")
	if err != nil {
		log.Fatal(err)
	}

	err = filepath.Walk("./docs", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		r, err := regexp.Compile("]\\((.*).md\\)")
		if err != nil {
			return err
		}

		contentStr := string(content)
		contentStr = strings.ReplaceAll(contentStr, "## kaytu", "# kaytu")
		contentStr = strings.ReplaceAll(contentStr, "{", "\\{")
		contentStr = strings.ReplaceAll(contentStr, "<", "&lt;")
		contentStr = strings.ReplaceAll(contentStr, ">", "&gt;")
		contentStr = strings.ReplaceAll(contentStr, "\"", "&quot;")
		contentStr = strings.ReplaceAll(contentStr, "'", "&apos;")
		contentStr = strings.ReplaceAll(contentStr, "&", "&amp;")
		contentStr = r.ReplaceAllString(contentStr, "]($1)")
		contentStr = strings.ReplaceAll(contentStr, "](kaytu)", "](.)")

		err = os.WriteFile(path, []byte(contentStr), os.ModePerm)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
