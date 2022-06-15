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
	"github.com/SENERGY-Platform/converter/lib/converter"
	"github.com/SENERGY-Platform/converter/lib/model"
	"reflect"
	"testing"
)

func TestExtensionFunctions(t *testing.T) {
	c, err := converter.New()
	if err != nil {
		t.Error(err)
		return
	}
	t.Run("atoi", func(t *testing.T) {
		out, err := c.CastWithExtension("42", "foo", "bar", []model.ConverterExtension{
			{
				From:            "foo",
				To:              "bar",
				Distance:        -1,
				Formula:         "atoi(x)",
				PlaceholderName: "x",
			},
		})
		if err != nil {
			t.Error(err)
			return
		}
		if !reflect.DeepEqual(out, int64(42)) {
			t.Error(out)
			return
		}
	})

	t.Run("atof", func(t *testing.T) {
		out, err := c.CastWithExtension("42.13", "foo", "bar", []model.ConverterExtension{
			{
				From:            "foo",
				To:              "bar",
				Distance:        -1,
				Formula:         "atof(x)",
				PlaceholderName: "x",
			},
		})
		if err != nil {
			t.Error(err)
			return
		}
		if !reflect.DeepEqual(out, float64(42.13)) {
			t.Error(out)
			return
		}
	})

	t.Run("ntoa_f", func(t *testing.T) {
		out, err := c.CastWithExtension(42.13, "foo", "bar", []model.ConverterExtension{
			{
				From:            "foo",
				To:              "bar",
				Distance:        -1,
				Formula:         "ntoa(x)",
				PlaceholderName: "x",
			},
		})
		if err != nil {
			t.Error(err)
			return
		}
		if !reflect.DeepEqual(out, "42.13") {
			t.Error(out)
			return
		}
	})

	t.Run("ntoa_i", func(t *testing.T) {
		out, err := c.CastWithExtension(42, "foo", "bar", []model.ConverterExtension{
			{
				From:            "foo",
				To:              "bar",
				Distance:        -1,
				Formula:         "ntoa(x)",
				PlaceholderName: "x",
			},
		})
		if err != nil {
			t.Error(err)
			return
		}
		if !reflect.DeepEqual(out, "42") {
			t.Error(out)
			return
		}
	})

	t.Run("atof literal with quotes", func(t *testing.T) {
		out, err := c.CastWithExtension("", "foo", "bar", []model.ConverterExtension{
			{
				From:            "foo",
				To:              "bar",
				Distance:        -1,
				Formula:         `atof("42.13")`,
				PlaceholderName: "x",
			},
		})
		if err != nil {
			t.Error(err)
			return
		}
		if !reflect.DeepEqual(out, float64(42.13)) {
			t.Error(out)
			return
		}
	})

	t.Run("atoi literal with single quotes", func(t *testing.T) {
		out, err := c.CastWithExtension("", "foo", "bar", []model.ConverterExtension{
			{
				From:            "foo",
				To:              "bar",
				Distance:        -1,
				Formula:         `atoi('42')`,
				PlaceholderName: "x",
			},
		})
		if err != nil {
			t.Error(err)
			return
		}
		if !reflect.DeepEqual(out, int64(42)) {
			t.Error(out)
			return
		}
	})
}
