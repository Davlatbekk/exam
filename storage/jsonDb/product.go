package jsonDb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type productRepo struct {
	fileName string
}

func NewProductRepo(fileName string) *productRepo {
	return &productRepo{
		fileName: fileName,
	}
}

func (p *productRepo) Create(req *models.CreateProduct) (string, error) {
	products, err := p.ReadWithCategory()
	if err != nil {
		return "", err
	}

	uuid := uuid.New().String()
	products = append(products, models.ProductWithCategory{
		Id:         uuid,
		Name:       req.Name,
		Price:      req.Price,
		CategoryID: req.CategoryID,
	})

	body, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func (p *productRepo) Delete(req *models.ProductPrimaryKey) error {
	products, err := p.Read()
	if err != nil {
		return err
	}
	flag := true
	for i, v := range products {
		if v.Id == req.Id {
			products = append(products[:i], products[i+1:]...)
			flag = false
			break
		}
	}

	if flag {
		return errors.New("There is no product with this id")
	}

	body, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepo) Update(req *models.UpdateProduct, productId string) error {
	products, err := p.Read()
	if err != nil {
		return err
	}

	flag := true
	for i, v := range products {
		if v.Id == productId {
			products[i].Name = req.Name
			products[i].Price = req.Price
			flag = false
		}
	}

	if flag {
		return errors.New("There is no product with this id")
	}

	body, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepo) GetByID(req *models.ProductPrimaryKey) (models.ProductWithCategory, error) {
	products, err := p.ReadWithCategory()
	if err != nil {
		return models.ProductWithCategory{}, err
	}

	for _, v := range products {
		if v.Id == req.Id {
			return v, nil
		}
	}

	return models.ProductWithCategory{}, errors.New("There is no product with this id")
}


func (p *productRepo) GetAllProduct(req *models.GetListRequestProduct) (models.GetListResponseProduct, error) {
	products, err := p.ReadWithCategory()
	if err != nil {
		return models.GetListResponseProduct{}, err
	}

	fProduct := []models.ProductWithCategory{}
	for i := 0; i < len(products); i++ {
		if products[i].CategoryID == req.CategoryID {
			fProduct = append(fProduct, products[i])
		}
	}

	if req.Limit+req.Offset > len(fProduct) {
		return models.GetListResponseProduct{}, errors.New("out of range")
	}

	fProduct = fProduct[req.Offset : req.Offset+req.Limit]
	return models.GetListResponseProduct{
		Products: fProduct,
		Count:    len(fProduct),
	}, nil
}

func (p *productRepo) Read() ([]models.Product, error) {
	data, err := ioutil.ReadFile(p.fileName)
	if err != nil {
		return []models.Product{}, err
	}

	var products []models.Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}

func (p *productRepo) ReadWithCategory() ([]models.ProductWithCategory, error) {
	data, err := ioutil.ReadFile(p.fileName)
	if err != nil {
		return []models.ProductWithCategory{}, err
	}

	var products []models.ProductWithCategory
	err = json.Unmarshal(data, &products)
	if err != nil {
		return []models.ProductWithCategory{}, err
	}
	return products, nil
}
