package escpos

import "errors"

// Init is used in initialising or resetting the printer.
type Init interface {
	// InitCommand should return the printer-specific command to
	// initialise the printer.
	InitCommand() (string, error)
}

// Cut allows for cutting the printer paper.
type Cut interface {
	// CutCommand should return the printer-specific command to cut the
	// printer paper.
	CutCommand() (string, error)
}

type End interface {
	EndCommand() (string, error)
}

// Font allows for different font support when printing.
type Font interface {
	// FontCommand should return the printer-specific command to set the
	// font based on the value found in the FormatConfig.
	FontCommand(*FormatConfig) (string, error)
}

// Justification allows for different text justification when printing.
type Justification interface {
	// JustificationCommand should return the printer-specific command to
	// set the justification based on the value found in the FormatConfig.
	JustificationCommand(*FormatConfig) (string, error)
}

// Emphasis allows for the emphasizing of text when printing.
type Emphasis interface {
	// EmphasisCommand should return the printer-specific command to set
	// the emphasis based on the value found in the FormatConfig.
	EmphasisCommand(*FormatConfig) (string, error)
}

// Profile represents a printer profile, surfacing commands for that specific
// printer. A profile should map from the generic FormatConfig to the specific
// commands for a particular printer.
type Profile interface {
	Init
	Cut
	End
	Font
	Justification
	Emphasis
}

type EpsonTMT20III struct {
}

func (EpsonTMT20III) InitCommand() (string, error) {
	return "\x1B@", nil
}

func (EpsonTMT20III) CutCommand() (string, error) {
	return "\x1DVA0", nil
}

func (EpsonTMT20III) EndCommand() (string, error) {
	return "\xFA", nil
}

func (EpsonTMT20III) FontCommand(fmtCfg *FormatConfig) (string, error) {
	switch fmtCfg.font {
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

func (EpsonTMT20III) JustificationCommand(fmtCfg *FormatConfig) (string, error) {
	switch fmtCfg.justification {
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

func (EpsonTMT20III) EmphasisCommand(fmtCfg *FormatConfig) (string, error) {
	if fmtCfg.emphasis {
		return "\x1BE1", nil
	} else {
		return "\x1BE0", nil
	}
}
