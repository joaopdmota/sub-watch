package services

import (
	"cnpj-finder/application/domain"
	httpclient "cnpj-finder/infra/http/client"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type ReceitaCPFService struct {
	client  httpclient.HTTPClient
}

func NewReceitaCPFService(client httpclient.HTTPClient) *ReceitaCPFService {

	return &ReceitaCPFService{
		client:  client,
	}
}

func (s *ReceitaCPFService) GetByCPF(ctx context.Context, cpf string) (*domain.CPF, error) {
	cpfClean := strings.ReplaceAll(cpf, ".", "")
	cpfClean = strings.ReplaceAll(cpfClean, "-", "")

	url := fmt.Sprintf("/v1/cpf/%s", cpfClean)

	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching CNPJ from ReceitaWS: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("ReceitaWS returned status %d", resp.StatusCode)
	}

	var result domain.CPF
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response from ReceitaWS: %w", err)
	}

	result.UltimaAtualizacao = time.Now()

	return &result, nil
}
