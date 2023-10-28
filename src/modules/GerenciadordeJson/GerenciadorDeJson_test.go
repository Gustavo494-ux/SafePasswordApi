package GerenciadordeJson_test

import (
	"fmt"
	"os"
	"safePasswordApi/src/modules/GerenciadordeJson"
	"safePasswordApi/src/modules/logger"
	"safePasswordApi/src/routines/inicializacao"
	"testing"
)

// TestMain:Função executada antes das demais
func TestMain(m *testing.M) {
	inicializacao.CarregarDotEnv()
	inicializacao.InicializarEncriptacao()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestInterfaceParaJsonString(t *testing.T) {
	type Pessoa struct {
		Nome  string `json:"nome"`
		Idade int    `json:"idade"`
	}
	p := Pessoa{"João", 30}
	jsonStr, err := GerenciadordeJson.InterfaceParaJsonString(p)
	if err != nil {
		t.Errorf("Erro ao converter interface{} para JSON: %v", err)
		logger.Logger().Error("Erro ao converter interface{} para JSON", err)
	}

	expected := `{"nome":"João","idade":30}`
	if jsonStr != expected {
		t.Errorf("JSON esperado: %s, JSON retornado: %s", expected, jsonStr)
		logger.Logger().Error(fmt.Sprintf("JSON esperado: %s, JSON retornado: %s", expected, jsonStr), err)
	}
	logger.Logger().Info("Teste TestInterfaceParaJsonString executado com sucesso!")
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
		logger.Logger().Error("Erro ao converter JSON para interface{}", err)
	}
	if !compareMaps(jsonData.(map[string]interface{}), expected) {
		t.Errorf("Interface{} esperada: %v, Interface{} retornada: %v", expected, jsonData)
		logger.Logger().Error(fmt.Sprintf("Interface{} esperada: %v, Interface{} retornada: %v", expected, jsonData), err)
	}
	logger.Logger().Info("Teste TestJsonStringParaInterface executado com sucesso!")
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
