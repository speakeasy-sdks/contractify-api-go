// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DossierCollection struct {
	Data []DossierRead `json:"data,omitempty"`
	Meta *Pagination   `json:"meta,omitempty"`
}

func (o *DossierCollection) GetData() []DossierRead {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *DossierCollection) GetMeta() *Pagination {
	if o == nil {
		return nil
	}
	return o.Meta
}
