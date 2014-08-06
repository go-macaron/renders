package renders

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func generateTemplateName(base, path string) string {
	//name := (r[0 : len(r)-len(ext)])
	return filepath.ToSlash(path[len(base)+1:])
}

func file_content(path string) (string, error) {
	// Read the file content of the template
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	s := string(b)

	if len(s) < 1 {
		return "", errors.New("render: template file is empty")
	}

	return s, nil
}

func getExt(s string) string {
	if strings.Index(s, ".") == -1 {
		return ""
	}
	return "." + strings.Join(strings.Split(s, ".")[1:], ".")
}
