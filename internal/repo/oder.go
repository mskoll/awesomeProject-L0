package repo

import (
	"awesomeProject-L0/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type OrderDB struct {
	db *sqlx.DB
}

func NewOrderDB(db *sqlx.DB) *OrderDB {
	return &OrderDB{db: db}
}

// CreateOrder добавление в БД order
func (r *OrderDB) CreateOrder(order model.Order) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var paymentId int
	paymentQuery := fmt.Sprintf("INSERT INTO %s (transaction, request_id, currency, provider, amount, "+
		"payment_dt, bank, delivery_cost, goods_total, custom_fee) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id", paymentTable)

	row := tx.QueryRow(paymentQuery, order.Payment.Transaction, order.Payment.RequestId, order.Payment.Currency,
		order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDt, order.Payment.Bank,
		order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee)
	if err := row.Scan(&paymentId); err != nil {
		tx.Rollback()
		return 0, err
	}

	var deliveryId int
	deliveryQuery := fmt.Sprintf("INSERT INTO %s (name, phone, zip, city, address, region, email) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", deliveryTable)

	row = tx.QueryRow(deliveryQuery, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip,
		order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err := row.Scan(&deliveryId); err != nil {
		tx.Rollback()
		return 0, err
	}

	var orderId int
	orderQuery := fmt.Sprintf("INSERT INTO %s (order_uid, track_number, entry, locale, internal_signature, "+
		"customer_id, delivery_service, shardkey, sm_id, oof_shard, payment_id, delivery_id) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id", orderTable)

	row = tx.QueryRow(orderQuery, order.OrderUid, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature,
		order.CustomerId, order.DeliveryService, order.Shardkey, order.SmId, order.OofShard, paymentId, deliveryId)
	if err := row.Scan(&orderId); err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, item := range order.Items {
		itemQuery := fmt.Sprintf("INSERT INTO %s (chrt_id, track_number, price, rid, name, sale, size, "+
			"total_price, nm_id, brand, status, order_id) "+
			"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)", itemTable)

		_, err := tx.Exec(itemQuery, item.ChrtId, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale,
			item.Size, item.TotalPrice, item.NmId, item.Brand, item.Status, orderId)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return orderId, tx.Commit()
}

// GetOrderById получение order из БД
func (r *OrderDB) GetOrderById(orderId int) (model.Order, error) {

	var orders model.Order
	orderQuery := fmt.Sprintf("SELECT ot.id, ot.order_uid, ot.track_number, ot.entry, ot.locale, "+
		"ot.internal_signature, ot.customer_id, ot.delivery_service, ot.shardkey, ot.sm_id, ot.oof_shard "+
		"FROM %s ot WHERE ot.id = $1", orderTable)
	err := r.db.Get(&orders, orderQuery, orderId)

	deliveryQuery := fmt.Sprintf("SELECT dt.name, dt.phone, dt.zip, dt.city, dt.address, dt.region, dt.email "+
		"FROM %s dt INNER JOIN %s ot ON dt.id = ot.delivery_id WHERE ot.id = $1", deliveryTable, orderTable)
	err = r.db.Get(&orders.Delivery, deliveryQuery, orderId)

	paymentQuery := fmt.Sprintf("SELECT pt.transaction, pt.request_id, pt.currency, pt.provider, pt.amount, "+
		"pt.payment_dt, pt.bank, pt.delivery_cost, pt.goods_total, pt.custom_fee "+
		"FROM %s pt INNER JOIN %s ot ON pt.id = ot.payment_id WHERE ot.id = $1", paymentTable, orderTable)
	err = r.db.Get(&orders.Payment, paymentQuery, orderId)

	itemsQuery := fmt.Sprintf("SELECT it.chrt_id, it.track_number, it.price, it.rid, it.name, it.sale, it.size,"+
		"it.total_price, it.nm_id, it.brand, it.status "+
		"FROM %s it INNER JOIN %s ot ON it.order_id = ot.id WHERE ot.id = $1", itemTable, orderTable)
	err = r.db.Select(&orders.Items, itemsQuery, orderId)

	return orders, err
}
