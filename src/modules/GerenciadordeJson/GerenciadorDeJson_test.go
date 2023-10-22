package GerenciadordeJson_test

import (
	"safePasswordApi/src/modules/GerenciadordeJson"
	"testing"
)

func TestInterfaceParaJsonString(t *testing.T) {
	type Pessoa struct {
		Nome  string `json:"nome"`
		Idade int    `json:"idade"`
	}
	p := Pessoa{"João", 30}
	jsonStr, err := GerenciadordeJson.InterfaceParaJsonString(p)
	if err != nil {
		t.Errorf("Erro ao converter interface{} para JSON: %v", err)
	}

	expected := `{"nome":"João","idade":30}`
	if jsonStr != expected {
		t.Errorf("JSON esperado: %s, JSON retornado: %s", expected, jsonStr)
	}
}

func TestJsonStringParaInterface(t *testing.T) {
	jsonStr := `{"nome":"João","idade":30}`
	expected := map[string]interface{}{
		"nome":  "João",
		"idade": float64(30),
	}
	jsonData, err := GerenciadordeJson.JsonStringParaInterface(jsonStr)
	if err != nil {
		t.Errorf("Erro ao converter JSON para interface{}: %v", err)
	}
	if !compareMaps(jsonData.(map[string]interface{}), expected) {
		t.Errorf("Interface{} esperada: %v, Interface{} retornada: %v", expected, jsonData)
	}
}

func compareMaps(m1, m2 map[string]interface{}) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}
