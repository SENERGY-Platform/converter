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

package converter

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/SENERGY-Platform/converter/lib/converter/register"
	"github.com/SENERGY-Platform/models/go/models"
)
import _ "github.com/SENERGY-Platform/converter/lib/converter/characteristics"

type Converter struct {
	register register.Register
}

func New() (converter *Converter, err error) {
	return NewFromRegisterEntries(register.List)
}

func NewFromRegisterEntries(entries []register.Entry) (converter *Converter, err error) {
	r, err := register.NewGraphRegister(entries)
	if err != nil {
		return converter, err
	}
	return NewFromRegister(r), nil
}

func NewFromRegister(register register.Register) (converter *Converter) {
	return &Converter{register: register}
}

func (this *Converter) UpdateRegister(castDescriptions []register.Entry) (err error) {
	return this.register.Update(append(register.List, castDescriptions...))
}

func (this *Converter) Cast(in interface{}, from string, to string) (out interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("recovered panic: " + fmt.Sprint(r))
		}
	}()
	if from == to {
		return in, nil
	}
	if in == nil {
		return nil, nil
	}

	casts, err := this.register.GetCasts(from, to)
	if err != nil {
		return out, err
	}

	out = in
	for _, cast := range casts {
		out, err = cast(out)
		if err != nil {
			return out, err
		}
	}
	return out, nil
}

func (this *Converter) CastWithExtension(in interface{}, from string, to string, extensions []models.ConverterExtension) (out interface{}, err error) {
	in, err = inputCleanup(in)
	if err != nil {
		return out, fmt.Errorf("cleanup error: %v", err)
	}
	rules := []register.Entry{}
	rules = append(rules, register.List...)
	for _, ruleDesc := range extensions {
		rules = append(rules, register.Entry{
			From:     ruleDesc.From,
			To:       ruleDesc.To,
			Distance: ruleDesc.Distance,
			Cast:     getExtensionCastFunction(ruleDesc),
		})
	}
	tempConverter, err := NewFromRegisterEntries(rules)
	if err != nil {
		return out, fmt.Errorf("unable to create extended converter: %v", err)
	}
	return tempConverter.Cast(in, from, to)
}

func (this *Converter) castWithExtension(in interface{}, from string, to string, extensions []models.ConverterExtension) (out interface{}, err error) {
	in, err = inputCleanup(in)
	if err != nil {
		return out, fmt.Errorf("cleanup error: %v", err)
	}
	rules := []register.Entry{}
	rules = append(rules, register.List...)
	for _, ruleDesc := range extensions {
		rules = append(rules, register.Entry{
			From:     ruleDesc.From,
			To:       ruleDesc.To,
			Distance: ruleDesc.Distance,
			Cast:     getExtensionCastFunction(ruleDesc),
		})
	}
	tempConverter, err := NewFromRegisterEntries(rules)
	if err != nil {
		return out, fmt.Errorf("unable to create extended converter: %v", err)
	}
	return tempConverter.Cast(in, from, to)
}

func (this *Converter) ValidateExtensions(nodes []string, extensions []models.ConverterExtension) (err error) {
	rules := []register.Entry{}
	rules = append(rules, register.List...)
	for _, ruleDesc := range extensions {
		rules = append(rules, register.Entry{
			From:     ruleDesc.From,
			To:       ruleDesc.To,
			Distance: ruleDesc.Distance,
			Cast:     getExtensionCastFunction(ruleDesc),
		})
	}
	tempConverter, err := NewFromRegisterEntries(rules)
	if err != nil {
		return fmt.Errorf("unable to create extended converter: %w", err)
	}
	for _, from := range nodes {
		for _, to := range nodes {
			if from != to {
				casts, err := tempConverter.register.GetCasts(from, to)
				if err != nil {
					return fmt.Errorf("error while searching conversion path from %v to %v: %w", from, to, err)
				}
				if len(casts) == 0 {
					return fmt.Errorf("no conversion path from %v to %v found", from, to)
				}
			}

		}
	}
	return nil
}

func (this *Converter) TryExtension(extension models.ConverterExtension, in interface{}) (out interface{}, err error) {
	return getExtensionCastFunction(extension)(in)
}

func getExtensionCastFunction(desc models.ConverterExtension) register.CastFunction {
	return func(in interface{}) (out interface{}, err error) {
		expression, err := govaluate.NewEvaluableExpressionWithFunctions(desc.Formula, getGovaluateExpressionFunctions())
		if err != nil {
			return out, fmt.Errorf("unable to parse extension expression (%v): %v", desc.Formula, err)
		}
		out, err = expression.Evaluate(map[string]interface{}{desc.PlaceholderName: in})
		if err != nil {
			return out, fmt.Errorf("unable to evaluate extension expression (%v) with input (%v = %#v): %v", desc.Formula, desc.PlaceholderName, in, err)
		}
		return out, nil
	}
}

func ensureMap(input interface{}) (result map[string]interface{}, ok bool) {
	switch v := input.(type) {
	case map[string]interface{}:
		return v, true
	case map[string]string:
		return castToInterfaceMap(v), true
	case map[string]int64:
		return castToInterfaceMap(v), true
	case map[string]int32:
		return castToInterfaceMap(v), true
	case map[string]int:
		return castToInterfaceMap(v), true
	case map[string]float64:
		return castToInterfaceMap(v), true
	case map[string]float32:
		return castToInterfaceMap(v), true
	case map[string]bool:
		return castToInterfaceMap(v), true
	}
	temp, err := json.Marshal(input)
	if err != nil {
		return result, false
	}
	err = json.Unmarshal(temp, &result)
	if err != nil {
		return result, false
	}
	return result, true
}

func castToInterfaceMap[T any](in map[string]T) (result map[string]interface{}) {
	result = map[string]interface{}{}
	for key, value := range in {
		result[key] = value
	}
	return result
}

func getInt(in interface{}) (int, error) {
	switch i := in.(type) {
	case int:
		return i, nil
	case int16:
		return int(i), nil
	case int32:
		return int(i), nil
	case int64:
		return int(i), nil
	case float32:
		return int(i), nil
	case float64:
		return int(i), nil
	default:
		return 0, errors.New("not number")
	}
}
