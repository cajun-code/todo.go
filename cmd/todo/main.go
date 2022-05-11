package main

import (
	"flag"
	"fmt"
	"os"

	"shadeauxmedia.com/tools/todo"
)

var todoFileName = ".todo.json"

const envFileNameKey = "TODO_FILENAME"

func main() {
	task := flag.String("task", "", "Task to be included in the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	completed := flag.Int("complete", 0, "Item to be completed")

	if os.Getenv(envFileNameKey) != "" {
		todoFileName = os.Getenv(envFileNameKey)
	}
	flag.Parse()

	l := &todo.List{}
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		fmt.Print(l)
		// for _, item := range *l {
		// 	if !item.Done {
		// 		fmt.Println(item.Task)
		// 	}
		// }
	case *completed > 0:
		if err := l.Complete(*completed); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		l.Add(*task)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid Option")
		os.Exit(1)
	}
}
