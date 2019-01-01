package ripo

var ErrorCodeByName = map[string]Code{
	"Canceled":           Canceled,           // 1
	"Unknown":            Unknown,            // 2
	"InvalidArgument":    InvalidArgument,    // 3
	"DeadlineExceeded":   DeadlineExceeded,   // 4
	"NotFound":           NotFound,           // 5
	"AlreadyExists":      AlreadyExists,      // 6
	"PermissionDenied":   PermissionDenied,   // 7
	"Unauthenticated":    Unauthenticated,    // 16
	"ResourceExhausted":  ResourceExhausted,  // 8
	"FailedPrecondition": FailedPrecondition, // 9
	"Aborted":            Aborted,            // 10
	"OutOfRange":         OutOfRange,         // 11
	"Unimplemented":      Unimplemented,      // 12
	"Internal":           Internal,           // 13
	"Unavailable":        Unavailable,        // 14
	"DataLoss":           DataLoss,           // 15
	"MissingArgument":    MissingArgument,    // 17 (extra code)
	"ResourceLocked":     ResourceLocked,     // 18 (extra code)
}
