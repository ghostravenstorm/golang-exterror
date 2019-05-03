/// Author:  GhostRavenstorm
/// Project: exterror
/// Date:    2019-05-01
/// Summary:
///   Simple and slightly more robust error handling extending from Go's builtin error interface.


package exterror

import "strconv"

type ExtError struct {
	StatusCode int     /// Http status code for REST APIs.
	ErrorCode  int     /// Application defined error codes.
	Message    string  /// Error message.
}

/// Interface for ExtError that implements builtin error interface.
type IExtError interface {
	GetStatusCode () int
	GetErrorCode  () int
	GetMessage    () string
	error
}

func (this ExtError) GetStatusCode () (int)    { return this.StatusCode }
func (this ExtError) GetErrorCode  () (int)    { return this.ErrorCode  }
func (this ExtError) GetMessage    () (string) { return this.Message    }

func (this ExtError) Error () (string) {
	return (`(` + strconv.Itoa(this.StatusCode) + `) ` + `(` + strconv.Itoa(this.ErrorCode) + `) ` + this.Message)
}

func New (scode, ecode int, msg string) (IExtError) {
	return &ExtError{
		StatusCode: scode,
		ErrorCode:  ecode,
		Message:    msg,
	}
}
