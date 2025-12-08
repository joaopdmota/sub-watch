package usecases

import (
	"cnpj-finder/application/domain"
	app_errors "cnpj-finder/application/errors"
	"cnpj-finder/application/services"
	"context"
	"errors"
	"fmt"
	"regexp"
)

type CNPJUseCaseOutput struct {
	CNPJ         string `json:"cnpj"`
	RazaoSocial  string `json:"razao_social"`
	NomeFantasia string `json:"nome_fantasia"`

	IdentificadorMatrizFilial    int    `json:"identificador_matriz_filial"`
	DescricaoIdentificadorMatriz string `json:"descricao_identificador_matriz_filial"`

	SituacaoCadastral          int    `json:"situacao_cadastral"`
	DescricaoSituacaoCadastral string `json:"descricao_situacao_cadastral"`
	MotivoSituacaoCadastral    int    `json:"motivo_situacao_cadastral"`
	DescricaoMotivoSituacao    string `json:"descricao_motivo_situacao_cadastral"`
	SituacaoEspecial           string `json:"situacao_especial"`
	DataSituacaoCadastral      string `json:"data_situacao_cadastral"`
	DataSituacaoEspecial       string `json:"data_situacao_especial"`

	CNAEFiscal          int                    `json:"cnae_fiscal"`
	CNAEFiscalDescricao string                 `json:"cnae_fiscal_descricao"`
	CNAESecundarios     []CNAESecundarioOutput `json:"cnaes_secundarios"`

	DataInicioAtividade string `json:"data_inicio_atividade"`

	DescricaoTipoDeLogradouro string `json:"descricao_tipo_de_logradouro"`
	Logradouro                string `json:"logradouro"`
	Numero                    string `json:"numero"`
	Complemento               string `json:"complemento"`
	Bairro                    string `json:"bairro"`
	CEP                       string `json:"cep"`
	UF                        string `json:"uf"`
	Municipio                 string `json:"municipio"`
	CodigoMunicipio           int    `json:"codigo_municipio"`
	CodigoMunicipioIBGE       int    `json:"codigo_municipio_ibge"`
	NomeCidadeNoExterior      string `json:"nome_cidade_no_exterior"`
	EnteFederativoResponsavel string `json:"ente_federativo_responsavel"`

	DDDtelefone1 string `json:"ddd_telefone_1"`
	DDDtelefone2 string `json:"ddd_telefone_2"`
	DDDFax       string `json:"ddd_fax"`
	Email        string `json:"email"`

	NaturezaJuridica         string  `json:"natureza_juridica"`
	CodigoNaturezaJuridica   int     `json:"codigo_natureza_juridica"`
	Porte                    string  `json:"porte"`
	CodigoPorte              int     `json:"codigo_porte"`
	CapitalSocial            float64 `json:"capital_social"`
	QualificacaoDoResponsavel int    `json:"qualificacao_do_responsavel"`

	RegimeTributario      []RegimeTributarioOutput `json:"regime_tributario"`
	OpcaoPeloSimples      bool                    `json:"opcao_pelo_simples"`
	DataOpcaoPeloSimples  string                  `json:"data_opcao_pelo_simples"`
	DataExclusaoDoSimples string                  `json:"data_exclusao_do_simples"`
	OpcaoPeloMEI          bool                    `json:"opcao_pelo_mei"`
	DataOpcaoPeloMEI      string                  `json:"data_opcao_pelo_mei"`
	DataExclusaoDoMEI     string                  `json:"data_exclusao_do_mei"`

	QSA []QSAOutput `json:"qsa"`
}

type QSAOutput struct {
	Pais                                 string `json:"pais"`
	CodigoPais                           int   `json:"codigo_pais"`
	NomeSocio                            string `json:"nome_socio"`
	CNPJCPFSocio                         string `json:"cnpj_cpf_do_socio"`
	FaixaEtaria                          string `json:"faixa_etaria"`
	CodigoFaixaEtaria                    int    `json:"codigo_faixa_etaria"`
	IdentificadorDeSocio                 int    `json:"identificador_de_socio"`
	QualificacaoSocio                    string `json:"qualificacao_socio"`
	CodigoQualificacaoSocio              int    `json:"codigo_qualificacao_socio"`
	DataEntradaSociedade                 string `json:"data_entrada_sociedade"`
	CPFRepresentanteLegal                string `json:"cpf_representante_legal"`
	NomeRepresentanteLegal               string `json:"nome_representante_legal"`
	QualificacaoRepresentanteLegal       string `json:"qualificacao_representante_legal"`
	CodigoQualificacaoRepresentanteLegal int    `json:"codigo_qualificacao_representante_legal"`
}

type CNAESecundarioOutput struct {
	Codigo    int    `json:"codigo"`
	Descricao string `json:"descricao"`
}

type RegimeTributarioOutput struct {
	Ano                       int     `json:"ano"`
	CNPJDaSCP                 string `json:"cnpj_da_scp"`
	FormaDeTributacao         string  `json:"forma_de_tributacao"`
	QuantidadeDeEscrituracoes int     `json:"quantidade_de_escrituracoes"`
}

var (
	ErrInvalidCNPJFormat = errors.New("invalid CNPJ format: must contain 14 digits")
)

type GetCNPJUseCase struct {
	service services.ReceitaCNPJService
}

func mapCNAESecundarios(src []domain.CNAE) []CNAESecundarioOutput {
	out := make([]CNAESecundarioOutput, len(src))
	for i, c := range src {
		out[i] = CNAESecundarioOutput{
			Codigo:    c.Codigo,
			Descricao: c.Descricao,
		}
	}
	return out
}

