package pointer

// "Gambiarra" to get a field as a pointer, without doing ugly things like &[]string{"str"}[0]
// https://www.urbandictionary.com/define.php?term=Gambiarra

func String(s string) *string {
	return &s
}
