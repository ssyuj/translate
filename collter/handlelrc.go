package collter

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Loadphonetic(path string) string {
	var plrc strings.Builder
	FileMarkWord(path)
	if Lrc != nil {
		for i := 1; i <= len(Lrc); i++ {
			if Plrc != nil {
				pv := Plrc[strconv.Itoa(int(i))+"_"]
				if len(strings.TrimSpace(pv)) > 0 {
					plrc.WriteString(pv + "\n")
				}
			}
			plrc.WriteString(Lrc[strconv.Itoa(int(i))] + "\n")
		}
	}
	return plrc.String()
}

func Savephonetic(path, content string) bool {
	var savepath string
	bytepath := []byte(path)
	bytesavepath := make([]byte, len(bytepath))
	copy(bytesavepath, bytepath)
	savepath = string(bytesavepath)
	savepath = savepath[:strings.LastIndex(savepath, ".")] + "_tmp" + savepath[strings.LastIndex(savepath, "."):]
	tmplrc, err := os.Create(savepath)
	if err != nil {
		panic(err)
	}
	defer tmplrc.Close()

	writebuffer := bufio.NewWriter(tmplrc)
	if err != nil {
		panic(err)
	}
	for _, v := range content {
		writebuffer.WriteString(string(v))
	}
	writebuffer.Flush()
	return true
}
