// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CustomFieldValueRead struct {
	CustomFieldID *int64      `json:"custom_field_id,omitempty"`
	ID            *int64      `json:"id,omitempty"`
	Value         interface{} `json:"value,omitempty"`
}

func (o *CustomFieldValueRead) GetCustomFieldID() *int64 {
	if o == nil {
		return nil
	}
	return o.CustomFieldID
}

func (o *CustomFieldValueRead) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CustomFieldValueRead) GetValue() interface{} {
	if o == nil {
		return nil
	}
	return o.Value
}
