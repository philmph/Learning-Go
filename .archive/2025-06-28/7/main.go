package main

func main() {
	u := &user{
		name: "Philipp",
		age:  33,
	}

	u.greet()
	u.greetNoPointer()

	uNoPointer := user{
		name: "Elsie",
		age:  3,
	}

	uNoPointer.greet()
	uNoPointer.greetNoPointer()

	// Style guide test
	// VSCode config test
}
