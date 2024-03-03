package option

import (
	"fmt"
	"strconv"
	"unicode"
)

type OptionsDto struct {
	Tasks map[string]TaskName
	Pages map[string]PageName
}

type OptionsResults struct {
	Task TaskName
	Page PageName
	Args any
}

type Options struct {
	tasks         map[rune]TaskName
	pages         map[rune]PageName
	tasksWithNums map[rune]TaskName
	pagesWithNums map[rune]PageName
	possibleArgs  []any
}

func NewOptions[E any](dto OptionsDto, args []E) Options {
	castedArgs := toAnySlice(args)

	options := Options{
		tasks:         make(map[rune]TaskName),
		pages:         make(map[rune]PageName),
		tasksWithNums: make(map[rune]TaskName),
		pagesWithNums: make(map[rune]PageName),
		possibleArgs:  castedArgs,
	}

	for letter, task := range dto.Tasks {
		if len(letter) == 1 {
			options.tasks[rune(letter[0])] = task
			continue
		}

		options.tasksWithNums[rune(letter[0])] = task
	}

	for letter, page := range dto.Pages {
		if len(letter) == 1 {
			options.pages[rune(letter[0])] = page
			continue
		}

		options.pagesWithNums[rune(letter[0])] = page
	}

	return options
}

func (ops Options) TakeInput() OptionsResults {
	var taskName TaskName
	var pageName PageName
	var num int

	for {
		var choice string
		fmt.Print(">")
		fmt.Scanf("%s", &choice)

		taskName, pageName, num = ops.parse(choice)
		if taskName != "" || pageName != "" {
			var args any
			if num > 0 {
				args = ops.possibleArgs[num-1]
			}

			return OptionsResults{
				Task: taskName,
				Page: pageName,
				Args: args,
			}
		}

		if num == -1 {
			fmt.Printf("Invalid number. Must be between 1 and %d", len(ops.possibleArgs))
			continue
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

		if num > len(ops.possibleArgs) || num <= 0 {
			return "", "", -1
		}
	}

	return taskName, pageName, num
}

func toAnySlice[S ~[]E, E any](s S) []any {
	anys := make([]any, len(s))
	for i, e := range s {
		anys[i] = e
	}

	return anys
}
