// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/api/v2/cluster/outlier_detection.proto

package envoy_api_v2_cluster

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// Validate checks the field values on OutlierDetection with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *OutlierDetection) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetConsecutive_5Xx()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return OutlierDetectionValidationError{
					field:  "Consecutive_5Xx",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	if d := m.GetInterval(); d != nil {
		dur, err := types.DurationFromProto(d)
		if err != nil {
			return OutlierDetectionValidationError{
				field:  "Interval",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return OutlierDetectionValidationError{
				field:  "Interval",
				reason: "value must be greater than 0s",
			}
		}

	}

	if d := m.GetBaseEjectionTime(); d != nil {
		dur, err := types.DurationFromProto(d)
		if err != nil {
			return OutlierDetectionValidationError{
				field:  "BaseEjectionTime",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return OutlierDetectionValidationError{
				field:  "BaseEjectionTime",
				reason: "value must be greater than 0s",
			}
		}

	}

	if wrapper := m.GetMaxEjectionPercent(); wrapper != nil {

		if wrapper.GetValue() > 100 {
			return OutlierDetectionValidationError{
				field:  "MaxEjectionPercent",
				reason: "value must be less than or equal to 100",
			}
		}

	}

	if wrapper := m.GetEnforcingConsecutive_5Xx(); wrapper != nil {

		if wrapper.GetValue() > 100 {
			return OutlierDetectionValidationError{
				field:  "EnforcingConsecutive_5Xx",
				reason: "value must be less than or equal to 100",
			}
		}

	}

	if wrapper := m.GetEnforcingSuccessRate(); wrapper != nil {

		if wrapper.GetValue() > 100 {
			return OutlierDetectionValidationError{
				field:  "EnforcingSuccessRate",
				reason: "value must be less than or equal to 100",
			}
		}

	}

	{
		tmp := m.GetSuccessRateMinimumHosts()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return OutlierDetectionValidationError{
					field:  "SuccessRateMinimumHosts",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetSuccessRateRequestVolume()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return OutlierDetectionValidationError{
					field:  "SuccessRateRequestVolume",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetSuccessRateStdevFactor()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return OutlierDetectionValidationError{
					field:  "SuccessRateStdevFactor",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetConsecutiveGatewayFailure()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return OutlierDetectionValidationError{
					field:  "ConsecutiveGatewayFailure",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	if wrapper := m.GetEnforcingConsecutiveGatewayFailure(); wrapper != nil {

		if wrapper.GetValue() > 100 {
			return OutlierDetectionValidationError{
				field:  "EnforcingConsecutiveGatewayFailure",
				reason: "value must be less than or equal to 100",
			}
		}

	}

	// no validation rules for SplitExternalLocalOriginErrors

	{
		tmp := m.GetConsecutiveLocalOriginFailure()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return OutlierDetectionValidationError{
					field:  "ConsecutiveLocalOriginFailure",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	if wrapper := m.GetEnforcingConsecutiveLocalOriginFailure(); wrapper != nil {

		if wrapper.GetValue() > 100 {
			return OutlierDetectionValidationError{
				field:  "EnforcingConsecutiveLocalOriginFailure",
				reason: "value must be less than or equal to 100",
			}
		}

	}

	if wrapper := m.GetEnforcingLocalOriginSuccessRate(); wrapper != nil {

		if wrapper.GetValue() > 100 {
			return OutlierDetectionValidationError{
				field:  "EnforcingLocalOriginSuccessRate",
				reason: "value must be less than or equal to 100",
			}
		}

	}

	return nil
}

// OutlierDetectionValidationError is the validation error returned by
// OutlierDetection.Validate if the designated constraints aren't met.
type OutlierDetectionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OutlierDetectionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OutlierDetectionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OutlierDetectionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OutlierDetectionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OutlierDetectionValidationError) ErrorName() string { return "OutlierDetectionValidationError" }

// Error satisfies the builtin error interface
func (e OutlierDetectionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOutlierDetection.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OutlierDetectionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OutlierDetectionValidationError{}
