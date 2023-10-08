package enum

type TipoFormatacao string

const (
	TipoFormatacao_Cadastro        TipoFormatacao = TipoFormatacao(TipoPreparacao_Cadastro)
	TipoFormatacao_Consulta        TipoFormatacao = TipoFormatacao(TipoPreparacao_Consulta)
	TipoFormatacao_RetornoConsulta TipoFormatacao = TipoFormatacao(TipoPreparacao_RetornoConsulta)
)
