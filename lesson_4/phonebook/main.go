package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"phonebook/book"
	"phonebook/logger"
	"strings"
	"time"
)

func main() {
	//points0 := make(map[string]int)

	//points0["alice"] = 4
	//points0["alex"] = 8

	points2 := map[string]int{
		"alice": 6,
		"alex":  14,
	}

	fmt.Println(points2)

	fmt.Println("Phonebook")
	fmt.Println("Available comands: add, get, delete, list, exit")

	scanner := bufio.NewScanner(os.Stdin)
	//phoneBook := make(map[string]string)
	phoneBook := make(book.PhoneBook)

	for {
		fmt.Print(" >")
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		command := parts[0]
		args := parts[1:]

		switch command {
		case "add":
			// RUN: add toad=123
			// kv := strings.SplitN(parts[1], "=", 2)
			// if len(kv) != 2 {
			//     fmt.Println("Invalid. Use name=number")
			//     continue
			// }
			// name, number := kv[0], kv[1]
			// phoneBook[name] = number
			// fmt.Printf("Add/Update: %s -> %s\n", name, number)
			handleCommand(doAdd, args, phoneBook)
		case "get":
			// name := parts[1]
			// number, exist := phoneBook[name]
			// if exist {
			// 	fmt.Printf("Number for %s is %s\n", name, number)
			// } else {
			// 	fmt.Printf("No result for %s is %s\n", name, number)
			// }
			handleCommand(doGet, args, phoneBook)
		case "delete":
			// name := parts[1]
			// _, exist := phoneBook[name]
			// if exist {
			// 	fmt.Printf("Delete %s\n", name)
			// } else {
			// 	fmt.Printf("No result for %s\n", name)
			// }
			handleCommand(doDelete, args, phoneBook)
		case "update":
			// kv := strings.SplitN(parts[1], "=", 2)
			// if len(kv) != 2 {
			// 	fmt.Println("Invalid format")
			// 	continue
			// }
			// name, newNumber := kv[0], kv[1]
			// _, exists := phoneBook[name]
			// if exists {
			// 	phoneBook[name] = newNumber
			// 	fmt.Printf("Updated: %s -> %s\n", name, newNumber)
			// } else {
			// 	fmt.Println("No Entry found")
			// }
			handleCommand(doUpdate, args, phoneBook)
		case "list":
			// if len(phoneBook) == 0 {
			// 	fmt.Println("PhoneBook is empty")
			// } else {
			// 	for name, number := range phoneBook {
			// 		fmt.Printf("%s -> %s", name, number)
			// 	}
			// }
			handleCommand(doList, args, phoneBook)
		case "exit":
			fmt.Println("Exiting phonebook")
			return
		}
	}
}

func handleCommand(cmd func([]string, book.PhoneBook) error, args []string, phoneBook book.PhoneBook) {
	if err := cmd(args, phoneBook); err != nil {
		logger.Warn(err, "cmd failed")
	}

}

func doUpdate(args []string, phoneBook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("missing parameters for 'update' command. Use: update name=new_number")
	}

	kv := strings.SplitN(args[0], "=", 2)
	if len(kv) != 2 {
		return errors.New("invalid format. Use: update name=number")
	}

	name, newNumber := kv[0], kv[1]

	err := phoneBook.Update(name, newNumber)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("Updated an entry: %s -> %s\n", name, newNumber))

	return nil
}
func doDelete(args []string, phoneBook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("missing parameters for 'delete' command. Use: delete name")
	}

	name := args[0]

	err := phoneBook.Delete(name)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("Deleted entry for %s\n", name))

	return nil
}

func doList(_ []string, phoneBook book.PhoneBook) error {
	if len(phoneBook) == 0 {
		return errors.New("phonebook is empty")
	} else {
		results := ""

		for name, number := range phoneBook {
			results += fmt.Sprintf("%s -> %s\n", name, number.Number)
		}

		logger.Info(results)
	}

	return nil
}

func doGet(args []string, phoneBook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("missing parameters for 'get' command. Use: get name")
	}

	name := args[0]

	numberData, err := phoneBook.Get(name)
	if err != nil {
		return err
	}

	unixUpdatedAt := time.Unix(numberData.LastUpdatedAt, 0)

	logger.Info(
		fmt.Sprintf("Number for %s is %s (last updated at %s)\n",
			name,
			numberData.Number,
			unixUpdatedAt.Format("2006-01-02 15:04:05"),
		),
	)

	return nil
}

func doAdd(args []string, phoneBook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("missing parameters for 'add' command. Use: add name=number")
	}

	kv := strings.SplitN(args[0], "=", 2)
	if len(kv) != 2 {
		return errors.New("invalid format. Use: add name=number")
	}

	name, number := kv[0], kv[1]
	err := phoneBook.Add(name, number)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("Added an entry: %s -> %s\n", name, number))

	return nil
}
