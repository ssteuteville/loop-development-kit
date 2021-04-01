package ldk_test

import (
	"github.com/open-olive/loop-development-kit/ldk/go/v2/service"
	"reflect"
	"runtime"
	"testing"
)

func TestKeyModifiers(t *testing.T) {
	type keyFunc func(service.KeyModifier) bool
	funcs := []keyFunc{
		service.KeyModifier.AltLeft,
		service.KeyModifier.AltRight,
		service.KeyModifier.Alt,
		service.KeyModifier.ControlLeft,
		service.KeyModifier.ControlRight,
		service.KeyModifier.Control,
		service.KeyModifier.MetaLeft,
		service.KeyModifier.MetaRight,
		service.KeyModifier.Meta,
		service.KeyModifier.ShiftLeft,
		service.KeyModifier.ShiftRight,
		service.KeyModifier.Shift,
	}
	tests := []struct {
		name string
		bit  int
	}{
		{name: "alt left", bit: 1 << 0},
		{name: "alt right", bit: 1 << 1},
		{name: "alt", bit: 1 << 2},
		{name: "control left", bit: 1 << 3},
		{name: "control right", bit: 1 << 4},
		{name: "control", bit: 1 << 5},
		{name: "meta left", bit: 1 << 6},
		{name: "meta right", bit: 1 << 7},
		{name: "meta", bit: 1 << 8},
		{name: "shift left", bit: 1 << 9},
		{name: "shift right", bit: 1 << 10},
		{name: "shift", bit: 1 << 11},
	}

	for _, test := range tests {
		test := test // grab new lexical scope
		t.Run(test.name, func(t *testing.T) {
			for i := 0; i <= 11; i++ {
				bit := 1 << i
				for j, fn := range funcs {
					fnName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
					if i == j {
						t.Logf("testing %d -> %s", bit, fnName)
						if !fn(service.KeyModifier(bit)) {
							t.Errorf("%s not detected for %d", test.name, bit)
						}
					} else {
						if fn(service.KeyModifier(bit)) {
							t.Errorf("%s should not be detected for %d (%s)", test.name, bit, fnName)
						}
					}
				}
			}
		})
	}
}
