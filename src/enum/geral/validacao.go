package enum

type TipoValidacao string

const (
	TipoValidacao_Cadastro        TipoValidacao = TipoValidacao(TipoPreparacao_Cadastro)
	TipoValidacao_Consulta        TipoValidacao = TipoValidacao(TipoPreparacao_Consulta)
	TipoValidacao_Atualizar       TipoValidacao = TipoValidacao(TipoPreparacao_Atualizar)
	TipoValidacao_RetornoConsulta TipoValidacao = TipoValidacao(TipoPreparacao_RetornoConsulta)
)
