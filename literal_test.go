package main

import "testing"

func BenchmarkLiteralVsStatement(b *testing.B) {
	b.Run("literal", func(sb *testing.B) {
		var m map[string]string
		for n := 0; n < sb.N; n++ {
			m = map[string]string{
				"a": "1",
				"b": "1",
				"c": "1",
				"d": "1",
				"e": "1",
				"f": "1",
				"g": "1",
				"h": "1",
				"i": "1",
				"j": "1",
				"k": "1",
				"l": "1",
				"m": "1",
				"n": "1",
				"o": "1",
				"p": "1",
				"q": "1",
				"r": "1",
				"s": "1",
				"t": "1",
				"u": "1",
				"v": "1",
				"w": "1",
				"x": "1",
				"y": "1",
				"z": "1",
			}
			if _, ok := m["z"]; !ok {
				b.Fatal(ok)
			}
		}
	})

	b.Run("statements", func(sb *testing.B) {
		var m map[string]string
		for n := 0; n < sb.N; n++ {
			m = map[string]string{}
			m["a"] = "1"
			m["b"] = "1"
			m["c"] = "1"
			m["d"] = "1"
			m["e"] = "1"
			m["f"] = "1"
			m["g"] = "1"
			m["h"] = "1"
			m["i"] = "1"
			m["j"] = "1"
			m["k"] = "1"
			m["l"] = "1"
			m["m"] = "1"
			m["n"] = "1"
			m["o"] = "1"
			m["p"] = "1"
			m["q"] = "1"
			m["r"] = "1"
			m["s"] = "1"
			m["t"] = "1"
			m["u"] = "1"
			m["v"] = "1"
			m["w"] = "1"
			m["x"] = "1"
			m["y"] = "1"
			m["z"] = "1"
			if _, ok := m["z"]; !ok {
				b.Fatal(ok)
			}
		}
	})
}
