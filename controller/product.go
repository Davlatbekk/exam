package controller

import "app/models"

func (c *Controller) CreateProduct(req *models.CreateProduct) (string, error) {
	id, err := c.store.Product().Create(req)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *Controller) DeleteProduct(req *models.ProductPrimaryKey) error {
	err := c.store.Product().Delete(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) UpdateProduct(req *models.UpdateProduct, productId string) error {
	err := c.store.Product().Update(req, productId)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) GetByIdProduct(req *models.ProductPrimaryKey) (models.Product, error) {
	product, err := c.store.Product().GetByID(req)
	if err != nil {
		return models.Product{}, err
	}

	category, err := c.store.Category().GetByID(&models.CategoryPrimaryKey{Id: product.CategoryID})
	if err != nil {
		return models.Product{}, err
	}

	return models.Product{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Category: category,
	}, nil
}

func (c *Controller) GetAllProduct(req *models.GetListRequestProduct) (models.GetListResponseProduct, error) {
	products, err := c.store.Product().GetAllProduct(req)
	if err != nil {
		return models.GetListResponseProduct{}, err
	}
	return products, nil
}
