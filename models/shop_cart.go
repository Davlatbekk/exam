package models

type ShopCartPrimaryKey struct {
	Id string `json:"id"`
}

type ShopCart struct {
	Id        string `json:"id"`
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
	Time      string `json:"time"`
}

type Add struct {
	ProductId string `json:"productId"`
	UserId    string `json:"userID"`
	Count     int    `json:"count"`
	Time      string `json:"time"`
}

type Remove struct {
	ProductId string `json:"productId"`
	UserId    string `json:"userID"`
}

type GetListResponsePro struct {
	Count    int        `json:"count"`
	Products []ShopCart `json:"products"`
}

type History struct {
	Name  string
	Price float64
	Count int
	Total int
	Time  string
}
