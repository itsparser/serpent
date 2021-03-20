package exception

// SerpentError will be the Base Custom Error Object structure
// maintained in serpent
type SerpentError struct {
	code string
	s    string
	args struct{}
}

// Error will be member for the Serpent Error will provide the error message for the Serpent error
func (e *SerpentError) Error() string {
	return e.s
}

// UnableToLocateRootDirectory - This error will be used when unable to locate root direct
func UnableToLocateRootDirectory() error {
	msg := "Unable to Locate the Root Drirectory of the Machine"
	return &SerpentError{code: "ER0100001", s: msg}
}

// FailedDriverInstanceCreation - This error will be used when unable to create the drive instance
func FailedDriverInstanceCreation(e error) error {
	msg := "Failed to Create Driver Instance "
	return &SerpentError{code: "ER0100002", s: msg}
}

// UnableToLocateDriver - Unable to locate the Driver cache in Machine
func UnableToLocateDriver() error {
	msg := "Unable to locate the Driver cache in Machine"
	return &SerpentError{code: "ER0100003", s: msg}
}
