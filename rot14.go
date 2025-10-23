package piscine

func Rot14(s string) string {
result := ""

for _, r := range s {
	if r >= 'a' && r <= 'z' {
		r = r + 14
		if r > 'z' {
			r = r - 26
		}
	} else if r >= 'A' && r <= 'Z' {
		r = r + 14
		if r > 'Z' {
			r = r - 26
		}
	}
	result = result + string(r)
}

return result


}

/*
func Rot14(s string) string {
	result := result + string (r)
for _, r := range s{
	if r >= 'a' && r<= 'z' || r >= 'A' && r <= 'Z'{
		r = r + 14
		if r > 'Z' {
			r =  r - 26
		}
		z01.PrintRune(r)
	} else {
		z01.PrintRune('0')
	}
	
}
}
*/