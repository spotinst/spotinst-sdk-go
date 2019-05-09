// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonutil

import (
	"encoding/json"
	"reflect"
	"testing"
)

type schema struct {
	// Basic types
	B    bool    `json:"b,omitempty"`
	F    float64 `json:"f,omitempty"`
	I    int64   `json:"i,omitempty"`
	Istr int64   `json:"istr,omitempty,string"`
	Str  string  `json:"str,omitempty"`

	// Pointers to basic types
	PB    *bool    `json:"pb,omitempty"`
	PF    *float64 `json:"pf,omitempty"`
	PI    *int64   `json:"pi,omitempty"`
	PIStr *int64   `json:"pistr,omitempty,string"`
	PStr  *string  `json:"pstr,omitempty"`

	// Other types
	S             []int                    `json:"s,omitempty"`
	M             map[string]string        `json:"m,omitempty"`
	Any           interface{}              `json:"any,omitempty"`
	Child         *child                   `json:"child,omitempty"`
	MapToAnyArray map[string][]interface{} `json:"maptoanyarray,omitempty"`
	Embedded

	forceSendFields []string
	nullFields      []string
}

type child struct {
	B bool `json:"childbool,omitempty"`
}

type Embedded struct {
	E bool `json:"embeddedbool,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type testCase struct {
	s    schema
	want string
}

func TestBasics(t *testing.T) {
	for _, tc := range []testCase{
		{
			s:    schema{},
			want: `{}`,
		},
		{
			s: schema{
				forceSendFields: []string{"B", "F", "I", "Istr", "Str", "PB", "PF", "PI", "PIStr", "PStr"},
			},
			want: `{"b":false,"f":0.0,"i":0,"istr":"0","str":""}`,
		},
		{
			s: schema{
				nullFields: []string{"B", "F", "I", "Istr", "Str", "PB", "PF", "PI", "PIStr", "PStr"},
			},
			want: `{"b":null,"f":null,"i":null,"istr":null,"str":null,"pb":null,"pf":null,"pi":null,"pistr":null,"pstr":null}`,
		},
		{
			s: schema{
				B:     true,
				F:     1.2,
				I:     1,
				Istr:  2,
				Str:   "a",
				PB:    boolPtr(true),
				PF:    float64Ptr(1.2),
				PI:    int64Ptr(int64(1)),
				PIStr: int64Ptr(int64(2)),
				PStr:  stringPtr("a"),
			},
			want: `{"b":true,"f":1.2,"i":1,"istr":"2","str":"a","pb":true,"pf":1.2,"pi":1,"pistr":"2","pstr":"a"}`,
		},
		{
			s: schema{
				B:     false,
				F:     0.0,
				I:     0,
				Istr:  0,
				Str:   "",
				PB:    boolPtr(false),
				PF:    float64Ptr(0.0),
				PI:    int64Ptr(int64(0)),
				PIStr: int64Ptr(int64(0)),
				PStr:  stringPtr(""),
			},
			want: `{"pb":false,"pf":0.0,"pi":0,"pistr":"0","pstr":""}`,
		},
		{
			s: schema{
				B:               false,
				F:               0.0,
				I:               0,
				Istr:            0,
				Str:             "",
				PB:              boolPtr(false),
				PF:              float64Ptr(0.0),
				PI:              int64Ptr(int64(0)),
				PIStr:           int64Ptr(int64(0)),
				PStr:            stringPtr(""),
				forceSendFields: []string{"B", "F", "I", "Istr", "Str", "PB", "PF", "PI", "PIStr", "PStr"},
			},
			want: `{"b":false,"f":0.0,"i":0,"istr":"0","str":"","pb":false,"pf":0.0,"pi":0,"pistr":"0","pstr":""}`,
		},
		{
			s: schema{
				B:          false,
				F:          0.0,
				I:          0,
				Istr:       0,
				Str:        "",
				PB:         boolPtr(false),
				PF:         float64Ptr(0.0),
				PI:         int64Ptr(int64(0)),
				PIStr:      int64Ptr(int64(0)),
				PStr:       stringPtr(""),
				nullFields: []string{"B", "F", "I", "Istr", "Str"},
			},
			want: `{"b":null,"f":null,"i":null,"istr":null,"str":null,"pb":false,"pf":0.0,"pi":0,"pistr":"0","pstr":""}`,
		},
	} {
		checkMarshalJSON(t, tc)
	}
}

func TestSliceFields(t *testing.T) {
	for _, tc := range []testCase{
		{
			s:    schema{},
			want: `{}`,
		},
		{
			s:    schema{S: []int{1}},
			want: `{"s":[1]}`,
		},
		{
			s: schema{
				forceSendFields: []string{"S"},
			},
			want: `{"s":[]}`,
		},
		{
			s: schema{
				S:               []int{},
				forceSendFields: []string{"S"},
			},
			want: `{"s":[]}`,
		},
		{
			s: schema{
				S:               []int{1},
				forceSendFields: []string{"S"},
			},
			want: `{"s":[1]}`,
		},
		{
			s: schema{
				nullFields: []string{"S"},
			},
			want: `{"s":null}`,
		},
	} {
		checkMarshalJSON(t, tc)
	}
}

func TestMapField(t *testing.T) {
	for _, tc := range []testCase{
		{
			s:    schema{},
			want: `{}`,
		},
		{
			s:    schema{M: make(map[string]string)},
			want: `{}`,
		},
		{
			s:    schema{M: map[string]string{"a": "b"}},
			want: `{"m":{"a":"b"}}`,
		},
		{
			s: schema{
				forceSendFields: []string{"M"},
			},
			want: `{"m":{}}`,
		},
		{
			s: schema{
				nullFields: []string{"M"},
			},
			want: `{"m":null}`,
		},
		{
			s: schema{
				M:               make(map[string]string),
				forceSendFields: []string{"M"},
			},
			want: `{"m":{}}`,
		},
		{
			s: schema{
				M:          make(map[string]string),
				nullFields: []string{"M"},
			},
			want: `{"m":null}`,
		},
		{
			s: schema{
				M:               map[string]string{"a": "b"},
				forceSendFields: []string{"M"},
			},
			want: `{"m":{"a":"b"}}`,
		},
	} {
		checkMarshalJSON(t, tc)
	}
}

func TestMapToAnyArray(t *testing.T) {
	for _, tc := range []testCase{
		{
			s:    schema{},
			want: `{}`,
		},
		{
			s:    schema{MapToAnyArray: make(map[string][]interface{})},
			want: `{}`,
		},
		{
			s: schema{
				MapToAnyArray: map[string][]interface{}{
					"a": {2, "b"},
				},
			},
			want: `{"maptoanyarray":{"a":[2, "b"]}}`,
		},
		{
			s: schema{
				MapToAnyArray: map[string][]interface{}{
					"a": nil,
				},
			},
			want: `{"maptoanyarray":{"a": null}}`,
		},
		{
			s: schema{
				MapToAnyArray: map[string][]interface{}{
					"a": {nil},
				},
			},
			want: `{"maptoanyarray":{"a":[null]}}`,
		},
		{
			s: schema{
				forceSendFields: []string{"MapToAnyArray"},
			},
			want: `{"maptoanyarray":{}}`,
		},
		{
			s: schema{
				nullFields: []string{"MapToAnyArray"},
			},
			want: `{"maptoanyarray":null}`,
		},
		{
			s: schema{
				MapToAnyArray:   make(map[string][]interface{}),
				forceSendFields: []string{"MapToAnyArray"},
			},
			want: `{"maptoanyarray":{}}`,
		},
		{
			s: schema{
				MapToAnyArray: map[string][]interface{}{
					"a": {2, "b"},
				},
				forceSendFields: []string{"MapToAnyArray"},
			},
			want: `{"maptoanyarray":{"a":[2, "b"]}}`,
		},
	} {
		checkMarshalJSON(t, tc)
	}
}

type anyType struct {
	Field int
}

func (a anyType) MarshalJSON() ([]byte, error) {
	return []byte(`"anyType value"`), nil
}

func TestAnyField(t *testing.T) {
	// forceSendFields has no effect on nil interfaces and interfaces that contain nil pointers.
	var nilAny *anyType
	for _, tc := range []testCase{
		{
			s:    schema{},
			want: `{}`,
		},
		{
			s:    schema{Any: nilAny},
			want: `{"any": null}`,
		},
		{
			s:    schema{Any: &anyType{}},
			want: `{"any":"anyType value"}`,
		},
		{
			s:    schema{Any: anyType{}},
			want: `{"any":"anyType value"}`,
		},
		{
			s: schema{
				forceSendFields: []string{"Any"},
			},
			want: `{}`,
		},
		{
			s: schema{
				nullFields: []string{"Any"},
			},
			want: `{"any":null}`,
		},
		{
			s: schema{
				Any:             nilAny,
				forceSendFields: []string{"Any"},
			},
			want: `{"any": null}`,
		},
		{
			s: schema{
				Any:             &anyType{},
				forceSendFields: []string{"Any"},
			},
			want: `{"any":"anyType value"}`,
		},
		{
			s: schema{
				Any:             anyType{},
				forceSendFields: []string{"Any"},
			},
			want: `{"any":"anyType value"}`,
		},
	} {
		checkMarshalJSON(t, tc)
	}
}

func TestSubschema(t *testing.T) {
	// Subschemas are always stored as pointers, so forceSendFields has no effect on them.
	for _, tc := range []testCase{
		{
			s:    schema{},
			want: `{}`,
		},
		{
			s: schema{
				forceSendFields: []string{"Child"},
			},
			want: `{}`,
		},
		{
			s: schema{
				nullFields: []string{"Child"},
			},
			want: `{"child":null}`,
		},
		{
			s:    schema{Child: &child{}},
			want: `{"child":{}}`,
		},
		{
			s: schema{
				Child:           &child{},
				forceSendFields: []string{"Child"},
			},
			want: `{"child":{}}`,
		},
		{
			s:    schema{Child: &child{B: true}},
			want: `{"child":{"childbool":true}}`,
		},
		{
			s: schema{
				Child:           &child{B: true},
				forceSendFields: []string{"Child"},
			},
			want: `{"child":{"childbool":true}}`,
		},
	} {
		checkMarshalJSON(t, tc)
	}
}

func TestEmbeddedSchema(t *testing.T) {
	// Subschemas are always stored as pointers, so forceSendFields has no effect on them.
	for _, tc := range []testCase{
		{
			s:    schema{},
			want: `{}`,
		},
		{
			s: schema{
				forceSendFields: []string{"Embedded"},
			},
			want: `{}`,
		},
		{
			s: schema{
				nullFields: []string{"Embedded"},
			},
			want: `{}`,
		},
		{
			s:    schema{Embedded: Embedded{}},
			want: `{}`,
		},
		{
			s: schema{
				Embedded:        Embedded{},
				forceSendFields: []string{"Embedded"},
			},
			want: `{}`,
		},
		{
			s:    schema{Embedded: Embedded{E: true}},
			want: `{"embeddedbool":true}`,
		},
		{
			s: schema{
				Embedded:        Embedded{E: false, forceSendFields: []string{"E"}},
				forceSendFields: []string{"Embedded"},
			},
			want: `{"embeddedbool":false}`,
		},
	} {
		checkMarshalJSON(t, tc)
	}
}

// checkMarshalJSON verifies that calling schemaToMap on tc.s yields a result which is equivalent to tc.want.
func checkMarshalJSON(t *testing.T, tc testCase) {
	doCheckMarshalJSON(t, tc.s, tc.s.forceSendFields, tc.s.nullFields, tc.want)
	if len(tc.s.forceSendFields) == 0 && len(tc.s.nullFields) == 0 {
		// verify that the code path used when forceSendFields and nullFields
		// are non-empty produces the same output as the fast path that is used
		// when they are empty.
		doCheckMarshalJSON(t, tc.s, []string{"dummy"}, []string{"dummy"}, tc.want)
	}
}

func doCheckMarshalJSON(t *testing.T, s schema, forceSendFields, nullFields []string, wantJSON string) {
	encoded, err := MarshalJSON(s, forceSendFields, nullFields)
	if err != nil {
		t.Fatalf("encoding json:\n got err: %v", err)
	}

	// The expected and obtained JSON can differ in field ordering, so unmarshal before comparing.
	var got interface{}
	var want interface{}
	err = json.Unmarshal(encoded, &got)
	if err != nil {
		t.Fatalf("decoding json:\n got err: %v", err)
	}
	err = json.Unmarshal([]byte(wantJSON), &want)
	if err != nil {
		t.Fatalf("decoding json:\n got err: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("schemaToMap:\ngot :%v\nwant: %v", got, want)
	}
}

func TestParseJSONTag(t *testing.T) {
	for _, tc := range []struct {
		tag  string
		want *jsonTag
	}{
		{
			tag:  "-",
			want: &jsonTag{ignore: true},
		}, {
			tag:  "name,omitempty",
			want: &jsonTag{apiName: "name"},
		}, {
			tag:  "name,omitempty,string",
			want: &jsonTag{apiName: "name", stringFormat: true},
		},
	} {
		got, err := parseJSONTag(tc.tag)
		if err != nil {
			t.Fatalf("parsing json:\n got err: %v\ntag: %q", err, tc.tag)
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("parseJSONTage:\ngot :%T\nwant:%T", got, tc.want)
		}
	}
}

func TestParseMalformedJSONTag(t *testing.T) {
	for _, tag := range []string{
		"",
		"name",
		"name,",
		"name,blah",
		"name,blah,string",
		",omitempty",
		",omitempty,string",
		"name,omitempty,string,blah",
	} {
		_, err := parseJSONTag(tag)
		if err == nil {
			t.Fatalf("parsing json: expected err, got nil for tag: %v", tag)
		}
	}
}

func int64Ptr(v int64) *int64       { return &v }
func float64Ptr(v float64) *float64 { return &v }
func boolPtr(v bool) *bool          { return &v }
func stringPtr(v string) *string    { return &v }
