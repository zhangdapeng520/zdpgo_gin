// Copyright 2022 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package render

import (
	"net/http"

	"github.com/zhangdapeng520/zdpgo_toml"
)

// TOML contains the given interface object.
type TOML struct {
	Data any
}

var TOMLContentType = []string{"application/toml; charset=utf-8"}

// Render (TOML) marshals the given interface object and writes data with custom ContentType.
func (r TOML) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	bytes, err := zdpgo_toml.Marshal(r.Data)
	if err != nil {
		return err
	}

	_, err = w.Write(bytes)
	return err
}

// WriteContentType (TOML) writes TOML ContentType for response.
func (r TOML) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, TOMLContentType)
}
