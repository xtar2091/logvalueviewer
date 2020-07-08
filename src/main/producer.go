package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
)

type Producer struct {
	regex *regexp.Regexp
}

var producer = &Producer{}

func (obj *Producer) Init() {
	obj.regex = nil
	if cmd.filterBy != "" {
		obj.regex = regexp.MustCompile(cmd.filterBy)
	}
}

func (obj *Producer) Start() {
	fs := &FileSystem{}
	for _, path := range cmd.fromList {
		if fs.IsFile(path) {
			obj.ProcessFile(path)
		} else {
			fmt.Fprintln(os.Stderr, "only file supported, path:", path)
			continue
		}
	}
	task.CloseWorkList()
}

func (obj *Producer) ProcessFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file failed, file:", fileName, ", error:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		line = bytes.Trim(line, "\n")
		obj.ProcessLine(string(line))
	}
}

func (obj *Producer) ProcessLine(line string) {
	ok := true
	if obj.regex != nil {
		ok = obj.regex.MatchString(line)
	}
	if ok {
		ok = !task.Push(line)
	}
}
