package repoerrors

import "errors"

var TaskNotFound = errors.New("task not found")
var TaskAlreadyDone = errors.New("task is alreay done")
