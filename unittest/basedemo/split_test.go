package basedemo

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestSplitWithComplexSep(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestSplitAll(t *testing.T) {
	t.Parallel() // 将 TLog 标记为能够与其他测试并行运行
	// 定义测试表格
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{"base case", "a:b:c", ":", []string{"a", "b", "c"}},
		{"wrong sep", "a:b:c", ",", []string{"a", "b", "c"}},
		{"more sep", "abcd", "bc", []string{"a", "d"}},
		{"leading sep", "沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // 标记每个子测试能够彼此并行运行
			got := Split(tt.input, tt.sep)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("expected: %#v, got: %#v", tt.want, got)
			//}
			assert.Equal(t, tt.want, got)
		})
	}
}
