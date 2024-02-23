package models

type SetKeyRequest struct {
	Key        string `json:"key" binding:"required"`
	Value      string `json:"value" binding:"required"`
	ExpiryTime int    `json:"expiry_time" binding:"required"`
}

type GetKeyRequest struct {
	Key string `json:"key" binding:"required"`
}
