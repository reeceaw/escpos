package escpos

import (
	"errors"
	"fmt"
)

const (
	qrCodeSymbol byte = 49
)

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
		return "", errors.New("invalid font option in FormatConfig")
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
		return "", errors.New("invalid justification option in FormatConfig")
	}
}

func (EpsonTMT20III) EmphasisCommand(fmtCfg *FormatConfig) (string, error) {
	if fmtCfg.emphasis {
		return "\x1BE1", nil
	} else {
		return "\x1BE0", nil
	}
}

func (EpsonTMT20III) UnderlineCommand(fmtCfg *FormatConfig) (string, error) {
	switch fmtCfg.underline {
	case "off":
		return "\x1B-0", nil
	case "1-dot":
		return "\x1B-1", nil
	case "2-dots":
		return "\x1B-2", nil
	default:
		return "", errors.New("invalid underline option in FormatConfig")
	}
}

func (EpsonTMT20III) CharSizeCommand(fmtCfg *FormatConfig) (string, error) {
	if fmtCfg.charWidth < 1 || fmtCfg.charHeight < 1 || fmtCfg.charWidth > 8 || fmtCfg.charHeight > 8 {
		message := fmt.Sprintf("invalid charsize options in FormatConfig: width %v, height %v\n", fmtCfg.charWidth, fmtCfg.charHeight)
		return "", errors.New(message)
	}

	sizeByte := ((fmtCfg.charWidth - 1) << 4) | (fmtCfg.charHeight - 1)

	return fmt.Sprintf("\x1D!%c", sizeByte), nil
}

func (EpsonTMT20III) SelectQrCodeModelCommand(cfg *QrCodeConfig) (string, error) {
	switch cfg.model {
	case "1":
		return string([]byte{'\x1D', '(', 'k', 4, 0, qrCodeSymbol, 65, 49, 0}), nil
	case "2":
		return string([]byte{'\x1D', '(', 'k', 4, 0, qrCodeSymbol, 65, 50, 0}), nil
	default:
		return "", errors.New(fmt.Sprintf("invalid model in QrCodeConfig: %v\n", cfg.model))
	}
}

func (EpsonTMT20III) SetQrCodeSizeCommand(cfg *QrCodeConfig) (string, error) {
	if cfg.size < 1 || cfg.size > 16 {
		return "", errors.New(fmt.Sprintf("invalid size option in QrCodeConfig: %v\n", cfg.size))
	}

	return string([]byte{'\x1D', '(', 'k', 3, 0, qrCodeSymbol, 67, byte(cfg.size)}), nil
}

func (EpsonTMT20III) SelectQrCodeErrorCorrectionLevelCommand(cfg *QrCodeConfig) (string, error) {
	switch cfg.errorCorrection {
	case "L":
		return string([]byte{'\x1D', '(', 'k', 3, 0, qrCodeSymbol, 69, 48}), nil
	case "M":
		return string([]byte{'\x1D', '(', 'k', 3, 0, qrCodeSymbol, 69, 49}), nil
	case "Q":
		return string([]byte{'\x1D', '(', 'k', 3, 0, qrCodeSymbol, 69, 50}), nil
	case "H":
		return string([]byte{'\x1D', '(', 'k', 3, 0, qrCodeSymbol, 69, 51}), nil
	default:
		return "", errors.New(fmt.Sprintf("invalid error correction level option in QrCodeConfig: %v\n", cfg.errorCorrection))
	}
}

func (EpsonTMT20III) StoreQrCodeDataCommand(data string) (string, error) {
	dataLength := len(data)

	if dataLength > 7086 {
		return "", errors.New(fmt.Sprintf("maximum data length exceeded: %v > 7086 (max)\n", dataLength))
	}

	var pH byte = 0
	bytesAfterPh := byte(dataLength + 3)

	return string(append([]byte{'\x1D', '(', 'k', bytesAfterPh, pH, qrCodeSymbol, 80, 48}, data...)), nil
}

func (EpsonTMT20III) PrintQrCodeDataCommand() (string, error) {
	return string([]byte{'\x1D', '(', 'k', 3, 0, qrCodeSymbol, 81, 48}), nil
}
