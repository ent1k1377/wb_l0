package dr

import (
	"database/sql"
	"github.com/ent1k1377/wb_l0/internal/model"
)

type OrderRepository struct {
	dr *DR
}

func (r OrderRepository) Create(o *model.Order) error {
	db := r.dr.db

	if err := CreateDelivery(db, o); err != nil {
		return err
	}
	if err := CreatePayment(db, o); err != nil {
		return err
	}
	if err := CreateOrder(db, o); err != nil {
		return err
	}
	if err := CreateItem(db, o); err != nil {
		return err
	}

	return nil
}

func CreateDelivery(db *sql.DB, o *model.Order) error {
	err := db.QueryRow(
		"INSERT INTO delivery (name, phone, zip, city, address, region, email) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING delivery_id",
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

func CreatePayment(db *sql.DB, o *model.Order) error {
	err := db.QueryRow(
		"INSERT INTO payment (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING payment_id",
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

func CreateOrder(db *sql.DB, o *model.Order) error {
	err := db.QueryRow(
		"INSERT INTO orders (track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING order_uid",
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

func CreateItem(db *sql.DB, o *model.Order) error {
	for _, item := range o.Items {
		err := db.QueryRow(
			"INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) "+
				"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING item_id",
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
