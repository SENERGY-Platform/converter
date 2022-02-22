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

import "github.com/SENERGY-Platform/converter/lib/converter/register"
import _ "github.com/SENERGY-Platform/converter/lib/converter/characteristics"

type Converter struct {
	register register.Register
}

func New() (converter *Converter, err error) {
	r, err := register.NewGraphRegister(register.List)
	if err != nil {
		return converter, err
	}
	return NewFromRegister(r), nil
}

func NewFromRegister(register register.Register) (converter *Converter) {
	return &Converter{register: register}
}

func (this *Converter) Cast(in interface{}, from string, to string) (out interface{}, err error) {
	if from == to {
		return in, nil
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
