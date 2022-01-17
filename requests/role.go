package requests

type RoleInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}
