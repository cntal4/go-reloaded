package testutils

import (
	"math/rand"
	"testing"
	"time"
)

// generateRandomText creates random text with various formatting markers
func generateRandomText(seed int64) string {
	rand.Seed(seed)

	words := []string{"hello", "world", "test", "amazing", "great", "simple"}
	markers := []string{"(cap)", "(up)", "(low)", "(hex)", "(bin)"}
	punctuation := []string{".", ",", "!", "?", "..."}

	var result string
	length := rand.Intn(10) + 5 // 5-15 elements

	for i := 0; i < length; i++ {
		if i > 0 {
			result += " "
		}

		switch rand.Intn(4) {
		case 0:
			result += words[rand.Intn(len(words))]
		case 1:
			result += markers[rand.Intn(len(markers))]
		case 2:
			result += punctuation[rand.Intn(len(punctuation))]
		case 3:
			result += "' " + words[rand.Intn(len(words))] + " '"
		}
	}

	return result
}

func TestPropertyIdempotency(t *testing.T) {
	// Test idempotency with randomly generated inputs
	for i := 0; i < 50; i++ {
		seed := time.Now().UnixNano() + int64(i)
		input := generateRandomText(seed)

		t.Run("random_input", func(t *testing.T) {
			first := FormatText(input)
			second := FormatText(first)

			if first != second {
				t.Errorf("Property test failed for seed %d:\n"+
					"Input: %q\n"+
					"First: %q\n"+
					"Second: %q", seed, input, first, second)
			}
		})
	}
}

func TestRoundTripConsistency(t *testing.T) {
	// Test that format -> format -> format is consistent
	inputs := []string{
		"hello (cap) world",
		"add 42 (hex) and 10 (bin)",
		"a apple and a orange",
		"wait ... what !? really",
		"he said ' hello world '",
	}

	for _, input := range inputs {
		t.Run("round_trip", func(t *testing.T) {
			pass1 := FormatText(input)
			pass2 := FormatText(pass1)
			pass3 := FormatText(pass2)

			if pass2 != pass3 {
				t.Errorf("Round trip consistency failed:\n"+
					"Input: %q\n"+
					"Pass2: %q\n"+
					"Pass3: %q", input, pass2, pass3)
			}
		})
	}
}
