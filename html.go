package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path"

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
	TermsOfUseURL:    "/info/terms.html",
	PrivacyPolicyURL: "/info/privacy.html",
	CookiePolicyURL:  "/info/cookies.html",
}

// HTML returns HTML for the policy.
func HTML(policy string, vars Variables) ([]byte, error) {
	path := path.Join(srcDir, policy) + ".md"

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	input, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	output, err := renderBytes(input, vars)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func renderBytes(input []byte, vars Variables) ([]byte, error) {
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
