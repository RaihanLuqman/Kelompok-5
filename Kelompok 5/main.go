package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		num1, _ := strconv.ParseFloat(r.Form.Get("num1"), 64)
		num2, _ := strconv.ParseFloat(r.Form.Get("num2"), 64)
		operator := r.Form.Get("operator")

		var result float64
		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 != 0 {
				result = num1 / num2
			} else {
				http.Error(w, "Division by zero", http.StatusBadRequest)
				return
			}
		default:
			http.Error(w, "Invalid operator", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Result: %.2f", result)
	} else {
		http.ServeFile(w, r, "calculator.html")
	}
}

func main() {
	http.HandleFunc("/", mainHandler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
