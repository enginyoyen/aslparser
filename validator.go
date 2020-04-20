/*
 * Copyright (c) 2020 Engin Yöyen.
 * Use of this source code is governed by an MIT
 * license that can be found in the LICENSE file.
 */

package aslparser

import "github.com/xeipuuv/gojsonschema"

// Loads the state-machine JSON file from provided path
// and validates it against states-language schema
// See https://states-language.net/spec.html
func Validate(payload []byte) (*gojsonschema.Result, error) {
	result, err := validateSchema(payload)
	if err != nil {
		return result, err
	}
	return result, nil
}

func validateSchema(payload []byte) (*gojsonschema.Result, error) {
	stateMachineSchema, assetError := stateMachineSchema()
	if assetError != nil {
		return nil, assetError
	}
	schemaLoader := gojsonschema.NewStringLoader(string(stateMachineSchema))
	documentLoader := gojsonschema.NewStringLoader(string(payload))

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	return result, err
}

func stateMachineSchema() ([]byte, error) {
	return Asset("schemas/state-machine.json")
}
