package jsonapi

type (
	Response struct {
		Data interface{} `json:"data"`
	}

	ResponseObject struct {
		Type      string      `json:"type"`
		ID        string      `json:"id"`
		Attribute interface{} `json:"attribute"`
	}
)
