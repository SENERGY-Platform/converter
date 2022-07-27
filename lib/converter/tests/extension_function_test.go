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

	t.Run("strlen", func(t *testing.T) {
		out, err := c.CastWithExtension("42", "foo", "bar", []model.ConverterExtension{
			{
				From:            "foo",
				To:              "bar",
				Distance:        -1,
				Formula:         "strlen(x)",
				PlaceholderName: "x",
			},
		})
		if err != nil {
			t.Error(err)
			return
		}
		if !reflect.DeepEqual(out, int(2)) {
			t.Error(out)
			return
		}
	})

	t.Run("strIndex", func(t *testing.T) {
		out, err := c.CastWithExtension("4.2", "foo", "bar", []model.ConverterExtension{
			{
				From:            "foo",
				To:              "bar",
				Distance:        -1,
				Formula:         "strIndex(x, \".\")",
				PlaceholderName: "x",
			},
		})
		if err != nil {
			t.Error(err)
			return
		}
		if !reflect.DeepEqual(out, int(1)) {
			t.Error(out)
			return
		}
	})

	t.Run("trimPrefix", func(t *testing.T) {
		out, err := c.CastWithExtension("foo:4.2", "foo", "bar", []model.ConverterExtension{
			{
				From:            "foo",
				To:              "bar",
				Distance:        -1,
				Formula:         "trimPrefix(x, \"foo:\")",
				PlaceholderName: "x",
			},
		})
		if err != nil {
			t.Error(err)
			return
		}
		if !reflect.DeepEqual(out, "4.2") {
			t.Error(out)
			return
		}
	})

	t.Run("trimSuffix", func(t *testing.T) {
		out, err := c.CastWithExtension("foo:4.2", "foo", "bar", []model.ConverterExtension{
			{
				From:            "foo",
				To:              "bar",
				Distance:        -1,
				Formula:         "trimSuffix(x, \":4.2\")",
				PlaceholderName: "x",
			},
		})
		if err != nil {
			t.Error(err)
			return
		}
		if !reflect.DeepEqual(out, "foo") {
			t.Error(out)
			return
		}
	})

	t.Run("replace", func(t *testing.T) {
		out, err := c.CastWithExtension("foo:4.2", "foo", "bar", []model.ConverterExtension{
			{
				From:            "foo",
				To:              "bar",
				Distance:        -1,
				Formula:         "replace(x, \":\", \"/\")",
				PlaceholderName: "x",
			},
		})
		if err != nil {
			t.Error(err)
			return
		}
		if !reflect.DeepEqual(out, "foo/4.2") {
			t.Error(out)
			return
		}
	})

	t.Run("substr", func(t *testing.T) {
		out, err := c.CastWithExtension("0123456789", "foo", "bar", []model.ConverterExtension{
			{
				From:            "foo",
				To:              "bar",
				Distance:        -1,
				Formula:         "substr(x, 2, 4)",
				PlaceholderName: "x",
			},
		})
		if err != nil {
			t.Error(err)
			return
		}
		if !reflect.DeepEqual(out, "23") {
			t.Error(out)
			return
		}
	})

	t.Run("toUpperCase", func(t *testing.T) {
		out, err := c.CastWithExtension("fooBAR", "foo", "bar", []model.ConverterExtension{
			{
				From:            "foo",
				To:              "bar",
				Distance:        -1,
				Formula:         "toUpperCase(x)",
				PlaceholderName: "x",
			},
		})
		if err != nil {
			t.Error(err)
			return
		}
		if !reflect.DeepEqual(out, "FOOBAR") {
			t.Error(out)
			return
		}
	})

	t.Run("toLowerCase", func(t *testing.T) {
		out, err := c.CastWithExtension("fooBAR", "foo", "bar", []model.ConverterExtension{
			{
				From:            "foo",
				To:              "bar",
				Distance:        -1,
				Formula:         "toLowerCase(x)",
				PlaceholderName: "x",
			},
		})
		if err != nil {
			t.Error(err)
			return
		}
		if !reflect.DeepEqual(out, "foobar") {
			t.Error(out)
			return
		}
	})

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

func ExampleNtoa() {
	c, err := converter.New()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	out, err := c.CastWithExtension(42.0, "foo", "bar", []model.ConverterExtension{
		{
			From:            "foo",
			To:              "bar",
			Distance:        -1,
			Formula:         `ntoa(x) == "42"`,
			PlaceholderName: "x",
		},
	})
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(out)
	//output:
	//true
}

func ExampleAtoiNtoa() {
	c, err := converter.New()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	out, err := c.CastWithExtension(42.0, "foo", "bar", []model.ConverterExtension{
		{
			From:            "foo",
			To:              "bar",
			Distance:        -1,
			Formula:         "atoi(ntoa(x))==42",
			PlaceholderName: "x",
		},
	})
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(out)
	//output:
	//true
}

func ExampleComplexExpression() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("PANIC:", r)
		}
	}()
	c, err := converter.New()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	expr := `(((x-(x%60))/60) >= 10 ? "" : "0") + ntoa(((x-(x%60))/60)) + ":" + (((x%60 >= 10) ? "" : "0" ) + ntoa(x%60))`

	out1, err := c.CastWithExtension(638, "foo", "bar", []model.ConverterExtension{
		{
			From:            "foo",
			To:              "bar",
			Distance:        -1,
			Formula:         expr,
			PlaceholderName: "x",
		},
	})
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(out1)

	out2, err := c.CastWithExtension(68, "foo", "bar", []model.ConverterExtension{
		{
			From:            "foo",
			To:              "bar",
			Distance:        -1,
			Formula:         expr,
			PlaceholderName: "x",
		},
	})
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(out2)

	//output:
	// 10:38
	// 01:08

}
