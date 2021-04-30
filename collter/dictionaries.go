package collter

import (
	"bufio"
	"io"
	"os"
	"strings"
	"translate/yamsutils"
)

var simpledict map[string]string
var traddict map[string]string

func Fileload(path string) error {
	check, err := yamsutils.CheckFiletype(path, ".txt")
	if !check && err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	size := stat.Size()
	if size > 0 {
		simpledict = make(map[string]string)
		traddict = make(map[string]string)
		buff := bufio.NewReader(file)
		for {
			line, err := buff.ReadString('\n')
			line = strings.TrimSpace(line)
			if len(line) > 0 {
				linerune := []rune(line)
				val := string(linerune[2:])
				if strings.Contains(val, ",") {
					val = strings.Replace(val, ",", "/", -1)
				}
				simpledict[string(linerune[0])] = val
				traddict[string(linerune[1])] = val
			}
			if err != nil {
				if err == io.EOF {
					break
				} else {
					return err
				}
			}
		}
	}
	return nil
}

func GetSimpleDictonaries(val string) string {
	return simpledict[val]
}

func GetTradDictonaries(val string) string {
	return traddict[val]
}
