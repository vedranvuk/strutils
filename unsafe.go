// Copyright 2025 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package strutils

import "unsafe"

// UnsafeStringBytes returns bytes of a string without allocating a slice.
//
// Result is a weak, non-garbage collected byte slice that is valid for the
// lifetime of s.
func UnsafeStringBytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// UnsafeString performs unholy acts to avoid allocations
// https://github.com/kubernetes/kubernetes/blob/e4b74dd12fa8cb63c174091d5536a10b8ec19d34/staging/src/k8s.io/apiserver/pkg/authentication/token/cache/cached_token_authenticator.go#L288-L297
func UnsafeString(b []byte) string {
	// unsafe.SliceData relies on cap whereas we want to rely on len
	if len(b) == 0 {
		return ""
	}
	// Copied from go 1.20.1 strings.Builder.String
	// https://github.com/golang/go/blob/202a1a57064127c3f19d96df57b9f9586145e21c/src/strings/builder.go#L48
	return unsafe.String(unsafe.SliceData(b), len(b))
}
