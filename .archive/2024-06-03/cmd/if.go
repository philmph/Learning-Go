package main

func ifhandling(n string) string {
	if o := ("Hello " + n + "!"); len(n) >= 3 {
		return o
	}

	return "Hello world!"
}
