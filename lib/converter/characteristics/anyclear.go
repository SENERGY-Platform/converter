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
	"github.com/SENERGY-Platform/converter/lib/converter/register"
	"log"
	"reflect"
	"runtime/debug"
	"strings"
)

const AnyClear = "urn:infai:ses:characteristic:8e520c73-7ee1-4a4b-954b-1c7c55139e27"

func init() {
	register.Labels[AnyClear] = "AnyClear"

	register.Add(Boolean, AnyClear, register.NoLosses, func(in interface{}) (out interface{}, err error) {
		b, ok := in.(bool)
		if !ok {
			debug.PrintStack()
			log.Println("ERROR: ", reflect.TypeOf(in).String(), in)
			return nil, errors.New("unable to interpret value as boolean; input type is " + reflect.TypeOf(in).String())
		}
		if b {
			return "not clear", nil
		} else {
			return "clear", nil
		}
	})

	register.Add(AnyClear, Boolean, register.NoLosses, func(in interface{}) (out interface{}, err error) {
		switch v := in.(type) {
		case string:
			if strings.ToLower(v) == "clear" {
				return false, nil
			} else {
				return true, nil
			}
		default:
			debug.PrintStack()
			return nil, errors.New("unable to interpret value as string; input type is " + reflect.TypeOf(in).String())
		}
	})
}
