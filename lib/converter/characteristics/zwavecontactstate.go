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

package characteristics

import (
	"errors"
	"github.com/SENERGY-Platform/converter/lib/converter/register"
	"log"
	"reflect"
	"runtime/debug"
)

const ZwaveContactState = "urn:infai:ses:characteristic:65e9a5db-348f-4888-9d95-588741d67dbf"

func init() {
	register.Labels[ZwaveContactState] = "ZwaveContactState"

	register.Add(ContactState, ZwaveContactState, register.NoLosses, func(in interface{}) (out interface{}, err error) {
		switch state := in.(type) {
		case bool:
			if state {
				return "Door/Window Closed", nil
			} else {
				return "Door/Window Open", nil
			}
		default:
			debug.PrintStack()
			log.Println("ERROR: ", reflect.TypeOf(in).String(), in)
			return nil, errors.New("unable to interpret value; input type is " + reflect.TypeOf(in).String())
		}
	})

	register.Add(ZwaveContactState, ContactState, register.NoLosses, func(in interface{}) (out interface{}, err error) {
		switch state := in.(type) {
		case string:
			switch state {
			case "Door/Window Open":
				return false, nil
			case "Door/Window Closed":
				return true, nil
			case "Clear":
				return false, nil
			default:
				return "", errors.New("unable to convert '" + state + "' to generic opening state")
			}
		default:
			debug.PrintStack()
			log.Println("ERROR: ", reflect.TypeOf(in).String(), in)
			return nil, errors.New("unable to interpret value; input type is " + reflect.TypeOf(in).String())
		}
	})
}
