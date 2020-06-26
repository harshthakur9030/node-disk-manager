/*
Copyright 2019 The OpenEBS Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package services

import (
	"reflect"
	"testing"
)

// TestAddStringstoSlice get the method which adds only unique devices
func TestAddStringstoSlice(t *testing.T) {
	var testCases = []struct {
		testName string
		name     string
		names    []string
		exp      []string
	}{
		{
			name:     "sdb1",
			names:    nil,
			testName: "Handling partitions of first disk",
			exp:      []string{"sdb1"},
		},
		{
			name:     "sdb1",
			names:    []string{"sdc1", "sdc2"},
			testName: "Handling partitions of two disks",
			exp:      []string{"sdc1", "sdc2", "sdb1"},
		},
		{
			name:     "sdd1",
			names:    []string{"sdb1", "sdc1", "sdc2"},
			testName: "Handling partitions of three disks",
			exp:      []string{"sdb1", "sdc1", "sdc2", "sdd1"},
		},
	}

	for _, e := range testCases {
		res := AddStringtoSlice(e.names, e.name)
		if !reflect.DeepEqual(res, e.exp) {
			t.Errorf(" Test failed : %v , expected : %v  , got : %v", e.testName, e.exp, res)
		}

	}

}

// TestFilterPartitions tests the FilterPartitions
func TestFilterPartitions(t *testing.T) {
	var testCases = []struct {
		testName string
		name     string
		names    []string
		exp      []string
	}{
		{
			name:     "sdb",
			names:    []string{"sdb1", "sdb2", "sdc1"},
			testName: "Filtering partitions of two disks",
			exp:      []string{"sdb1", "sdb2"},
		},
		{
			name:     "sdd",
			names:    []string{"sdb1", "sdc1", "sdc2", "sdd1"},
			testName: "Filtering partitions of three disks",
			exp:      []string{"sdd1"},
		},
	}

	for _, e := range testCases {
		res := FilterParitions(e.name, e.names)
		if !reflect.DeepEqual(res, e.exp) {
			t.Errorf(" Test failed : %v , expected : %v  , got : %v", e.testName, e.exp, res)
		}

	}

}
