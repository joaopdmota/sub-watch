package services

import (
	"cnpj-finder/application/domain"
	httpclient "cnpj-finder/infra/http/client"
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

type ReceitaCNPJService struct {
	client  httpclient.HTTPClient
}

func NewReceitaCNPJService(client httpclient.HTTPClient) *ReceitaCNPJService {
	return &ReceitaCNPJService{
		client: client,
	}
}

func (s *ReceitaCNPJService) GetByCNPJ(ctx context.Context, cnpj string) (*domain.CNPJ, error) {
	cnpjClean := strings.ReplaceAll(cnpj, ".", "")
	cnpjClean = strings.ReplaceAll(cnpjClean, "/", "")
	cnpjClean = strings.ReplaceAll(cnpjClean, "-", "")

	url := fmt.Sprintf("/api/cnpj/v1/%s", cnpjClean)

	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching CNPJ from ReceitaWS: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("ReceitaWS returned status %d", resp.StatusCode)
	}

	var result domain.CNPJ
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response from ReceitaWS: %w", err)
	}

	return &result, nil
}
