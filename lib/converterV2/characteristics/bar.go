/*
 * Copyright 2019 InfAI (CC SES)
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

package characteristics

import (
	"errors"
	"github.com/SENERGY-Platform/converter/lib/converterV2/register"
	"log"
	"reflect"
	"runtime/debug"
)

const Bar = "urn:infai:ses:characteristic:2871c1c4-1c8c-4169-81f6-fb5313df08b1"

// 1 bar = 1000 hPa

func init() {
	register.Labels[Bar] = "Bar"

	register.Add(HPa, Bar, register.NoLosses, func(in interface{}) (out interface{}, err error) {
		switch hpa := in.(type) {
		case int:
			return float64(hpa / 1000.0), nil
		case int32:
			return float64(hpa / 1000.0), nil
		case int64:
			return float64(hpa / 1000.0), nil
		case float32:
			return float64(hpa / 1000.0), nil
		case float64:
			return float64(hpa / 1000.0), nil
		default:
			debug.PrintStack()
			log.Println("ERROR: ", reflect.TypeOf(in).String(), in)
			return nil, errors.New("unable to interpret value; input type is " + reflect.TypeOf(in).String())
		}
	})

	register.Add(Bar, HPa, register.NoLosses, func(in interface{}) (out interface{}, err error) {
		switch bar := in.(type) {
		case int:
			return float64(bar * 1000.0), nil
		case int32:
			return float64(bar * 1000.0), nil
		case int64:
			return float64(bar * 1000.0), nil
		case float32:
			return float64(bar * 1000.0), nil
		case float64:
			return float64(bar * 1000.0), nil
		default:
			debug.PrintStack()
			log.Println("ERROR: ", reflect.TypeOf(in).String(), in)
			return nil, errors.New("unable to interpret value; input type is " + reflect.TypeOf(in).String())
		}
	})

}
