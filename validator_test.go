/*
 * Copyright (c) 2020 Engin Yöyen.
 * Use of this source code is governed by an MIT
 * license that can be found in the LICENSE file.
 */

package aslparser

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

func Test(t *testing.T) {
	fileName, filePath := collectFileInfo()
	r, _ := regexp.Compile("^(valid|invalid)-(.+)\\.json$")
	for i, name := range fileName {
		path := filePath[i]
		subMatch := r.FindStringSubmatch(name)
		if subMatch[1] == "valid" {
			runValidCase(t, path, subMatch[2])
		} else if subMatch[1] == "invalid" {
			runInvalidCase(t, path, subMatch[2])
		}
	}
}

func runValidCase(t *testing.T, path string, name string) {
	payload, _ := ioutil.ReadFile(path)
	validate, _ := Validate(payload, true)
	if !validate.Valid() {
		t.Errorf("Failed validation, where as suppose to pass for input %s", name)
	}
}

func runInvalidCase(t *testing.T, path string, name string) {
	payload, _ := ioutil.ReadFile(path)
	validate, _ := Validate(payload, true)
	if validate.Valid() {
		t.Errorf("Validation pass, where as suppose to fail for input %s", name)
	}
}

func collectFileInfo() ([]string, []string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	wd = filepath.Join(wd, "testdata", "definitions")
	// exclusion list runs in the other test
	var exclude = []string{"invalid-unreachable-state.json", "invalid-inexistant-state.json", "invalid-json-path.json"}
	var filePath []string
	var fileName []string
	fileErr := filepath.Walk(wd, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || filepath.Ext(info.Name()) != ".json" {
			return nil
		}
		for _, s := range exclude {
			if s == info.Name() {
				return nil
			}
		}
		fileName = append(fileName, info.Name())
		filePath = append(filePath, path)
		return nil
	})
	if fileErr != nil {
		panic(fileErr)
	}
	return fileName, filePath
}


func TestStrictMode(t *testing.T) {
	name := "task-alias-function"

	wd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	path := filepath.Join(wd, "testdata", "definitions","invalid-task-alias-function.json")
	payload, _ := ioutil.ReadFile(path)

	strictValidationResult, _ := Validate(payload, true)
	if strictValidationResult.Valid() {
		t.Errorf("ARN validation pass, where as suppose to fail for input %s", name)
	}
	///
	loseValidationResult, _ := Validate(payload, false)
	if !loseValidationResult.Valid() {
		t.Errorf("Failed validation, where as suppose to pass for input %s", name)
	}
}