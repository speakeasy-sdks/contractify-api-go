// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ContractCollection - OK
type ContractCollection struct {
	Data []ContractRead `json:"data,omitempty"`
	Meta *Pagination    `json:"meta,omitempty"`
}

func (o *ContractCollection) GetData() []ContractRead {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *ContractCollection) GetMeta() *Pagination {
	if o == nil {
		return nil
	}
	return o.Meta
}
