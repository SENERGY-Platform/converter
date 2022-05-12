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
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/SENERGY-Platform/converter/lib/converter/register"
	"github.com/SENERGY-Platform/converter/lib/model"
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

func (this *Converter) CastWithExtension(in interface{}, from string, to string, extensions []model.ConverterExtension) (out interface{}, err error) {
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

func getExtensionCastFunction(desc model.ConverterExtension) register.CastFunction {
	return func(in interface{}) (out interface{}, err error) {
		expression, err := govaluate.NewEvaluableExpression(desc.F)
		if err != nil {
			return out, fmt.Errorf("unable to parse extension expression (%v): %v", desc.F, err)
		}
		out, err = expression.Evaluate(map[string]interface{}{desc.PlaceholderName: in})
		if err != nil {
			return out, fmt.Errorf("unable to evaluate extension expression (%v) with input (%v = %v): %v", desc.F, desc.PlaceholderName, in, err)
		}
		return out, nil
	}
}
