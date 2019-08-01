// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/http/ip_tagging/v2/ip_tagging.proto

package v2

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

// Validate checks the field values on IPTagging with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *IPTagging) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RequestType

	for idx, item := range m.GetIpTags() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IPTaggingValidationError{
					field:  fmt.Sprintf("IpTags[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// IPTaggingValidationError is the validation error returned by
// IPTagging.Validate if the designated constraints aren't met.
type IPTaggingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IPTaggingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IPTaggingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IPTaggingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IPTaggingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IPTaggingValidationError) ErrorName() string { return "IPTaggingValidationError" }

// Error satisfies the builtin error interface
func (e IPTaggingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIPTagging.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IPTaggingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IPTaggingValidationError{}

// Validate checks the field values on IPTagging_IPTag with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *IPTagging_IPTag) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for IpTagName

	for idx, item := range m.GetIpList() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IPTagging_IPTagValidationError{
					field:  fmt.Sprintf("IpList[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// IPTagging_IPTagValidationError is the validation error returned by
// IPTagging_IPTag.Validate if the designated constraints aren't met.
type IPTagging_IPTagValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IPTagging_IPTagValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IPTagging_IPTagValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IPTagging_IPTagValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IPTagging_IPTagValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IPTagging_IPTagValidationError) ErrorName() string { return "IPTagging_IPTagValidationError" }

// Error satisfies the builtin error interface
func (e IPTagging_IPTagValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIPTagging_IPTag.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IPTagging_IPTagValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IPTagging_IPTagValidationError{}
