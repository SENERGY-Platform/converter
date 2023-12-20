/*
 * Copyright 2023 InfAI (CC SES)
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

package converter

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type InterfaceTypeWrapper struct {
	InterfaceTypeField interface{}
}

func getInterfaceType() reflect.Type {
	field, ok := reflect.TypeOf(InterfaceTypeWrapper{}).FieldByName("InterfaceTypeField")
	if !ok {
		panic("expect InterfaceTypeField in InterfaceTypeWrapper")
	}
	return field.Type
}

func inputCleanup(in interface{}) (result interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("catch panic: %v", r))
		}
	}()

	if in == nil {
		return nil, nil
	}
	if reflect.TypeOf(in).Kind() == reflect.Map && reflect.TypeOf(in).Key().Kind() == reflect.String {
		m := map[string]interface{}{}
		iter := reflect.ValueOf(in).MapRange()
		for iter.Next() {
			key := iter.Key()
			value, err := inputCleanup(iter.Value().Interface())
			if err != nil {
				return nil, err
			}
			m[key.String()] = value
		}
		return m, nil
	}
	//in the odd case that the map key is not a string
	if reflect.TypeOf(in).Kind() == reflect.Map {
		iter := reflect.ValueOf(in).MapRange()
		m := reflect.MakeMap(reflect.MapOf(reflect.TypeOf(in).Key(), getInterfaceType()))
		for iter.Next() {
			key := iter.Key()
			value, err := inputCleanup(iter.Value().Interface())
			if err != nil {
				return nil, err
			}
			m.SetMapIndex(key, reflect.ValueOf(value))
		}
		return m.Interface(), nil
	}
	if reflect.TypeOf(in).Kind() == reflect.Slice {
		l := reflect.ValueOf(in).Len()
		slice := []interface{}{}
		for i := 0; i < l; i++ {
			value, err := inputCleanup(reflect.ValueOf(in).Index(i).Interface())
			if err != nil {
				return nil, err
			}
			slice = append(slice, value)
		}
		return slice, nil
	}
	switch v := in.(type) {
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case string:
		return v, nil
	case bool:
		return v, nil
	default:
		return jsonCast(v)
	}
}

func jsonCast(in interface{}) (out interface{}, err error) {
	temp, err := json.Marshal(in)
	if err != nil {
		return out, err
	}
	err = json.Unmarshal(temp, &out)
	return out, err
}
