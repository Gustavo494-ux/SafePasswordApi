package tipo

import (
	"github.com/duaneking/coalesce"
)

type coalesceType struct{}

// CoalesceStr: recebe um slice de strings e retorna o primeiro valor não nulo
func (c coalesceType) Str(args ...string) string {
	return (*coalesce.Coalesce(&args))[0]
}

// CoalesceInt: recebe um slice de inteiros e retorna o primeiro valor não nulo
func (c coalesceType) Int(args ...int) int {
	return (*coalesce.Coalesce(&args))[0]
}

// CoalesceFloat64: recebe um slice de valores float64 e retorna o primeiro valor não nulo
func (c coalesceType) Float64(args ...float64) float64 {
	return (*coalesce.Coalesce(&args))[0]
}

// CoalesceBool recebe um slice de valores booleanos e retorna o primeiro valor não nulo
func (c coalesceType) Bool(args ...bool) bool {
	return (*coalesce.Coalesce(&args))[0]
}

// CoalesceByte: recebe um slice de bytes e retorna o primeiro valor não nulo
func (c coalesceType) Byte(args ...byte) byte {
	return (*coalesce.Coalesce(&args))[0]
}

// Coalesce: cria uma instância de coalesceType
func Coalesce() coalesceType {
	return coalesceType{}
}
