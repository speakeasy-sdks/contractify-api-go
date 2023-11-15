// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ContractAutomaticRenewal struct {
	NumberOfRenewals *int64  `json:"number_of_renewals,omitempty"`
	RenewalPeriod    *string `json:"renewal_period,omitempty"`
}

func (o *ContractAutomaticRenewal) GetNumberOfRenewals() *int64 {
	if o == nil {
		return nil
	}
	return o.NumberOfRenewals
}

func (o *ContractAutomaticRenewal) GetRenewalPeriod() *string {
	if o == nil {
		return nil
	}
	return o.RenewalPeriod
}
