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
)

const Seconds = "urn:infai:ses:characteristic:9e1024da-3b60-4531-9f29-464addccb13c"
const SecondsName = "second"

func init() {
	register.Labels[Seconds] = SecondsName
	register.Add(MilliSeconds, Seconds, register.RoundingLoss, func(in interface{}) (out interface{}, err error) {
		var msAsFloat float64
		switch ms := in.(type) {
		case int:
			msAsFloat = float64(ms)
		case int32:
			msAsFloat = float64(ms)
		case int64:
			msAsFloat = float64(ms)
		case float32:
			msAsFloat = float64(ms)
		case float64:
			msAsFloat = ms
		default:
			debug.PrintStack()
			log.Println("ERROR: ", reflect.TypeOf(in).String(), in)
			return nil, errors.New("unable to interpret value; input type is " + reflect.TypeOf(in).String())
		}
		return msAsFloat / 1000.0, nil
	})
	register.Add(Seconds, MilliSeconds, register.NoLosses, func(in interface{}) (out interface{}, err error) {
		var secondsAsFloat float64
		switch ms := in.(type) {
		case int:
			secondsAsFloat = float64(ms)
		case int32:
			secondsAsFloat = float64(ms)
		case int64:
			secondsAsFloat = float64(ms)
		case float32:
			secondsAsFloat = float64(ms)
		case float64:
			secondsAsFloat = ms
		default:
			debug.PrintStack()
			log.Println("ERROR: ", reflect.TypeOf(in).String(), in)
			return nil, errors.New("unable to interpret value; input type is " + reflect.TypeOf(in).String())
		}
		return secondsAsFloat * 1000.0, nil
	})

}
