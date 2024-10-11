package simple_query_commission_for_tariff_handler

import (
	"context"
	"fmt"
	"github.com/ansakharov/few_options_of_code_composition/server_dto"
)

type Handler struct {
	repo TariffRepo
}

func New(repo TariffRepo) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Handle(_ context.Context, in *server_dto.QueryDtoIn, out *server_dto.QueryDtoOut) error {
	tariff, err := h.repo.GetTariffByID(in.TariffID)

	if err != nil {
		fmt.Println(err)

		return nil
	}

	out.AmountCommission = float64(tariff.Price)

	return nil
}
