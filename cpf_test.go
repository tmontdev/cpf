package cpf

import (
	"testing"
)

func Test(t *testing.T) {
	if IsValid("99") || IsValid("999999999999") {
		t.Error("should not be valid")
	}
	if Must("39483350875").Format() != "394.833.508-75" {
		t.Error("should not be valid")
	}
}
