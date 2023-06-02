package sqlstorage

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ent1k1377/wb_l0/internal/cache"
	"github.com/ent1k1377/wb_l0/internal/model"
)

type OrderRepository struct {
	storage *Storage
	cache   cache.Cache
}

func (r OrderRepository) Create(o *model.Order) error {
	db := r.storage.db

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if err := CreateDelivery(tx, o); err != nil {
		tx.Rollback()
		return err
	}
	if err := CreatePayment(tx, o); err != nil {
		tx.Rollback()
		return err
	}
	if err := CreateOrder(tx, o); err != nil {
		tx.Rollback()
		return err
	}
	if err := CreateItem(tx, o); err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Println("create-order")
	return nil
}

func CreateDelivery(tx *sql.Tx, o *model.Order) error {
	err := tx.QueryRow("SELECT insert_delivery($1, $2, $3, $4, $5, $6, $7)",
		o.Delivery.Name,
		o.Delivery.Phone,
		o.Delivery.Zip,
		o.Delivery.City,
		o.Delivery.Address,
		o.Delivery.Region,
		o.Delivery.Email).Scan(&o.Delivery.ID)
	if err != nil {
		return err
	}

	return nil
}

func CreatePayment(tx *sql.Tx, o *model.Order) error {
	err := tx.QueryRow(
		"SELECT insert_payment($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		o.Payment.Transaction,
		o.Payment.RequestID,
		o.Payment.Currency,
		o.Payment.Provider,
		o.Payment.Amount,
		o.Payment.PaymentDT,
		o.Payment.Bank,
		o.Payment.DeliveryCost,
		o.Payment.GoodsTotal,
		o.Payment.CustomFee).Scan(&o.Payment.ID)
	if err != nil {
		return err
	}

	return nil
}

func CreateOrder(tx *sql.Tx, o *model.Order) error {
	err := tx.QueryRow("SELECT insert_order($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
		o.TrackNumber,
		o.Entry,
		o.Delivery.ID,
		o.Payment.ID,
		o.Locale,
		o.InternalSignature,
		o.CustomerID,
		o.DeliveryService,
		o.Shardkey,
		o.SmID,
		o.DateCreated,
		o.OofShard).Scan(&o.OrderUID)
	if err != nil {
		return err
	}
	return nil
}

func CreateItem(tx *sql.Tx, o *model.Order) error {
	for _, item := range o.Items {
		err := tx.QueryRow("SELECT insert_item($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
			o.OrderUID,
			item.ChrtID,
			item.TrackNumber,
			item.Price,
			item.RID,
			item.Name,
			item.Sale,
			item.Size,
			item.TotalPrice,
			item.NmID,
			item.Brand,
			item.Status).Scan(&item.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r OrderRepository) Get(id int) (string, error) {
	rows, err := r.storage.db.Query("SELECT * FROM get_order($1)", id)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	order := &model.Order{}

	order.Items = make([]model.Item, 0)
	isEmpty := true
	for rows.Next() {
		isEmpty = false
		item := model.Item{}
		err := rows.Scan(
			&order.Delivery.ID,
			&order.Delivery.Name,
			&order.Delivery.Phone,
			&order.Delivery.Zip,
			&order.Delivery.City,
			&order.Delivery.Address,
			&order.Delivery.Region,
			&order.Delivery.Email,
			&order.Payment.ID,
			&order.Payment.Transaction,
			&order.Payment.RequestID,
			&order.Payment.Currency,
			&order.Payment.Provider,
			&order.Payment.Amount,
			&order.Payment.PaymentDT,
			&order.Payment.Bank,
			&order.Payment.DeliveryCost,
			&order.Payment.GoodsTotal,
			&order.Payment.CustomFee,
			&order.OrderUID,
			&order.TrackNumber,
			&order.Entry,
			&order.Delivery.ID,
			&order.Payment.ID,
			&order.Locale,
			&order.InternalSignature,
			&order.CustomerID,
			&order.DeliveryService,
			&order.Shardkey,
			&order.SmID,
			&order.DateCreated,
			&order.OofShard,
			&item.ID,
			&item.ChrtID,
			&item.TrackNumber,
			&item.Price,
			&item.RID,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NmID,
			&item.Brand,
			&item.Status,
		)
		if err != nil {
			return "", err
		}

		order.Items = append(order.Items, item)
	}

	if isEmpty {
		return "", errors.New("id not found")
	}

	if err := rows.Err(); err != nil {
		return "", err
	}

	orderJson, err := json.Marshal(order)
	if err != nil {
		return "", err
	}
	return string(orderJson), nil
}
