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
}
