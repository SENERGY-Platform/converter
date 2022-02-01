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

package timestamp

import (
	"errors"
	"log"
	"reflect"
	"runtime/debug"
	"strconv"
	"time"
)

const UnixMilliSecondsId = "urn:infai:ses:characteristic:64691f8d-4909-470f-a1fa-e977ebe28684"
const UnixMilliSecondsName = "unix_format_milliseconds"

func init() {
	conceptToCharacteristic.Set(UnixMilliSecondsId, func(concept interface{}) (out interface{}, err error) {
		iso, ok := concept.(string)
		if !ok {
			log.Println("ERROR: conceptToCharacteristic", UnixMilliSecondsName, reflect.TypeOf(concept).String(), concept)
			return nil, errors.New("unable to interpret value; input type is " + reflect.TypeOf(concept).String())
		}
		t, err := time.Parse(time.RFC3339, iso)
		if err != nil {
			debug.PrintStack()
			log.Println("ERROR: ", err)
			return nil, err
		}
		return t.UnixMilli(), nil
	})

	characteristicToConcept.Set(UnixMilliSecondsId, func(in interface{}) (concept interface{}, err error) {
		var milliseconds int64
		switch m := in.(type) {
		case int:
			milliseconds = int64(m)
		case int32:
			milliseconds = int64(m)
		case int64:
			milliseconds = m
		case float32:
			milliseconds = int64(m)
		case float64:
			milliseconds = int64(m)
		case string:
			milliseconds, err = strconv.ParseInt(m, 10, 64)
			if err != nil {
				debug.PrintStack()
				log.Println("ERROR: ", err)
				return nil, err
			}
		default:
			debug.PrintStack()
			log.Println("ERROR: characteristicToConcept", UnixMilliSecondsName, reflect.TypeOf(in).String(), in)
			return nil, errors.New("unable to interpret value; input type is " + reflect.TypeOf(in).String())
		}
		return time.UnixMilli(milliseconds).Format(time.RFC3339), nil
	})
}