func mapQSA(src []domain.Socio) []QSAOutput {
	out := make([]QSAOutput, len(src))
	for i, s := range src {
		out[i] = QSAOutput{
			Pais:                                 s.Pais,
			CodigoPais:                           s.CodigoPais,
			NomeSocio:                            s.NomeSocio,
			CNPJCPFSocio:                         s.CNPJCPFSocio,
			FaixaEtaria:                          s.FaixaEtaria,
			CodigoFaixaEtaria:                    s.CodigoFaixaEtaria,
			IdentificadorDeSocio:                 s.IdentificadorDeSocio,
			QualificacaoSocio:                    s.QualificacaoSocio,
			CodigoQualificacaoSocio:              s.CodigoQualificacaoSocio,
			DataEntradaSociedade:                 s.DataEntradaSociedade,
			CPFRepresentanteLegal:                s.CPFRepresentanteLegal,
			NomeRepresentanteLegal:               s.NomeRepresentanteLegal,
			QualificacaoRepresentanteLegal:       s.QualificacaoRepresentanteLegal,
			CodigoQualificacaoRepresentanteLegal: s.CodigoQualificacaoRepresentanteLegal,
		}
	}
	return out
}

func mapRegimeTributario(src []domain.RegimeTributario) []RegimeTributarioOutput {
	out := make([]RegimeTributarioOutput, len(src))
	for i, r := range src {
		out[i] = RegimeTributarioOutput{
			Ano:                       r.Ano,
			CNPJDaSCP:                 r.CNPJDaSCP,
			FormaDeTributacao:         r.FormaDeTributacao,
			QuantidadeDeEscrituracoes: r.QuantidadeDeEscrituracoes,
		}
	}
	return out
}

func NewGetCNPJUseCase(service services.ReceitaCNPJService) *GetCNPJUseCase {
	return &GetCNPJUseCase{
		service: service,
	}
}

func (uc *GetCNPJUseCase) Execute(ctx context.Context, cnpj string) (*CNPJUseCaseOutput, *app_errors.Error) {
	if err := validateCNPJ(cnpj); err != nil {
		return nil, &app_errors.Error{
			Code:    400,
			Type:    app_errors.ERROR_BAD_REQUEST,
			Message: err.Error(),
		}
	}

	result, err := uc.service.GetByCNPJ(ctx, cnpj)
	if err != nil {
		fmt.Println(err)
		return nil, &app_errors.Error{
			Code:    500,
			Type:    app_errors.ERROR_UNKNOW,
			Message: "error fetching CNPJ data",
		}
	}

	output := &CNPJUseCaseOutput{
		CNPJ:         result.CNPJ,
		RazaoSocial:  result.RazaoSocial,
		NomeFantasia: result.NomeFantasia,

		IdentificadorMatrizFilial:    result.IdentificadorMatrizFilial,
		DescricaoIdentificadorMatriz: result.DescricaoMatrizFilial,

		SituacaoCadastral:          result.SituacaoCadastral,
		DescricaoSituacaoCadastral: result.DescricaoSituacaoCadastral,
		MotivoSituacaoCadastral:    result.MotivoSituacaoCadastral,
		DescricaoMotivoSituacao:    result.DescricaoMotivoSituacao,
		SituacaoEspecial:           result.SituacaoEspecial,
		DataSituacaoCadastral:      result.DataSituacaoCadastral,
		DataSituacaoEspecial:       result.DataSituacaoEspecial,

		DataInicioAtividade: result.DataInicioAtividade,

		CNAEFiscal:          result.CNAEFiscal,
		CNAEFiscalDescricao: result.CNAEFiscalDescricao,
		CNAESecundarios:     mapCNAESecundarios(result.CNAESecundarios),

		DescricaoTipoDeLogradouro: result.DescricaoTipoDeLogradouro,
		Logradouro:                result.Logradouro,
		Numero:                    result.Numero,
		Complemento:               result.Complemento,
		Bairro:                    result.Bairro,
		CEP:                       result.CEP,
		UF:                        result.UF,
		Municipio:                 result.Municipio,
		CodigoMunicipio:           result.CodigoMunicipio,
		CodigoMunicipioIBGE:       result.CodigoMunicipioIBGE,
		NomeCidadeNoExterior:      result.NomeCidadeNoExterior,
		EnteFederativoResponsavel: result.EnteFederativoResponsavel,

		DDDtelefone1: result.DDDtelefone1,
		DDDtelefone2: result.DDDtelefone2,
		DDDFax:       result.DDDFax,
		Email:        result.Email,

		NaturezaJuridica:       result.NaturezaJuridica,
		CodigoNaturezaJuridica: result.CodigoNaturezaJuridica,
		Porte:                  result.Porte,
		CodigoPorte:            result.CodigoPorte,
		CapitalSocial:          result.CapitalSocial,

		QualificacaoDoResponsavel: result.QualificacaoDoResponsavel,

		QSA: mapQSA(result.QSA),

		RegimeTributario:      mapRegimeTributario(result.RegimeTributario),
		OpcaoPeloSimples:      result.OpcaoPeloSimples,
		DataOpcaoPeloSimples:  result.DataOpcaoPeloSimples,
		DataExclusaoDoSimples: result.DataExclusaoDoSimples,
		OpcaoPeloMEI:          result.OpcaoPeloMEI,
		DataOpcaoPeloMEI:      result.DataOpcaoPeloMEI,
		DataExclusaoDoMEI:     result.DataExclusaoDoMEI,
	}

	return output, nil
}

func validateCNPJ(cnpj string) error {
	cnpjClean := regexp.MustCompile(`[^0-9]`).ReplaceAllString(cnpj, "")

	if len(cnpjClean) != 14 {
		return ErrInvalidCNPJFormat
	}

	return nil
}
