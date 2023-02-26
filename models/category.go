package models

type Category struct {
	Id            string     `json:"id"`
	Name          string     `json:"name"`
	ParentID      string     `json:"parent_id"`
	// SubCategories []Category `json:"sub_categories"`
}

type CategoryPrimaryKey struct {
	Id string `json:"id"`
}

type CreateCategory struct {
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
}

type UpdateCategory struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
}

type GetListCategoryRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetListCategoryResponse struct {
	Count      int        `json:"count"`
	Categories []Category `json:"categories"`
}
