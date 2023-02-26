package models

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type Product struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
	Category Category `json:"category"`
}

type ProductWithCategory struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID string  `json:"category_id"`
}

type CreateProduct struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID string  `json:"category_id"`
}

type UpdateProduct struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID string  `json:"category_id"`
}

type GetListResponseProduct struct {
	Count    int                   `json:"count"`
	Products []ProductWithCategory `json:"products"`
}



type GetListRequestProduct struct {
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	CategoryID string `json:"category_id"`
}
