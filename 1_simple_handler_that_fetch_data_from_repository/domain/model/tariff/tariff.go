package tariff

// Tariff является бизнес-сущностью
type Tariff struct {
	ID    int64  // Уникальный ID тарифа
	Type  Type   // Тип услуги, чтобы уметь их различить
	Level int64  // Уровень услуги: Простой, Прокачанный, Мега, Гига и тп
	Slug  string // Строковый идентификатор
	Price int64  // Сколько стоит тариф в копейках
}

type Type int64

const (
	PayPerTransaction Type = 1
	PayPerAction      Type = 2
	PayPerListing     Type = 3
	Free              Type = 4
)
