package core

// EmptyResponse is an empty struct. The main reason behind this
// dummy response is that there are endpoints where the return
// body is empty, but they have 2xx response code, therefore
// it was successful.
type EmptyResponse struct{}
