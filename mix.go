// MIT License
//
// Copyright (c) 2021 Lack
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package gkvm

type Unit string

const (
	UnitBytes = "bytes"
	UnitM     = "M"
	UnitG     = "G"
	UnitT     = "T"
)

type Capacity struct {
	Unit Unit   `xml:"unit,attr" json:"unit,omitempty"`
	Data uint64 `xml:",chardata" json:"data,omitempty"`
}

type Allocation struct {
	Unit Unit   `xml:"unit,attr" json:"unit,omitempty"`
	Data uint64 `xml:",chardata" json:"data,omitempty"`
}

type Available struct {
	Unit Unit   `xml:"unit,attr" json:"unit,omitempty"`
	Data uint64 `xml:",chardata" json:"data,omitempty"`
}

type Permissions struct {
	Model int32  `xml:"model" json:"model"`
	Owner int32  `xml:"owner" json:"owner"`
	Group int32  `xml:"group" json:"group"`
	Label string `xml:"label" json:"label"`
}

type ButtonState string

const (
	ButtonStateYes = "yes"
	ButtonStateNo  = "no"
)
