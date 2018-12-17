package ripo

const (
	// MissingArgument: added by Saeed Rasooli
	// MissingArgument means that an input parameter is required but not given
	// either missing in the input/request, or has empty value
	MissingArgument Code = 17

	// ResourceLocked: added by Saeed Rasooli
	// ResourceLocked means that the give resource is currently busy or temporarily locked
	// by abother request (either by the same user or another user)
	ResourceLocked Code = 18
)
