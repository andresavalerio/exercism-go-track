package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starsLine := strings.Repeat("*", numStarsPerLine) 
	return starsLine + "\n" + welcomeMsg + "\n" + starsLine
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	lines := strings.Split(oldMsg, "\n")
	return strings.Trim(lines[1], "* ")
}
