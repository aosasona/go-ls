package gols

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

const (
	RESET = "\033[0m"
	CYAN  = "\033[36m"
)

type Config struct {
	ShowHidden bool
	Path       string
}

func Run(config Config) error {
	if !dirExists(config.Path) {
		return fmt.Errorf("%v does not exist or is not a directory", config.Path)
	}

	files, err := os.ReadDir(config.Path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if len(file.Name()) > 0 {
			if file.Name()[:1] == "." && !config.ShowHidden {
				continue
			}
		}

		output := fmt.Sprintf("%v", file.Name())
		if file.IsDir() {
			output = CYAN + output + RESET
		}

		fmt.Println(output)
	}

	return nil
}

func dirExists(dir string) bool {
	result, err := os.Stat(dir)
	if errors.Is(fs.ErrNotExist, err) || !result.IsDir() {
		return false
	}
	return true
}
