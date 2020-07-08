package main

import (
	"flag"
	"fmt"
	"os"
)

type CommandLine struct {
	selectList		[]string
	fromList		[]string
	where			[]string
	splitByList		[]string
	filterBy		string
}

var cmd = &CommandLine{
	selectList: make([]string, 0),
	fromList: make([]string, 0),
	where: make([]string, 0),
	splitByList: make([]string, 0),
	filterBy: "",
}

func (obj *CommandLine) Init() bool {
	obj.ParseQuery()

	if len(obj.selectList) == 1 && obj.selectList[0] == "*" {
		obj.selectList = make([]string, 0)
	}
	if len(obj.fromList) == 0 {
		fmt.Fprintln(os.Stderr, "from field missed")
		return false
	}
	if len(obj.where) != 0 && len(obj.where) != 3 {
		fmt.Fprintln(os.Stderr, "invalid where field")
		return false
	}
	fmt.Println(obj.ToString())
	return true
}

func (obj *CommandLine) ClearTagMap(m map[string]bool) {
	for k, _ := range m {
		m[k] = false
	}
}

func (obj *CommandLine) ParseQuery() {
	flag.Parse()
	tagMap := map[string]bool {
		"isSelect": false,
		"isFrom": false,
		"isWhere": false,
		"isSplitBy": false,
		"isFilterBy": false,
	}
	for i := 0; i < flag.NArg(); i++ {
		if flag.Arg(i) == "select" {
			obj.ClearTagMap(tagMap)
			tagMap["isSelect"] = true
		} else if flag.Arg(i) == "from" {
			obj.ClearTagMap(tagMap)
			tagMap["isFrom"] = true
		} else if flag.Arg(i) == "where" {
			obj.ClearTagMap(tagMap)
			tagMap["isWhere"] = true
		} else if flag.Arg(i) == "splitby" {
			obj.ClearTagMap(tagMap)
			tagMap["isSplitBy"] = true
		} else if flag.Arg(i) == "filterby" {
			obj.ClearTagMap(tagMap)
			tagMap["isFilterBy"] = true
		}

		if tagMap["isSelect"] {
			if flag.Arg(i) != "select" {
				obj.selectList = append(obj.selectList, flag.Arg(i))
			}
		} else if tagMap["isFrom"] {
			if flag.Arg(i) != "from" {
				obj.fromList = append(obj.fromList, flag.Arg(i))
			}
		} else if tagMap["isWhere"] {
			if flag.Arg(i) != "where" {
				obj.where = append(obj.where, flag.Arg(i))
			}
		} else if tagMap["isSplitBy"] {
			if flag.Arg(i) != "splitby" {
				obj.splitByList = append(obj.splitByList, flag.Arg(i))
			}
		} else if tagMap["isFilterBy"] {
			if flag.Arg(i) != "filterby" {
				obj.filterBy = flag.Arg(i)
			}
		}
	}
}

func (obj *CommandLine) ToString() string {
	return fmt.Sprintf("command line strings\n" +
		"select list: %v\n" +
		"from list: %v\n" +
		"where: %v\n" +
		"split by: %v\n" +
		"filter by: %v\n",
		obj.selectList, obj.fromList, obj.where, obj.splitByList, obj.filterBy)
}
