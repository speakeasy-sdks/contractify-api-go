// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ContractTypeCollection - OK
type ContractTypeCollection struct {
	Data []ContractTypeRead `json:"data,omitempty"`
}

func (o *ContractTypeCollection) GetData() []ContractTypeRead {
	if o == nil {
		return nil
	}
	return o.Data
}
