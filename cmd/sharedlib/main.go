/*
 * Copyright 2021 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"C"
	"encoding/json"
	"fmt"
	"github.com/SENERGY-Platform/converter/lib/converter"
	"sync"
)

type CastResult struct {
	Err    string      `json:"err,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

var globalConverter *converter.Converter
var converterStartupErr error
var startupOnce = sync.Once{}

func getConverter() (*converter.Converter, error) {
	get := func() {
		globalConverter, converterStartupErr = converter.New()
	}
	startupOnce.Do(get)

	return globalConverter, converterStartupErr
}

func getMarshaledCastResult(result interface{}, err error) *C.char {
	errStr := ""
	if err != nil {
		errStr = err.Error()
	}
	temp, err := json.Marshal(CastResult{Result: result, Err: errStr})
	if err != nil {
		fmt.Println("ERROR:", err)
		return C.CString("")
	}
	return C.CString(string(temp))
}

//export Cast
func Cast(inJson *C.char, from *C.char, to *C.char) (resultJson *C.char) {
	var in interface{}
	err := json.Unmarshal([]byte(C.GoString(inJson)), &in)
	if err != nil {
		return getMarshaledCastResult(nil, err)
	}
	conv, err := getConverter()
	if err != nil {
		return getMarshaledCastResult(nil, err)
	}
	result, err := conv.Cast(in, C.GoString(from), C.GoString(to))
	return getMarshaledCastResult(result, err)
}

func main() {}
