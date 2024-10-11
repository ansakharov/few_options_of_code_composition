package simple_query_commission_for_tariff_handler

import "github.com/ansakharov/few_options_of_code_composition/1_simple_handler_that_fetch_data_from_repository/domain/model/tariff"

type TariffRepo interface {
	GetTariffByID(id int64) (tariff.Tariff, error)
}
