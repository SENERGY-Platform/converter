/*
 * Copyright 2022 InfAI (CC SES)
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

const Kmh = "urn:infai:ses:characteristic:becfc22a-7f24-44df-9a46-45acebc56d00"
const KmhName = "km/h"

func init() {
	register.Labels[Kmh] = KmhName

	register.Add(MetersPerSecond, Kmh, register.RoundingLoss, func(in interface{}) (out interface{}, err error) {
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
		return msAsFloat * 3.6, nil
	})

	register.Add(Kmh, MetersPerSecond, register.NoLosses, func(in interface{}) (out interface{}, err error) {
		var kmhAsFloat float64
		switch kmh := in.(type) {
		case int:
			kmhAsFloat = float64(kmh)
		case int32:
			kmhAsFloat = float64(kmh)
		case int64:
			kmhAsFloat = float64(kmh)
		case float32:
			kmhAsFloat = float64(kmh)
		case float64:
			kmhAsFloat = kmh
		default:
			debug.PrintStack()
			log.Println("ERROR: ", reflect.TypeOf(in).String(), in)
			return nil, errors.New("unable to interpret value; input type is " + reflect.TypeOf(in).String())
		}
		return kmhAsFloat / 3.6, nil
	})

}
