package uid

import "testing"

const (
	caseOk    = "ok"
	caseError = "error"
)

func TestGenUID(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{
			name: caseOk,
			in:   10,
			want: 10,
		},
		{
			name: caseError,
			in:   10,
			want: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u, err := GenUID(test.in)
			if err != nil {
				t.Fatal(err)
			}

			switch test.name {
			case caseOk:
				if len(u) != test.want {
					t.Fatalf("got %d, want %d", len(u), test.want)
				}
			case caseError:
				if len(u) == test.want {
					t.Fatalf("got %d, want %d", len(u), test.want)
				}
			}
		})
	}
}

func TestGenSpecialUID(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{
			name: caseOk,
			in:   10,
			want: 10,
		},
		{
			name: caseError,
			in:   10,
			want: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u, err := GenSpecialUID(test.in)
			if err != nil {
				t.Fatal(err)
			}

			switch test.name {
			case caseOk:
				if len(u) != test.want {
					t.Fatalf("got %d, want %d", len(u), test.want)
				}
			case caseError:
				if len(u) == test.want {
					t.Fatalf("got %d, want %d", len(u), test.want)
				}
			}
		})
	}
}
