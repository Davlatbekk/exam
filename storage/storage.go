package storage

import (
	"app/models"
)

type StorageI interface {
	CloseDb()
	User() UserRepoI
	Product() ProductRepoI
	ShopCart() ShopCartRepoI
	Commission() CommissionRepoI
	Category() CategoryRepoI
	Order() OrderRepoI
	Branch() BranchRepoI
}

type UserRepoI interface {
	Create(*models.CreateUser) (string, error)
	Delete(*models.UserPrimaryKey) error
	Update(*models.UpdateUser, string) error
	GetByID(*models.UserPrimaryKey) (models.User, error)
	GetAll(*models.GetListRequest) (models.GetListResponse, error)
}

type ProductRepoI interface {
	Create(*models.CreateProduct) (string, error)
	GetByID(*models.ProductPrimaryKey) (models.ProductWithCategory, error)
	GetAllProduct(req *models.GetListRequestProduct) (models.GetListResponseProduct, error)

	Update(*models.UpdateProduct, string) error
	Delete(*models.ProductPrimaryKey) error
}

type ShopCartRepoI interface {
	AddShopCart(*models.Add) (string, error)
	RemoveShopCart(*models.Remove) error
	GetUserShopCart(*models.UserPrimaryKey) ([]models.ShopCart, error)
	UpdateShopCart(string) error
	GetAll() (models.GetListResponsePro, error)
}

type CommissionRepoI interface {
	AddCommission(*models.Commission) error
}

type CategoryRepoI interface {
	Create(*models.CreateCategory) (string, error)
	GetByID(*models.CategoryPrimaryKey) (models.Category, error)
	GetAll(*models.GetListCategoryRequest) (models.GetListCategoryResponse, error)
	Update(*models.UpdateCategory, string) error
	Delete(*models.CategoryPrimaryKey) error
}

type OrderRepoI interface {
	Create(req *models.CreateOrder, total int) (string, error)
	GetByID(req *models.OrderPrimaryKey) (models.GetOrder, error)
	Delete(req *models.OrderPrimaryKey) error
	Update(req *models.UpdateOrder, orderID string) error
}

type BranchRepoI interface {
	Create(req *models.CreateBranch) (string, error)
	GetByIDBranch(req *models.BranchPrimaryKey) (models.Branch, error)
	GetAllBranch(req *models.GetListBranchRequest) (models.GetListBranchResponse, error)
	Update(req *models.UpdateBranch, orderID string) error
	DeleteBranch(req *models.BranchPrimaryKey) error
}
