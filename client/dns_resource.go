//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// DNSResource implements api.DNSResource
type DNSResource struct {
	APIClient APIClient
}

func (d *DNSResource) client(id int) APIClient {
	return d.APIClient.GetSubObject(fmt.Sprintf("dnsresources/%v", id))
}

// Get fetches a given DNSResource
func (d *DNSResource) Get(id int) (dnsResource *entity.DNSResource, err error) {
	dnsResource = new(entity.DNSResource)
	err = d.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, dnsResource)
	})

	return
}

// Update updates a given DNSResource
func (d *DNSResource) Update(id int, params *entity.DNSResourceParams) (dnsResource *entity.DNSResource, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}

	dnsResource = new(entity.DNSResource)
	err = d.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, dnsResource)
	})

	return
}

// Delete deletes a given DNSResource
func (d *DNSResource) Delete(id int) error {
	return d.client(id).Delete()
}
