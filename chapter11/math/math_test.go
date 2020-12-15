package math

import "testing"

func TestAverage(t *testing.T) {
	type args struct {
		xs []float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				xs: []float32{3,4,3,2},
			},
			want: 3,
		},
		{
			name: "test 2",
			args: args{
				xs: []float32{1.5, 6.7, 2.6},
			},
			want: 3.5999997,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.args.xs...); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		xs []float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				xs: []float32{3,5,7,10},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.xs...); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		xs []float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
		{
			name: "min test 1",
			args: args{
				xs: []float32{4.5 , 7, 9, 3.4},
			},
			want: 3.4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.xs...); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}