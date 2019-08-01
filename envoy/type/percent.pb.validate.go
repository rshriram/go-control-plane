// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/type/percent.proto

package envoy_type

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

// Validate checks the field values on Percent with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Percent) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Value

	return nil
}

// PercentValidationError is the validation error returned by Percent.Validate
// if the designated constraints aren't met.
type PercentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PercentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PercentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PercentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PercentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PercentValidationError) ErrorName() string { return "PercentValidationError" }

// Error satisfies the builtin error interface
func (e PercentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPercent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PercentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PercentValidationError{}

// Validate checks the field values on FractionalPercent with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *FractionalPercent) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Numerator

	// no validation rules for Denominator

	return nil
}

// FractionalPercentValidationError is the validation error returned by
// FractionalPercent.Validate if the designated constraints aren't met.
type FractionalPercentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FractionalPercentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FractionalPercentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FractionalPercentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FractionalPercentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FractionalPercentValidationError) ErrorName() string {
	return "FractionalPercentValidationError"
}

// Error satisfies the builtin error interface
func (e FractionalPercentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFractionalPercent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FractionalPercentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FractionalPercentValidationError{}
