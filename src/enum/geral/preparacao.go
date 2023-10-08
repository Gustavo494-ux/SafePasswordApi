package enum

type TipoPreparacao string

const (
	TipoPreparacao_Cadastro        TipoPreparacao = "Cadastro"
	TipoPreparacao_Atualizar       TipoPreparacao = "Atualizar"
	TipoPreparacao_Consulta        TipoPreparacao = "Consultar"
	TipoPreparacao_RetornoConsulta TipoPreparacao = "RetornarDados"
)
