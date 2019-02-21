package legal

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/recentralized/legal/src"
	"github.com/russross/blackfriday"
)

// Variables are the placeholder values used by the markdown templates.
type Variables struct {

	// Email address that users should contact.
	ContactEmail string

	// URL pointing to the Terms of Use.
	TermsOfUseURL string

	// URL pointing to the Privacy Policy.
	PrivacyPolicyURL string

	// URL pointing to the Cookie Policy.
	CookiePolicyURL string
}

// DefaultVariables is the standard values for variables.
var DefaultVariables = Variables{
	ContactEmail:     "legal@recentalized.org",
	TermsOfUseURL:    "terms.html",
	PrivacyPolicyURL: "privacy.html",
	CookiePolicyURL:  "cookies.html",
}

// HTML returns HTML for the policy.
func HTML(policy string, vars Variables) ([]byte, error) {
	input, err := read(policy)
	if err != nil {
		return nil, err
	}
	output, err := render(input, vars)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func render(input []byte, vars Variables) ([]byte, error) {
	// First render as golang template.
	tmpl, err := template.New("content").Parse(string(input))
	if err != nil {
		return nil, err
	}
	var out bytes.Buffer
	err = tmpl.Execute(&out, vars)
	if err != nil {
		return nil, err
	}
	// Then render as markdown.
	return blackfriday.Run(out.Bytes()), nil
}

func read(name string) ([]byte, error) {
	path := filepath.Join(src.GetPath(), name) + ".md"
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}
