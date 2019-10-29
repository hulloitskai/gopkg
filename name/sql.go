package name

import "strings"

// SQLTable snake-cases the given parts to derive an SQL table name.
func SQLTable(parts ...string) string {
	for i, part := range parts {
		parts[i] = strings.ToLower(part)
	}
	return strings.Join(parts, "_")
}
