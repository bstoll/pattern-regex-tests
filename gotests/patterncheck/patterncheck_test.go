// Copyright 2020 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package patterncheck

import (
	"testing"

	"github.com/openconfig/gnmi/errdiff"
)

func TestCheckRegexps(t *testing.T) {
	tests := []struct {
		desc             string
		inFiles          []string
		inPaths          []string
		wantErrSubstring string
	}{{
		desc:    "passing cases",
		inFiles: []string{"testdata/passing.yang"},
		inPaths: []string{"../../testdata"},
	}, {
		desc:             "simple leaf fail",
		inFiles:          []string{"testdata/simple-leaf-fail.yang"},
		inPaths:          []string{"../../testdata"},
		wantErrSubstring: `"ipv4" matches type string (leaf ipv-0), fail: "ipv6" doesn't match type string (leaf ipv-0)`,
	}, {
		desc:             "union leaf fail",
		inFiles:          []string{"testdata/union-leaf-fail.yang"},
		inPaths:          []string{"../../testdata"},
		wantErrSubstring: `fail: "ipv4" matches type ip-string-typedef (leaf ipv-0), fail: "ipv5" doesn't match type ip-string-typedef (leaf ipv-0)`,
	}, {
		desc:             "derived string type fail",
		inFiles:          []string{"testdata/derived-string-fail.yang"},
		inPaths:          []string{"../../testdata"},
		wantErrSubstring: `fail: "ipV4" doesn't match type ipv4-address-str (leaf ipv-0), fail: "ipV4-address" matches type ipv4-address-str (leaf ipv-0)`,
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := CheckRegexps(tt.inFiles, tt.inPaths)
			if diff := errdiff.Substring(got, tt.wantErrSubstring); diff != "" {
				t.Errorf("(-got, +want):\n%s", diff)
			}
		})
	}
}
