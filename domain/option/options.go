package option

import (
	"fmt"
	"strconv"
	"unicode"
)

type OptionsResults struct {
	CmdName CmdName
	Num     int
}

type OptionsArgs struct {
	Options map[string]CmdName
}

type Options struct {
	options         map[rune]CmdName
	optionsWithNums map[rune]CmdName
}

func NewOptions(args OptionsArgs) Options {
	options := Options{
		options:         make(map[rune]CmdName),
		optionsWithNums: make(map[rune]CmdName),
	}

	for letter, cmd := range args.Options {
		if len(letter) == 1 {
			options.options[rune(letter[0])] = cmd
			continue
		}

		options.optionsWithNums[rune(letter[0])] = cmd
	}

	return options
}

func (ops Options) TakeInput() OptionsResults {
	var choice string
	var cmdName CmdName
	var num int

	for {
		fmt.Print(">")
		fmt.Scanf("%s", &choice)

		cmdName, num = ops.parse(choice)
		if cmdName != "" {
			return OptionsResults{
				CmdName: cmdName,
				Num:     num,
			}
		}

		fmt.Printf("The hell are you saying? Use one of the available options.\n")
	}
}

func (ops Options) parse(choice string) (CmdName, int) {
	if len(choice) == 0 {
		return "", 0
	}

	chosenLetter := unicode.ToUpper(rune(choice[0]))

	if cmdName, exists := ops.options[chosenLetter]; exists {
		return cmdName, 0
	}

	if cmdName, exists := ops.optionsWithNums[chosenLetter]; exists {
		if len(choice) == 1 {
			return "", 0
		}

		num, err := strconv.Atoi(choice[1:])
		if err != nil {
			return "", 0
		}

		return cmdName, num
	}

	return "", 0
}
