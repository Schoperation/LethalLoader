package option

import (
	"fmt"
	"strconv"
	"unicode"
)

type OptionsResults struct {
	Task TaskName
	Page PageName
	Num  int
}

type NewOptionsArgs struct {
	Tasks map[string]TaskName
	Pages map[string]PageName
}

type Options struct {
	tasks         map[rune]TaskName
	pages         map[rune]PageName
	tasksWithNums map[rune]TaskName
	pagesWithNums map[rune]PageName
}

func NewOptions(args NewOptionsArgs) Options {
	options := Options{
		tasks:         make(map[rune]TaskName),
		pages:         make(map[rune]PageName),
		tasksWithNums: make(map[rune]TaskName),
		pagesWithNums: make(map[rune]PageName),
	}

	for letter, task := range args.Tasks {
		if len(letter) == 1 {
			options.tasks[rune(letter[0])] = task
			continue
		}

		options.tasksWithNums[rune(letter[0])] = task
	}

	for letter, page := range args.Pages {
		if len(letter) == 1 {
			options.pages[rune(letter[0])] = page
			continue
		}

		options.pagesWithNums[rune(letter[0])] = page
	}

	return options
}

func (ops Options) TakeInput() OptionsResults {
	var choice string
	var taskName TaskName
	var pageName PageName
	var num int

	for {
		fmt.Print(">")
		fmt.Scanf("%s", &choice)

		taskName, pageName, num = ops.parse(choice)
		if taskName != "" || pageName != "" {
			return OptionsResults{
				Task: taskName,
				Page: pageName,
				Num:  num,
			}
		}

		fmt.Printf("The hell are you saying? Use one of the available options.\n")
	}
}

func (ops Options) parse(choice string) (TaskName, PageName, int) {
	if len(choice) == 0 {
		return "", "", 0
	}

	var taskName TaskName
	var pageName PageName
	num := 0
	hasNum := false

	chosenLetter := unicode.ToUpper(rune(choice[0]))

	if task, exists := ops.tasks[chosenLetter]; exists {
		taskName = task
	}

	if task, exists := ops.tasksWithNums[chosenLetter]; exists {
		taskName = task
		hasNum = true
	}

	if page, exists := ops.pages[chosenLetter]; exists {
		pageName = page
	}

	if page, exists := ops.pagesWithNums[chosenLetter]; exists {
		pageName = page
		hasNum = true
	}

	if hasNum {
		if len(choice) == 1 {
			return "", "", 0
		}

		var err error
		num, err = strconv.Atoi(choice[1:])
		if err != nil {
			return "", "", 0
		}
	}

	return taskName, pageName, num
}
