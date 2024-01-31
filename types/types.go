package types

type SecretRequestModel struct {
	PlainText string `json:"plain_text"`
}

type GetSecretResponse struct {
	Secret string
}

type CreateSecretResponse struct {
	Id string
}
