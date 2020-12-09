package foo

import "testing"

func TestFoo(t *testing.T) {
	want := "Foo"
	if got := Foo(); got != want {
		t.Errorf("Foo() = %v, want %v", got, want)
	}
}
