package jsonDb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type shopCartRepo struct {
	fileName string
}

func NewShopCartRepo(fileName string) *shopCartRepo {
	return &shopCartRepo{
		fileName: fileName,
	}
}

func (s *shopCartRepo) Read() ([]models.ShopCart, error) {
	data, err := ioutil.ReadFile(s.fileName)
	if err != nil {
		return []models.ShopCart{}, err
	}

	var shopCarts []models.ShopCart
	err = json.Unmarshal(data, &shopCarts)
	if err != nil {
		return []models.ShopCart{}, err
	}
	return shopCarts, nil
}

func (s *shopCartRepo) AddShopCart(req *models.Add) (string, error) {
	shopCarts, err := s.Read()
	if err != nil {
		return "", err
	}

	uuid := uuid.New().String()
	shopCarts = append(shopCarts, models.ShopCart{
		Id:        uuid,
		ProductId: req.ProductId,
		UserId:    req.UserId,
		Count:     req.Count,
		Status:    false,
		// Time:      req.Time,
	})

	body, err := json.MarshalIndent(shopCarts, "", " ")
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(s.fileName, body, os.ModePerm)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func (s *shopCartRepo) RemoveShopCart(req *models.Remove) error {
	shopCarts, err := s.Read()
	if err != nil {
		return err
	}

	flag := true
	for i, v := range shopCarts {
		if req.ProductId == v.ProductId && req.UserId == v.UserId {
			shopCarts = append(shopCarts[:i], shopCarts[i+1:]...)
			flag = false
			break
		}
	}

	if flag {
		return errors.New("UserId or ProductId doesn't exist")
	}

	body, err := json.MarshalIndent(shopCarts, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (s *shopCartRepo) GetUserShopCart(req *models.UserPrimaryKey) ([]models.ShopCart, error) {
	shopCarts, err := s.Read()
	if err != nil {
		return []models.ShopCart{}, err
	}

	userShopCarts := []models.ShopCart{}
	for _, v := range shopCarts {
		if v.UserId == req.Id && v.Status == false {
			userShopCarts = append(userShopCarts, v)
		}
	}

	if len(userShopCarts) == 0 {
		return []models.ShopCart{}, errors.New("There are no unpaid products")
	}

	return userShopCarts, nil
}

func (s *shopCartRepo) UpdateShopCart(userId string) error {
	shopCarts, err := s.Read()
	if err != nil {
		return err
	}

	for i, v := range shopCarts {
		if v.UserId == userId && v.Status == false {
			shopCarts[i].Status = true
		}
	}

	body, err := json.MarshalIndent(shopCarts, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (p *shopCartRepo) GetAll() (models.GetListResponsePro, error) {
	products, err := p.Read()
	if err != nil {
		return models.GetListResponsePro{}, err
	}
	return models.GetListResponsePro{
		Products: products,
		Count:    len(products),
	}, nil
}
