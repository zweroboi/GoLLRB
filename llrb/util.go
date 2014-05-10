// Copyright 2010 Petar Maymounkov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package llrb

type Int int

type IntItem interface {
	Key() Int
}

func (k Int) Key() Int {
	return k
}

func (x Int) Less(than Item) bool {
	return x < than.(IntItem).Key()
}

type String string

type StringItem interface {
	Key() String
}

func (k String) Key() String {
	return k
}

func (x String) Less(than Item) bool {
	return x < than.(StringItem).Key()
}

type Float64 float64

type Float64Item interface {
	Key() Float64
}

func (k Float64) Key() Float64 {
	return k
}

func (k Float64) Less(than Item) bool {
	return k < than.(Float64Item).Key()
}
