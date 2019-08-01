// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/api/v2/cluster/outlier_detection.proto

package cluster

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

	if v, ok := interface{}(m.GetConsecutive_5Xx()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OutlierDetectionValidationError{
				field:  "Consecutive_5Xx",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OutlierDetectionValidationError{
				field:  "Interval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetBaseEjectionTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OutlierDetectionValidationError{
				field:  "BaseEjectionTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMaxEjectionPercent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OutlierDetectionValidationError{
				field:  "MaxEjectionPercent",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetEnforcingConsecutive_5Xx()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OutlierDetectionValidationError{
				field:  "EnforcingConsecutive_5Xx",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetEnforcingSuccessRate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OutlierDetectionValidationError{
				field:  "EnforcingSuccessRate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSuccessRateMinimumHosts()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OutlierDetectionValidationError{
				field:  "SuccessRateMinimumHosts",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSuccessRateRequestVolume()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OutlierDetectionValidationError{
				field:  "SuccessRateRequestVolume",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSuccessRateStdevFactor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OutlierDetectionValidationError{
				field:  "SuccessRateStdevFactor",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetConsecutiveGatewayFailure()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OutlierDetectionValidationError{
				field:  "ConsecutiveGatewayFailure",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetEnforcingConsecutiveGatewayFailure()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OutlierDetectionValidationError{
				field:  "EnforcingConsecutiveGatewayFailure",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for SplitExternalLocalOriginErrors

	if v, ok := interface{}(m.GetConsecutiveLocalOriginFailure()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OutlierDetectionValidationError{
				field:  "ConsecutiveLocalOriginFailure",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetEnforcingConsecutiveLocalOriginFailure()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OutlierDetectionValidationError{
				field:  "EnforcingConsecutiveLocalOriginFailure",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetEnforcingLocalOriginSuccessRate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OutlierDetectionValidationError{
				field:  "EnforcingLocalOriginSuccessRate",
				reason: "embedded message failed validation",
				cause:  err,
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
