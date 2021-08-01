package main

import "fmt"

/*
Constants should improve performance of your application as it saves you
creating the "Hello, " string instance every time Hello is called.
To be clear, the performance boost is incredibly negligible for this example!
But it's worth thinking about creating constants to capture the meaning of values and
sometimes to aid performance.
*/

const enHelloPrefix = "Hello "
const deHelloPrefix = "Hallo "
const enHelloSuffix = "world"

func hello(name string, language string) string {

	if name == "" {
		name = enHelloSuffix
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case "EN":
		prefix = enHelloPrefix
	case "DE":
		prefix = deHelloPrefix
	default:
		prefix = enHelloPrefix
	}
	return
}

func main() {
	fmt.Println(hello("world", "EN"))
}
