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

const UnixSecondsId = "urn:infai:ses:characteristic:8b9411f3-bf56-4e57-8fd4-728bdcd451cd"
const UnixSecondsName = "unix_format_seconds"

func init() {
	conceptToCharacteristic.Set(UnixSecondsId, func(concept interface{}) (out interface{}, err error) {
		iso, ok := concept.(string)
		if !ok {
			log.Println("ERROR: ", reflect.TypeOf(concept).String(), concept)
			return nil, errors.New("unable to interpret value; input type is " + reflect.TypeOf(concept).String())
		}
		t, err := time.Parse(time.RFC3339, iso)
		if err != nil {
			debug.PrintStack()
			log.Println("ERROR: ", err)
			return nil, err
		}
		return t.Unix(), nil
	})

	characteristicToConcept.Set(UnixSecondsId, func(in interface{}) (concept interface{}, err error) {
		var seconds int64
		switch s := concept.(type) {
		case int:
			seconds = int64(s)
		case int32:
			seconds = int64(s)
		case int64:
			seconds = s
		case float32:
			seconds = int64(s)
		case float64:
			seconds = int64(s)
		case string:
			seconds, err = strconv.ParseInt(s, 10, 64)
			if err != nil {
				debug.PrintStack()
				log.Println("ERROR: ", err)
				return nil, err
			}
		default:
			debug.PrintStack()
			log.Println("ERROR: ", reflect.TypeOf(concept).String(), concept)
			return nil, errors.New("unable to interpret value; input type is " + reflect.TypeOf(concept).String())
		}
		return time.Unix(seconds, 0).Format(time.RFC3339), nil
	})
}
