package misc

import "fmt"

func ByteSliceTutorial() {
	fmt.Println("------------")
	fmt.Println("Byte Slice")
	data := []byte{65, 66, 67} //ASCII values A,B,C

	// Append more bytes
	data = append(data, 68, 69)

	// Convert to a string
	str := string(data)

	fmt.Printf("Byte slice: %v\n", data)
	fmt.Printf("String: %s\n", str)
	fmt.Println("------------")
}

/*
A "byte slice" in Go is a fundamental data structure used
to represent a sequence of bytes. It is a dynamic, resizable
collection of bytes and is similar to what is commonly
referred to as a "byte array" in other programming languages.
Byte slices are often used for tasks involving binary data,
file I/O, and data serialization.

Here are some key characteristics and information about byte
slices in Go:

Dynamic Size: Unlike arrays in Go, which have a fixed size,
byte slices can dynamically grow or shrink as needed.
This dynamic behavior makes them more flexible for working
with data of unknown or varying lengths.

Slice Literal: You can create a byte slice using a slice
literal, which is a compact way of specifying the initial
contents of the slice. For example:

data := []byte{65, 66, 67} // Creates a byte slice with ASCII values 'A', 'B', 'C'

Byte Conversion: Byte slices can be easily converted to
strings and vice versa. This is useful when working with
text data or binary data that needs to be converted to a
human-readable format.

Common Uses:

Reading and writing binary data from files or network sockets.
Encoding and decoding data in various formats (e.g., JSON, XML, binary).
Manipulating and processing binary data, such as images, audio, or cryptographic data.
Working with raw network packets or protocol data.
Built-in Functions: Go provides many built-in functions for working with byte slices,
such as append(), copy(), len(), and cap(), which allow you to
manipulate and manage the contents of byte slices efficiently.

Slicing: Byte slices can be sliced to create new slices that
reference a portion of the original data. This is done without
copying the data, which makes it efficient for working with large
data sets.

Byte and Rune: Be aware that in Go, a byte represents a single byte
of data, which can range from 0 to 255. If you're working with
Unicode characters or strings, you may also encounter the
term "rune," which represents a Unicode code point and can be more
than one byte in size.
*/
