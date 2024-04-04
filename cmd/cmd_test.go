package cmd

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"testing"
)

func TestGowitnesses(t *testing.T) {
	GoWitnessess([]string{os.Getenv("DOMAIN")}, 4)
}
