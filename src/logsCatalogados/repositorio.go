package logsCatalogados

import "errors"

var (
	ErroRepositorio_DadosNaoEncontrados = errors.New("nenhum registro encontrado, verifique os dados fornecidos")
	ErroRepositorio_DadosNaoAfetados    = errors.New("nenhum registro foi afetado, verifique os dados fornecidos")
)
