package whosonfirst

import (
	"encoding/json"
	"github.com/whosonfirst/go-whosonfirst-brands"
	"io"
	"io/ioutil"
	_ "log"
	"os"
	"path/filepath"
)

type WOFBrand struct {
	brands.Brand      `json:",omitempty"`
	BrandId           int64    `json:"wof:brand_id"`
	BrandName         string   `json:"wof:brand_name"`
	BrandCategories   []string `json:"wof:categories,omitempty"`
	BrandTags         []string `json:"wof:tags,omitempty"`
	BrandSize         string   `json:"wof:brand_size"`
	BrandLastModified int      `json:"wof:lastmodified"`
}

func (b *WOFBrand) Id() int64 {
	return b.BrandId
}

func (b *WOFBrand) Name() string {
	return b.BrandName
}

func (b *WOFBrand) Size() string {
	return b.BrandSize
}

func (b *WOFBrand) String() string {
	return b.Name()
}

func LoadWOFBrandFromFile(path string) (brands.Brand, error) {

	abs_path, err := filepath.Abs(path)

	if err != nil {
		return nil, err
	}

	fh, err := os.Open(abs_path)

	if err != nil {
		return nil, err
	}

	defer fh.Close()

	return LoadWOFBrandFromReader(fh)
}

func LoadWOFBrandFromReader(fh io.ReadCloser) (brands.Brand, error) {

	body, err := ioutil.ReadAll(fh)

	if err != nil {
		return nil, err
	}

	var br WOFBrand
	err = json.Unmarshal(body, &br)

	if err != nil {
		return nil, err
	}

	return &br, nil
}

func UnmarshalBrand(body []byte) ([]byte, error) {

	var stub interface{}
	err := json.Unmarshal(body, &stub)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func UnmarshalBrandFromReader(fh io.Reader) ([]byte, error) {

	body, err := ioutil.ReadAll(fh)

	if err != nil {
		return nil, err
	}

	return UnmarshalBrand(body)
}

func UnmarshalBrandFromFile(path string) ([]byte, error) {

	fh, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer fh.Close()

	return UnmarshalBrandFromReader(fh)
}


/*
func NewWOFBrand(name string) (brands.Brand, error) {

	client := api.NewAPIClient()
	brand_id, err := client.CreateInteger()

	if err != nil {
		return nil, err
	}

	br := Brand{
		WOFId:   brand_id,
		WOFName: name,
		WOFSize: "",
	}

	return &br, nil
}
*/
