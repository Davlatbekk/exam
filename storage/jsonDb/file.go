package jsonDb

import (
	"app/config"
	"app/storage"
)

type Store struct {
	user       *userRepo
	product    *productRepo
	shopCart   *shopCartRepo
	commission *commissionRepo
	category   *categoryRepo
}

func NewFileJson(cfg *config.Config) (storage.StorageI, error) {
	return &Store{
		user:       NewUserRepo(cfg.UserFileName),
		product:    NewProductRepo(cfg.ProductFileName),
		shopCart:   NewShopCartRepo(cfg.ShopCartFileName),
		commission: NewCommissionRepo(cfg.CommissionFileName),
		category:   NewCategoryRepo(cfg.CategoryName),
	}, nil
}

func (s *Store) CloseDb() {}

func (s *Store) User() storage.UserRepoI {
	return s.user
}

func (s *Store) Product() storage.ProductRepoI {
	return s.product
}

func (s *Store) ShopCart() storage.ShopCartRepoI {
	return s.shopCart
}

func (s *Store) Commission() storage.CommissionRepoI {
	return s.commission
}

func (s *Store) Category() storage.CategoryRepoI {
	return s.category
}
