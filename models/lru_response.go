package models

type SetKeyResponse struct {
}

type GetKeyResponse struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	LastUsed string `json:"last_used"`
}
