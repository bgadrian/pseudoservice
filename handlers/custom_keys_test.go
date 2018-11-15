package handlers

import (
	"fmt"
	"testing"
)

func TestGenerateCustom(t *testing.T) {
	for param := range keys {
		template := fmt.Sprintf("~%s~", param)
		gotList := GenerateCustom(42, template, 5)

		for _, got := range gotList {
			if template == got {
				t.Errorf("template var did not got replaced for '%s'", template)
				break
			}

			if len(got) == 0 {
				t.Errorf("result is empty for '%s'", template)
				break
			}
		}
	}
}
