package util

import (
	"os"
	"fmt"
	// "sync"
	"path"
	"regexp"
	// "time"
	"strings"
	// "math/rand"
	"io/ioutil"
	"encoding/csv"
	"encoding/json"
	"../config"
)

func Split(s string) []string {
	r := csv.NewReader(strings.NewReader(s))
	r.Comma = ' ' // space
	fields, _ := r.Read()
    return fields
}

func Strip(str, chr string) string {
    return strings.Map(func(r rune) rune {
        if strings.IndexRune(chr, r) < 0 {
            return r
        }
        return -1
    }, str)
}

func Contains(input string, words []string) (bool, string) {
	for _, word := range words {
		if strings.Index(input, word) > -1 {
			return true, word
			break
		}
	}
	return false, ""
}

func SearchRegex(words []string, input string) []string {
	s := []string{}
	for _, v := range words {
		pattern := fmt.Sprintf("(?i)%s", input)
		if match, _ := regexp.MatchString(pattern, v); match {
			s = append(s, v)
		}
	}
	return s
}

func FormatFolderName(name string) (string, string) {
	var s []string
	if ok := strings.Contains(name, " - "); ok {
		s = strings.Split(name, " - ")
	} else {
		s = []string{name, ""}
	}
	return s[0], s[1]
}

func FormatPath(str []string) string {
	return path.Join(str...)
}

func CreateFolder(folders []string) bool {
	_path := path.Join(folders...)
	if _, err := os.Stat(_path); os.IsNotExist(err) {
		os.MkdirAll(_path, os.ModePerm)
	} else {
		return false
	}
	return true
}

func GetDirectoryList(params ...string) []string {
	directory := config.C_BASE
	if len(params) > 0 {
		directory = params[0]
	} 
	var list []string
	items, _ := ioutil.ReadDir(directory)
    for _, item := range items {
        if item.IsDir() {
			list = append(list, item.Name())
		}
	}

	return list
}

func GetFileList(params ...string) []string {
	directory := config.C_BASE
	if len(params) > 0 {
		directory = params[0]
	} 
	var list []string
	items, _ := ioutil.ReadDir(directory)
    for _, item := range items {
        if !item.IsDir() {
			list = append(list, item.Name())
		}
	}

	return list
}

func Save(directory string, file string, result interface{}) bool {
	data, _ := json.Marshal(result)
	_path := path.Join(directory, file)
	err := ioutil.WriteFile(_path, data, 0644)
	if err != nil {
		return false
	}
	return true
}

func Load(directory string, file string) interface{} {
	var _content interface{}
	_path := path.Join(directory, file)
	content, _ := ioutil.ReadFile(_path)
	json.Unmarshal(content, &_content)
	return _content
}