package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	type Todo struct {
		id          int
		description string
		status      bool
	}

	counter := 0
	data := make([]Todo, 0)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Commands : show create remove done")

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		switch split[0] {

		case "create":
			data = append(data, Todo{
				id:          counter,
				description: strings.Join(split[1:], " "),
				status:      false,
			})
			counter++
			fmt.Println("* Created the Item")
			fmt.Printf("* The ID of Item is %v \n ", counter-1)

		case "remove":
			index, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			for i, todo := range data {
				if todo.id == index {
					data = slices.Delete(data, i, i+1)
					break
				}
			}

		case "show":
			for _, todo := range data {
				fmt.Println(todo)
			}

		case "done":
			index, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			for i, todo := range data {
				if todo.id == index {
					data[i].status = true

				}
			}
		}
	}

}
