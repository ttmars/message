package core

import "testing"

// 要注释defer wg1.Done()
func TestGithub(t *testing.T) {
	F("https://github.com/trending/go?since=daily", "Github0")
}
