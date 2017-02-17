// Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	iconv "github.com/djimenez/iconv-go"
	"github.com/saintfish/chardet"
)

func detectEnc(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	arr := []byte(dat)
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(arr)
	check(err)
	res := result.Charset
	return res
}

func toUTF8(path, outpath, charset string) {
	fmt.Println("--> Converting " + path + " from " + charset + " to UTF-8 ...")

	f, err := os.Open(path)
	check(err)
	defer f.Close()

	reader, err := iconv.NewReader(f, charset, "utf-8")
	check(err)

	fo, err := os.Create(outpath)
	check(err)
	defer fo.Close()

	io.Copy(fo, reader)
}
