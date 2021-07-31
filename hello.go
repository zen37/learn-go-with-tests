package main

/*
Constants should improve performance of your application as it saves you
creating the "Hello, " string instance every time Hello is called.
To be clear, the performance boost is incredibly negligible for this example!
But it's worth thinking about creating constants to capture the meaning of values and
sometimes to aid performance.
*/

const enHelloPrefix = "Hello "

func hello(name string) string {
	//return "hello, " + name

	return enHelloPrefix + name

}

func main() {
	hello("world")
}
