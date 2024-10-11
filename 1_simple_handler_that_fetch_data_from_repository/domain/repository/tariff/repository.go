package repository

import (
	"github.com/ansakharov/few_options_of_code_composition/1_simple_handler_that_fetch_data_from_repository/domain/model/tariff"
	"math/rand/v2"
)

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) GetTariffByID(id int64) (tariff.Tariff, error) {
	return tariff.Tariff{
		ID: id,
		// Получить случайный тип тарифа
		Type: tariff.Type(rand.IntN(int(tariff.Free) + 1)),
		// Получить уровень тарифа от 1 до 3
		Level: int64(rand.IntN(3) + 1),
		Slug:  "some_static_slug",
		Price: int64(rand.IntN(100000000)),
	}, nil
}
