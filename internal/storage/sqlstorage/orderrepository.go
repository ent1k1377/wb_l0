package sqlstorage

import (
	"database/sql"
	"github.com/ent1k1377/wb_l0/internal/model"
)

type OrderRepository struct {
	store *Storage
}

func (r OrderRepository) Create(o *model.Order) error {
	db := r.store.db

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
