package utils

import (
	"bufio"
	"os"
	"strings"
)

func ScanfString() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	str = strings.TrimSpace(str)

	return str, nil
}
