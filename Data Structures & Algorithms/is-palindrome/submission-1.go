
func isPalindrome(s string) bool {
	if len(s) == 1{
		return true
	}

	reg := regexp.MustCompile("[^a-zA-Z0-9 ]+")
	regexFormated := reg.ReplaceAllString(s, "")

	formatedString := strings.Split(regexFormated, " ")
	regexFormated = strings.Join(formatedString, "")
	regexFormated = strings.ToLower(regexFormated)
	formatedString = strings.Split(regexFormated, "")

	pointer1 := 0
	pointer2 := len(formatedString) - 1

	for pointer2 > pointer1{
		if formatedString[pointer1] != formatedString[pointer2]{
			return false
		}
		pointer1 ++
		pointer2 --
	}

	return true

}