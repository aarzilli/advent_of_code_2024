package util

type SplitsimilarFlags uint8

const (
	SSHasNegatives SplitsimilarFlags = 1 << iota
	SSRemoveSymbols
	SSRemoveAlpha
)

func Splitsimilar(in string, flags SplitsimilarFlags) []string {
	r := []string{}
	start := 0

	category := func(s string) int {
		switch {
		case (s[0] >= 'a' && s[0] <= 'z') || (s[0] >= 'A' && s[0] <= 'Z'):
			return 0
		case s[0] >= '0' && s[0] <= '9':
			return 1
		case s[0] == '-':
			if (flags&SSHasNegatives != 0) && len(s) > 1 && (s[1] >= '0' && s[0] <= '9') {
				return 1
			} else {
				return 2
			}
		default:
			return 2
		}
	}

	add := func(s string) {
		switch category(s) {
		case 1:
			r = append(r, s)
		case 0:
			if flags&SSRemoveAlpha == 0 {
				r = append(r, s)
			}
		case 2:
			if flags&SSRemoveSymbols == 0 {
				r = append(r, s)
			}
		}
	}

	for i := range in {
		if category(in[i:]) != category(in[start:]) {
			add(in[start:i])
			start = i
		}
	}

	add(in[start:])

	return r
}
