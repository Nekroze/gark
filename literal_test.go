package main

import "testing"

func BenchmarkLiteralVsStatement(b *testing.B) {
	b.Run("literal", func(sb *testing.B) {
		var m map[string]string
		for n := 0; n < sb.N; n++ {
			m = map[string]string{
				"a": "",
				"b": "",
				"c": "",
				"d": "",
				"e": "",
				"f": "",
				"g": "",
				"h": "",
				"i": "",
				"j": "",
				"k": "",
				"l": "",
				"m": "",
				"n": "",
				"o": "",
				"p": "",
				"q": "",
				"r": "",
				"s": "",
				"t": "",
				"u": "",
				"v": "",
				"w": "",
				"x": "",
				"y": "",
				"z": "",
			}
			if _, ok := m["z"]; ok {
				b.Fatal(ok)
			}
		}
	})
	b.Run("statements", func(sb *testing.B) {
		var m map[string]string
		for n := 0; n < sb.N; n++ {
			m = map[string]string{}
			m["a"] = ""
			m["b"] = ""
			m["c"] = ""
			m["d"] = ""
			m["e"] = ""
			m["f"] = ""
			m["g"] = ""
			m["h"] = ""
			m["i"] = ""
			m["j"] = ""
			m["k"] = ""
			m["l"] = ""
			m["m"] = ""
			m["n"] = ""
			m["o"] = ""
			m["p"] = ""
			m["q"] = ""
			m["r"] = ""
			m["s"] = ""
			m["t"] = ""
			m["u"] = ""
			m["v"] = ""
			m["w"] = ""
			m["x"] = ""
			m["y"] = ""
			m["z"] = ""
			if _, ok := m["z"]; ok {
				b.Fatal(ok)
			}
		}
	})
}
