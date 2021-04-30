package collter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func GetFileContent(path string) (string, error) {
	lrc := ""
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return lrc, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	size := stat.Size()
	if size > 0 {
		buff := bufio.NewReader(file)
		for {
			line, err := buff.ReadString('\n')
			line = strings.TrimSpace(line)
			if len(line) > 0 {
				lrc = lrc + line + "\n"
			}
			if err != nil {
				if err == io.EOF {
					break
				} else {
					return lrc, err
				}
			}
		}
	}
	return lrc, nil
}
