// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: portal/member.proto

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

// Validate checks the field values on MemberRegisterReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MemberRegisterReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MemberRegisterReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MemberRegisterReqMultiError, or nil if none found.
func (m *MemberRegisterReq) ValidateAll() error {
	return m.validate(true)
}

func (m *MemberRegisterReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUsername()) < 1 {
		err := MemberRegisterReqValidationError{
			field:  "Username",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 1 {
		err := MemberRegisterReqValidationError{
			field:  "Password",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetTelephone()) < 1 {
		err := MemberRegisterReqValidationError{
			field:  "Telephone",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetAuthCode()) < 1 {
		err := MemberRegisterReqValidationError{
			field:  "AuthCode",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return MemberRegisterReqMultiError(errors)
	}

	return nil
}

// MemberRegisterReqMultiError is an error wrapping multiple validation errors
// returned by MemberRegisterReq.ValidateAll() if the designated constraints
// aren't met.
type MemberRegisterReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MemberRegisterReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MemberRegisterReqMultiError) AllErrors() []error { return m }

// MemberRegisterReqValidationError is the validation error returned by
// MemberRegisterReq.Validate if the designated constraints aren't met.
type MemberRegisterReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MemberRegisterReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MemberRegisterReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MemberRegisterReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MemberRegisterReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MemberRegisterReqValidationError) ErrorName() string {
	return "MemberRegisterReqValidationError"
}

// Error satisfies the builtin error interface
func (e MemberRegisterReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMemberRegisterReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MemberRegisterReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MemberRegisterReqValidationError{}

// Validate checks the field values on MemberLoginReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MemberLoginReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MemberLoginReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MemberLoginReqMultiError,
// or nil if none found.
func (m *MemberLoginReq) ValidateAll() error {
	return m.validate(true)
}

func (m *MemberLoginReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUsername()) < 1 {
		err := MemberLoginReqValidationError{
			field:  "Username",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 1 {
		err := MemberLoginReqValidationError{
			field:  "Password",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return MemberLoginReqMultiError(errors)
	}

	return nil
}

// MemberLoginReqMultiError is an error wrapping multiple validation errors
// returned by MemberLoginReq.ValidateAll() if the designated constraints
// aren't met.
type MemberLoginReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MemberLoginReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MemberLoginReqMultiError) AllErrors() []error { return m }

// MemberLoginReqValidationError is the validation error returned by
// MemberLoginReq.Validate if the designated constraints aren't met.
type MemberLoginReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MemberLoginReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MemberLoginReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MemberLoginReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MemberLoginReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MemberLoginReqValidationError) ErrorName() string { return "MemberLoginReqValidationError" }

// Error satisfies the builtin error interface
func (e MemberLoginReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMemberLoginReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MemberLoginReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MemberLoginReqValidationError{}

// Validate checks the field values on MemberLoginRsp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MemberLoginRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MemberLoginRsp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MemberLoginRspMultiError,
// or nil if none found.
func (m *MemberLoginRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *MemberLoginRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	// no validation rules for TokenHead

	if len(errors) > 0 {
		return MemberLoginRspMultiError(errors)
	}

	return nil
}

// MemberLoginRspMultiError is an error wrapping multiple validation errors
// returned by MemberLoginRsp.ValidateAll() if the designated constraints
// aren't met.
type MemberLoginRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MemberLoginRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MemberLoginRspMultiError) AllErrors() []error { return m }

// MemberLoginRspValidationError is the validation error returned by
// MemberLoginRsp.Validate if the designated constraints aren't met.
type MemberLoginRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MemberLoginRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MemberLoginRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MemberLoginRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MemberLoginRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MemberLoginRspValidationError) ErrorName() string { return "MemberLoginRspValidationError" }

// Error satisfies the builtin error interface
func (e MemberLoginRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMemberLoginRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MemberLoginRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MemberLoginRspValidationError{}

// Validate checks the field values on MemberInfoReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MemberInfoReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MemberInfoReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MemberInfoReqMultiError, or
// nil if none found.
func (m *MemberInfoReq) ValidateAll() error {
	return m.validate(true)
}

func (m *MemberInfoReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return MemberInfoReqMultiError(errors)
	}

	return nil
}

// MemberInfoReqMultiError is an error wrapping multiple validation errors
// returned by MemberInfoReq.ValidateAll() if the designated constraints
// aren't met.
type MemberInfoReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MemberInfoReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MemberInfoReqMultiError) AllErrors() []error { return m }

// MemberInfoReqValidationError is the validation error returned by
// MemberInfoReq.Validate if the designated constraints aren't met.
type MemberInfoReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MemberInfoReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MemberInfoReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MemberInfoReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MemberInfoReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MemberInfoReqValidationError) ErrorName() string { return "MemberInfoReqValidationError" }

