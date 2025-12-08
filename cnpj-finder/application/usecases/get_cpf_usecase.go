package usecases

import (
	app_errors "cnpj-finder/application/errors"
	"cnpj-finder/application/services"
	"context"
	"errors"
	"fmt"
	"regexp"
)

type CPFOutput struct {
	CPF               string `json:"cpf" example:"12345678900"`
	Nome              string `json:"nome" example:"João da Silva"`
	Situacao          string `json:"situacao" example:"Regular"`
	DataNascimento    string `json:"data_nascimento" example:"01/01/1990"`
	Sexo              string `json:"sexo" example:"M"`
	TituloEleitor     string `json:"titulo_eleitor" example:"123456789012"`
}

var (
	ErrInvalidCPFFormat = errors.New("invalid CPF format: must contain 11 digits")
)

type GetCPFUseCase struct {
	service services.ReceitaCPFService
}

func NewGetCPFUseCase(service services.ReceitaCPFService) *GetCPFUseCase {
	return &GetCPFUseCase{
		service: service,
	}
}

func (uc *GetCPFUseCase) Execute(ctx context.Context, cpf string) (*CPFOutput, *app_errors.Error) {
	if err := validateCPF(cpf); err != nil {
		return nil, &app_errors.Error{
			Code:    400,
			Type:    app_errors.ERROR_BAD_REQUEST,
			Message: err.Error(),
		}
	}

	result, err := uc.service.GetByCPF(ctx, cpf)
	if err != nil {
		fmt.Println(err)
		return nil, &app_errors.Error{
			Code:    500,
			Type:    app_errors.ERROR_UNKNOW,
			Message: "error fetching CPF data",
		}
	}

	output := &CPFOutput{
		CPF:            result.CPF,
		Nome:           result.Nome,
		Situacao:       result.Situacao,
		DataNascimento: result.DataNascimento,
		Sexo:           result.Sexo,
		TituloEleitor:  result.TituloEleitor,
	}

	return output, nil
}

func validateCPF(cpf string) error {
	cpfClean := regexp.MustCompile(`[^0-9]`).ReplaceAllString(cpf, "")

	if len(cpfClean) != 11 {
		return ErrInvalidCPFFormat
	}

	return nil
}
