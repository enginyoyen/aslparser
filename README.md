# Amazon States Language Parser
Package `aslparser` validates and parses the Amazon States Language, that is used for step functions.
See [https://states-language.net/spec.html](https://states-language.net/spec.html) for details.

[![Build Status](https://travis-ci.com/enginyoyen/aslparser.svg?branch=master)](https://travis-ci.com/enginyoyen/strmetric)


## Installation 
```
go get github.com/enginyoyen/aslparser 
```

## Usage 
First argument is the file content to be validated, while the second argument is whether it should use strict mode for validation
```
stateMachine, err := aslparser.Parse(filePath, true)
if !stateMachine.Valid() {
	for _, e := range stateMachine.Errors(){
		fmt.Print(e.Description())
	}
}
```

Alternatively, `ParseFile` method uses file path to load and validate 

```
stateMachine, err := aslparser.ParseFile(filePath, true)
if !stateMachine.Valid() {
	for _, e := range stateMachine.Errors(){
		fmt.Print(e.Description())
	}
}
```

## Converting JSON Schema to a static file
JSON schema is converted to an static go file to be included as an executable.
```
go-bindata -o state_machine_bin.go schemas/state-machine.json schemas/state-machine-strict-arn.json

```

## JSON Schema 
JSON Schema file modified and original JSON schemas based on [asl-validator](https://github.com/airware/asl-validator) by AirWare (https://www.airware.com/). 
Additionally, test input file are also copied from the repo. 
See [schema files](https://github.com/airware/asl-validator/tree/master/src/schemas)
See [test files](https://github.com/airware/asl-validator/tree/master/src/__tests__/definitions)



## TODO 
- Validate json-path in InputPath, OutputPath and ResultPath
- Extend `StateMachine` to include complete spec


# Licence 
Use of this source code is governed by an MIT license that can be found in the LICENSE file.
