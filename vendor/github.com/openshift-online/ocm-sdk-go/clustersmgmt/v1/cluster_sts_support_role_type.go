/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// ClusterStsSupportRole represents the values of the 'cluster_sts_support_role' type.
//
// Isolated STS support role created per organization.
type ClusterStsSupportRole struct {
	bitmap_ uint32
	roleArn string
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *ClusterStsSupportRole) Empty() bool {
	return o == nil || o.bitmap_ == 0
}

// RoleArn returns the value of the 'role_arn' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// ARN of the support role created in the SRE jump account.
func (o *ClusterStsSupportRole) RoleArn() string {
	if o != nil && o.bitmap_&1 != 0 {
		return o.roleArn
	}
	return ""
}

// GetRoleArn returns the value of the 'role_arn' attribute and
// a flag indicating if the attribute has a value.
//
// ARN of the support role created in the SRE jump account.
func (o *ClusterStsSupportRole) GetRoleArn() (value string, ok bool) {
	ok = o != nil && o.bitmap_&1 != 0
	if ok {
		value = o.roleArn
	}
	return
}

// ClusterStsSupportRoleListKind is the name of the type used to represent list of objects of
// type 'cluster_sts_support_role'.
const ClusterStsSupportRoleListKind = "ClusterStsSupportRoleList"

// ClusterStsSupportRoleListLinkKind is the name of the type used to represent links to list
// of objects of type 'cluster_sts_support_role'.
const ClusterStsSupportRoleListLinkKind = "ClusterStsSupportRoleListLink"

// ClusterStsSupportRoleNilKind is the name of the type used to nil lists of objects of
// type 'cluster_sts_support_role'.
const ClusterStsSupportRoleListNilKind = "ClusterStsSupportRoleListNil"

// ClusterStsSupportRoleList is a list of values of the 'cluster_sts_support_role' type.
type ClusterStsSupportRoleList struct {
	href  string
	link  bool
	items []*ClusterStsSupportRole
}

// Len returns the length of the list.
func (l *ClusterStsSupportRoleList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *ClusterStsSupportRoleList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *ClusterStsSupportRoleList) Get(i int) *ClusterStsSupportRole {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *ClusterStsSupportRoleList) Slice() []*ClusterStsSupportRole {
	var slice []*ClusterStsSupportRole
	if l == nil {
		slice = make([]*ClusterStsSupportRole, 0)
	} else {
		slice = make([]*ClusterStsSupportRole, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ClusterStsSupportRoleList) Each(f func(item *ClusterStsSupportRole) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *ClusterStsSupportRoleList) Range(f func(index int, item *ClusterStsSupportRole) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
