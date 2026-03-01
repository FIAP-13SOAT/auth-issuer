package validator

import "testing"

func TestIsValidDocument(t *testing.T) {
	tests := []struct {
		name     string
		document string
		want     bool
	}{
		{"CPF válido com pontuação", "033.326.420-73", true},
		{"CPF válido sem pontuação", "03332642073", true},
		{"CPF inválido", "111.111.111-11", false},
		{"CPF inválido dígito", "033.326.420-72", false},
		{"CNPJ válido", "11.222.333/0001-81", true},
		{"CNPJ inválido", "11.222.333/0001-82", false},
		{"Documento vazio", "", false},
		{"Documento curto", "123", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidDocument(tt.document); got != tt.want {
				t.Errorf("IsValidDocument(%q) = %v, want %v", tt.document, got, tt.want)
			}
		})
	}
}
