package repository

import (
	"context"

	"github.com/MochamadAkbar/ordent-test/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository interface {
	FindByID(ctx context.Context, id int) (entity.Product, bool)
}

type ProductRepositoryImpl struct {
	Conn *pgxpool.Pool
}

func (repository *ProductRepositoryImpl) FindByID(ctx context.Context, id int) (entity.Product, bool) {
	var row entity.Product

	statement := `SELECT "id", "sku", "name", "description", "price" FROM "products" WHERE "id" = $1;`

	err := repository.Conn.QueryRow(ctx, statement, id).Scan(&row.ID, &row.SKU, &row.Name, &row.Description, &row.Price)
	if err != nil {
		return row, false
	}

	return row, true
}
