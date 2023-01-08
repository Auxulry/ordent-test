package repository

import (
	"context"

	"github.com/MochamadAkbar/ordent-test/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CartRepository interface {
	AddToCart(ctx context.Context, cart *entity.Cart) (entity.Cart, bool)
}

type CartRepositoryImpl struct {
	Conn *pgxpool.Pool
}

func (repository *CartRepositoryImpl) AddToCart(ctx context.Context, cart *entity.Cart) (entity.Cart, bool) {
	var row entity.Cart
	
	statement := `INSERT INTO "carts" ("cart_session_id", "product_id", "quantity") VALUES($1, $2, $3) RETURNING "id";`

	err := repository.Conn.QueryRow(ctx, statement, cart.CartSessionID, cart.ProductID, cart.Quantity).Scan(&row.ID)
	if err != nil {
		return row, false
	}

	return row, true
}
