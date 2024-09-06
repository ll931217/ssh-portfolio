package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/charmbracelet/glamour"
)

type page struct {
	name     string
	path     string
	contents string
	rendered string
}

type pages map[int]page

func sortPages(files []os.DirEntry) {
	sort.Slice(files, func(i, j int) bool {
		// welcome.md comes first
		if files[i].Name() == "welcome.md" {
			return true
		}

		// Sort by name
		return files[i].Name() < files[j].Name()
	})
}

func validatePages() error {
	markdownDir, err := filepath.Abs("markdown")
	if err != nil {
		return err
	}

	if _, err := os.Stat(markdownDir); os.IsNotExist(err) {
		return errors.New("markdown directory does not exist")
	}

	files, err := os.ReadDir(markdownDir)
	if err != nil {
		return fmt.Errorf("unable to list files in %s: %w", markdownDir, err)
	}

	hasWelcome := false
	for _, f := range files {
		if f.Name() == "welcome.md" {
			hasWelcome = true
			break
		}
	}

	if !hasWelcome {
		return errors.New("no welcome.md file found")
	}

	return nil
}

func loadPages() (pages, error) {
	pages := make(map[int]page)

	markdownDir, err := filepath.Abs("markdown")
	if err != nil {
		return pages, err
	}

	files, err := os.ReadDir(markdownDir)
	if err != nil {
		return pages, fmt.Errorf("unable to list files in %s: %w", markdownDir, err)
	}

	sortPages(files)

	for i, f := range files {
		filepath := filepath.Join(markdownDir, f.Name())
		contents, err := os.ReadFile(filepath)
		if err != nil {
			return pages, fmt.Errorf("unable to read file %s: %w", filepath, err)
		}

		rendered, err := glamour.Render(string(contents), "dark")
		if err != nil {
			return pages, fmt.Errorf("unable to render markdown: %w", err)
		}

		pages[i] = page{
			name:     f.Name(),
			path:     filepath,
			contents: string(contents),
			rendered: rendered,
		}
	}

	return pages, nil
}
