package main

import (
    "fmt"
    "github.com/username/base" // Replace with the correct import path
)

func main() {
    categories := []string{"Books", "Movies", "Music"}
    fileName := "categories.txt"

    success, err := base.CreateCategory(categories, fileName)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if success {
        fmt.Println("Categories created successfully!")
    }
}
