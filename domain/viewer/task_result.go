package viewer

type TaskResult struct {
	nextPage Page
	nextArgs any
}

func NewTaskResult(nextPage Page, nextArgs any) TaskResult {
	return TaskResult{
		nextPage: nextPage,
		nextArgs: nextArgs,
	}
}

func (result TaskResult) HasNextPage() bool {
	return result.nextPage != ""
}

func (result TaskResult) NextPage() Page {
	return result.nextPage
}

func (result TaskResult) NextArgs() any {
	return result.nextArgs
}
