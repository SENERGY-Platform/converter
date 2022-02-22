/*
 * Copyright 2021 InfAI (CC SES)
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
	"github.com/SENERGY-Platform/converter/lib/converterV2/characteristics"
	"reflect"
	"testing"
)

func TestPressureBarToHPa(t *testing.T) {
	t.Parallel()
	table := []struct {
		Bar interface{}
		HPa float64
	}{
		{Bar: int(1), HPa: 1000},
		{Bar: float32(1), HPa: 1000},
		{Bar: float64(1), HPa: 1000},
		{Bar: int64(1), HPa: 1000},
	}
	for i, test := range table {
		out, err := Cast(test.Bar, characteristics.Bar, characteristics.HPa)
		if err != nil {
			t.Fatal(err)
		}
		result, ok := out.(float64)
		if !ok {
			t.Error(out, reflect.TypeOf(out))
			continue
		}

		if result != test.HPa {
			t.Error(i, result, test.Bar)
		}
	}
}

func TestPressureHPaToBar(t *testing.T) {
	t.Parallel()
	table := []struct {
		Bar float64
		HPa interface{}
	}{
		{Bar: 1, HPa: int(1000)},
		{Bar: 1, HPa: int64(1000)},
		{Bar: 1, HPa: float32(1000)},
		{Bar: 1, HPa: float64(1000)},
	}
	for i, test := range table {
		out, err := Cast(test.HPa, characteristics.HPa, characteristics.Bar)
		if err != nil {
			t.Fatal(err)
		}
		result, ok := out.(float64)
		if !ok {
			t.Fatal(out, reflect.TypeOf(out))
		}

		if result != test.Bar {
			t.Fatal(i, result, test.Bar)
		}
	}
}
