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
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	"strconv"
	"strings"
)

func getGovaluateExpressionFunctions() map[string]govaluate.ExpressionFunction {
	return map[string]govaluate.ExpressionFunction{
		"substr": func(arguments ...interface{}) (result interface{}, err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("substr: catch panic %v", r)
				}
			}()
			if len(arguments) != 3 {
				return nil, errors.New("substr: expect exactly 3 arguments")
			}
			str, ok := arguments[0].(string)
			if !ok {
				return nil, errors.New("substr: expect argument 1 to be a string")
			}
			from, err := getInt(arguments[1])
			if err != nil {
				return nil, errors.New("substr: expect argument 2 to be a number")
			}
			to, err := getInt(arguments[2])
			if err != nil {
				return nil, errors.New("substr: expect argument 3 to be a number")
			}
			return str[from:to], nil
		},
		"replace": func(arguments ...interface{}) (result interface{}, err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("replace: catch panic %v", r)
				}
			}()
			if len(arguments) != 3 {
				return nil, errors.New("replace: expect exactly 3 arguments")
			}
			str, ok := arguments[0].(string)
			if !ok {
				return nil, errors.New("replace: expect argument 1 to be a string")
			}
			oldStr, ok := arguments[1].(string)
			if !ok {
				return nil, errors.New("replace: expect argument 2 to be a string")
			}
			newStr, ok := arguments[2].(string)
			if !ok {
				return nil, errors.New("replace: expect argument 3 to be a string")
			}
			return strings.ReplaceAll(str, oldStr, newStr), nil
		},
		"trimSuffix": func(arguments ...interface{}) (result interface{}, err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("trimSuffix: catch panic %v", r)
				}
			}()
			if len(arguments) != 2 {
				return nil, errors.New("trimSuffix: expect exactly 2 arguments")
			}
			str, ok := arguments[0].(string)
			if !ok {
				return nil, errors.New("trimSuffix: expect argument 1 to be a string")
			}
			suffix, ok := arguments[1].(string)
			if !ok {
				return nil, errors.New("trimSuffix: expect argument 2 to be a string")
			}
			return strings.TrimSuffix(str, suffix), nil
		},
		"trimPrefix": func(arguments ...interface{}) (result interface{}, err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("trimPrefix: catch panic %v", r)
				}
			}()
			if len(arguments) != 2 {
				return nil, errors.New("trimPrefix: expect exactly 2 arguments")
			}
			str, ok := arguments[0].(string)
			if !ok {
				return nil, errors.New("trimPrefix: expect argument 1 to be a string")
			}
			suffix, ok := arguments[1].(string)
			if !ok {
				return nil, errors.New("trimPrefix: expect argument 2 to be a string")
			}
			return strings.TrimPrefix(str, suffix), nil
		},
		"strIndex": func(arguments ...interface{}) (interface{}, error) {
			if len(arguments) != 2 {
				return nil, errors.New("strIndex: expect exactly 2 arguments")
			}
			str, ok := arguments[0].(string)
			if !ok {
				return nil, errors.New("strIndex: expect argument 1 to be a string")
			}
			sub, ok := arguments[1].(string)
			if !ok {
				return nil, errors.New("strIndex: expect argument 2 to be a string")
			}
			return strings.Index(str, sub), nil
		},
		"strlen": func(arguments ...interface{}) (interface{}, error) {
			if len(arguments) != 1 {
				return nil, errors.New("strlen: expect exactly one argument")
			}
			str, ok := arguments[0].(string)
			if !ok {
				return nil, errors.New("strlen: expect argument 1 to be a string")
			}
			return len(str), nil
		},
		"toUpperCase": func(arguments ...interface{}) (interface{}, error) {
			if len(arguments) != 1 {
				return nil, errors.New("toUpperCase: expect exactly 1 argument")
			}
			str, ok := arguments[0].(string)
			if !ok {
				return nil, errors.New("toUpperCase: expect argument 1 to be a string")
			}
			return strings.ToUpper(str), nil
		},
		"toLowerCase": func(arguments ...interface{}) (interface{}, error) {
			if len(arguments) != 1 {
				return nil, errors.New("toLowerCase: expect exactly 1 argument")
			}
			str, ok := arguments[0].(string)
			if !ok {
				return nil, errors.New("toLowerCase: expect argument 1 to be a string")
			}
			return strings.ToLower(str), nil
		},
		"atoi": func(arguments ...interface{}) (interface{}, error) {
			if len(arguments) != 1 {
				return nil, errors.New("atoi: expect exactly one argument")
			}
			str, ok := arguments[0].(string)
			if !ok {
				return nil, errors.New("atoi: expect argument 1 to be a string")
			}
			return strconv.ParseInt(str, 10, 64)
		},
		"atof": func(arguments ...interface{}) (interface{}, error) {
			if len(arguments) != 1 {
				return nil, errors.New("atof: expect exactly one argument")
			}
			str, ok := arguments[0].(string)
			if !ok {
				return nil, errors.New("atof: expect argument 1 to be a string")
			}
			return strconv.ParseFloat(str, 64)
		},
		"mapSet": func(arguments ...interface{}) (interface{}, error) {
			if len(arguments) != 3 {
				return nil, errors.New("mapSet: expect exactly 3 arguments")
			}
			m, ok := ensureMap(arguments[0])
			if !ok {
				return nil, errors.New("mapSet: expect argument 1 to be a map[string]interface{}")
			}
			key, ok := arguments[1].(string)
			if !ok {
				return nil, errors.New("strIndex: expect argument 2 to be a string")
			}
			value := arguments[2]

			result := map[string]interface{}{}
			for k, v := range m {
				result[k] = v
			}
			result[key] = value
			return result, nil
		},
		"mapGet": func(arguments ...interface{}) (interface{}, error) {
			if len(arguments) != 2 {
				return nil, errors.New("mapGet: expect exactly 2 arguments")
			}
			m, ok := ensureMap(arguments[0])
			if !ok {
				return nil, errors.New("mapGet: expect argument 1 to be a map[string]interface{}")
			}
			key, ok := arguments[1].(string)
			if !ok {
				return nil, errors.New("mapGet: expect argument 2 to be a string")
			}
			result, ok := m[key]
			if !ok {
				return nil, errors.New("mapGet: key not found")
			}
			return result, nil
		},
		"mapDelete": func(arguments ...interface{}) (interface{}, error) {
			if len(arguments) != 2 {
				return nil, errors.New("mapDelete: expect exactly 2 arguments")
			}
			m, ok := ensureMap(arguments[0])
			if !ok {
				return nil, errors.New("mapDelete: expect argument 1 to be a map[string]interface{}")
			}
			key, ok := arguments[1].(string)
			if !ok {
				return nil, errors.New("mapDelete: expect argument 2 to be a string")
			}
			result := map[string]interface{}{}
			for k, v := range m {
				result[k] = v
			}
			delete(result, key)
			return result, nil
		},
		"ntoa": func(arguments ...interface{}) (interface{}, error) {
			if len(arguments) != 1 {
				return nil, errors.New("ntoa: expect exactly one argument")
			}
			switch v := arguments[0].(type) {
			case int:
				return strconv.FormatInt(int64(v), 10), nil
			case int16:
				return strconv.FormatInt(int64(v), 10), nil
			case int32:
				return strconv.FormatInt(int64(v), 10), nil
			case int64:
				return strconv.FormatInt(v, 10), nil
			case float32:
				return strconv.FormatFloat(float64(v), 'f', -1, 32), nil
			case float64:
				return strconv.FormatFloat(v, 'f', -1, 64), nil
			default:
				return "", errors.New("ntoa: expect argument 1 to have one of the following types: int, int16, int32, int64, float32, float64")
			}
		},
	}
}
