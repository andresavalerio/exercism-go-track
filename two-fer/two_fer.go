/*
Two Fer package is the solution for the Exercism Two Fer problem.
We want to return the sentence "One for <name>, one for me."
If you don't know the persons name, it should substitute <name> for "you".
Developed by Andresa Valério.
*/
package twofer

// ShareWith returns a string showing who we're sharing with.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	} 
	return "One for " + name + ", one for me."
}
