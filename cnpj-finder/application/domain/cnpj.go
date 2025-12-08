package domain

type CNPJ struct {
	UF                        string                `json:"uf"`
	CEP                       string                `json:"cep"`
	QSA                       []Socio            `json:"qsa"`
	CNPJ                      string                `json:"cnpj"`
	Pais                      string               `json:"pais"`
	Email                     string               `json:"email"`
	Porte                     string                `json:"porte"`
	Bairro                    string                `json:"bairro"`
	Numero                    string                `json:"numero"`
	DDDFax                    string                `json:"ddd_fax"`
	Municipio                 string                `json:"municipio"`
	Logradouro                string                `json:"logradouro"`
	CNAEFiscal                int                   `json:"cnae_fiscal"`
	CodigoPais                int                  `json:"codigo_pais"`
	Complemento               string                `json:"complemento"`
	CodigoPorte               int                   `json:"codigo_porte"`
	RazaoSocial               string                `json:"razao_social"`
	NomeFantasia              string                `json:"nome_fantasia"`
	CapitalSocial             float64               `json:"capital_social"`
	DDDtelefone1              string                `json:"ddd_telefone_1"`
	DDDtelefone2              string                `json:"ddd_telefone_2"`
	OpcaoPeloMEI              bool                 `json:"opcao_pelo_mei"`
	CodigoMunicipio           int                   `json:"codigo_municipio"`
	CNAESecundarios           []CNAE   `json:"cnaes_secundarios"`
	NaturezaJuridica          string                `json:"natureza_juridica"`
	RegimeTributario          []RegimeTributario `json:"regime_tributario"`
	SituacaoEspecial          string                `json:"situacao_especial"`
	OpcaoPeloSimples          bool                 `json:"opcao_pelo_simples"`
	SituacaoCadastral         int                   `json:"situacao_cadastral"`
	DataOpcaoPeloMEI          string               `json:"data_opcao_pelo_mei"`
	DataExclusaoDoMEI         string               `json:"data_exclusao_do_mei"`
	CNAEFiscalDescricao       string                `json:"cnae_fiscal_descricao"`
	CodigoMunicipioIBGE       int                   `json:"codigo_municipio_ibge"`
	DataInicioAtividade       string                `json:"data_inicio_atividade"`
	DataSituacaoEspecial      string               `json:"data_situacao_especial"`
	DataOpcaoPeloSimples      string               `json:"data_opcao_pelo_simples"`
	DataSituacaoCadastral     string               `json:"data_situacao_cadastral"`
	NomeCidadeNoExterior      string                `json:"nome_cidade_no_exterior"`
	CodigoNaturezaJuridica    int                   `json:"codigo_natureza_juridica"`
	DataExclusaoDoSimples     string               `json:"data_exclusao_do_simples"`
	MotivoSituacaoCadastral   int                   `json:"motivo_situacao_cadastral"`
	EnteFederativoResponsavel string                `json:"ente_federativo_responsavel"`
	IdentificadorMatrizFilial int                   `json:"identificador_matriz_filial"`
	QualificacaoDoResponsavel int                   `json:"qualificacao_do_responsavel"`
	DescricaoSituacaoCadastral string               `json:"descricao_situacao_cadastral"`
	DescricaoTipoDeLogradouro string               `json:"descricao_tipo_de_logradouro"`
	DescricaoMotivoSituacao   string               `json:"descricao_motivo_situacao_cadastral"`
	DescricaoMatrizFilial     string               `json:"descricao_identificador_matriz_filial"`
}

type Socio struct {
	Pais                                 string `json:"pais"`
	NomeSocio                            string  `json:"nome_socio"`
	CodigoPais                           int    `json:"codigo_pais"`
	FaixaEtaria                          string  `json:"faixa_etaria"`
	CNPJCPFSocio                         string  `json:"cnpj_cpf_do_socio"`
	QualificacaoSocio                    string  `json:"qualificacao_socio"`
	CodigoFaixaEtaria                    int     `json:"codigo_faixa_etaria"`
	DataEntradaSociedade                 string  `json:"data_entrada_sociedade"`
	IdentificadorDeSocio                 int     `json:"identificador_de_socio"`
	CPFRepresentanteLegal                string  `json:"cpf_representante_legal"`
	NomeRepresentanteLegal               string  `json:"nome_representante_legal"`
	CodigoQualificacaoSocio              int     `json:"codigo_qualificacao_socio"`
	QualificacaoRepresentanteLegal       string  `json:"qualificacao_representante_legal"`
	CodigoQualificacaoRepresentanteLegal int     `json:"codigo_qualificacao_representante_legal"`
}

type CNAE struct {
	Codigo    int    `json:"codigo"`
	Descricao string `json:"descricao"`
}

type RegimeTributario struct {
	Ano                       int     `json:"ano"`
	CNPJDaSCP                 string `json:"cnpj_da_scp"`
	FormaDeTributacao         string  `json:"forma_de_tributacao"`
	QuantidadeDeEscrituracoes int     `json:"quantidade_de_escrituracoes"`
}
