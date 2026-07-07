package grid

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func GetAbsPath(path string) (string, error) {

	abs, err := filepath.Abs(path)

	if err != nil {
		return "", err
	}

	return abs, nil

}

// Check the file is not too large.
func CheckFileSizeOK(abs_filepath string) (bool, error) {

	fileStat, err := os.Stat(abs_filepath)

	if err != nil {
		return false, err
	}

	if int(fileStat.Size()) > MAX_INPUTFILE_SIZE_INBYTES {
		return false, errors.New("File too large. Change grid/setup.go to allow more memory usage.")
	}

	return true, nil
}

// Read the whole file and return the contents as string array
func ReadFileToStringRows(abs_filepath string) ([]string, error) {

	_, err := CheckFileSizeOK(abs_filepath)

	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(abs_filepath)

	if err != nil {
		return nil, err
	}

	content := string(data)

	lines := strings.Split(content, "\n")

	return lines, nil

}
