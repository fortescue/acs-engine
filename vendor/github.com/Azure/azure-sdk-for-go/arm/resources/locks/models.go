package locks

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// LockLevel enumerates the values for lock level.
type LockLevel string

const (
	// CanNotDelete specifies the can not delete state for lock level.
	CanNotDelete LockLevel = "CanNotDelete"
	// NotSpecified specifies the not specified state for lock level.
	NotSpecified LockLevel = "NotSpecified"
	// ReadOnly specifies the read only state for lock level.
	ReadOnly LockLevel = "ReadOnly"
)

// ManagementLockListResult is the list of locks.
type ManagementLockListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ManagementLockObject `json:"value,omitempty"`
	NextLink          *string                 `json:"nextLink,omitempty"`
}

// ManagementLockListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ManagementLockListResult) ManagementLockListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ManagementLockObject is the lock information.
type ManagementLockObject struct {
	autorest.Response         `json:"-"`
	*ManagementLockProperties `json:"properties,omitempty"`
	ID                        *string `json:"id,omitempty"`
	Type                      *string `json:"type,omitempty"`
	Name                      *string `json:"name,omitempty"`
}

// ManagementLockOwner is lock owner properties.
type ManagementLockOwner struct {
	ApplicationID *string `json:"applicationId,omitempty"`
}

// ManagementLockProperties is the lock properties.
type ManagementLockProperties struct {
	Level  LockLevel              `json:"level,omitempty"`
	Notes  *string                `json:"notes,omitempty"`
	Owners *[]ManagementLockOwner `json:"owners,omitempty"`
}
