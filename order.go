package main

type Payment struct {
	Transaction  string `json:"transaction,omitempty"`
	RequestID    string `json:"request_id,omitempty"`
	Currency     string `json:"currency,omitempty"`
	Provider     string `json:"provider,omitempty"`
	Amount       int    `json:"amount,omitempty"`
	PaymentDt    int    `json:"payment_dt,omitempty"`
	Bank         string `json:"bank,omitempty"`
	DeliveryCost int    `json:"delivery_cost,omitempty"`
	GoodsTotal   int    `json:"goods_total,omitempty"`
	CustomFee    int    `json:"custom_fee,omitempty"`
}

type Delivery struct {
	Name    string `json:"name,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Zip     string `json:"zip,omitempty"`
	City    string `json:"city,omitempty"`
	Address string `json:"address,omitempty"`
	Region  string `json:"region,omitempty"`
	Email   string `json:"email,omitempty"`
}

type Item struct {
	ChrtID     string `json:"chrt_id,omitempty"`
	TrackNum   string `json:"track_number,omitempty"`
	Price      int    `json:"price,omitempty"`
	RID        string `json:"rid,omitempty"`
	Name       string `json:"name,omitempty"`
	Sale       int    `json:"sale,omitempty"`
	Size       string `json:"size,omitempty"`
	TotalPrice int    `json:"total_price,omitempty"`
	NmID       int    `json:"nm_id,omitempty"`
	Brand      string `json:"brand,omitempty"`
	Status     int    `json:"status,omitempty"`
}
type Order struct {
	UID               string `json:"order_uid,omitempty"`
	TrackNum          string `json:"track_number,omitempty"`
	Entry             string `json:"entry,omitempty"`
	Locale            string `json:"locale,omitempty"`
	InternalSignature string `json:"internal_signature,omitempty"`
	CustomerID        string `json:"customer_id,omitempty"`
	DeliveryService   string `json:"delivery_service,omitempty"`
	ShardKey          string `json:"shard_key,omitempty"`
	SmID              int    `json:"sm_id,omitempty"`
	DateCreated       string `json:"date_created,omitempty"`
	OofShard          string `json:"oof_shard,omitempty"`
	Payment           `json:"payment"`
	Delivery          `json:"delivery"`
	Items             []Item `json:"items,omitempty"`
}