// Error satisfies the builtin error interface
func (e MemberInfoReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMemberInfoReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MemberInfoReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MemberInfoReqValidationError{}

// Validate checks the field values on MemberGetAuthCodeReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MemberGetAuthCodeReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MemberGetAuthCodeReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MemberGetAuthCodeReqMultiError, or nil if none found.
func (m *MemberGetAuthCodeReq) ValidateAll() error {
	return m.validate(true)
}

func (m *MemberGetAuthCodeReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Telephone

	if len(errors) > 0 {
		return MemberGetAuthCodeReqMultiError(errors)
	}

	return nil
}

// MemberGetAuthCodeReqMultiError is an error wrapping multiple validation
// errors returned by MemberGetAuthCodeReq.ValidateAll() if the designated
// constraints aren't met.
type MemberGetAuthCodeReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MemberGetAuthCodeReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MemberGetAuthCodeReqMultiError) AllErrors() []error { return m }

// MemberGetAuthCodeReqValidationError is the validation error returned by
// MemberGetAuthCodeReq.Validate if the designated constraints aren't met.
type MemberGetAuthCodeReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MemberGetAuthCodeReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MemberGetAuthCodeReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MemberGetAuthCodeReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MemberGetAuthCodeReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MemberGetAuthCodeReqValidationError) ErrorName() string {
	return "MemberGetAuthCodeReqValidationError"
}

// Error satisfies the builtin error interface
func (e MemberGetAuthCodeReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMemberGetAuthCodeReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MemberGetAuthCodeReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MemberGetAuthCodeReqValidationError{}

// Validate checks the field values on MemberGetAuthCodeRsp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MemberGetAuthCodeRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MemberGetAuthCodeRsp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MemberGetAuthCodeRspMultiError, or nil if none found.
func (m *MemberGetAuthCodeRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *MemberGetAuthCodeRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AuthCode

	if len(errors) > 0 {
		return MemberGetAuthCodeRspMultiError(errors)
	}

	return nil
}

// MemberGetAuthCodeRspMultiError is an error wrapping multiple validation
// errors returned by MemberGetAuthCodeRsp.ValidateAll() if the designated
// constraints aren't met.
type MemberGetAuthCodeRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MemberGetAuthCodeRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MemberGetAuthCodeRspMultiError) AllErrors() []error { return m }

// MemberGetAuthCodeRspValidationError is the validation error returned by
// MemberGetAuthCodeRsp.Validate if the designated constraints aren't met.
type MemberGetAuthCodeRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MemberGetAuthCodeRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MemberGetAuthCodeRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MemberGetAuthCodeRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MemberGetAuthCodeRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MemberGetAuthCodeRspValidationError) ErrorName() string {
	return "MemberGetAuthCodeRspValidationError"
}

// Error satisfies the builtin error interface
func (e MemberGetAuthCodeRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMemberGetAuthCodeRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MemberGetAuthCodeRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MemberGetAuthCodeRspValidationError{}

// Validate checks the field values on MemberUpdatePasswordReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MemberUpdatePasswordReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MemberUpdatePasswordReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MemberUpdatePasswordReqMultiError, or nil if none found.
func (m *MemberUpdatePasswordReq) ValidateAll() error {
	return m.validate(true)
}

func (m *MemberUpdatePasswordReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTelephone()) < 1 {
		err := MemberUpdatePasswordReqValidationError{
			field:  "Telephone",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 1 {
		err := MemberUpdatePasswordReqValidationError{
			field:  "Password",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetAuthCode()) < 1 {
		err := MemberUpdatePasswordReqValidationError{
			field:  "AuthCode",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return MemberUpdatePasswordReqMultiError(errors)
	}

	return nil
}

// MemberUpdatePasswordReqMultiError is an error wrapping multiple validation
// errors returned by MemberUpdatePasswordReq.ValidateAll() if the designated
// constraints aren't met.
type MemberUpdatePasswordReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MemberUpdatePasswordReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MemberUpdatePasswordReqMultiError) AllErrors() []error { return m }

// MemberUpdatePasswordReqValidationError is the validation error returned by
// MemberUpdatePasswordReq.Validate if the designated constraints aren't met.
type MemberUpdatePasswordReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MemberUpdatePasswordReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MemberUpdatePasswordReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MemberUpdatePasswordReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MemberUpdatePasswordReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MemberUpdatePasswordReqValidationError) ErrorName() string {
	return "MemberUpdatePasswordReqValidationError"
}

// Error satisfies the builtin error interface
func (e MemberUpdatePasswordReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMemberUpdatePasswordReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MemberUpdatePasswordReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MemberUpdatePasswordReqValidationError{}

// Validate checks the field values on MemberUpdatePasswordRsp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MemberUpdatePasswordRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MemberUpdatePasswordRsp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MemberUpdatePasswordRspMultiError, or nil if none found.
func (m *MemberUpdatePasswordRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *MemberUpdatePasswordRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return MemberUpdatePasswordRspMultiError(errors)
	}

	return nil
}

// MemberUpdatePasswordRspMultiError is an error wrapping multiple validation
// errors returned by MemberUpdatePasswordRsp.ValidateAll() if the designated
// constraints aren't met.
type MemberUpdatePasswordRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MemberUpdatePasswordRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MemberUpdatePasswordRspMultiError) AllErrors() []error { return m }

// MemberUpdatePasswordRspValidationError is the validation error returned by
// MemberUpdatePasswordRsp.Validate if the designated constraints aren't met.
type MemberUpdatePasswordRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MemberUpdatePasswordRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MemberUpdatePasswordRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MemberUpdatePasswordRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MemberUpdatePasswordRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MemberUpdatePasswordRspValidationError) ErrorName() string {
	return "MemberUpdatePasswordRspValidationError"
}

