package viewer

import (
	"fmt"
	"strconv"
	"unicode"
)

type Options struct {
	ops map[rune]Option
}

func NewOptions(options []Option) Options {
	opsCollection := make(map[rune]Option, len(options))

	for _, option := range options {
		opsCollection[option.Letter()] = option
	}

	return Options{
		ops: opsCollection,
	}
}

func (ops Options) TakeInput() OptionsResult {
	for {
		var choice string
		fmt.Print(">")
		fmt.Scanf("%s", &choice)

		option, num, err := ops.parse(choice)
		if err != nil {
			fmt.Printf("Bruh we got a problem: %v\n", err)
			continue
		}

		result, err := NewOptionsResult(option, num)
		if err != nil {
			fmt.Printf("Bruh we got a problem: %v\n", err)
			continue
		}

		return result
	}
}

func (ops Options) parse(choice string) (Option, int, error) {
	if len(choice) == 0 {
		return Option{}, -1, fmt.Errorf("no choice detected")
	}

	chosenOption := unicode.ToUpper(rune(choice[0]))

	option, exists := ops.ops[chosenOption]
	if !exists {
		return Option{}, -1, fmt.Errorf("invalid option")
	}

	if !option.TakesNum() {
		return option, -1, nil
	}

	if len(choice) == 1 {
		return Option{}, -1, fmt.Errorf("no number detected")
	}

	num, err := strconv.Atoi(choice[1:])
	if err != nil {
		return Option{}, -1, fmt.Errorf("not a number")
	}

	return option, num, nil
}
