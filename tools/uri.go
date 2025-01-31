// Copyright 2022 The Gidari Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
package tools

import (
	"fmt"
	"net/http"
	"strings"
)

// ErrParsingURL is returned when there is an error parsing the url.
var ErrParsingURL = fmt.Errorf("error parsing url")

// SplitURLPath will return the endpoint parts from the request.
func SplitURLPath(req http.Request) []string {
	parts := strings.Split(strings.TrimPrefix(req.URL.EscapedPath(), "/"), "/")
	if len(parts) == 1 && parts[0] == "" {
		return []string{}
	}

	return parts
}

// ParseDBTableFromURL will return the table name from the request.
func ParseDBTableFromURL(req http.Request) (string, error) {
	endpointParts := SplitURLPath(req)
	if len(endpointParts) == 0 {
		return "", ErrParsingURL
	}

	return endpointParts[len(endpointParts)-1], nil
}
