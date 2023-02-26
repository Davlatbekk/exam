package models

type BranchPrimaryKey struct {
	Id string `json:"id"`
}

type Branch struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateBranch struct {
	Name string `json:"name"`
}

type UpdateBranch struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetListBranchRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetListBranchResponse struct {
	Count  int      `json:"count"`
	Branch []Branch `json:"branch"`
}
