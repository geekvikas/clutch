// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: sourcecontrol/v1/sourcecontrol.proto

package sourcecontrolv1

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _sourcecontrol_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CreateRepositoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateRepositoryRequest) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetOwner()) < 1 {
		return CreateRepositoryRequestValidationError{
			field:  "Owner",
			reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetName()) < 1 {
		return CreateRepositoryRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	// no validation rules for Description

	switch m.Options.(type) {

	case *CreateRepositoryRequest_CustomOptions:

		if v, ok := interface{}(m.GetCustomOptions()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateRepositoryRequestValidationError{
					field:  "CustomOptions",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CreateRepositoryRequest_GithubOptions:

		if v, ok := interface{}(m.GetGithubOptions()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateRepositoryRequestValidationError{
					field:  "GithubOptions",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return CreateRepositoryRequestValidationError{
			field:  "Options",
			reason: "value is required",
		}

	}

	return nil
}

// CreateRepositoryRequestValidationError is the validation error returned by
// CreateRepositoryRequest.Validate if the designated constraints aren't met.
type CreateRepositoryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRepositoryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRepositoryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRepositoryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRepositoryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRepositoryRequestValidationError) ErrorName() string {
	return "CreateRepositoryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRepositoryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRepositoryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRepositoryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRepositoryRequestValidationError{}

// Validate checks the field values on CreateRepositoryResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateRepositoryResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Url

	return nil
}

// CreateRepositoryResponseValidationError is the validation error returned by
// CreateRepositoryResponse.Validate if the designated constraints aren't met.
type CreateRepositoryResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRepositoryResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRepositoryResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRepositoryResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRepositoryResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRepositoryResponseValidationError) ErrorName() string {
	return "CreateRepositoryResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRepositoryResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRepositoryResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRepositoryResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRepositoryResponseValidationError{}