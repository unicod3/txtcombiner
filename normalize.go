package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func normalize(path string) {
	data, _ := ioutil.ReadFile(path)
	input := string(data)
	// Create replacer with pairs as arguments.
	r := strings.NewReplacer(
		".", " . ",
		",", " , ",
		"?", " ? ",
		";", " ; ",
		"-", " - ",
		"+", " + ",
		"(", " ( ",
		")", " ) ",
		"*", " * ",
		"@", " @ ",
		"#", " # ",
		"$", " $ ",
		"%", " % ",
		"=", " = ",
		"<", " < ",
		">", " > ",
		"}", " } ",
		"{", " { ",
		"/", " / ",
		"\\", " \\ ",
		"_", " _ ",
		"|", " | ",
		"]", " ] ",
		"[", " [ ",
		"~", " ~ ")

	input = r.Replace(input)

	filename := "clean_" + filepath.Base(path)
	filedir := filepath.Dir(path)
	newFile := filepath.Join(filedir, filename)

	fo, err := os.Create(newFile)
	check(err)
	fo.WriteString(input)
	defer fo.Close()
}
