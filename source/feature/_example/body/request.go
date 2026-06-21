package body

// {{FEATURE_CAMEL}}Request example request body
type {{FEATURE_CAMEL}}Request struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}