//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// IPRanges implements api.IPRanges
type IPRanges struct {
	APIClient APIClient
}

func (i *IPRanges) client() APIClient {
	return i.APIClient.GetSubObject("ipranges")
}

// Get fetches a list of IPRange objects
func (i *IPRanges) Get() (ipRanges []entity.IPRange, err error) {
	err = i.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &ipRanges)
	})

	return
}

// Create creates a new IPRange object
func (i *IPRanges) Create(params *entity.IPRangeParams) (ipRange *entity.IPRange, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}

	ipRange = new(entity.IPRange)
	err = i.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, ipRange)
	})

	return
}
