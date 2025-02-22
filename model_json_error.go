/*
 * Ory Oathkeeper API
 *
 * Documentation for all of Ory Oathkeeper's APIs. 
 *
 * API version: v1.11.7
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// JsonError Error responses are sent when an error (e.g. unauthorized, bad request, ...) occurred.
type JsonError struct {
	// Name is the error name.
	Error *string `json:"error,omitempty"`
	// Debug contains debug information. This is usually not available and has to be enabled.
	ErrorDebug *string `json:"error_debug,omitempty"`
	// Description contains further information on the nature of the error.
	ErrorDescription *string `json:"error_description,omitempty"`
	// Code represents the error status code (404, 403, 401, ...).
	StatusCode *int64 `json:"status_code,omitempty"`
}

// NewJsonError instantiates a new JsonError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonError() *JsonError {
	this := JsonError{}
	return &this
}

// NewJsonErrorWithDefaults instantiates a new JsonError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonErrorWithDefaults() *JsonError {
	this := JsonError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *JsonError) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonError) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *JsonError) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *JsonError) SetError(v string) {
	o.Error = &v
}

// GetErrorDebug returns the ErrorDebug field value if set, zero value otherwise.
func (o *JsonError) GetErrorDebug() string {
	if o == nil || o.ErrorDebug == nil {
		var ret string
		return ret
	}
	return *o.ErrorDebug
}

// GetErrorDebugOk returns a tuple with the ErrorDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonError) GetErrorDebugOk() (*string, bool) {
	if o == nil || o.ErrorDebug == nil {
		return nil, false
	}
	return o.ErrorDebug, true
}

// HasErrorDebug returns a boolean if a field has been set.
func (o *JsonError) HasErrorDebug() bool {
	if o != nil && o.ErrorDebug != nil {
		return true
	}

	return false
}

// SetErrorDebug gets a reference to the given string and assigns it to the ErrorDebug field.
func (o *JsonError) SetErrorDebug(v string) {
	o.ErrorDebug = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *JsonError) GetErrorDescription() string {
	if o == nil || o.ErrorDescription == nil {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonError) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || o.ErrorDescription == nil {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *JsonError) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription != nil {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *JsonError) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *JsonError) GetStatusCode() int64 {
	if o == nil || o.StatusCode == nil {
		var ret int64
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonError) GetStatusCodeOk() (*int64, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *JsonError) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int64 and assigns it to the StatusCode field.
func (o *JsonError) SetStatusCode(v int64) {
	o.StatusCode = &v
}

func (o JsonError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.ErrorDebug != nil {
		toSerialize["error_debug"] = o.ErrorDebug
	}
	if o.ErrorDescription != nil {
		toSerialize["error_description"] = o.ErrorDescription
	}
	if o.StatusCode != nil {
		toSerialize["status_code"] = o.StatusCode
	}
	return json.Marshal(toSerialize)
}

type NullableJsonError struct {
	value *JsonError
	isSet bool
}

func (v NullableJsonError) Get() *JsonError {
	return v.value
}

func (v *NullableJsonError) Set(val *JsonError) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonError) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonError(val *JsonError) *NullableJsonError {
	return &NullableJsonError{value: val, isSet: true}
}

func (v NullableJsonError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


