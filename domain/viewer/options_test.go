package viewer

import (
	"errors"
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	const (
		TaskA Task = "a"
		TaskB Task = "b"
		TaskC Task = "c"
		TaskD Task = "d"
	)

	const (
		PageA Page = "a"
		PageB Page = "b"
		PageC Page = "c"
		PageD Page = "d"
	)

	// A is task w/o num
	// B is task w/ num
	// C is page w/o num
	// D is page w/ num

	optionA := NewOption(OptionDto{
		Letter:   'a',
		Task:     TaskA,
		TakesNum: false,
	}, []string{})

	optionB := NewOption(OptionDto{
		Letter:   'b',
		Task:     TaskB,
		TakesNum: true,
	}, []string{"arg1", "arg2", "arg3"})

	optionC := NewOption(OptionDto{
		Letter:   'c',
		Page:     PageC,
		TakesNum: false,
	}, []string{})

	optionD := NewOption(OptionDto{
		Letter:   'd',
		Page:     PageD,
		TakesNum: true,
	}, []string{"arg1", "arg2", "arg3"})

	options := NewOptions([]Option{optionA, optionB, optionC, optionD})

	testCases := []struct {
		name           string
		choice         string
		expectedOption Option
		expectedNum    int
		expectedError  error
	}{
		{
			name:           "with_valid_numless_task_passes",
			choice:         "a",
			expectedOption: optionA,
			expectedNum:    -1,
			expectedError:  nil,
		},
		{
			name:           "with_valid_numful_task_passes",
			choice:         "b1",
			expectedOption: optionB,
			expectedNum:    1,
			expectedError:  nil,
		},
		{
			name:           "with_valid_numless_page_passes",
			choice:         "c",
			expectedOption: optionC,
			expectedNum:    -1,
			expectedError:  nil,
		},
		{
			name:           "with_valid_numful_page_passes",
			choice:         "d1",
			expectedOption: optionD,
			expectedNum:    1,
			expectedError:  nil,
		},
		{
			name:           "with_multiple_digits_passes",
			choice:         "b11",
			expectedOption: optionB,
			expectedNum:    11,
			expectedError:  nil,
		},
		{
			name:           "with_invalid_option_fails",
			choice:         "z",
			expectedOption: Option{},
			expectedNum:    -1,
			expectedError:  fmt.Errorf("invalid option"),
		},
		{
			name:           "with_invalid_num_fails",
			choice:         "brwehdq",
			expectedOption: Option{},
			expectedNum:    -1,
			expectedError:  fmt.Errorf("not a number"),
		},
		{
			name:           "with_no_num_fails",
			choice:         "b",
			expectedOption: Option{},
			expectedNum:    -1,
			expectedError:  fmt.Errorf("no number detected"),
		},
		{
			name:           "with_num_on_numless_option_passes",
			choice:         "a10",
			expectedOption: optionA,
			expectedNum:    -1,
			expectedError:  nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			option, num, err := options.parse(tc.choice)

			if errors.Unwrap(err) != errors.Unwrap(tc.expectedError) {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}

			if option.Letter() != tc.expectedOption.Letter() {
				t.Errorf("Expected option %c, got %c", tc.expectedOption.Letter(), option.Letter())
			}

			if num != tc.expectedNum {
				t.Errorf("Expected num %d, got %d", tc.expectedNum, num)
			}
		})
	}
}
