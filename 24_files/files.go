package main

import (
	// "bufio"
	// "fmt"
	"os"
)

func main() {
	// f,err :=os.Open("example.txt") // Open a file named example.txt

	// if err != nil {
	// 	panic(err) // Handle the error if the file cannot be opened
	// }

	// fileInfo, err := f.Stat()

	// if err != nil {
	// 	panic(err) // Handle the error if file information cannot be retrieved
	// }

	// fmt.Println("File Name:", fileInfo.Name())
	// fmt.Println("File or Foldere:", fileInfo.IsDir())
	// fmt.Println("File Size:", fileInfo.Size())
	// fmt.Println("File Mode:", fileInfo.Mode())
	// fmt.Println("Last Modified:", fileInfo.ModTime())

	// read file

	// f, err := os.Open("example.txt") // Open a file named example.txt
	// if err != nil {
	// 	panic(err) // Handle the error if the file cannot be opened
	// }

	// defer f.Close() // Ensure the file is closed after reading

	// buf := make([]byte, 10) // Create a buffer to hold the file content

	// d, err := f.Read(buf)

	// if err != nil {
	// 	panic(err) // Handle the error if reading the file fails
	// }

	// for i :=0;i<len(buf); i++ {
	// 	println("data",d ,string(buf[i])) // Print the data read from the file
	// }
	// println(d,buf)

	// f, err := os.ReadFile("example.txt") // Read the entire file content into memory
	// if err != nil {
	// 	panic(err) // Handle the error if the file cannot be read
	// }

	// fmt.Println(string(f))

	// dir, err := os.Open("../") // Open the current directory

	// if err != nil {
	// 	panic(err) // Handle the error if the directory cannot be opened
	// }

	// defer dir.Close() // Ensure the directory is closed after reading

	// fileInfo, err := dir.ReadDir(-1)

	// for _, file := range fileInfo {
	// 	if file.IsDir() {
	// 		println("Directory:", file.Name()) // Print the name of the directory
	// 	} else {
	// 		println("File:", file.Name()) // Print the name of the file
	// 	}
	// }

	// create file
	// f, err := os.Create("example2.txt") // Create a new file named newfile.txt
	// if err != nil {
	// 	panic(err) // Handle the error if the file cannot be created
	// }

	// defer f.Close() // Ensure the file is closed after writing

	// // f.WriteString("Hello, World!") // Write a string to the file
	// // f.WriteString("Hello, World!2") // Write a string to the file

	// bytes := []byte("Hello, World!") // Create a byte slice to write to the file
	// f.Write(bytes)

	// read and write to another file (streaming fashion)

	// sourceFile, err := os.Open("example.txt") // Open the source file
	// if err != nil {
	// 	panic(err) // Handle the error if the file cannot be opened
	// }

	// defer sourceFile.Close() // Ensure the source file is closed after reading

	// destFile, err := os.Create("example2.txt") // Create the destination file

	// if err != nil {
	// 	panic(err) // Handle the error if the file cannot be created
	// }

	// defer destFile.Close() // Ensure the destination file is closed after writing

	// reader := bufio.NewReader(sourceFile) // Create a buffered reader for the source file
	// writer := bufio.NewWriter(destFile)   // Create a buffered writer for the destination file

	// for {
	// 	b, err := reader.ReadByte() // Read bytes from the source file until a newline character is encountered
	// 	if err != nil {
	// 		if err.Error() == "EOF" { // Check if the end of the file is reached
	// 			break // Exit the loop if the end of the file is reached
	// 		}
	// 		panic(err) // Handle any other error that occurs during reading
	// 	}

	// 	e := writer.WriteByte(b)

	// 	if e != nil {
	// 		panic(e) // Handle any error that occurs during writing
	// 	}
	// }
	// writer.Flush()                          // Flush the buffered writer to ensure all data is written to the destination file
	// fmt.Println("File copied successfully") // Print a success message

	// sourceFile, err := os.Open("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()

	err := os.Remove("example2.txt") // Remove the file named example2.txt
	if err != nil {
		panic(err) // Handle the error if the file cannot be removed
	}

	println("File removed successfully") // Print a success message

}
