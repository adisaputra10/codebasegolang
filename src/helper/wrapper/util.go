package wrapper

// FindOneErrorResult will return one error of variadic, if the result is nil it means no error
func FindOneErrorResult(results ...*Result) (errorResult *Result) {
	for _, result := range results {
		if result.Err != nil {
			errorResult = result
			return
		}
	}
	return
}
