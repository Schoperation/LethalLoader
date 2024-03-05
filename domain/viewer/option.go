package viewer

import (
	"fmt"
	"unicode"
)

type OptionDto struct {
	Letter   rune
	Task     Task
	Page     Page
	TakesNum bool
}

type Option struct {
	letter       rune
	task         Task
	page         Page
	takesNum     bool
	possibleArgs []any
}

func NewOption[E any](dto OptionDto, possibleArgs []E) Option {
	possibleArgsAsAnys := toAnySlice(possibleArgs)
	return Option{
		letter:       unicode.ToUpper(dto.Letter),
		task:         dto.Task,
		page:         dto.Page,
		takesNum:     dto.TakesNum,
		possibleArgs: possibleArgsAsAnys,
	}
}

func (op Option) Letter() rune {
	return op.letter
}

func (op Option) Task() Task {
	return op.task
}

func (op Option) Page() Page {
	return op.page
}

func (op Option) TakesNum() bool {
	return op.takesNum
}

func (op Option) DoesTask() bool {
	return op.task != ""
}

func (op Option) Arg(i int) (any, error) {
	if !op.takesNum {
		return nil, nil
	}

	if i > len(op.possibleArgs) || i <= 0 {
		return nil, fmt.Errorf("arg choice must be between 1 and %d", len(op.possibleArgs))
	}

	return op.possibleArgs[i-1], nil
}

func toAnySlice[S ~[]E, E any](s S) []any {
	anys := make([]any, len(s))
	for i, e := range s {
		anys[i] = e
	}

	return anys
}
