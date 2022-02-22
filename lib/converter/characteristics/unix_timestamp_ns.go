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
	"strconv"
	"time"
)

const UnixNanoSeconds = "urn:infai:ses:characteristic:c7dfcb86-2733-4917-a5ca-0a150a458eed"
const UnixNanoSecondsName = "unix_format_nanoseconds"

func init() {
	register.Labels[UnixNanoSeconds] = UnixNanoSecondsName

	register.Add(IsoTimestamp, UnixNanoSeconds, register.NoLosses, func(in interface{}) (out interface{}, err error) {
		iso, ok := in.(string)
		if !ok {
			log.Println("ERROR: ", reflect.TypeOf(in).String(), in)
			return nil, errors.New("unable to interpret value; input type is " + reflect.TypeOf(in).String())
		}
		t, err := time.Parse(time.RFC3339, iso)
		if err != nil {
			debug.PrintStack()
			log.Println("ERROR: ", err)
			return nil, err
		}
		return t.UnixNano(), nil
	})
	register.Add(UnixNanoSeconds, IsoTimestamp, register.RoundingLoss, func(in interface{}) (out interface{}, err error) {
		var nanoseconds int64
		switch m := in.(type) {
		case int:
			nanoseconds = int64(m)
		case int32:
			nanoseconds = int64(m)
		case int64:
			nanoseconds = m
		case float32:
			nanoseconds = int64(m)
		case float64:
			nanoseconds = int64(m)
		case string:
			nanoseconds, err = strconv.ParseInt(m, 10, 64)
			if err != nil {
				debug.PrintStack()
				log.Println("ERROR: ", err)
				return nil, err
			}
		default:
			debug.PrintStack()
			log.Println("ERROR: ", reflect.TypeOf(in).String(), in)
			return nil, errors.New("unable to interpret value; input type is " + reflect.TypeOf(in).String())
		}
		return time.Unix(0, nanoseconds).Format(time.RFC3339), nil
	})

}
