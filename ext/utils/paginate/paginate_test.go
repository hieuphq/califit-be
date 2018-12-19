package paginate

import (
	"reflect"
	"testing"
)

func TestMakePaginate(t *testing.T) {
	type args struct {
		total  int
		offset int
		limit  int
	}
	tests := []struct {
		name string
		args args
		want *Paginate
	}{
		{
			name: "case 1",
			args: args{total: 10, offset: 9, limit: 4},
			want: &Paginate{
				CurrentPage: 3,
				TotalPage:   3,
				TotalItem:   10,
			},
		},
		{
			name: "case 2",
			args: args{total: 10, offset: 3, limit: 10},
			want: &Paginate{
				CurrentPage: 1,
				TotalPage:   1,
				TotalItem:   10,
			},
		},
		{
			name: "case 3",
			args: args{total: 20, offset: 10, limit: 10},
			want: &Paginate{
				CurrentPage: 2,
				TotalPage:   2,
				TotalItem:   20,
			},
		},
		{
			name: "case 4",
			args: args{total: 20, offset: 10, limit: 0},
			want: &Paginate{
				CurrentPage: 0,
				TotalPage:   0,
				TotalItem:   0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakePaginate(tt.args.total, tt.args.offset, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakePaginate() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
