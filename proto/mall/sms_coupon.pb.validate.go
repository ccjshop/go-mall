// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: model/sms_coupon.proto

package mall

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Coupon with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Coupon) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Coupon with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CouponMultiError, or nil if none found.
func (m *Coupon) ValidateAll() error {
	return m.validate(true)
}

func (m *Coupon) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Amount

	// no validation rules for Note

	// no validation rules for Code

	// no validation rules for Count

	// no validation rules for PublishCount

	// no validation rules for UseCount

	// no validation rules for ReceiveCount

	// no validation rules for Type

	// no validation rules for UseType

	// no validation rules for PerLimit

	// no validation rules for EnableTime

	// no validation rules for MemberLevel

	// no validation rules for MinPoint

	// no validation rules for Platform

	// no validation rules for StartTime

	// no validation rules for EndTime

	if len(errors) > 0 {
		return CouponMultiError(errors)
	}

	return nil
}

// CouponMultiError is an error wrapping multiple validation errors returned by
// Coupon.ValidateAll() if the designated constraints aren't met.
type CouponMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CouponMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CouponMultiError) AllErrors() []error { return m }

// CouponValidationError is the validation error returned by Coupon.Validate if
// the designated constraints aren't met.
type CouponValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CouponValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CouponValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CouponValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CouponValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CouponValidationError) ErrorName() string { return "CouponValidationError" }

// Error satisfies the builtin error interface
func (e CouponValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCoupon.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CouponValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CouponValidationError{}
