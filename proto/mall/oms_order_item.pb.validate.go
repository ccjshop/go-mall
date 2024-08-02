// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: model/oms_order_item.proto

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

// Validate checks the field values on OrderItem with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrderItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderItem with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrderItemMultiError, or nil
// if none found.
func (m *OrderItem) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for OrderId

	// no validation rules for OrderSn

	// no validation rules for ProductId

	// no validation rules for ProductPic

	// no validation rules for ProductName

	// no validation rules for ProductBrand

	// no validation rules for ProductPrice

	// no validation rules for ProductSn

	// no validation rules for ProductAttr

	// no validation rules for ProductQuantity

	// no validation rules for ProductCategoryId

	// no validation rules for ProductSkuId

	// no validation rules for ProductSkuCode

	// no validation rules for PromotionName

	// no validation rules for PromotionAmount

	// no validation rules for CouponAmount

	// no validation rules for IntegrationAmount

	// no validation rules for RealAmount

	// no validation rules for GiftIntegration

	// no validation rules for GiftGrowth

	if len(errors) > 0 {
		return OrderItemMultiError(errors)
	}

	return nil
}

// OrderItemMultiError is an error wrapping multiple validation errors returned
// by OrderItem.ValidateAll() if the designated constraints aren't met.
type OrderItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderItemMultiError) AllErrors() []error { return m }

// OrderItemValidationError is the validation error returned by
// OrderItem.Validate if the designated constraints aren't met.
type OrderItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderItemValidationError) ErrorName() string { return "OrderItemValidationError" }

// Error satisfies the builtin error interface
func (e OrderItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderItemValidationError{}