// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: sequin/v1/sequin.proto

package sequinv1

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

// Validate checks the field values on ExecRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ExecRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExecRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ExecRequestMultiError, or
// nil if none found.
func (m *ExecRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ExecRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RequestId

	if all {
		switch v := interface{}(m.GetOperation()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ExecRequestValidationError{
					field:  "Operation",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ExecRequestValidationError{
					field:  "Operation",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOperation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecRequestValidationError{
				field:  "Operation",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ExecRequestValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ExecRequestValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecRequestValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ExecRequestMultiError(errors)
	}

	return nil
}

// ExecRequestMultiError is an error wrapping multiple validation errors
// returned by ExecRequest.ValidateAll() if the designated constraints aren't met.
type ExecRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExecRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExecRequestMultiError) AllErrors() []error { return m }

// ExecRequestValidationError is the validation error returned by
// ExecRequest.Validate if the designated constraints aren't met.
type ExecRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecRequestValidationError) ErrorName() string { return "ExecRequestValidationError" }

// Error satisfies the builtin error interface
func (e ExecRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecRequestValidationError{}

// Validate checks the field values on StartRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StartRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StartRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StartRequestMultiError, or
// nil if none found.
func (m *StartRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StartRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RequestId

	if all {
		switch v := interface{}(m.GetOperation()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StartRequestValidationError{
					field:  "Operation",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StartRequestValidationError{
					field:  "Operation",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOperation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StartRequestValidationError{
				field:  "Operation",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StartRequestValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StartRequestValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StartRequestValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return StartRequestMultiError(errors)
	}

	return nil
}

// StartRequestMultiError is an error wrapping multiple validation errors
// returned by StartRequest.ValidateAll() if the designated constraints aren't met.
type StartRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StartRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StartRequestMultiError) AllErrors() []error { return m }

// StartRequestValidationError is the validation error returned by
// StartRequest.Validate if the designated constraints aren't met.
type StartRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StartRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StartRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StartRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StartRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StartRequestValidationError) ErrorName() string { return "StartRequestValidationError" }

// Error satisfies the builtin error interface
func (e StartRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStartRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StartRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StartRequestValidationError{}

// Validate checks the field values on StartResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StartResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StartResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StartResponseMultiError, or
// nil if none found.
func (m *StartResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StartResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return StartResponseMultiError(errors)
	}

	return nil
}

// StartResponseMultiError is an error wrapping multiple validation errors
// returned by StartResponse.ValidateAll() if the designated constraints
// aren't met.
type StartResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StartResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StartResponseMultiError) AllErrors() []error { return m }

// StartResponseValidationError is the validation error returned by
// StartResponse.Validate if the designated constraints aren't met.
type StartResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StartResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StartResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StartResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StartResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StartResponseValidationError) ErrorName() string { return "StartResponseValidationError" }

// Error satisfies the builtin error interface
func (e StartResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStartResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StartResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StartResponseValidationError{}

// Validate checks the field values on GetRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetRequestMultiError, or
// nil if none found.
func (m *GetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RequestId

	// no validation rules for LastUpdateId

	if len(errors) > 0 {
		return GetRequestMultiError(errors)
	}

	return nil
}

// GetRequestMultiError is an error wrapping multiple validation errors
// returned by GetRequest.ValidateAll() if the designated constraints aren't met.
type GetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRequestMultiError) AllErrors() []error { return m }

// GetRequestValidationError is the validation error returned by
// GetRequest.Validate if the designated constraints aren't met.
type GetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRequestValidationError) ErrorName() string { return "GetRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRequestValidationError{}

// Validate checks the field values on GetResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetResponseMultiError, or
// nil if none found.
func (m *GetResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetResponseValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetResponseValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetResponseValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetResponseValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetResponseValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetResponseValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for LastUpdateId

	if len(errors) > 0 {
		return GetResponseMultiError(errors)
	}

	return nil
}

// GetResponseMultiError is an error wrapping multiple validation errors
// returned by GetResponse.ValidateAll() if the designated constraints aren't met.
type GetResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetResponseMultiError) AllErrors() []error { return m }

// GetResponseValidationError is the validation error returned by
// GetResponse.Validate if the designated constraints aren't met.
type GetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetResponseValidationError) ErrorName() string { return "GetResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetResponseValidationError{}

// Validate checks the field values on FuncOperation with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FuncOperation) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FuncOperation with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FuncOperationMultiError, or
// nil if none found.
func (m *FuncOperation) ValidateAll() error {
	return m.validate(true)
}

func (m *FuncOperation) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	for idx, item := range m.GetArgs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FuncOperationValidationError{
						field:  fmt.Sprintf("Args[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FuncOperationValidationError{
						field:  fmt.Sprintf("Args[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FuncOperationValidationError{
					field:  fmt.Sprintf("Args[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return FuncOperationMultiError(errors)
	}

	return nil
}

// FuncOperationMultiError is an error wrapping multiple validation errors
// returned by FuncOperation.ValidateAll() if the designated constraints
// aren't met.
type FuncOperationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FuncOperationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FuncOperationMultiError) AllErrors() []error { return m }

// FuncOperationValidationError is the validation error returned by
// FuncOperation.Validate if the designated constraints aren't met.
type FuncOperationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FuncOperationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FuncOperationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FuncOperationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FuncOperationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FuncOperationValidationError) ErrorName() string { return "FuncOperationValidationError" }

// Error satisfies the builtin error interface
func (e FuncOperationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFuncOperation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FuncOperationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FuncOperationValidationError{}

// Validate checks the field values on ExecResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ExecResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExecResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ExecResponseMultiError, or
// nil if none found.
func (m *ExecResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ExecResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ExecResponseValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ExecResponseValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExecResponseValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ExecResponseValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ExecResponseValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecResponseValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ExecResponseMultiError(errors)
	}

	return nil
}

// ExecResponseMultiError is an error wrapping multiple validation errors
// returned by ExecResponse.ValidateAll() if the designated constraints aren't met.
type ExecResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExecResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExecResponseMultiError) AllErrors() []error { return m }

// ExecResponseValidationError is the validation error returned by
// ExecResponse.Validate if the designated constraints aren't met.
type ExecResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecResponseValidationError) ErrorName() string { return "ExecResponseValidationError" }

