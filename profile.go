package escpos

import "errors"

type Font interface {
	FontCommand(*FormatConfig) (string, error)
}

type Justification interface {
	JustificationCommand(*FormatConfig) (string, error)
}

type Emphasis interface {
	EmphasisCommand(*FormatConfig) (string, error)
}

type Profile interface {
	Font
	Justification
	Emphasis
}

type EpsonTMT20III struct {
}

func (EpsonTMT20III) FontCommand(fmtCfg *FormatConfig) (string, error) {
	switch fmtCfg.font.selection {
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
	switch fmtCfg.justify.justification {
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
	if fmtCfg.emphasize.enabled {
		return "\x1BE1", nil
	} else {
		return "\x1BE0", nil
	}
}
