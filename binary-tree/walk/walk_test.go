package walk

import (
	"regexp"
	"testing"

	"golang.org/x/tour/tree"
)

func TestString(t *testing.T) {
	var tests = []struct {
		t    *tree.Tree
		expr string
	}{
		{tree.New(1), "^[(]+1[)]* [(]*2[)]* [(]*3[)]* [(]*4[)]* [(]*5[)]* [(]*6[)]* [(]*7[)]* [(]*8[)]* [(]*9[)]* [(]*10[)]+$"},
		{tree.New(2), "^[(]+2[)]* [(]*4[)]* [(]*6[)]* [(]*8[)]* [(]*10[)]* [(]*12[)]* [(]*14[)]* [(]*16[)]* [(]*18[)]* [(]*20[)]+$"},
	}
	for _, c := range tests {
		_, err := regexp.MatchString(c.expr, c.t.String())
		if err != nil {
			t.Errorf("t.String() does not match %q", c.expr)
		}
	}
}

func TestWalk(t *testing.T) {
	var tests = []struct {
		t1, t2 *tree.Tree
		want   bool
	}{
		{tree.New(1), tree.New(1), true},
		{tree.New(1), tree.New(2), false},
	}
	for _, c := range tests {
		got := Same(c.t1, c.t2)
		if got != c.want {
			t.Errorf("Same(%q, %q) == %q, want %q", c.t1.String(), c.t2.String(), got, c.want)
		}
	}
}
