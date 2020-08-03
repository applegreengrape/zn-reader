package znr_test

import (
	"reflect"
	"strings"
	"testing"

	znr "github.com/billglover/zn-reader"
)

func TestKnownPhrases(t *testing.T) {
	vl := znr.VocabList{
		znr.Vocab{Writing: "你"},
		znr.Vocab{Writing: "是"},
		znr.Vocab{Writing: "好"},
		znr.Vocab{Writing: "友"},
	}
	cases := []struct {
		in  string
		out string
	}{
		{in: "你好世界", out: "你,好"},
		{in: "你是谁？", out: "你,是"},
		{in: "我是你的好朋友", out: "是,你,好,友"},
	}

	for _, c := range cases {
		known, err := znr.KnownPhrases(c.in, vl)
		if err != nil {
			t.Error(err)
		}

		if reflect.DeepEqual(known, strings.Split(c.out, ",")) == false {
			t.Errorf("%v != %v", known, c.out)
		}
	}
}
