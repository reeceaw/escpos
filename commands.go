package escpos

import "errors"

type toCommander interface {
	toCommand() (string, error)
}

// justify is used to track the desired text justification.
type justify struct {
	justification string
}

func (j justify) toCommand() (string, error) {
	switch j.justification {
	case "left":
		return "\x1Ba0", nil
	case "center":
		return "\x1Ba1", nil
	case "right":
		return "\x1Ba2", nil
	default:
		return "", errors.New("invalid justification")
	}
}

// emphasize is used to track whether text should be emphasized.
type emphasize struct {
	enabled bool
}

func (e emphasize) toCommand() (string, error) {
	if e.enabled {
		return "\x1BE1", nil
	} else {
		return "\x1BE0", nil
	}
}

// font is used to track the font selection
type font struct {
	selection string
}

func (f font) toCommand() (string, error) {
	switch f.selection {
	case "A":
		return "\x1BM0", nil
	case "B":
		return "\x1BM1", nil
	case "C":
		return "\x1BM2", nil
	case "D":
		return "\x1BM3", nil
	default:
		return "", errors.New("invalid font selection")
	}
}
