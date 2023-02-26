package jsonDb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type branchRepo struct {
	fileName string
}

func NewBranchRepo(fileName string) *branchRepo {
	return &branchRepo{
		fileName: fileName,
	}
}

func (p *branchRepo) Create(req *models.CreateBranch) (string, error) {
	branch, err := p.Read()
	if err != nil {
		return "", err
	}

	uuid := uuid.New().String()
	branch = append(branch, models.Branch{
		Id:   uuid,
		Name: req.Name,
	})

	body, err := json.MarshalIndent(branch, "", " ")
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)
	if err != nil {
		return "", err
	}
	return uuid, nil
}
func (p *branchRepo) GetByIDBranch(req *models.BranchPrimaryKey) (models.Branch, error) {
	branch, err := p.Read()
	if err != nil {
		return models.Branch{}, err
	}

	for _, v := range branch {
		if v.Id == req.Id {
			return v, nil
		}
	}

	return models.Branch{}, errors.New("There is no product with this id")
}

func (p *branchRepo) Read() ([]models.Branch, error) {
	data, err := ioutil.ReadFile(p.fileName)
	if err != nil {
		return []models.Branch{}, err
	}

	var branch []models.Branch
	err = json.Unmarshal(data, &branch)
	if err != nil {
		return []models.Branch{}, err
	}
	return branch, nil
}

func (p *branchRepo) DeleteBranch(req *models.BranchPrimaryKey) error {
	branch, err := p.Read()
	if err != nil {
		return err
	}
	flag := true
	for i, v := range branch {
		if v.Id == req.Id {
			branch = append(branch[:i], branch[i+1:]...)
			flag = false
			break
		}
	}

	if flag {
		return errors.New("There is no product with this id")
	}

	body, err := json.MarshalIndent(branch, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (p *branchRepo) Update(req *models.UpdateBranch, orderID string) error {
	branch, err := p.Read()
	if err != nil {
		return err
	}

	flag := true
	for i, v := range branch {
		if v.Id == orderID {
			branch[i].Name = req.Name

			flag = false
		}
	}

	if flag {
		return errors.New("There is no product with this id")
	}

	body, err := json.MarshalIndent(branch, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (u *branchRepo) GetAllBranch(req *models.GetListBranchRequest) (models.GetListBranchResponse, error) {
	branch, err := u.Read()
	if err != nil {
		return models.GetListBranchResponse{}, err
	}

	if req.Limit+req.Offset > len(branch) {
		return models.GetListBranchResponse{}, errors.New("out of range")
	}

	Branchs := []models.Branch{}
	for i := req.Offset; i < req.Offset+req.Limit; i++ {
		Branchs = append(Branchs, branch[i])
	}
	return models.GetListBranchResponse{
		Branch: Branchs,
		Count:  len(Branchs),
	}, nil
}
