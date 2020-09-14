package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main()  {
	fmt.Println("Write a number")	
	operador := "-"
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operation := scanner.Text()
	fmt.Println(operation)	
	valors := strings.Split(operation,"+")
	fmt.Println(valors)	
	fmt.Println(valors[0]+valors[1])	
	operation_a, err1 := strconv.Atoi(valors[0])
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(operation_a)
	}
	operation_b, _ := strconv.Atoi(valors[1])
	fmt.Println(operation_a + operation_b)	
	switch operador {
	case "+":
		fmt.Println(operation_a + operation_b)
	
	case "-":
		fmt.Println(operation_a - operation_b)
	case "*":
		fmt.Println(operation_a * operation_b)
	case "/":
		fmt.Println(operation_a / operation_b)
	default:
		fmt.Println("Nope :(")
	}
	
}