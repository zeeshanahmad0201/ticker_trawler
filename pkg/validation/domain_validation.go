package validation

import (
	"regexp"
	"strings"
)

// ParseDomain ensures that the domain string has a proper HTTP/HTTPS prefix and a trailing slash.
// It adds "http://" or "https://" if missing and ensures a trailing slash is present.
func ParseDomain(domain string) string {
	if !ValidDomain(domain) {
		return ""
	}

	// Check if the domain starts with "http://" or "https://"
	if !strings.HasPrefix(domain, "http://") && !strings.HasPrefix(domain, "https://") {
		// Default to "http://"
		domain = "http://" + domain
	}

	// Ensure the domain ends with a trailing slash
	if !strings.HasSuffix(domain, "/") {
		domain += "/"
	}

	return domain
}

// ValidDomain checks if the provided domain string is a valid domain name.
// A valid domain name consists of alphanumeric characters and hyphens, with specific rules for hyphen placement.
func ValidDomain(domain string) bool {
	// Regular expression to match a valid domain
	const domainPattern = `^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`

	// Compile the regular expression
	reg := regexp.MustCompile(domainPattern)

	// Validate the domain
	return reg.MatchString(domain)
}
