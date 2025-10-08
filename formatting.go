package escpos

import "fmt"

type FormatConfig struct {
	justify   justify
	emphasize emphasize
	font      font
}

func (fmtCfg FormatConfig) apply(client *Client) {
	applyCommand(fmtCfg.justify, client)
	applyCommand(fmtCfg.emphasize, client)
	applyCommand(fmtCfg.font, client)
}

func (fmtCfg FormatConfig) applyRefactored(client *Client, profile Profile) {
	command, err := profile.FontCommand(&fmtCfg)
	if err != nil {
		fmt.Printf("invalid command: %v\n", err)
	}
	client.writeString(command)

	command, err = profile.JustificationCommand(&fmtCfg)
	if err != nil {
		fmt.Printf("invalid command: %v\n", err)
	}
	client.writeString(command)

	command, err = profile.EmphasisCommand(&fmtCfg)
	if err != nil {
		fmt.Printf("invalid command: %v\n", err)
	}
	client.writeString(command)
}

func applyCommand(option toCommander, client *Client) {
	command, err := option.toCommand()
	if err != nil {
		fmt.Printf("invalid command: %v\n", err)
	}
	client.writeString(command)
}

func DefaultFormatConfig() FormatConfig {
	return FormatConfig{
		justify:   justify{"left"},
		emphasize: emphasize{false},
		font:      font{"A"},
	}
}

// Justification setters
func (fmtCfg FormatConfig) JustifyLeft() FormatConfig {
	fmtCfg.justify.justification = "left"
	return fmtCfg
}

func (fmtCfg FormatConfig) JustifyRight() FormatConfig {
	fmtCfg.justify.justification = "right"
	return fmtCfg
}

func (fmtCfg FormatConfig) JustifyCenter() FormatConfig {
	fmtCfg.justify.justification = "center"
	return fmtCfg
}

// Emphasize setters
func (fmtCfg FormatConfig) Emphasize(enabled bool) FormatConfig {
	fmtCfg.emphasize.enabled = enabled
	return fmtCfg
}

// Font setters
func (fmtCfg FormatConfig) FontA() FormatConfig {
	fmtCfg.font.selection = "A"
	return fmtCfg
}

func (fmtCfg FormatConfig) FontB() FormatConfig {
	fmtCfg.font.selection = "B"
	return fmtCfg
}

func (fmtCfg FormatConfig) FontC() FormatConfig {
	fmtCfg.font.selection = "C"
	return fmtCfg
}

func (fmtCfg FormatConfig) FontD() FormatConfig {
	fmtCfg.font.selection = "D"
	return fmtCfg
}
