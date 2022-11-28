package handler

type responseWithId struct {
	ID int `json:"id" binding:"required"`
}

type responseWithToken struct {
	Token string `json:"token" binding:"required"`
}
