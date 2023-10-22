package GerenciadordeJson

import (
	"encoding/json"
)

// InterfaceParaJsonString: Converte uma interface genérica em um json em formato string.
func InterfaceParaJsonString(i any) (string, error) {
	jsonBytes, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	jsonStr := string(jsonBytes)
	return jsonStr, nil
}

// JsonStringParaInterface: Converte um json em formato string para uma interface genérica.
func JsonStringParaInterface(jsonStr string) (any, error) {
	var i interface{}
	err := json.Unmarshal([]byte(jsonStr), &i)
	if err != nil {
		return nil, err
	}
	return i, nil
}
