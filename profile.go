package escpos

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

// Underline allows for the underlining of text, usually with multiple
// modes to set thickness.
type Underline interface {
	// UnderlineCommand should return the printer-specific command to set
	// the underline mode based on the value found in the FormatConfig.
	UnderlineCommand(*FormatConfig) (string, error)
}

// CharSize allows for the scaling of character size.
type CharSize interface {
	// CharSizeCommand should return the printer-specific command to set
	// the character size scaling based on the values in the FormatConfig.
	CharSizeCommand(*FormatConfig) (string, error)
}

// SelectQrCodeModel selects the QR code model.
type SelectQrCodeModel interface {
	// SelectQrCodeModelCommand should return the printer-specific command
	// to set the QR code model based on the given QrCodeConfig.
	SelectQrCodeModelCommand(*QrCodeConfig) (string, error)
}

// SetQrCodeSize sets the QR code size.
type SetQrCodeSize interface {
	// SetQrCodeSizeCommand should return the printer-specific command
	// to set the QR code size based on the given QrCodeConfig.
	SetQrCodeSizeCommand(*QrCodeConfig) (string, error)
}

// SelectQrCodeErrorCorrectionLevel selects the QR code error
// correction level.
type SelectQrCodeErrorCorrectionLevel interface {
	// SelectQrCodeErrorCorrectionLevelCommand should return the
	// printer-specific command to set the QR error correction
	// level based on the given QrCodeConfig.
	SelectQrCodeErrorCorrectionLevelCommand(*QrCodeConfig) (string, error)
}

// StoreQrCodeData stores the QR code data in the symbol storage area.
type StoreQrCodeData interface {
	// StoreQrCodeDataCommand should return the printer-specific command
	// to store the QR code data in the symbol storage area.
	StoreQrCodeDataCommand(string) (string, error)
}

// PrintQrCodeData prints the QR code.
type PrintQrCodeData interface {
	// PrintQrCodeDataCommand should return the printer-specific command
	// to print the QR code.
	PrintQrCodeDataCommand() (string, error)
}

// Beep allows for the sounding of the buzzer.
type Beep interface {
	// BeepCommand should return the printer-specific command for sounding
	// the buzzer based on the given map of parameters. This command does
	// vary for printer models hence the usage of a parameter map to take
	// in values which control sound pattern, repetitions and more.
	BeepCommand(map[string]uint8) (string, error)
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
	Underline
	CharSize
	SelectQrCodeModel
	SetQrCodeSize
	SelectQrCodeErrorCorrectionLevel
	StoreQrCodeData
	PrintQrCodeData
	Beep
}
