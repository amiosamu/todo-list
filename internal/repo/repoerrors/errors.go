package repoerrors

import "errors"

var TaskNotFound = errors.New("task not found")
var TaskAlreadyDone = errors.New("task is already done")
var TaskTitleTooLong = errors.New("task title is too long. Max - 200 symbols")
