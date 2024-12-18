// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfconfig

// Variable represents a single variable from a Terraform module.
type Variable struct {
	Name        string       `json:"name"`
	Type        string       `json:"type,omitempty"`
	Description string       `json:"description,omitempty"`
	Validation  []Validation `json:"validation,omitempty"`

	// Default is an approximate representation of the default value in
	// the native Go type system. The conversion from the value given in
	// configuration may be slightly lossy. Only values that can be
	// serialized by json.Marshal will be included here.
	Default   interface{} `json:"default"`
	Required  bool        `json:"required"`
	Sensitive bool        `json:"sensitive,omitempty"`

	Pos SourcePos `json:"pos"`
}

type Validation struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
	Condition    string `json:"condition,omitempty"`
}
