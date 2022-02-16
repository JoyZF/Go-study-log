package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{name: "case1", args: args{x: 1}, wantResult: 3},
		{name: "case2", args: args{x: 2}, wantResult: 4},
		{name: "case3", args: args{x: 3}, wantResult: 5},
		{name: "case4", args: args{x: 4}, wantResult: 6},
		{name: "case5", args: args{x: 5}, wantResult: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, Calculate(tt.args.x), tt.wantResult)
		})
	}
}
