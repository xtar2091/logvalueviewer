package main

import (
	"fmt"
	"strings"
)

type Consumer struct {
	logLine				string
	selectFieldsMap 	map[string]string
	whereFilter			WhereFilter
}

var consumer = &Consumer{}

func (obj *Consumer) Start() {
	obj.init()
	for line := range task.WorkList {
		obj.logLine = line
		obj.resetSelectFieldsMap()
		if obj.processSplit(line) == false {
			continue
		}
		if obj.processWhere() == false {
			continue
		}
		obj.processOutput()
	}
	task.Stop()
}

func (obj *Consumer) init() {
	obj.selectFieldsMap = make(map[string]string, 0)
	for _, field := range cmd.selectList {
		obj.selectFieldsMap[field] = ""
	}
	obj.selectFieldsMap[cmd.where[0]] = ""
	obj.whereFilter = CreateWhereFilter(cmd.where[1])
}

func (obj *Consumer) resetSelectFieldsMap() {
	for key, _ := range obj.selectFieldsMap {
		obj.selectFieldsMap[key] = ""
	}
}

/*
将line按照splitby给定的字符分割，select字段存入selectFieldsMap
 */
func (obj *Consumer) processSplit(line string) bool {
	for _, separator := range cmd.splitByList {
		line = strings.ReplaceAll(line, separator, "\t")
	}
	blocks := strings.Split(line, "\t")
	ok := false
	ind := 0
	fieldCount := 0
	for _, block := range blocks {
		ind += 1
		_, ok = obj.selectFieldsMap[block]
		if ok {
			fieldCount += 1
			obj.selectFieldsMap[block] = blocks[ind]
		}
	}

	return fieldCount == len(obj.selectFieldsMap)
}

/*
进行where条件过滤。目前只支持一个条件，且过滤字段必须存在于select中
 */
func (obj *Consumer) processWhere() bool {
	return obj.whereFilter.DoFilter(obj.selectFieldsMap[cmd.where[0]], cmd.where[2])
}

/*
输出
 */
func (obj *Consumer) processOutput() {
	if len(cmd.selectList) == 0 {
		fmt.Println(obj.logLine)
	} else {
		fmt.Print(cmd.selectList[0], "\t", obj.selectFieldsMap[cmd.selectList[0]])
		for i := 1; i < len(cmd.selectList); i++ {
			fmt.Print("\t", cmd.selectList[i], "\t", obj.selectFieldsMap[cmd.selectList[i]])
		}
		fmt.Println()
	}
}
