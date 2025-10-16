package escpos

type QrCodeConfig struct {
	model           string
	size            uint
	errorCorrection string
	justification   string
}

// DefaultQrCodeConfig creates a QrCodeConfig containing sensible
// default values for QR code printing.
func DefaultQrCodeConfig() QrCodeConfig {
	return QrCodeConfig{
		model:           "2",
		size:            3,
		errorCorrection: "L",
		justification:   "center",
	}
}

// Model sets the QR code model. The default is model 2.
func (cfg QrCodeConfig) Model(model string) QrCodeConfig {
	cfg.model = model
	return cfg
}

// Size sets the QR code size. The default is 3.
func (cfg QrCodeConfig) Size(size uint) QrCodeConfig {
	cfg.size = size
	return cfg
}

// Justify sets the QR code justification. The default is center.
func (cfg QrCodeConfig) Justify(justification string) QrCodeConfig {
	cfg.justification = justification
	return cfg
}

// ErrorCorrection sets the error correction level. The default is
// level L.
func (cfg QrCodeConfig) ErrorCorrection(level string) QrCodeConfig {
	cfg.errorCorrection = level
	return cfg
}
