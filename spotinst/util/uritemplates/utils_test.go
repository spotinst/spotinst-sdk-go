package uritemplates

import (
	"testing"
)

func TestExpand(t *testing.T) {
	var tt = map[string]struct {
		in       string
		values   Values
		expected string
	}{
		"#0: no expansions": {
			"http://www.golang.org/",
			Values{},
			"http://www.golang.org/",
		},
		"#1: one expansion, no escaping": {
			"http://www.golang.org/{bucket}/delete",
			Values{
				"bucket": "red",
			},
			"http://www.golang.org/red/delete",
		},
		"#2: one expansion, with hex escapes": {
			"http://www.golang.org/{bucket}/delete",
			Values{
				"bucket": "red/blue",
			},
			"http://www.golang.org/red%2Fblue/delete",
		},
		"#3: one expansion, with space": {
			"http://www.golang.org/{bucket}/delete",
			Values{
				"bucket": "red or blue",
			},
			"http://www.golang.org/red%20or%20blue/delete",
		},
		"#4: expansion not found": {
			"http://www.golang.org/{object}/delete",
			Values{
				"bucket": "red or blue",
			},
			"http://www.golang.org//delete",
		},
		"#5: multiple expansions": {
			"http://www.golang.org/{one}/{two}/{three}/get",
			Values{
				"one":   "ONE",
				"two":   "TWO",
				"three": "THREE",
			},
			"http://www.golang.org/ONE/TWO/THREE/get",
		},
		"#6: utf-8 characters": {
			"http://www.golang.org/{bucket}/get",
			Values{
				"bucket": "£100",
			},
			"http://www.golang.org/%C2%A3100/get",
		},
		"#7: punctuations": {
			"http://www.golang.org/{bucket}/get",
			Values{
				"bucket": `/\@:,.*~`,
			},
			"http://www.golang.org/%2F%5C%40%3A%2C.%2A~/get",
		},
		"#8: mis-matched brackets": {
			"http://www.golang.org/{bucket/get",
			Values{
				"bucket": "red",
			},
			"",
		},
		//
		"#9: prefix for suppressing escape": {
			"http://www.golang.org/{+topic}",
			Values{
				"topic": "/topics/myproject/mytopic",
			},
			// The double slashes here look weird, but it's intentional
			"http://www.golang.org//topics/myproject/mytopic",
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			actual, _ := Expand(test.in, test.values)
			if actual != test.expected {
				t.Errorf("Expected %q, got %q", test.expected, actual)
			}
		})
	}
}
