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
	"time"
)

// AcceptConsentRequest struct for AcceptConsentRequest
type AcceptConsentRequest struct {
	GrantAccessTokenAudience []string `json:"grant_access_token_audience,omitempty"`
	GrantScope []string `json:"grant_scope,omitempty"`
	HandledAt *time.Time `json:"handled_at,omitempty"`
	// Remember, if set to true, tells ORY Hydra to remember this consent authorization and reuse it if the same client asks the same user for the same, or a subset of, scope.
	Remember *bool `json:"remember,omitempty"`
	// RememberFor sets how long the consent authorization should be remembered for in seconds. If set to `0`, the authorization will be remembered indefinitely.
	RememberFor *int64 `json:"remember_for,omitempty"`
	Session *ConsentRequestSession `json:"session,omitempty"`
}

// NewAcceptConsentRequest instantiates a new AcceptConsentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptConsentRequest() *AcceptConsentRequest {
	this := AcceptConsentRequest{}
	return &this
}

// NewAcceptConsentRequestWithDefaults instantiates a new AcceptConsentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptConsentRequestWithDefaults() *AcceptConsentRequest {
	this := AcceptConsentRequest{}
	return &this
}

// GetGrantAccessTokenAudience returns the GrantAccessTokenAudience field value if set, zero value otherwise.
func (o *AcceptConsentRequest) GetGrantAccessTokenAudience() []string {
	if o == nil || o.GrantAccessTokenAudience == nil {
		var ret []string
		return ret
	}
	return o.GrantAccessTokenAudience
}

// GetGrantAccessTokenAudienceOk returns a tuple with the GrantAccessTokenAudience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptConsentRequest) GetGrantAccessTokenAudienceOk() ([]string, bool) {
	if o == nil || o.GrantAccessTokenAudience == nil {
		return nil, false
	}
	return o.GrantAccessTokenAudience, true
}

// HasGrantAccessTokenAudience returns a boolean if a field has been set.
func (o *AcceptConsentRequest) HasGrantAccessTokenAudience() bool {
	if o != nil && o.GrantAccessTokenAudience != nil {
		return true
	}

	return false
}

// SetGrantAccessTokenAudience gets a reference to the given []string and assigns it to the GrantAccessTokenAudience field.
func (o *AcceptConsentRequest) SetGrantAccessTokenAudience(v []string) {
	o.GrantAccessTokenAudience = v
}

// GetGrantScope returns the GrantScope field value if set, zero value otherwise.
func (o *AcceptConsentRequest) GetGrantScope() []string {
	if o == nil || o.GrantScope == nil {
		var ret []string
		return ret
	}
	return o.GrantScope
}

// GetGrantScopeOk returns a tuple with the GrantScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptConsentRequest) GetGrantScopeOk() ([]string, bool) {
	if o == nil || o.GrantScope == nil {
		return nil, false
	}
	return o.GrantScope, true
}

// HasGrantScope returns a boolean if a field has been set.
func (o *AcceptConsentRequest) HasGrantScope() bool {
	if o != nil && o.GrantScope != nil {
		return true
	}

	return false
}

// SetGrantScope gets a reference to the given []string and assigns it to the GrantScope field.
func (o *AcceptConsentRequest) SetGrantScope(v []string) {
	o.GrantScope = v
}

// GetHandledAt returns the HandledAt field value if set, zero value otherwise.
func (o *AcceptConsentRequest) GetHandledAt() time.Time {
	if o == nil || o.HandledAt == nil {
		var ret time.Time
		return ret
	}
	return *o.HandledAt
}

// GetHandledAtOk returns a tuple with the HandledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptConsentRequest) GetHandledAtOk() (*time.Time, bool) {
	if o == nil || o.HandledAt == nil {
		return nil, false
	}
	return o.HandledAt, true
}

// HasHandledAt returns a boolean if a field has been set.
func (o *AcceptConsentRequest) HasHandledAt() bool {
	if o != nil && o.HandledAt != nil {
		return true
	}

	return false
}

// SetHandledAt gets a reference to the given time.Time and assigns it to the HandledAt field.
func (o *AcceptConsentRequest) SetHandledAt(v time.Time) {
	o.HandledAt = &v
}

// GetRemember returns the Remember field value if set, zero value otherwise.
func (o *AcceptConsentRequest) GetRemember() bool {
	if o == nil || o.Remember == nil {
		var ret bool
		return ret
	}
	return *o.Remember
}

// GetRememberOk returns a tuple with the Remember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptConsentRequest) GetRememberOk() (*bool, bool) {
	if o == nil || o.Remember == nil {
		return nil, false
	}
	return o.Remember, true
}

// HasRemember returns a boolean if a field has been set.
func (o *AcceptConsentRequest) HasRemember() bool {
	if o != nil && o.Remember != nil {
		return true
	}

	return false
}

// SetRemember gets a reference to the given bool and assigns it to the Remember field.
func (o *AcceptConsentRequest) SetRemember(v bool) {
	o.Remember = &v
}

// GetRememberFor returns the RememberFor field value if set, zero value otherwise.
func (o *AcceptConsentRequest) GetRememberFor() int64 {
	if o == nil || o.RememberFor == nil {
		var ret int64
		return ret
	}
	return *o.RememberFor
}

// GetRememberForOk returns a tuple with the RememberFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptConsentRequest) GetRememberForOk() (*int64, bool) {
	if o == nil || o.RememberFor == nil {
		return nil, false
	}
	return o.RememberFor, true
}

// HasRememberFor returns a boolean if a field has been set.
func (o *AcceptConsentRequest) HasRememberFor() bool {
	if o != nil && o.RememberFor != nil {
		return true
	}

	return false
}

// SetRememberFor gets a reference to the given int64 and assigns it to the RememberFor field.
func (o *AcceptConsentRequest) SetRememberFor(v int64) {
	o.RememberFor = &v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *AcceptConsentRequest) GetSession() ConsentRequestSession {
	if o == nil || o.Session == nil {
		var ret ConsentRequestSession
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptConsentRequest) GetSessionOk() (*ConsentRequestSession, bool) {
	if o == nil || o.Session == nil {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *AcceptConsentRequest) HasSession() bool {
	if o != nil && o.Session != nil {
		return true
	}

	return false
}

// SetSession gets a reference to the given ConsentRequestSession and assigns it to the Session field.
func (o *AcceptConsentRequest) SetSession(v ConsentRequestSession) {
	o.Session = &v
}

func (o AcceptConsentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GrantAccessTokenAudience != nil {
		toSerialize["grant_access_token_audience"] = o.GrantAccessTokenAudience
	}
	if o.GrantScope != nil {
		toSerialize["grant_scope"] = o.GrantScope
	}
	if o.HandledAt != nil {
		toSerialize["handled_at"] = o.HandledAt
	}
	if o.Remember != nil {
		toSerialize["remember"] = o.Remember
	}
	if o.RememberFor != nil {
		toSerialize["remember_for"] = o.RememberFor
	}
	if o.Session != nil {
		toSerialize["session"] = o.Session
	}
	return json.Marshal(toSerialize)
}

type NullableAcceptConsentRequest struct {
	value *AcceptConsentRequest
	isSet bool
}

func (v NullableAcceptConsentRequest) Get() *AcceptConsentRequest {
	return v.value
}

func (v *NullableAcceptConsentRequest) Set(val *AcceptConsentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptConsentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptConsentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptConsentRequest(val *AcceptConsentRequest) *NullableAcceptConsentRequest {
	return &NullableAcceptConsentRequest{value: val, isSet: true}
}

func (v NullableAcceptConsentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptConsentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


