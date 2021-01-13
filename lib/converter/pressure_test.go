package converter

import (
	"github.com/SENERGY-Platform/converter/lib/converter/pressure"
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
		out, err := Cast(test.Bar, pressure.Bar, pressure.HPa)
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
		out, err := Cast(test.HPa, pressure.HPa, pressure.Bar)
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
