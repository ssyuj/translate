package collter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var hallmake = "/"

// 半角空格，英文空格
var espace int32 = 32

// 全角空格，中文空格
var cspace int32 = 12288

var Lrc map[string]string
var Plrc map[string]string

func FileMarkWord(path string) {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	size := stat.Size()
	if size > 0 {
		Lrc = make(map[string]string)
		Plrc = make(map[string]string)
		var index int32 = 1
		buff := bufio.NewReader(file)
		for {
			line, err := buff.ReadString('\n')
			line = strings.TrimSpace(line)
			if len(line) > 0 {
				sindex := strconv.Itoa(int(index))
				plrcstr, lrcstr := MarkWord(line, sindex)
				Lrc[sindex] = lrcstr
				if len(strings.TrimSpace(plrcstr)) > 0 {
					pindex := sindex + "_"
					Plrc[pindex] = plrcstr
				}
				index++
			}
			if err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Println(err)
					return
				}
			}
		}
	}
}

func MarkWord(value string, sindex string) (string, string) {
	transmit := make([]byte, len(value))
	copy(transmit, value)
	target := string(transmit)
	roperation := []rune(target)
	// voperation := make([]rune, len(roperation))
	// copy(voperation, roperation)
	var pbuff strings.Builder
	var lbuff strings.Builder
	var sign bool
	for index, value := range roperation {
		// 判断是否是中文
		if unicode.Is(unicode.Han, value) {
			// vrow = make(map[string]string)
			phonetic := SeparateWord(string(value))

			pbuff.WriteString("[")
			if strings.Contains(phonetic, hallmake) {
				pbuff.WriteString("#")
				pbuff.WriteString(sindex + "$" + strconv.Itoa(index) + ":")
				sign = true
			}
			pbuff.WriteString(phonetic)
			pbuff.WriteString("]")
		} else {
			if value == espace || value == cspace {
				pbuff.WriteString(string(value))
			} else if unicode.Is(unicode.Letter, value) {
				pbuff.WriteString(string(espace))
			} else {
				pbuff.WriteString(string(value))
			}
		}
		if sign {
			lbuff.WriteString("[")
			lbuff.WriteString("#" + sindex + "$" + strconv.Itoa(index) + ":")
			lbuff.WriteString(string(value))
			lbuff.WriteString("]")
			sign = false
		} else {
			lbuff.WriteString(string(value))
		}
	}
	vrow := pbuff.String()
	lrow := lbuff.String()
	return vrow, lrow
}

func SeparateWord(val string) string {
	val = strings.TrimSpace(val)
	phonetic := simpledict[val]
	phonetic = strings.TrimSpace(phonetic)
	if phonetic == "" || len(phonetic) <= 0 {
		phonetic = traddict[val]
		phonetic = strings.TrimSpace(phonetic)
	}
	return phonetic
}

func SimpleSeparateWord(val string) string {
	val = strings.TrimSpace(val)
	return strings.TrimSpace(simpledict[val])
}

func TradSeparateWord(val string) string {
	val = strings.TrimSpace(val)
	return strings.TrimSpace(traddict[val])
}

func GetLrc() {

}
