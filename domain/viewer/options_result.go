package viewer

type OptionsResult struct {
	nextTask Task
	nextPage Page
	nextArgs any
}

func NewOptionsResult(option Option, num int) (OptionsResult, error) {
	nextArgs, err := option.Arg(num)
	if err != nil {
		return OptionsResult{}, err
	}

	if option.DoesTask() {
		return OptionsResult{
			nextTask: option.Task(),
			nextArgs: nextArgs,
		}, nil
	}

	return OptionsResult{
		nextPage: option.Page(),
		nextArgs: nextArgs,
	}, nil
}

func (result OptionsResult) HasNextTask() bool {
	return result.nextTask != ""
}

func (result OptionsResult) NextTask() Task {
	return result.nextTask
}

func (result OptionsResult) NextPage() Page {
	return result.nextPage
}

func (result OptionsResult) NextArgs() any {
	return result.nextArgs
}
