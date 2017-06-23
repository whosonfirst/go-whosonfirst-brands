package brands

import (
	"github.com/whosonfirst/go-brooklynintegers-api"
)

type Brand struct {
	WOFBrandId int64  `json:"wof:brand_id"`
	WOFName    string `json:"wof:name"`
}

func NewBrand(name string) (*Brand, error) {

	client := api.NewAPIClient()
	brand_id, err := client.CreateInteger()

	if err != nil {
		return nil, err
	}

	br := Brand{
		WOFBrandId: brand_id,
		WOFName:    name,
	}

	return &br, nil
}
