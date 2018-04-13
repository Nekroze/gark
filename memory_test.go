package main

import "testing"

func BenchmarkMemoryStoreVsCall(b *testing.B) {
	getter := func() string {
		return "TEST"
	}
	user := func(_ string) {
		return
	}
	b.Run("Store Short Assignment", func(sb *testing.B) {
		for n := 0; n < sb.N; n++ {
			v := getter()
			user(v)
			user(v)
		}
	})
	b.Run("Store Var Assignment", func(sb *testing.B) {
		var v string
		for n := 0; n < sb.N; n++ {
			v = getter()
			user(v)
			user(v)
		}
	})
	b.Run("Call", func(sb *testing.B) {
		for n := 0; n < sb.N; n++ {
			user(getter())
			user(getter())
		}
	})
	b.Run("Store Var Assignment Single Use", func(sb *testing.B) {
		for n := 0; n < sb.N; n++ {
			var v string = getter()
			user(v)
		}
	})
	b.Run("Call Single Use", func(sb *testing.B) {
		for n := 0; n < sb.N; n++ {
			user(getter())
		}
	})
}
