package calculator_testing

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		number_a int
		number_b int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "success 1 : positive input",
			args: args{
				number_a: 2,
				number_b: 14,
			},
			want: 16,
		},
		{
			name: "success 2 : negative input",
			args: args{
				number_a: 2,
				number_b: -14,
			},
			want: -12,
		},
		{
			name: "success 3 : zero input",
			args: args{
				number_a: 2,
				number_b: 0,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.number_a, tt.args.number_b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	type args struct {
		number_a int
		number_b int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "success 1 : positive input",
			args: args{
				number_a: 14,
				number_b: 2,
			},
			want: 12,
		},
		{
			name: "success 2 : negative input",
			args: args{
				number_a: -2,
				number_b: 14,
			},
			want: -16,
		},
		{
			name: "success 3 : zero input",
			args: args{
				number_a: 2,
				number_b: 0,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sub(tt.args.number_a, tt.args.number_b); got != tt.want {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMult(t *testing.T) {
	type args struct {
		number_a int
		number_b int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "success 1 : positive input",
			args: args{
				number_a: 2,
				number_b: 14,
			},
			want: 28,
		},
		{
			name: "success 2 : negative input",
			args: args{
				number_a: 2,
				number_b: -14,
			},
			want: -28,
		},
		{
			name: "success 3 : zero input",
			args: args{
				number_a: 2,
				number_b: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mult(tt.args.number_a, tt.args.number_b); got != tt.want {
				t.Errorf("Mult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	type args struct {
		number_a int
		number_b int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "success 1 : positive input",
			args: args{
				number_a: 12,
				number_b: 4,
			},
			want: 3,
		},
		{
			name: "fail 1 : negative input",
			args: args{
				number_a: 2,
				number_b: -14,
			},
			want: 0,
		},
		{
			name: "fail 2 : zero input",
			args: args{
				number_a: 0,
				number_b: 2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Div(tt.args.number_a, tt.args.number_b); got != tt.want {
				t.Errorf("Div() = %v, want %v", got, tt.want)
			}
		})
	}
}
