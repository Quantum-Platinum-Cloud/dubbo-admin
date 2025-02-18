// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

// AuthorizationPolicySourceApplyConfiguration represents an declarative configuration of the AuthorizationPolicySource type for use
// with apply.
type AuthorizationPolicySourceApplyConfiguration struct {
	Namespaces    []string                                      `json:"namespaces,omitempty"`
	NotNamespaces []string                                      `json:"notNamespaces,omitempty"`
	IpBlocks      []string                                      `json:"ipBlocks,omitempty"`
	NotIpBlocks   []string                                      `json:"notIpBlocks,omitempty"`
	Principals    []string                                      `json:"principals,omitempty"`
	NotPrincipals []string                                      `json:"notPrincipals,omitempty"`
	Extends       []AuthorizationPolicyExtendApplyConfiguration `json:"extends,omitempty"`
	NotExtends    []AuthorizationPolicyExtendApplyConfiguration `json:"notExtends,omitempty"`
}

// AuthorizationPolicySourceApplyConfiguration constructs an declarative configuration of the AuthorizationPolicySource type for use with
// apply.
func AuthorizationPolicySource() *AuthorizationPolicySourceApplyConfiguration {
	return &AuthorizationPolicySourceApplyConfiguration{}
}

// WithNamespaces adds the given value to the Namespaces field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Namespaces field.
func (b *AuthorizationPolicySourceApplyConfiguration) WithNamespaces(values ...string) *AuthorizationPolicySourceApplyConfiguration {
	for i := range values {
		b.Namespaces = append(b.Namespaces, values[i])
	}
	return b
}

// WithNotNamespaces adds the given value to the NotNamespaces field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NotNamespaces field.
func (b *AuthorizationPolicySourceApplyConfiguration) WithNotNamespaces(values ...string) *AuthorizationPolicySourceApplyConfiguration {
	for i := range values {
		b.NotNamespaces = append(b.NotNamespaces, values[i])
	}
	return b
}

// WithIpBlocks adds the given value to the IpBlocks field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the IpBlocks field.
func (b *AuthorizationPolicySourceApplyConfiguration) WithIpBlocks(values ...string) *AuthorizationPolicySourceApplyConfiguration {
	for i := range values {
		b.IpBlocks = append(b.IpBlocks, values[i])
	}
	return b
}

// WithNotIpBlocks adds the given value to the NotIpBlocks field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NotIpBlocks field.
func (b *AuthorizationPolicySourceApplyConfiguration) WithNotIpBlocks(values ...string) *AuthorizationPolicySourceApplyConfiguration {
	for i := range values {
		b.NotIpBlocks = append(b.NotIpBlocks, values[i])
	}
	return b
}

// WithPrincipals adds the given value to the Principals field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Principals field.
func (b *AuthorizationPolicySourceApplyConfiguration) WithPrincipals(values ...string) *AuthorizationPolicySourceApplyConfiguration {
	for i := range values {
		b.Principals = append(b.Principals, values[i])
	}
	return b
}

// WithNotPrincipals adds the given value to the NotPrincipals field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NotPrincipals field.
func (b *AuthorizationPolicySourceApplyConfiguration) WithNotPrincipals(values ...string) *AuthorizationPolicySourceApplyConfiguration {
	for i := range values {
		b.NotPrincipals = append(b.NotPrincipals, values[i])
	}
	return b
}

// WithExtends adds the given value to the Extends field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Extends field.
func (b *AuthorizationPolicySourceApplyConfiguration) WithExtends(values ...*AuthorizationPolicyExtendApplyConfiguration) *AuthorizationPolicySourceApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithExtends")
		}
		b.Extends = append(b.Extends, *values[i])
	}
	return b
}

// WithNotExtends adds the given value to the NotExtends field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NotExtends field.
func (b *AuthorizationPolicySourceApplyConfiguration) WithNotExtends(values ...*AuthorizationPolicyExtendApplyConfiguration) *AuthorizationPolicySourceApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithNotExtends")
		}
		b.NotExtends = append(b.NotExtends, *values[i])
	}
	return b
}
