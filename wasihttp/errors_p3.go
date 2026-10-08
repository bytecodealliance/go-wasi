//go:build componentizego_async

package wasihttp

import (
	"fmt"
	"strings"

	types "go.bytecodealliance.org/wasi/imports/wasi_http_0_3_0_types"
)

func headerErrorString(err types.HeaderError) string {
	switch err.Tag() {
	case types.HeaderError_InvalidSyntax:
		return "invalid syntax"
	case types.HeaderError_Forbidden:
		return "forbidden header"
	case types.HeaderError_Immutable:
		return "immutable header fields"
	case types.HeaderError_SizeExceeded:
		return "size exceeded"
	case types.HeaderError_Other:
		if v := err.Other(); v.IsSome() {
			return v.Some()
		}
		return "unknown header error"
	default:
		return "unknown header error"
	}
}

func errorCodeString(code types.ErrorCode) string {
	switch code.Tag() {
	case types.ErrorCode_DnsTimeout:
		return "DNS timeout"
	case types.ErrorCode_DnsError:
		p := code.DnsError()
		var parts []string
		if p.Rcode.IsSome() {
			parts = append(parts, fmt.Sprintf("rcode=%s", p.Rcode.Some()))
		}
		if p.InfoCode.IsSome() {
			parts = append(parts, fmt.Sprintf("info-code=%d", p.InfoCode.Some()))
		}
		if len(parts) == 0 {
			return "DNS error"
		}
		return "DNS error (" + strings.Join(parts, ", ") + ")"
	case types.ErrorCode_DestinationNotFound:
		return "destination not found"
	case types.ErrorCode_DestinationUnavailable:
		return "destination unavailable"
	case types.ErrorCode_DestinationIpProhibited:
		return "destination IP prohibited"
	case types.ErrorCode_DestinationIpUnroutable:
		return "destination IP unroutable"
	case types.ErrorCode_ConnectionRefused:
		return "connection refused"
	case types.ErrorCode_ConnectionTerminated:
		return "connection terminated"
	case types.ErrorCode_ConnectionTimeout:
		return "connection timeout"
	case types.ErrorCode_ConnectionReadTimeout:
		return "connection read timeout"
	case types.ErrorCode_ConnectionWriteTimeout:
		return "connection write timeout"
	case types.ErrorCode_ConnectionLimitReached:
		return "connection limit reached"
	case types.ErrorCode_TlsProtocolError:
		return "TLS protocol error"
	case types.ErrorCode_TlsCertificateError:
		return "TLS certificate error"
	case types.ErrorCode_TlsAlertReceived:
		p := code.TlsAlertReceived()
		var parts []string
		if p.AlertId.IsSome() {
			parts = append(parts, fmt.Sprintf("alert-id=%d", p.AlertId.Some()))
		}
		if p.AlertMessage.IsSome() {
			parts = append(parts, fmt.Sprintf("alert-message=%s", p.AlertMessage.Some()))
		}
		if len(parts) == 0 {
			return "TLS alert received"
		}
		return "TLS alert received (" + strings.Join(parts, ", ") + ")"
	case types.ErrorCode_HttpRequestDenied:
		return "HTTP request denied"
	case types.ErrorCode_HttpRequestLengthRequired:
		return "HTTP request length required"
	case types.ErrorCode_HttpRequestBodySize:
		if v := code.HttpRequestBodySize(); v.IsSome() {
			return fmt.Sprintf("HTTP request body size: %d", v.Some())
		}
		return "HTTP request body size"
	case types.ErrorCode_HttpRequestMethodInvalid:
		return "HTTP request method invalid"
	case types.ErrorCode_HttpRequestUriInvalid:
		return "HTTP request URI invalid"
	case types.ErrorCode_HttpRequestUriTooLong:
		return "HTTP request URI too long"
	case types.ErrorCode_HttpRequestHeaderSectionSize:
		if v := code.HttpRequestHeaderSectionSize(); v.IsSome() {
			return fmt.Sprintf("HTTP request header section size: %d", v.Some())
		}
		return "HTTP request header section size"
	case types.ErrorCode_HttpRequestHeaderSize:
		if v := code.HttpRequestHeaderSize(); v.IsSome() {
			return "HTTP request header size " + fieldSizeString(v.Some())
		}
		return "HTTP request header size"
	case types.ErrorCode_HttpRequestTrailerSectionSize:
		if v := code.HttpRequestTrailerSectionSize(); v.IsSome() {
			return fmt.Sprintf("HTTP request trailer section size: %d", v.Some())
		}
		return "HTTP request trailer section size"
	case types.ErrorCode_HttpRequestTrailerSize:
		return "HTTP request trailer size " + fieldSizeString(code.HttpRequestTrailerSize())
	case types.ErrorCode_HttpResponseIncomplete:
		return "HTTP response incomplete"
	case types.ErrorCode_HttpResponseHeaderSectionSize:
		if v := code.HttpResponseHeaderSectionSize(); v.IsSome() {
			return fmt.Sprintf("HTTP response header section size: %d", v.Some())
		}
		return "HTTP response header section size"
	case types.ErrorCode_HttpResponseHeaderSize:
		return "HTTP response header size " + fieldSizeString(code.HttpResponseHeaderSize())
	case types.ErrorCode_HttpResponseBodySize:
		if v := code.HttpResponseBodySize(); v.IsSome() {
			return fmt.Sprintf("HTTP response body size: %d", v.Some())
		}
		return "HTTP response body size"
	case types.ErrorCode_HttpResponseTrailerSectionSize:
		if v := code.HttpResponseTrailerSectionSize(); v.IsSome() {
			return fmt.Sprintf("HTTP response trailer section size: %d", v.Some())
		}
		return "HTTP response trailer section size"
	case types.ErrorCode_HttpResponseTrailerSize:
		return "HTTP response trailer size " + fieldSizeString(code.HttpResponseTrailerSize())
	case types.ErrorCode_HttpResponseTransferCoding:
		if v := code.HttpResponseTransferCoding(); v.IsSome() {
			return fmt.Sprintf("HTTP response transfer coding: %s", v.Some())
		}
		return "HTTP response transfer coding"
	case types.ErrorCode_HttpResponseContentCoding:
		if v := code.HttpResponseContentCoding(); v.IsSome() {
			return fmt.Sprintf("HTTP response content coding: %s", v.Some())
		}
		return "HTTP response content coding"
	case types.ErrorCode_HttpResponseTimeout:
		return "HTTP response timeout"
	case types.ErrorCode_HttpUpgradeFailed:
		return "HTTP upgrade failed"
	case types.ErrorCode_HttpProtocolError:
		return "HTTP protocol error"
	case types.ErrorCode_LoopDetected:
		return "loop detected"
	case types.ErrorCode_ConfigurationError:
		return "configuration error"
	case types.ErrorCode_InternalError:
		if v := code.InternalError(); v.IsSome() {
			return "internal error: " + v.Some()
		}
		return "internal error"
	default:
		return fmt.Sprintf("unknown error code: %d", code.Tag())
	}
}

func fieldSizeString(p types.FieldSizePayload) string {
	var parts []string
	if p.FieldName.IsSome() {
		parts = append(parts, fmt.Sprintf("field-name=%s", p.FieldName.Some()))
	}
	if p.FieldSize.IsSome() {
		parts = append(parts, fmt.Sprintf("field-size=%d", p.FieldSize.Some()))
	}
	if len(parts) == 0 {
		return ""
	}
	return "(" + strings.Join(parts, ", ") + ")"
}
