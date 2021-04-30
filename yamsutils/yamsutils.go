package yamsutils

import (
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// 验证文件类型(文件后缀)
func CheckFiletype(path, ftype string) (bool, error) {
	lastindex := strings.LastIndexAny(path, ".")
	if lastindex < 0 {
		return false, errors.New("请选择以.txt结尾的文本文件.\n当前选择路径是:" + path)
	}
	filetype := path[lastindex:]
	if !strings.EqualFold(filetype, ftype) {
		return false, errors.New("请选择以.txt结尾的文本文件.\n当前文件路径是:" + path)
	}
	return true, nil
}

// 三个数组合并，多数组合并这个程序暂时用不到，因此暂且搁置主要是rune的合并
func ArrayMergeAndWordReplace(array0, array1, array2 []rune) (array []rune) {
	array = make([]rune, len(array0)+len(array1)+len(array2))
	copy(array, array0)
	copy(array[len(array0):], array1)
	copy(array[len(array0)+len(array1):], array2)
	return
}

// func ArrayMergeAndWordReplace(array0, array1, array2 []interface{}) (array []interface{}) {
// 	array = make([]interface{}, len(array0)+len(array1)+len(array2))
// 	copy(array, array0)
// 	copy(array[len(array0):], array1)
// 	copy(array[len(array0)+len(array1):], array2)
// 	return
// }

// 用于krc文本内容解压，krc是对原文本进行zip压缩后，进行了异或运算。目前这个方法与kugou使用的zip压缩时使用的偏移量值不同，因此暂且搁置krc转txt文本功能
// 由于go中的zlib是由go团队自己开发的,并不是原始zlib库的，因算法与其他语言的zlib不同，其他语言主要使用C的zlib算法，在go中对zlib算法修正的包是youtube团队开发的cgzip
func UnZip(sourcedata []byte) ([]byte, error) {
	in := bytes.NewReader(sourcedata)
	var out bytes.Buffer
	r, err := zlib.NewReader(in)
	if err != nil {
		return nil, err
	}
	// fmt.Println(r)
	// defer r.Close()

	io.Copy(&out, r)
	fmt.Println(out.Len())
	return out.Bytes(), nil
}

func DoZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	// w := zlib.NewWriter(&in)
	w, _ := zlib.NewWriterLevel(&in, zlib.DefaultCompression)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

func FileToBytes(path string) ([]byte, error) {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	size := stat.Size()
	if size > 0 {
		var buff []byte
		for {
			buffer := make([]byte, 1024)
			length, err := file.Read(buffer)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					return nil, err
				}
			}
			if length == 0 {
				break
			}
			buff = append(buff, buffer[:length]...)
		}
		return buff, nil
	}
	return nil, errors.New("读取失败。")
}
