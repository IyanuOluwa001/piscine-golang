package piscine

func Rot14(s string) string {
	convertedString := ""

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
		convertedString = convertedString + string(r)
	}

	return convertedString
}

/*
func Rot14(s string) string {
	convertedString := ""

	for _, givenText := range s {
		if givenText >= 'a' && givenText <= 'z' {
			givenText = givenText + 14
			if givenText > 'z' {
				givenText = givenText - 26
			}
		} else if givenText >= 'A' && givenText <= 'Z' {
			givenText = givenText + 14
			if givenText > 'Z' {
				givenText = givenText - 26
			}
		}
		convertedString = convertedString + string(givenText)
	}

	return convertedString
}
*/

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
