package handlers

import (
	"fmt"
	"testing"
)

func TestGenerateCustomSimple(t *testing.T) {
	for _, test := range []struct {
		template, shoud string
	}{
		{"", ""},
		{"~", "~"},
		{"~~", "~~"},
		{"~notfound~", "~notfound~"},
		{"Alfa", "Alfa"},
		{"Al~fa", "Al~fa"},
		{"~name~", "Jeromy Schmeler"},
		{"~name~~", "Jeromy Schmeler~"},
		{"~~name~~", "~Jeromy Schmeler~"},
		{"~~name~", "~Jeromy Schmeler"},
		{"~name~!", "Jeromy Schmeler!"},
		{"X~name~X", "XJeromy SchmelerX"},
		{"X~name~X~name~X", "XJeromy SchmelerXKim SteuberX"},
		{"~name", "~name"},
		{"name~", "name~"},
		{"name~", "name~"},
	} {
		got := GenerateCustom(42, test.template, 1)

		if len(got) != 1 {
			t.Errorf("the result should have 1 result")
		}
		if got[0] == test.shoud {
			continue
		}

		t.Errorf("fot template '%s', got '%s' should '%s'",
			test.template, got[0], test.shoud)
	}
}

func TestGenerateCustomJson(t *testing.T) {
	for _, test := range []struct {
		template, shoud string
	}{
		{`{name:"~name~"}`, `{name:"Jeromy Schmeler"}`},
		{`["~name~", "~name~"]`, `["Jeromy Schmeler", "Kim Steuber"]`},

		{`{name:"~name~", age: ~digit~}`, `{name:"Jeromy Schmeler", age: 8}`},
	} {
		got := GenerateCustom(42, test.template, 1)

		if len(got) != 1 {
			t.Errorf("the result should have 1 result")
		}
		if got[0] == test.shoud {
			continue
		}

		t.Errorf("fot template '%s', got '%s' should '%s'",
			test.template, got[0], test.shoud)
	}
}

func ExampleGenerateCustom() {
	template := "Hello ~name~!"
	seed := 42 //for each seed value will generate a different result
	count := 2

	randomData := GenerateCustom(int64(seed), template, int32(count))

	for _, result := range randomData {
		fmt.Printf("%s\n", result)
	}
	// Output:Hello Jeromy Schmeler!
	// Hello Kim Steuber!
}

func ExampleGenerateCustomJson() {
	template := `{name:"~name~", age: ~digit~}`
	seed := 42 //for each seed value will generate a different result
	count := 1

	randomData := GenerateCustom(int64(seed), template, int32(count))

	for _, result := range randomData {
		fmt.Printf("%s\n", result)
	}
	// Output:{name:"Jeromy Schmeler", age: 8}
}
