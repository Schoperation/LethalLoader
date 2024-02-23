package option

import "testing"

func TestParse(t *testing.T) {
	const (
		TaskA CmdName = "a_task"
		TaskB CmdName = "b_task"
		TaskC CmdName = "c_task"
	)

	options := NewOptions(OptionsArgs{
		Options: map[string]CmdName{
			"An": TaskA,
			"B":  TaskB,
			"Cn": TaskC,
		},
	})

	testCases := []struct {
		name         string
		input        string
		expectedTask CmdName
		expectedNum  int
	}{
		{
			name:         "with_valid_numless_task_passes",
			input:        "b",
			expectedTask: TaskB,
			expectedNum:  0,
		},
		{
			name:         "with_valid_numful_task_passes",
			input:        "a1",
			expectedTask: TaskA,
			expectedNum:  1,
		},
		{
			name:         "with_multiple_digits_passes",
			input:        "a77",
			expectedTask: TaskA,
			expectedNum:  77,
		},
		{
			name:         "with_invalid_task_fails",
			input:        "d",
			expectedTask: "",
			expectedNum:  0,
		},
		{
			name:         "with_invalid_num_fails",
			input:        "arwehdq",
			expectedTask: "",
			expectedNum:  0,
		},
		{
			name:         "with_no_num_fails",
			input:        "a",
			expectedTask: "",
			expectedNum:  0,
		},
		{
			name:         "with_num_on_numless_task_passes",
			input:        "b10",
			expectedTask: TaskB,
			expectedNum:  0,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			task, num := options.parse(tc.input)

			if task != tc.expectedTask {
				t.Errorf("Expected task %s, got %s", tc.expectedTask, task)
			}

			if num != tc.expectedNum {
				t.Errorf("Expected num %d, got %d", tc.expectedNum, num)
			}
		})
	}
}
