package day01

import "testing"

func Test_searchForTwoEntriesEqualTo(t *testing.T) {
	type args struct {
		expenses    []int
		expectedSum int
	}
	tests := []struct {
		name       string
		args       args
		wantFirst  int
		wantSecond int
		wantErr    bool
	}{
		{
			name: "Pair resulting 2020 is found",
			args: args{
				expenses:    []int{299, 366, 675, 979, 1456, 1721},
				expectedSum: 2020,
			},
			wantFirst:  299,
			wantSecond: 1721,
			wantErr:    false,
		},
		{
			name: "No pair resulting 2020 is found",
			args: args{
				expenses:    []int{295, 366, 675, 979, 1456, 1721},
				expectedSum: 2020,
			},
			wantFirst:  0,
			wantSecond: 0,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirst, gotSecond, err := searchForTwoEntriesEqualTo(tt.args.expenses, tt.args.expectedSum)
			if (err != nil) != tt.wantErr {
				t.Errorf("searchForTwoEntriesEqualTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotFirst != tt.wantFirst {
				t.Errorf("searchForTwoEntriesEqualTo() gotFirst = %v, want %v", gotFirst, tt.wantFirst)
			}
			if gotSecond != tt.wantSecond {
				t.Errorf("searchForTwoEntriesEqualTo() gotSecond = %v, want %v", gotSecond, tt.wantSecond)
			}
		})
	}
}
