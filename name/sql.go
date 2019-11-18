package name

import "strings"

// SQLObject snake-cases the given parts to derive an SQLObject object name.
func SQLObject(parts ...string) string {
	for i, part := range parts {
		parts[i] = strings.ToLower(part)
	}
	return strings.Join(parts, "_")
}
