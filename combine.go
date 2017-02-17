// Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func walkRecursive(folderPath, outputPath string) {
	fmt.Println("Combining : " + folderPath)
	files, _ := ioutil.ReadDir(folderPath)
	var counter int32
	for _, f := range files {
		fullpath := filepath.Join(folderPath, f.Name())
		if f.IsDir() {
			walkRecursive(fullpath, outputPath)
		} else {
			if filepath.Ext(f.Name()) != ".txt" {
				continue
			}
			//combine
			combine(fullpath, outputPath)
			counter++
		}
	}
	if counter > 0 {
		count := fmt.Sprint(counter)
		fmt.Println(count + " file(s) combined!")
	}
}

func combine(path, outputPath string) {
	//fmt.Println(" --> Converting " + path + " from " + fromCharset + " to UTF-8 ...")

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	charset := detectEnc(path)
	if charset != "UTF-8" {
		filename := "utf8_" + filepath.Base(path)
		filedir := filepath.Dir(path)
		newFile := filepath.Join(filedir, filename)
		toUTF8(path, newFile, charset)
		file, err = os.Open(newFile)
		check(err)
		defer file.Close()
	}
	b, err := ioutil.ReadAll(file)
	check(err)
	text := string(b)

	//Create outputFile if it is not exist
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		fo, err := os.Create(outputPath)
		check(err)
		defer fo.Close()
	}

	outputFile, err := os.OpenFile(outputPath, os.O_APPEND|os.O_WRONLY, 0600)
	check(err)
	defer outputFile.Close()

	_, err = outputFile.WriteString(text)
	check(err)
}
