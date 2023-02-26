package models

type OrderPrimaryKey struct {
	Id string
}
type Order struct {
	Id               string
	User_id          string
	Cutomer_name     string
	Customer_address string
	Customer_phone   string
	Total            int
}

type CreateOrder struct {
	User_id          string
	Cutomer_name     string
	Customer_address string
	Customer_phone   string
	OrderItems       []CreateOrderItems
}
type CreateOrderItems struct {
	Product_id string
	Count      int
}

type OrderItems struct {
	Product_id string
	Count      int
	Order_id   string
}

type GetOrder struct {
	Id               string
	User_id          string
	Cutomer_name     string
	Customer_address string
	Customer_phone   string
	OrderItems       []CreateOrderItems
	Total            int
}

type UpdateOrder struct {
	Cutomer_name     string
	Customer_address string
	Customer_phone   string
}
