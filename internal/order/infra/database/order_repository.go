package database

import (
	"database/sql"

	"github.com/cesarsturmer/gointensivo/internal/order/entity"
)

type OrderRepository struct {
	Db *sql.DB
}


func NewORderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (repository *OrderRepository) Save(order *entity.Order) error {

	conection, err := repository.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES(?,?,?,?)")

	if err != nil {
		return err
	}
	
	// ignorar a necessidade declarar uma var
	_, err = conection.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	
	if err != nil {
		return err
	}

	return nil

}