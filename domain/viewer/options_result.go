package viewer

type OptionsResult struct {
	Task    Task
	Page    Page
	NextArg any
}

func NewOptionsResult(option Option, num int) (OptionsResult, error) {
	nextArg, err := option.Arg(num)
	if err != nil {
		return OptionsResult{}, err
	}

	if option.DoesTask() {
		return OptionsResult{
			Task:    option.Task(),
			NextArg: nextArg,
		}, nil
	}

	return OptionsResult{
		Page:    option.Page(),
		NextArg: nextArg,
	}, nil
}
