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
	"log/slog"
	"reflect"
	"runtime/debug"

	"github.com/SENERGY-Platform/converter/lib/converter/register"
)

const Lux = "urn:infai:ses:characteristic:0419856b-d198-480e-9998-bf990f226844"
const LuxName = "lux"

func init() {
	register.Labels[Lux] = LuxName

	register.Add(Lux, WattPerSquareMeter, register.MajorLoss, func(in interface{}) (out interface{}, err error) {
		//https://ambientweather.com/faqs/question/view/id/1452/
		var luxAsFloat float64
		switch lux := in.(type) {
		case int:
			luxAsFloat = float64(lux)
		case int32:
			luxAsFloat = float64(lux)
		case int64:
			luxAsFloat = float64(lux)
		case float32:
			luxAsFloat = float64(lux)
		case float64:
			luxAsFloat = lux
		default:
			debug.PrintStack()
			slog.Info("unable to interpret value", "input-type", reflect.TypeOf(in).String(), "input-value", in)
			return nil, errors.New("unable to interpret value; input type is " + reflect.TypeOf(in).String())
		}
		return luxAsFloat / 126.7, nil
	})

	register.Add(WattPerSquareMeter, Lux, register.MajorLoss, func(in interface{}) (out interface{}, err error) {
		//https://ambientweather.com/faqs/question/view/id/1452/
		var wAsFloat float64
		switch w := in.(type) {
		case int:
			wAsFloat = float64(w)
		case int32:
			wAsFloat = float64(w)
		case int64:
			wAsFloat = float64(w)
		case float32:
			wAsFloat = float64(w)
		case float64:
			wAsFloat = w
		default:
			debug.PrintStack()
			slog.Info("unable to interpret value", "input-type", reflect.TypeOf(in).String(), "input-value", in)
			return nil, errors.New("unable to interpret value; input type is " + reflect.TypeOf(in).String())
		}
		return wAsFloat * 126.7, nil
	})

}
