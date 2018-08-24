// Code generated by protoc-gen-validate
// source: envoy/config/filter/http/ip_tagging/v2/ip_tagging.proto
// DO NOT EDIT!!!

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

	if _, ok := IPTagging_RequestType_name[int32(m.GetRequestType())]; !ok {
		return IPTaggingValidationError{
			Field:  "RequestType",
			Reason: "value must be one of the defined enum values",
		}
	}

	if len(m.GetIpTags()) < 1 {
		return IPTaggingValidationError{
			Field:  "IpTags",
			Reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetIpTags() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return IPTaggingValidationError{
					Field:  fmt.Sprintf("IpTags[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// IPTaggingValidationError is the validation error returned by
// IPTagging.Validate if the designated constraints aren't met.
type IPTaggingValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e IPTaggingValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIPTagging.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = IPTaggingValidationError{}

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

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return IPTagging_IPTagValidationError{
					Field:  fmt.Sprintf("IpList[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// IPTagging_IPTagValidationError is the validation error returned by
// IPTagging_IPTag.Validate if the designated constraints aren't met.
type IPTagging_IPTagValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e IPTagging_IPTagValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIPTagging_IPTag.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = IPTagging_IPTagValidationError{}