package contract

// Contract является бизнес сущностью.
type Contract struct {
	ID       int64 // ID уникальный идентификатор контракта
	Revision int64 // Revision в какой ревизии был контракт
}
