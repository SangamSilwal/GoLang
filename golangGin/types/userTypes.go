package NewUserType

type User struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"required,min=18,max=100"`
	Phone string `json:"phone"`
	Role  string `json:"role" binding:"required,oneof=admin user guest"`
}
