// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: publicapi/publicapi.proto

package publicapi

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

// Validate checks the field values on ChatResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ChatResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ChatResponseMultiError, or
// nil if none found.
func (m *ChatResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.ChatResponseOptions.(type) {

	case *ChatResponse_Command:

		if all {
			switch v := interface{}(m.GetCommand()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ChatResponseValidationError{
						field:  "Command",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ChatResponseValidationError{
						field:  "Command",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCommand()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ChatResponseValidationError{
					field:  "Command",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ChatResponse_KeepAlive:

		if all {
			switch v := interface{}(m.GetKeepAlive()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ChatResponseValidationError{
						field:  "KeepAlive",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ChatResponseValidationError{
						field:  "KeepAlive",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetKeepAlive()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ChatResponseValidationError{
					field:  "KeepAlive",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ChatResponseMultiError(errors)
	}

	return nil
}

// ChatResponseMultiError is an error wrapping multiple validation errors
// returned by ChatResponse.ValidateAll() if the designated constraints aren't met.
type ChatResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatResponseMultiError) AllErrors() []error { return m }

// ChatResponseValidationError is the validation error returned by
// ChatResponse.Validate if the designated constraints aren't met.
type ChatResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatResponseValidationError) ErrorName() string { return "ChatResponseValidationError" }

// Error satisfies the builtin error interface
func (e ChatResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatResponseValidationError{}

// Validate checks the field values on ChatResponse_ChatKeepAlive with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatResponse_ChatKeepAlive) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatResponse_ChatKeepAlive with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatResponse_ChatKeepAliveMultiError, or nil if none found.
func (m *ChatResponse_ChatKeepAlive) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatResponse_ChatKeepAlive) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ChatResponse_ChatKeepAliveMultiError(errors)
	}

	return nil
}

// ChatResponse_ChatKeepAliveMultiError is an error wrapping multiple
// validation errors returned by ChatResponse_ChatKeepAlive.ValidateAll() if
// the designated constraints aren't met.
type ChatResponse_ChatKeepAliveMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatResponse_ChatKeepAliveMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatResponse_ChatKeepAliveMultiError) AllErrors() []error { return m }

// ChatResponse_ChatKeepAliveValidationError is the validation error returned
// by ChatResponse_ChatKeepAlive.Validate if the designated constraints aren't met.
type ChatResponse_ChatKeepAliveValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatResponse_ChatKeepAliveValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatResponse_ChatKeepAliveValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatResponse_ChatKeepAliveValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatResponse_ChatKeepAliveValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatResponse_ChatKeepAliveValidationError) ErrorName() string {
	return "ChatResponse_ChatKeepAliveValidationError"
}

// Error satisfies the builtin error interface
func (e ChatResponse_ChatKeepAliveValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatResponse_ChatKeepAlive.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatResponse_ChatKeepAliveValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatResponse_ChatKeepAliveValidationError{}
