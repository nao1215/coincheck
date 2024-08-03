package coincheck

import "testing"

// printDiff prints the gocmp diff.
func printDiff(t *testing.T, diff string) {
	t.Helper()
	t.Errorf("differs: (-want +got)\n%s", diff)
}
