package tgbot

import "errors"

var (
	ErrSendMessage       = errors.New("failed to send message")
	ErrDeleteMessage     = errors.New("failed to delete message")
	ErrEditMessage       = errors.New("failed to edit message")
	ErrOperationCanceled = errors.New("operation was cancel")
	ErrUnknown           = errors.New("error")
)
