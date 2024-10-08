package marketmath

import (
	"socialpredict/models"
	"testing"
)

// TestGetMarketVolume tests that the total volume of trades is returned correctly
func TestGetMarketVolume(t *testing.T) {
	tests := []struct {
		Name           string
		Bets           []models.Bet
		ExpectedVolume int64
	}{
		{
			Name:           "empty slice",
			Bets:           []models.Bet{},
			ExpectedVolume: 0,
		},
		{
			Name: "positive Bets",
			Bets: []models.Bet{
				{Amount: 100},
				{Amount: 200},
				{Amount: 300},
			},
			ExpectedVolume: 600,
		},
		{
			Name: "negative Bets",
			Bets: []models.Bet{
				{Amount: -100},
				{Amount: -200},
				{Amount: -300},
			},
			ExpectedVolume: -600,
		},
		{
			Name: "mixed Bets",
			Bets: []models.Bet{
				{Amount: 100},
				{Amount: -50},
				{Amount: 200},
				{Amount: -150},
			},
			ExpectedVolume: 100,
		},
		{
			Name: "large numbers",
			Bets: []models.Bet{
				{Amount: 9223372036854775807},
				{Amount: -9223372036854775806},
			},
			ExpectedVolume: 1,
		},
		{
			Name: "infinity avoidance",
			Bets: []models.Bet{
				{Amount: 1, Outcome: "YES"},
				{Amount: -1, Outcome: "YES"},
				{Amount: 1, Outcome: "NO"},
				{Amount: -1, Outcome: "NO"},
				{Amount: 1, Outcome: "NO"},
			},
			ExpectedVolume: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			volume := GetMarketVolume(test.Bets)
			if volume != test.ExpectedVolume {
				t.Errorf("%s: expected %d, got %d", test.Name, test.ExpectedVolume, volume)
			}
		})
	}
}
