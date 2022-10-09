package logger

type loggerState struct {
	owner string
	state bool
}

func formattedName(owner string) string {
	own := ""
	if len(owner) <= sourceFieldLength {
		own = owner
	} else if owner == "" {
		own = "Anonymous"
	} else {
		own = owner[:15]
	}

	return own
}
