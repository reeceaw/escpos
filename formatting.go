package escpos

import "fmt"

type FormatConfig struct {
	justification string
	emphasis      bool
	font          string
	underline     string
}

func (fmtCfg FormatConfig) apply(client *Client, profile Profile) {
	commands := []func(*FormatConfig) (string, error){
		profile.FontCommand,
		profile.JustificationCommand,
		profile.EmphasisCommand,
		profile.UnderlineCommand,
	}

	for _, command := range commands {
		value, err := command(&fmtCfg)
		if err != nil {
			fmt.Printf("failed building command: %v\n", err)
			return
		}
		client.writeString(value)
	}
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
// include A, B, C, D and E but this may differ based on printer profile.
func (fmtCfg FormatConfig) Font(font string) FormatConfig {
	fmtCfg.font = font
	return fmtCfg
}

// Underline sets the style of underline. This varies based on printer
// but usually supports off, 1-dot and 2-dots.
func (fmtCfg FormatConfig) Underline(mode string) FormatConfig {
	fmtCfg.underline = mode
	return fmtCfg
}
