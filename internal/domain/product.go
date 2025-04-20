package domain

type Product struct {
	ID          int64
	Description string
	Tags        []string
	Quantity    int
}
