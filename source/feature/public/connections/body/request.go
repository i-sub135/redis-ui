package body

type CreateConnectionRequest struct {
	Name     string `json:"name" binding:"required"`
	Addr     string `json:"addr" binding:"required"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type UpdateConnectionRequest struct {
	Name     string `json:"name" binding:"required"`
	Addr     string `json:"addr" binding:"required"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}
