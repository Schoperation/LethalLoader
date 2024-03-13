package task

import (
	"bufio"
	"fmt"
	"os"
	"schoperation/lethalloader/domain/input"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
	"strings"
)

type SearchTermTask struct {
}

func NewSearchTermTask() SearchTermTask {
	return SearchTermTask{}
}

func (task SearchTermTask) Do(args any) (viewer.TaskResult, error) {
	pf, ok := args.(profile.Profile)
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

	if strings.ToLower(term) == "q" {
		return viewer.NewTaskResult(viewer.PageProfileViewer, pf), nil
	}

	return viewer.NewTaskResult(viewer.PageModSearchResults, input.ModSearchResultsPageInput{Profile: pf, Term: term}), nil
}
