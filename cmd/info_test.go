// Copyright © 2017 dnxware Team
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"testing"
)

func TestRepoLocation(t *testing.T) {
	var repoTests = []struct {
		s        string // input
		expected string // expected result
	}{
		{"git@github.com:dnxware/promu.git", "github.com/dnxware/promu"},
		{"https://github.com/dnxware/promu.git", "github.com/dnxware/promu"},
		{"ssh://git@gitlab.fr:22443/dnxware/promu.git", "gitlab.fr/dnxware/promu"},
		{"https://sdurrheimer@gitlab.fr/dnxware/promu.git", "gitlab.fr/dnxware/promu"},
	}

	for _, tc := range repoTests {
		actual, err := repoLocation(tc.s)
		if err != nil {
			t.Errorf("repoLocation(%s): %+v", tc.s, err)
		}
		if actual != tc.expected {
			t.Errorf("repoLocation(%s): expected %s, got %s", tc.s, tc.expected, actual)
		}
	}
}
