package file

import (
	"bufio"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	logger "aoc/internal/utils/logger"
)

func ReadFile(filename string) ([]string) {
	log := logger.New("aoc2021", "ReadFile")
	path := absPathify(filename)
	log.Info(path)
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal("failed to open file", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := []string{}

	for scanner.Scan() {
		nextLine := scanner.Text()
		lines = append(lines, nextLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("failed to read file", err)
	}

	return lines
}

func absPathify(inPath string) string {

	if inPath == "$HOME" || strings.HasPrefix(inPath, "$HOME"+string(os.PathSeparator)) {
		inPath = userHomeDir() + inPath[5:]
	}

	inPath = os.ExpandEnv(inPath)

	if filepath.IsAbs(inPath) {
		return filepath.Clean(inPath)
	}

	p, err := filepath.Abs(inPath)
	if err == nil {
		return filepath.Clean(p)
	}

	return ""
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}