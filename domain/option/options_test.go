package option

import "testing"

func TestParse(t *testing.T) {
	const (
		TaskA TaskName = "a"
		TaskB TaskName = "b"
		TaskC TaskName = "c"
		TaskD TaskName = "d"
	)

	const (
		PageA PageName = "a"
		PageB PageName = "b"
		PageC PageName = "c"
		PageD PageName = "d"
		PageE PageName = "e"
	)

	options := NewOptions(OptionsDto{
		Tasks: map[string]TaskName{
			"An": TaskA,
			"B":  TaskB,
			"Cn": TaskC,
			"D":  TaskD,
		},
		Pages: map[string]PageName{
			"Cn": PageC,
			"D":  PageD,
			"E":  PageE,
		},
	}, []string{"arg1", "arg2", "arg3", "arg4", "arg5", "arg6", "arg7", "arg8", "arg9", "arg10", "arg11"})

	testCases := []struct {
		name         string
		input        string
		expectedTask TaskName
		expectedPage PageName
		expectedNum  int
	}{
		{
			name:         "with_valid_numless_task_passes",
			input:        "b",
			expectedTask: TaskB,
			expectedPage: "",
			expectedNum:  0,
		},
		{
			name:         "with_valid_numful_task_passes",
			input:        "a1",
			expectedTask: TaskA,
			expectedPage: "",
			expectedNum:  1,
		},
		{
			name:         "with_multiple_digits_passes",
			input:        "a11",
			expectedTask: TaskA,
			expectedPage: "",
			expectedNum:  11,
		},
		{
			name:         "with_invalid_option_fails",
			input:        "z",
			expectedTask: "",
			expectedPage: "",
			expectedNum:  0,
		},
		{
			name:         "with_invalid_num_fails",
			input:        "arwehdq",
			expectedTask: "",
			expectedPage: "",
			expectedNum:  0,
		},
		{
			name:         "with_no_num_fails",
			input:        "a",
			expectedTask: "",
			expectedPage: "",
			expectedNum:  0,
		},
		{
			name:         "with_out_of_bounds_num_fails",
			input:        "a12",
			expectedTask: "",
			expectedPage: "",
			expectedNum:  -1,
		},
		{
			name:         "with_negative_out_of_bounds_num_fails",
			input:        "a-12",
			expectedTask: "",
			expectedPage: "",
			expectedNum:  -1,
		},
		{
			name:         "with_zero_num_fails",
			input:        "a0",
			expectedTask: "",
			expectedPage: "",
			expectedNum:  -1,
		},
		{
			name:         "with_num_on_numless_task_passes",
			input:        "b10",
			expectedTask: TaskB,
			expectedPage: "",
			expectedNum:  0,
		},
		{
			name:         "with_num_on_numless_page_passes",
			input:        "e2",
			expectedTask: "",
			expectedPage: PageE,
			expectedNum:  0,
		},
		{
			name:         "with_valid_page_passes",
			input:        "e",
			expectedTask: "",
			expectedPage: PageE,
			expectedNum:  0,
		},
		{
			name:         "with_valid_page_and_task_passes",
			input:        "d",
			expectedTask: TaskD,
			expectedPage: PageD,
			expectedNum:  0,
		},
		{
			name:         "with_valid_page_and_task_and_num_passes",
			input:        "c3",
			expectedTask: TaskC,
			expectedPage: PageC,
			expectedNum:  3,
		},
		{
			name:         "with_valid_page_and_task_and_multiple_digits_passes",
			input:        "c11",
			expectedTask: TaskC,
			expectedPage: PageC,
			expectedNum:  11,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			task, page, num := options.parse(tc.input)

			if task != tc.expectedTask {
				t.Errorf("Expected task %s, got %s", tc.expectedTask, task)
			}

			if page != tc.expectedPage {
				t.Errorf("Expected page %s, got %s", tc.expectedPage, page)
			}

			if num != tc.expectedNum {
				t.Errorf("Expected num %d, got %d", tc.expectedNum, num)
			}
		})
	}
}
