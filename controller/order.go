package controller

import "app/models"

func (c *Controller) CreateOrder(req *models.CreateOrder) (string, error) {
	total := 0
	for _, v := range req.OrderItems {
		product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: v.Product_id})
		if err != nil {
			return "", err
		}
		total += int(product.Price) * v.Count
	}

	id, err := c.store.Order().Create(req, total)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *Controller) GetByIdOrder(req *models.OrderPrimaryKey) (models.GetOrder, error) {
	orders, err := c.store.Order().GetByID(req)
	if err != nil {
		return models.GetOrder{}, err
	}
	return orders, nil

}

func (c *Controller) DeleteOrder(req *models.OrderPrimaryKey) error {
	err := c.store.Order().Delete(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) UpdateOrder(req *models.UpdateOrder, orderID string) error {
	err := c.store.Order().Update(req, orderID)
	if err != nil {
		return err
	}
	return nil
}
