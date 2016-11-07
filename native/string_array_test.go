package native

import "testing"

func TestNativeStringArray_Each(t *testing.T) {
	var out string
	type fields struct {
		List []string
	}
	type args struct {
		fn func(string)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"iterates", fields{[]string{"1", "2", "3"}}, args{func(s string) { out += s }}},
	}
	for _, tt := range tests {
		out = ""
		a := &NativeStringArray{
			List: &tt.fields.List,
		}
		a.Each(tt.args.fn)
		if out != "123" {
			t.Errorf("expected each to fill out")
		}
	}
}

func TestNativeStringArray_Get(t *testing.T) {
	type fields struct {
		List []string
	}
	type args struct {
		idx int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"gets index", fields{[]string{"1"}}, args{0}, "1"},
	}
	for _, tt := range tests {
		a := &NativeStringArray{
			List: &tt.fields.List,
		}
		if got := a.Get(tt.args.idx); got != tt.want {
			t.Errorf("%q. NativeStringArray.Get() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNativeStringArray_Set(t *testing.T) {
	type fields struct {
		List []string
	}
	type args struct {
		idx   int
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"sets index", fields{[]string{"0"}}, args{0, "1"}},
	}
	for _, tt := range tests {
		a := &NativeStringArray{
			List: &tt.fields.List,
		}
		a.Set(tt.args.idx, tt.args.value)
		if a.Get(tt.args.idx) != tt.args.value {
			t.Errorf("%q. NativeStringArray.Get(%v) != %v", tt.name, tt.args.idx, tt.args.value)
		}
	}
}

func TestNativeStringArray_Push(t *testing.T) {
	type fields struct {
		List []string
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"pushes", fields{[]string{"1", "2", "3"}}, args{"4"}},
	}
	for _, tt := range tests {
		a := &NativeStringArray{
			List: &tt.fields.List,
		}
		sz := a.Size()
		a.Push(tt.args.value)
		if a.Size() != sz+1 {
			t.Fatalf("expected value to be pushed (%d != %d)", a.Size(), sz+1)
		}
		if a.Get(sz) != tt.args.value {
			t.Errorf("expected last index to be %q", tt.args.value)
		}
	}
}

func TestNativeStringArray_Size(t *testing.T) {
	type fields struct {
		List []string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"size check", fields{[]string{"1", "2", "3"}}, 3},
	}
	for _, tt := range tests {
		a := &NativeStringArray{
			List: &tt.fields.List,
		}
		if got := a.Size(); got != tt.want {
			t.Errorf("%q. NativeStringArray.Size() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNativeStringArray_Length(t *testing.T) {
	type fields struct {
		List []string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"size check", fields{[]string{"1", "2", "3"}}, 3},
	}
	for _, tt := range tests {
		a := &NativeStringArray{
			List: &tt.fields.List,
		}
		if got := a.Length(); got != tt.want {
			t.Errorf("%q. NativeStringArray.Length() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
