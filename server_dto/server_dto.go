package server_dto

type QueryDtoIn struct {
	TariffID int64 `json:"tariff_id"`
}

type QueryDtoOut struct {
	AmountCommission float64 `json:"amount_commission"`
}
