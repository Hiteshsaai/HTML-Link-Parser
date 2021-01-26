package link

// This Link represents the links in the HTML tag
type Link struct {
	href string
	Text string
}

// Parse will take an HTML document and it will
// return slices
func Parser(r io.reader) ([]Link, error) {
	return nil, nil
}
