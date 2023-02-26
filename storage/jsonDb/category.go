package jsonDb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type categoryRepo struct {
	fileName string
}

func NewCategoryRepo(fileName string) *categoryRepo {
	return &categoryRepo{
		fileName: fileName,
	}
}

func (c *categoryRepo) Create(req *models.CreateCategory) (string, error) {
	data, err := ioutil.ReadFile(c.fileName)
	if err != nil {
		return "", err
	}

	var categories []models.Category
	err = json.Unmarshal(data, &categories)
	if err != nil {
		return "", err
	}

	uuid := uuid.New().String()

	categories = append(categories, models.Category{
		Id:       uuid,
		Name:     req.Name,
		ParentID: req.ParentID,
	})

	body, err := json.MarshalIndent(categories, "", " ")
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(c.fileName, body, os.ModePerm)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func (u *categoryRepo) GetByID(req *models.CategoryPrimaryKey) (models.Category, error) {
	categories, err := u.Read()
	if err != nil {
		return models.Category{}, err
	}

	for _, v := range categories {
		if v.Id == req.Id {

			// for _, subCategory := range categories {
			// 	if v.Id == subCategory.ParentID {
			// 		v.SubCategories = append(v.SubCategories, subCategory)
			// 	}
			// }

			return v, nil
		}
	}

	return models.Category{}, errors.New("There is no user with this id")
}

func (u *categoryRepo) GetAll(req *models.GetListCategoryRequest) (models.GetListCategoryResponse, error) {
	categories, err := u.Read()
	if err != nil {
		return models.GetListCategoryResponse{}, err
	}

	if req.Limit+req.Offset > len(categories) {
		return models.GetListCategoryResponse{}, errors.New("out of range")
	}

	Categories := []models.Category{}
	for i := req.Offset; i < req.Offset+req.Limit; i++ {
		Categories = append(Categories, categories[i])
	}
	return models.GetListCategoryResponse{
		Categories: Categories,
		Count:      len(Categories),
	}, nil
}

func (u *categoryRepo) Update(req *models.UpdateCategory, userId string) error {
	categories, err := u.Read()
	if err != nil {
		return err
	}

	flag := true
	for i, v := range categories {
		if v.Id == userId {
			categories[i].Name = req.Name
			categories[i].ParentID = req.ParentID
			flag = false
		}
	}

	if flag {
		return errors.New("There is no user with this id")
	}

	body, err := json.MarshalIndent(categories, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (u *categoryRepo) Delete(req *models.CategoryPrimaryKey) error {
	categories, err := u.Read()
	if err != nil {
		return err
	}
	flag := true
	for i, v := range categories {
		if v.Id == req.Id {
			categories = append(categories[:i], categories[i+1:]...)
			flag = false
			break
		}
	}

	if flag {
		return errors.New("There is no user with this id")
	}

	body, err := json.MarshalIndent(categories, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (u *categoryRepo) Read() ([]models.Category, error) {
	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return []models.Category{}, err
	}

	var categories []models.Category
	err = json.Unmarshal(data, &categories)
	if err != nil {
		return []models.Category{}, err
	}
	return categories, nil
}
