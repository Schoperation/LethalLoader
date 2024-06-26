package task

import (
	"bufio"
	"fmt"
	"os"
	"schoperation/lethalloader/domain/input"
	"schoperation/lethalloader/domain/viewer"
	"strings"
)

type SearchTermTask struct {
}

func NewSearchTermTask() SearchTermTask {
	return SearchTermTask{}
}

func (task SearchTermTask) Do(args any) (viewer.TaskResult, error) {
	taskInput, ok := args.(input.SearchTermTaskInput)
	if !ok {
		return viewer.TaskResult{}, fmt.Errorf("could not parse profile")
	}

	fmt.Print("\n")
	fmt.Print("Enter search term (Q to cancel):\n")

	term := ""
	var err error
	reader := bufio.NewReader(os.Stdin)
	for {
		term, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf("The hell is this? %v\n", err)
			continue
		}

		if strings.TrimSpace(term) == "" {
			fmt.Printf("Bruh where's your query?\n")
			continue
		}

		break
	}

	term = strings.TrimSuffix(term, "\n")
	term = strings.TrimSuffix(term, "\r")

	if strings.ToLower(term) == "q" {
		return viewer.NewTaskResult(viewer.PageProfileViewer, taskInput.Profile), nil
	}

	return viewer.NewTaskResult(viewer.PageModSearchResults, input.ModSearchResultsPageInput{
		Profile:         taskInput.Profile,
		Term:            term,
		SkipCacheSearch: taskInput.SkipCacheSearch,
	}), nil
}
