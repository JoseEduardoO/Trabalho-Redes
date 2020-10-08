package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "#Teste 1",
			args: []int{10, -1, 0, 50, 99, -5},
			want: []int{-5, -1, 0, 10, 50, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "#Teste 1",
			args: []int{10, -1, 0, 50, 99, -5},
			want: []int{-5, -1, 0, 10, 50, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSort(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "#Teste 1",
			args: []int{10, -1, 0, 50, 99, -5},
			want: []int{-5, -1, 0, 10, 50, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectionSort(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
