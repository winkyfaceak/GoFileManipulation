package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		switch text {
		case "exit":
			fmt.Println("Exiting the program.")
			return
		case "hello":
			fmt.Println("Hello there!")
		case "create":
			fmt.Println("Enter the filename:")
			scanner.Scan()
			filename := scanner.Text()
			fmt.Println("Enter the content:")
			scanner.Scan()
			content := scanner.Text()
			createFile(filename, content)
		case "read":
			fmt.Println("Enter the filename to read:")
			scanner.Scan()
			filename := scanner.Text()
			readDataFromFile(filename)
		default:
			saveBasedOnInput(text)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

func saveBasedOnInput(input string) {
	switch input {
	case "file1":
		save("file1.txt", "Data for file 1.")
	case "file2":
		save("file2.txt", "Data for file 2.")
	default:
		fmt.Println("Unknown command:", input)
	}
}

func save(filename, in string) {
	data := []byte(in)
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Data saved to", filename, "successfully.")
}

func readDataFromFile(filename string) {
	readData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	fmt.Println("Data read from file:", string(readData))

	err = os.Remove(filename)
	if err != nil {
		fmt.Println("Error removing file:", err)
		return
	}
	fmt.Println("File removed successfully.")
}

func createFile(filename, content string) {
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	fmt.Printf("File '%s' created successfully.\n", filename)
}