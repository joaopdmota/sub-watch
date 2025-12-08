package handlers

import (
	app_errors "cnpj-finder/application/errors"
	"cnpj-finder/application/usecases"
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DocumentHandler struct {
	getCNPJUseCase *usecases.GetCNPJUseCase
}

func NewDocumentHandler(getCNPJUseCase *usecases.GetCNPJUseCase) *DocumentHandler {
	return &DocumentHandler{
		getCNPJUseCase: getCNPJUseCase,
	}
}

type DocumentResponse struct {
	Tipo string `json:"tipo" example:"CNPJ"`

	CNPJ                       *string `json:"cnpj,omitempty" example:"06990590000123"`
	RazaoSocial                *string `json:"razao_social,omitempty" example:"GOOGLE BRASIL INTERNET LTDA"`
	NomeFantasia               *string `json:"nome_fantasia,omitempty" example:"Google"`
	SituacaoCadastral          *string `json:"situacao_cadastral,omitempty" example:"2"`
	DescricaoSituacaoCadastral *string `json:"descricao_situacao_cadastral,omitempty" example:"ATIVA"`
	DataInicioAtividade        *string `json:"data_inicio_atividade,omitempty" example:"2004-09-01"`
	CNAEFiscal                 *string `json:"cnae_fiscal,omitempty" example:"6319400"`
	CNAEFiscalDescricao        *string `json:"cnae_fiscal_descricao,omitempty" example:"Portais, provedores de conteúdo e outros serviços de informação na internet"`
	Logradouro                 *string `json:"logradouro,omitempty" example:"AV BRIGADEIRO FARIA LIMA"`
	Numero                     *string `json:"numero,omitempty" example:"3477"`
	Bairro                     *string `json:"bairro,omitempty" example:"ITAIM BIBI"`
	CEP                        *string `json:"cep,omitempty" example:"04538133"`
	UF                         *string `json:"uf,omitempty" example:"SP"`
	Municipio                  *string `json:"municipio,omitempty" example:"SAO PAULO"`
	Telefone                   *string `json:"telefone,omitempty" example:"1123958400"`
	CapitalSocial              *string `json:"capital_social,omitempty" example:"200000000.00"`
	Porte                      *string `json:"porte,omitempty" example:"DEMAIS"`

	// CPF (futuro)
	CPF            *string `json:"cpf,omitempty" example:"12345678900"`
	Nome           *string `json:"nome,omitempty" example:"João da Silva"`
	Situacao       *string `json:"situacao,omitempty" example:"Regular"`
	DataNascimento *string `json:"data_nascimento,omitempty" example:"01/01/1990"`
	Sexo           *string `json:"sexo,omitempty" example:"M"`
}

// Get retrieves information for a CNPJ automatically
// @Summary Search CNPJ
// @Description Queries CNPJ (14 digits) data from the Federal Revenue. The type is detected automatically.
// @Tags Documents
// @Accept json
// @Produce json
// @Param document path string true "CPF or CNPJ number (with or without formatting)" example:"06990590000123"
// @Success 200 {object} DocumentResponse "Data found"
// @Failure 400 {object} ErrorResponse "Invalid document"
// @Failure 404 {object} ErrorResponse "Document not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /consulta/{document} [get]
func (h *DocumentHandler) Get(c echo.Context) error {
	numero := c.Param("document")

	numeroClean := regexp.MustCompile(`[^0-9]`).ReplaceAllString(numero, "")

	if len(numeroClean) == 14 {
		return h.handleCNPJ(c, numeroClean)
	} else {
		return c.JSON(http.StatusBadRequest, app_errors.CreateErrors(app_errors.Error{
			Code:    400,
			Type:    app_errors.ERROR_BAD_REQUEST,
			Message: "invalid number: must contain 11 digits (CPF) or 14 digits (CNPJ)",
		}))
	}
}

func (h *DocumentHandler) handleCNPJ(c echo.Context, cnpj string) error {
	result, err := h.getCNPJUseCase.Execute(c.Request().Context(), cnpj)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	if result == nil {
		return c.JSON(http.StatusNotFound, app_errors.CreateErrors(app_errors.Error{
			Code:    404,
			Type:    app_errors.ERROR_NOT_FOUND,
			Message: "document not found",
		}))
	}

	tipo := "CNPJ"

	situacaoStr := strconv.Itoa(result.SituacaoCadastral)
	cnaeStr := strconv.Itoa(result.CNAEFiscal)
	capitalStr := strconv.FormatFloat(result.CapitalSocial, 'f', 2, 64)
	telefone := result.DDDtelefone1 

	response := DocumentResponse{
		Tipo: tipo,

		CNPJ:                       &result.CNPJ,
		RazaoSocial:                &result.RazaoSocial,
		NomeFantasia:               &result.NomeFantasia,
		SituacaoCadastral:          &situacaoStr,
		DescricaoSituacaoCadastral: &result.DescricaoSituacaoCadastral,
		DataInicioAtividade:        &result.DataInicioAtividade,
		CNAEFiscal:                 &cnaeStr,
		CNAEFiscalDescricao:        &result.CNAEFiscalDescricao,
		Logradouro:                 &result.Logradouro,
		Numero:                     &result.Numero,
		Bairro:                     &result.Bairro,
		CEP:                        &result.CEP,
		UF:                         &result.UF,
		Municipio:                  &result.Municipio,
		Telefone:                   &telefone,
		CapitalSocial:              &capitalStr,
		Porte:                      &result.Porte,
	}

	return c.JSON(http.StatusOK, response)
}
