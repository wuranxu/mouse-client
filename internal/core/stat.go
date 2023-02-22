package core

type RequestStat struct {
	Elapsed    int64
	FailedNum  int64
	SuccessNum int64
	success    chan *TestResult
	failure    chan *TestResult
}

type TestResult struct {
	// result
	Result    bool
	Exception string
}

func Success() *TestResult {
	return &TestResult{
		Result: true,
	}
}

func Failed(err error) *TestResult {
	return &TestResult{
		Result:    false,
		Exception: err.Error(),
	}
}
