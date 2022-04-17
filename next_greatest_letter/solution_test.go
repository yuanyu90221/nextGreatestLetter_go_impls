package next_greatest_letter

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "Example1",
			args: args{letters: []byte{'c', 'f', 'j'}, target: 'a'},
			want: 'c',
		},
		{
			name: "Example2",
			args: args{letters: []byte{'c', 'f', 'j'}, target: 'c'},
			want: 'f',
		},
		{
			name: "Example3",
			args: args{letters: []byte{'c', 'f', 'j'}, target: 'd'},
			want: 'f',
		},
		{
			name: "Example4",
			args: args{letters: []byte{'c', 'f', 'j'}, target: 'j'},
			want: 'c',
		},
		{
			name: "Example5",
			args: args{letters: []byte{'c', 'f', 'j'}, target: 'k'},
			want: 'c',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %c, want %c", got, tt.want)
			}
		})
	}
}
