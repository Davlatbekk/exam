package jsonDb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type orderRepo struct {
	fileName string
}

func NewOrderRepo(fileName string) *orderRepo {
	return &orderRepo{
		fileName: fileName,
	}
}

func (p *orderRepo) Create(req *models.CreateOrder, total int) (string, error) {
	orders, err := p.Read()
	if err != nil {
		return "", err
	}

	uuid := uuid.New().String()
	orders = append(orders, models.Order{
		Id:               uuid,
		User_id:          req.User_id,
		Cutomer_name:     req.Cutomer_name,
		Customer_address: req.Customer_address,
		Customer_phone:   req.Customer_phone,
		Total:            total,
	})

	data, err := ioutil.ReadFile("./data/orders_item.json")
	if err != nil {
		return "", err
	}

	var oreders_item []models.OrderItems
	err = json.Unmarshal(data, &oreders_item)
	if err != nil {
		return "", err
	}
	for _, val := range req.OrderItems {
		oreders_item = append(oreders_item, models.OrderItems{
			Product_id: val.Product_id,
			Count:      val.Count,
			Order_id:   uuid,
		})
	}
	item, err := json.MarshalIndent(oreders_item, "", " ")
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile("./data/orders_item.json", item, os.ModePerm)
	if err != nil {
		return "", err
	}

	body, err := json.MarshalIndent(orders, "", " ")
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func (p *orderRepo) GetByID(req *models.OrderPrimaryKey) (models.GetOrder, error) {
	orders, err := p.Read()
	if err != nil {
		return models.GetOrder{}, err
	}

	data, err := ioutil.ReadFile("./data/orders_item.json")
	if err != nil {
		return models.GetOrder{}, err
	}

	response := models.GetOrder{}

	var oreders_item []models.OrderItems
	err = json.Unmarshal(data, &oreders_item)
	if err != nil {
		return models.GetOrder{}, err
	}

	flag := false
	for _, v := range orders {
		if v.Id == req.Id {
			response.Id = v.Id
			response.User_id = v.User_id
			response.Cutomer_name = v.Cutomer_name
			response.Customer_address = v.Customer_address
			response.Customer_phone = v.Customer_phone
			response.Total = v.Total

			flag = true
		}
	}
	if !flag {
		return models.GetOrder{}, errors.New("There is no product with this id")
	}
	for _, val := range oreders_item {
		if req.Id == val.Order_id {
			response.OrderItems = append(response.OrderItems, models.CreateOrderItems{

				Product_id: val.Product_id,
				Count:      val.Count,
			})
		}
	}
	return response, nil
}

func (p *orderRepo) Read() ([]models.Order, error) {
	data, err := ioutil.ReadFile(p.fileName)
	if err != nil {
		return []models.Order{}, err
	}

	var orders []models.Order
	err = json.Unmarshal(data, &orders)
	if err != nil {
		return []models.Order{}, err
	}
	return orders, nil
}

func (p *orderRepo) Delete(req *models.OrderPrimaryKey) error {
	orders, err := p.Read()
	if err != nil {
		return err
	}
	flag := true
	for i, v := range orders {
		if v.Id == req.Id {
			orders = append(orders[:i], orders[i+1:]...)
			flag = false
			break
		}
	}

	if flag {
		return errors.New("There is no product with this id")
	}

	body, err := json.MarshalIndent(orders, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (p *orderRepo) Update(req *models.UpdateOrder, orderID string) error {
	orders, err := p.Read()
	if err != nil {
		return err
	}

	flag := true
	for i, v := range orders {
		if v.Id == orderID {
			orders[i].Cutomer_name = req.Cutomer_name
			orders[i].Customer_address = req.Customer_address
			orders[i].Customer_phone = req.Customer_phone
			flag = false
		}
	}

	if flag {
		return errors.New("There is no product with this id")
	}

	body, err := json.MarshalIndent(orders, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// // func (p *productRepo) GetAll() (models.GetListProduct, error) {
// // 	products, err := p.Read()
// // 	if err != nil {
// // 		return models.GetListProduct{}, err
// // 	}
// // 	return models.GetListProduct{
// // 		Products: products,
// // 		Count:    len(products),
// // 	}, nil
// // }

// func (p *orderRepo) GetAllProduct(req *models.GetListRequestProduct) (models.GetListResponseProduct, error) {
// 	products, err := p.ReadWithCategory()
// 	if err != nil {
// 		return models.GetListResponseProduct{}, err
// 	}

// 	fProduct := []models.ProductWithCategory{}
// 	for i:=0; i<len(products); i++ {
// 		if products[i].CategoryID == req.CategoryID {
// 			fProduct = append(fProduct, products[i])
// 		}
// 	}

// 	if req.Limit+req.Offset > len(fProduct) {
// 		return models.GetListResponseProduct{}, errors.New("out of range")
// 	}

// 	fProduct = fProduct[req.Offset:req.Offset+req.Limit]
// 	return models.GetListResponseProduct{
// 		Products: fProduct,
// 		Count:    len(fProduct),
// 	}, nil
// }
