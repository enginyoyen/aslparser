/*
 * Copyright (c) 2020 Engin Yöyen.
 * Use of this source code is governed by an MIT
 * license that can be found in the LICENSE file.
 */

package aslparser

import (
	"github.com/enginyoyen/aslparser/static"
	"github.com/xeipuuv/gojsonschema"
)

// Loads the state-machine JSON file from provided path
// and validates it against states-language schema
// strict argument defines whether Resource name must be AWS ARN pattern or not
// See https://states-language.net/spec.html
func Validate(payload []byte, strict bool) (*gojsonschema.Result, error) {
	result, err := validateSchema(payload, strict)
	if err != nil {
		return result, err
	}
	return result, nil
}

func validateSchema(payload []byte, strict bool) (*gojsonschema.Result, error) {
	stateMachineSchema, assetError := stateMachineSchema(strict)
	if assetError != nil {
		return nil, assetError
	}
	schemaLoader := gojsonschema.NewStringLoader(string(stateMachineSchema))
	documentLoader := gojsonschema.NewStringLoader(string(payload))

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	return result, err
}

func stateMachineSchema(strict bool) ([]byte, error) {
	if strict {
		return static.Asset("schemas/state-machine-strict-arn.json")
	} else {
		return static.Asset("schemas/state-machine.json")
	}

}