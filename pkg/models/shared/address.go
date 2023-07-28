// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Address struct {
	AddressLine1 *string `json:"address_line_1,omitempty"`
	AddressLine2 *string `json:"address_line_2,omitempty"`
	City         *string `json:"city,omitempty"`
	Country      *string `json:"country,omitempty"`
	PostalCode   *string `json:"postal_code,omitempty"`
}

func (o *Address) GetAddressLine1() *string {
	if o == nil {
		return nil
	}
	return o.AddressLine1
}

func (o *Address) GetAddressLine2() *string {
	if o == nil {
		return nil
	}
	return o.AddressLine2
}

func (o *Address) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *Address) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *Address) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}
