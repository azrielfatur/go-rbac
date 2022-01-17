package requests

type MenuInput struct {
	Name     string `json:"name" binding:"required"`
	Url      string `json:"url" binding:"required"`
	Icon     string `json:"icon"`
	Order    int    `json:"order"`
	ParentID int    `json:"parent_id"`
}
