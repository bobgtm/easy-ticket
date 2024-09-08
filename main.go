package main

func main() {
	filename := ReadDir()

	ParseJson(filename, "act.csv")
}
