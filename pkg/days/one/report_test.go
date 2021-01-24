package one

import "testing"

func TestReportRepair(t *testing.T) {
	type args struct {
		expenses []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Pair resulting 2020 is found",
			args:    args{expenses: []int{1721, 979, 366, 299, 675, 1456}},
			want:    514579,
			wantErr: false,
		},
		{
			name:    "Exhausting pair resulting 2020 is found",
			args:    args{expenses: []int{1456, 979, 366, 299, 675, 1721}},
			want:    514579,
			wantErr: false,
		},
		{
			name:    "No pair resulting 2020 is found",
			args:    args{expenses: []int{1456, 979, 366, 294, 675, 1721}},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReportRepair(tt.args.expenses)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReportRepair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReportRepair() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_check(t *testing.T) {
	type args struct {
		expenses []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name:  "Pair resulting 2020 is found",
			args:  args{expenses: []int{1721, 979, 366, 299, 675, 1456}},
			want:  0,
			want1: 3,
		},
		{
			name:  "Exhausting pair resulting 2020 is found",
			args:  args{expenses: []int{1456, 979, 366, 299, 675, 1721}},
			want:  3,
			want1: 5,
		},
		{
			name:  "No pair resulting 2020 is found",
			args:  args{expenses: []int{1456, 979, 366, 295, 675, 1721}},
			want:  -1,
			want1: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := check(tt.args.expenses)
			if got != tt.want {
				t.Errorf("check() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
