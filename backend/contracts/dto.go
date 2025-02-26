package contracts

type PurchaseTokenRequest struct {
	Buyer    string `json:"buyer"`
	Quantity int64  `json:"quantity"`
}