// Error satisfies the builtin error interface
func (e MemberUpdatePasswordRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMemberUpdatePasswordRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MemberUpdatePasswordRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MemberUpdatePasswordRspValidationError{}

// Validate checks the field values on MemberRefreshTokenReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MemberRefreshTokenReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MemberRefreshTokenReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MemberRefreshTokenReqMultiError, or nil if none found.
func (m *MemberRefreshTokenReq) ValidateAll() error {
	return m.validate(true)
}

func (m *MemberRefreshTokenReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return MemberRefreshTokenReqMultiError(errors)
	}

	return nil
}

// MemberRefreshTokenReqMultiError is an error wrapping multiple validation
// errors returned by MemberRefreshTokenReq.ValidateAll() if the designated
// constraints aren't met.
type MemberRefreshTokenReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MemberRefreshTokenReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MemberRefreshTokenReqMultiError) AllErrors() []error { return m }

// MemberRefreshTokenReqValidationError is the validation error returned by
// MemberRefreshTokenReq.Validate if the designated constraints aren't met.
type MemberRefreshTokenReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MemberRefreshTokenReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MemberRefreshTokenReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MemberRefreshTokenReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MemberRefreshTokenReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MemberRefreshTokenReqValidationError) ErrorName() string {
	return "MemberRefreshTokenReqValidationError"
}

// Error satisfies the builtin error interface
func (e MemberRefreshTokenReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMemberRefreshTokenReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MemberRefreshTokenReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MemberRefreshTokenReqValidationError{}

// Validate checks the field values on MemberRefreshTokenRsp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MemberRefreshTokenRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MemberRefreshTokenRsp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MemberRefreshTokenRspMultiError, or nil if none found.
func (m *MemberRefreshTokenRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *MemberRefreshTokenRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return MemberRefreshTokenRspMultiError(errors)
	}

	return nil
}

// MemberRefreshTokenRspMultiError is an error wrapping multiple validation
// errors returned by MemberRefreshTokenRsp.ValidateAll() if the designated
// constraints aren't met.
type MemberRefreshTokenRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MemberRefreshTokenRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MemberRefreshTokenRspMultiError) AllErrors() []error { return m }

// MemberRefreshTokenRspValidationError is the validation error returned by
// MemberRefreshTokenRsp.Validate if the designated constraints aren't met.
type MemberRefreshTokenRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MemberRefreshTokenRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MemberRefreshTokenRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MemberRefreshTokenRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MemberRefreshTokenRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MemberRefreshTokenRspValidationError) ErrorName() string {
	return "MemberRefreshTokenRspValidationError"
}

// Error satisfies the builtin error interface
func (e MemberRefreshTokenRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMemberRefreshTokenRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MemberRefreshTokenRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MemberRefreshTokenRspValidationError{}

// Validate checks the field values on MemberInfoRsp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MemberInfoRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MemberInfoRsp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MemberInfoRspMultiError, or
// nil if none found.
func (m *MemberInfoRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *MemberInfoRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Nickname

	// no validation rules for Icon

	// no validation rules for Gender

	// no validation rules for Birthday

	// no validation rules for PersonalizedSignature

	// no validation rules for Phone

	// no validation rules for City

	// no validation rules for Job

	// no validation rules for MemberLevelId

	// no validation rules for SourceType

	// no validation rules for Integration

	// no validation rules for Growth

	// no validation rules for LuckeyCount

	// no validation rules for HistoryIntegration

	// no validation rules for Status

	// no validation rules for CreateTime

	if len(errors) > 0 {
		return MemberInfoRspMultiError(errors)
	}

	return nil
}

// MemberInfoRspMultiError is an error wrapping multiple validation errors
// returned by MemberInfoRsp.ValidateAll() if the designated constraints
// aren't met.
type MemberInfoRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MemberInfoRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MemberInfoRspMultiError) AllErrors() []error { return m }

// MemberInfoRspValidationError is the validation error returned by
// MemberInfoRsp.Validate if the designated constraints aren't met.
type MemberInfoRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MemberInfoRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MemberInfoRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MemberInfoRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MemberInfoRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MemberInfoRspValidationError) ErrorName() string { return "MemberInfoRspValidationError" }

// Error satisfies the builtin error interface
func (e MemberInfoRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMemberInfoRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MemberInfoRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MemberInfoRspValidationError{}