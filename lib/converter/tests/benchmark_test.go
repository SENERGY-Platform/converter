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
	"fmt"
	"github.com/SENERGY-Platform/converter/lib/converter"
	"github.com/SENERGY-Platform/converter/lib/converter/characteristics"
	"github.com/SENERGY-Platform/converter/lib/model"
	"testing"
)

func BenchmarkCast(b *testing.B) {
	c, err := converter.New()
	if err != nil {
		b.Error(err)
		return
	}

	b.Run("build_in", func(b *testing.B) {
		out, err := c.Cast(42.0, characteristics.Celsius, characteristics.Kelvin)
		if err != nil {
			b.Error(err)
			return
		}
		if fmt.Sprint(out) != "315.15" {
			b.Error(out)
			return
		}
	})

	b.Run("extension", func(b *testing.B) {
		out, err := c.CastWithExtension(42, "foo", "bar", []model.ConverterExtension{
			{
				From:            "foo",
				To:              "bar",
				Distance:        -1,
				Formula:         "x+2.4",
				PlaceholderName: "x",
			},
		})
		if err != nil {
			b.Error(err)
			return
		}
		if fmt.Sprint(out) != "44.4" {
			b.Error(out)
			return
		}
	})

	b.Run("extension_overwrite", func(b *testing.B) {
		out, err := c.CastWithExtension(42, characteristics.Celsius, characteristics.Kelvin, []model.ConverterExtension{
			{
				From:            characteristics.Celsius,
				To:              characteristics.Kelvin,
				Distance:        -1,
				Formula:         "x+2.4",
				PlaceholderName: "x",
			},
		})
		if err != nil {
			b.Error(err)
			return
		}
		if fmt.Sprint(out) != "44.4" {
			b.Error(out)
			return
		}
	})

}