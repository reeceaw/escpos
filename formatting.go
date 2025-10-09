package escpos

import "fmt"

type FormatConfig struct {
	justification string
	emphasis      bool
	font          string
	underline     string
}

func (fmtCfg FormatConfig) apply(client *Client, profile Profile) {
	applyCommand(
		func() (string, error) { return profile.FontCommand(&fmtCfg) }, client)

	applyCommand(
		func() (string, error) { return profile.JustificationCommand(&fmtCfg) }, client)

	applyCommand(
		func() (string, error) { return profile.EmphasisCommand(&fmtCfg) }, client)

	applyCommand(
		func() (string, error) { return profile.UnderlineCommand(&fmtCfg) }, client)
}

func applyCommand(getCommand func() (string, error), client *Client) {
	command, err := getCommand()
	if err != nil {
		fmt.Printf("invalid command: %v\n", err)
	}
	client.writeString(command)
}

func DefaultFormatConfig() FormatConfig {
	return FormatConfig{
		justification: "left",
		emphasis:      false,
		font:          "A",
		underline:     "off",
	}
}

// Justify justifies the text based on the given value. Supported values
// usually include 'left', 'center' and 'right'.
func (fmtCfg FormatConfig) Justify(justification string) FormatConfig {
	fmtCfg.justification = justification
	return fmtCfg
}

// Emphasize sets the text emphasis to the given bool value.
func (fmtCfg FormatConfig) Emphasize(enabled bool) FormatConfig {
	fmtCfg.emphasis = enabled
	return fmtCfg
}

// Font sets the font to the specific value. Supported values usually
// include A, B, C, D and E but this may differ per printer.
func (fmtCfg FormatConfig) Font(font string) FormatConfig {
	fmtCfg.font = font
	return fmtCfg
}

func (fmtCfg FormatConfig) Underline(mode string) FormatConfig {
	fmtCfg.underline = mode
	return fmtCfg
}
