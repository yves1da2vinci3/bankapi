package utils

import (
	"strings"
)

// IsEmptyString vérifie si une chaîne est vide ou composée uniquement d'espaces.
func IsEmptyString(s string) bool {
	return strings.TrimSpace(s) == ""
}

// ContainsString vérifie si une chaîne contient une sous-chaîne donnée.
func ContainsString(s, substr string) bool {
	return strings.Contains(s, substr)
}

// JoinStrings concatène une liste de chaînes avec un séparateur donné.
func JoinStrings(strings []string, separator string) string {
	return strings.Join(strings, separator)
}

// SplitString divise une chaîne en une liste de sous-chaînes en utilisant un séparateur donné.
func SplitString(s, separator string) []string {
	return strings.Split(s, separator)
}

// ToLowerString convertit une chaîne en minuscules.
func ToLowerString(s string) string {
	return strings.ToLower(s)
}

// ToUpperString convertit une chaîne en majuscules.
func ToUpperString(s string) string {
	return strings.ToUpper(s)
}
