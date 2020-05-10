/*
 * Copyright (c) 2020 Engin Yöyen.
 * Use of this source code is governed by an MIT
 * license that can be found in the LICENSE file.
 */

package aslparser

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	wd = filepath.Join(wd, "testdata", "definitions")
	unreachableState := filepath.Join(wd, "invalid-unreachable-state.json")
	inexistantState := filepath.Join(wd, "invalid-inexistant-state.json")
	//jsonPath := filepath.Join(wd, "invalid-json-path.json")

	runInvalidParseCase(t, unreachableState, "unreachable-state")
	runInvalidParseCase(t, inexistantState, "inexistant-state")
	//runInvalidParseCase(t, jsonPath,"json-path")
}

func runInvalidParseCase(t *testing.T, path string, name string) {
	stateMachine, _ := ParseFile(path, true)
	if stateMachine.Valid() {
		t.Errorf("Validation passed, where as suppose to fail for input %s", name)
	}
}
