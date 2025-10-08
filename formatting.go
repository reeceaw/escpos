package escpos

import "fmt"

type FormatConfig struct {
	justification string
	emphasis      bool
	font          string
}

func (fmtCfg FormatConfig) apply(client *Client, profile Profile) {
	applyCommand(
		func() (string, error) { return profile.FontCommand(&fmtCfg) }, client)

	applyCommand(
		func() (string, error) { return profile.JustificationCommand(&fmtCfg) }, client)

	applyCommand(
		func() (string, error) { return profile.EmphasisCommand(&fmtCfg) }, client)
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
	}
}

// Justification setters
func (fmtCfg FormatConfig) JustifyLeft() FormatConfig {
	fmtCfg.justification = "left"
	return fmtCfg
}

func (fmtCfg FormatConfig) JustifyRight() FormatConfig {
	fmtCfg.justification = "right"
	return fmtCfg
}

func (fmtCfg FormatConfig) JustifyCenter() FormatConfig {
	fmtCfg.justification = "center"
	return fmtCfg
}

// Emphasize setters
func (fmtCfg FormatConfig) Emphasize(enabled bool) FormatConfig {
	fmtCfg.emphasis = enabled
	return fmtCfg
}

// Font setters
func (fmtCfg FormatConfig) FontA() FormatConfig {
	fmtCfg.font = "A"
	return fmtCfg
}

func (fmtCfg FormatConfig) FontB() FormatConfig {
	fmtCfg.font = "B"
	return fmtCfg
}

func (fmtCfg FormatConfig) FontC() FormatConfig {
	fmtCfg.font = "C"
	return fmtCfg
}

func (fmtCfg FormatConfig) FontD() FormatConfig {
	fmtCfg.font = "D"
	return fmtCfg
}
