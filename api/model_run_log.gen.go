/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RunLog struct for RunLog
type RunLog struct {
	RunID   *string `json:"runID,omitempty"`
	Time    *string `json:"time,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewRunLog instantiates a new RunLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunLog() *RunLog {
	this := RunLog{}
	return &this
}

// NewRunLogWithDefaults instantiates a new RunLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunLogWithDefaults() *RunLog {
	this := RunLog{}
	return &this
}

// GetRunID returns the RunID field value if set, zero value otherwise.
func (o *RunLog) GetRunID() string {
	if o == nil || o.RunID == nil {
		var ret string
		return ret
	}
	return *o.RunID
}

// GetRunIDOk returns a tuple with the RunID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunLog) GetRunIDOk() (*string, bool) {
	if o == nil || o.RunID == nil {
		return nil, false
	}
	return o.RunID, true
}

// HasRunID returns a boolean if a field has been set.
func (o *RunLog) HasRunID() bool {
	if o != nil && o.RunID != nil {
		return true
	}

	return false
}

// SetRunID gets a reference to the given string and assigns it to the RunID field.
func (o *RunLog) SetRunID(v string) {
	o.RunID = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *RunLog) GetTime() string {
	if o == nil || o.Time == nil {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunLog) GetTimeOk() (*string, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *RunLog) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *RunLog) SetTime(v string) {
	o.Time = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RunLog) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunLog) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RunLog) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RunLog) SetMessage(v string) {
	o.Message = &v
}

func (o RunLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RunID != nil {
		toSerialize["runID"] = o.RunID
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableRunLog struct {
	value *RunLog
	isSet bool
}

func (v NullableRunLog) Get() *RunLog {
	return v.value
}

func (v *NullableRunLog) Set(val *RunLog) {
	v.value = val
	v.isSet = true
}

func (v NullableRunLog) IsSet() bool {
	return v.isSet
}

func (v *NullableRunLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunLog(val *RunLog) *NullableRunLog {
	return &NullableRunLog{value: val, isSet: true}
}

func (v NullableRunLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}