// Error satisfies the builtin error interface
func (e ExecResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecResponseValidationError{}

// Validate checks the field values on RequestMetadata with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RequestMetadata) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RequestMetadata with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RequestMetadataMultiError, or nil if none found.
func (m *RequestMetadata) ValidateAll() error {
	return m.validate(true)
}

func (m *RequestMetadata) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ClientVersion

	// no validation rules for Labels

	if len(errors) > 0 {
		return RequestMetadataMultiError(errors)
	}

	return nil
}

// RequestMetadataMultiError is an error wrapping multiple validation errors
// returned by RequestMetadata.ValidateAll() if the designated constraints
// aren't met.
type RequestMetadataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RequestMetadataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RequestMetadataMultiError) AllErrors() []error { return m }

// RequestMetadataValidationError is the validation error returned by
// RequestMetadata.Validate if the designated constraints aren't met.
type RequestMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestMetadataValidationError) ErrorName() string { return "RequestMetadataValidationError" }

// Error satisfies the builtin error interface
func (e RequestMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequestMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestMetadataValidationError{}

// Validate checks the field values on RunMetadata with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RunMetadata) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RunMetadata with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RunMetadataMultiError, or
// nil if none found.
func (m *RunMetadata) ValidateAll() error {
	return m.validate(true)
}

func (m *RunMetadata) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Labels

	// no validation rules for Submitter

	if all {
		switch v := interface{}(m.GetSubmittedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RunMetadataValidationError{
					field:  "SubmittedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RunMetadataValidationError{
					field:  "SubmittedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSubmittedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RunMetadataValidationError{
				field:  "SubmittedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetStartedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RunMetadataValidationError{
					field:  "StartedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RunMetadataValidationError{
					field:  "StartedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStartedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RunMetadataValidationError{
				field:  "StartedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetFinishedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RunMetadataValidationError{
					field:  "FinishedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RunMetadataValidationError{
					field:  "FinishedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFinishedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RunMetadataValidationError{
				field:  "FinishedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ServerVersion

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RunMetadataValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RunMetadataValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RunMetadataValidationError{
				field:  "Status",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RunMetadataMultiError(errors)
	}

	return nil
}

// RunMetadataMultiError is an error wrapping multiple validation errors
// returned by RunMetadata.ValidateAll() if the designated constraints aren't met.
type RunMetadataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RunMetadataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RunMetadataMultiError) AllErrors() []error { return m }

// RunMetadataValidationError is the validation error returned by
// RunMetadata.Validate if the designated constraints aren't met.
type RunMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RunMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RunMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RunMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RunMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RunMetadataValidationError) ErrorName() string { return "RunMetadataValidationError" }

// Error satisfies the builtin error interface
func (e RunMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRunMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RunMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RunMetadataValidationError{}
