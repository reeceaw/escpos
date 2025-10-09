package escpos

import "errors"

// EpsonTMT20III implements the ESC/POS commands specific to the Epson
// TM-T20III printer.
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
