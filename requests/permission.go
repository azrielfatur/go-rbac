package requests

type PermissionInput struct {
	RoleID      int16  `json:"role_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}
