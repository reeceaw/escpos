package escpos

import (
	"fmt"
	"io"
)

// Client is an ESC/POS client that can be used to interact with an
// ESC/POS printer such as the Epson TM-T20II.
type Client struct {
	writer  io.Writer
	profile Profile
}

// NewClient creates an ESC/POS client which takes an io.Writer as
// the target to write ESC/POS commands to.
func NewClient(writer io.Writer, profile Profile) Client {
	client := Client{
		writer,
		profile,
	}
	client.Init()
	return client
}

func (client *Client) writeRaw(data []byte) {
	_, err := client.writer.Write(data)
	if err != nil {
		fmt.Printf("error writing data: %v", err)
	}
}

func (client *Client) writeString(s string) {
	client.writeRaw([]byte(s))
}

// Init clears the data in the print buffer and resets the printer
// modes to the modes that were in effect when the power was turned on.
func (client *Client) Init() {
	command, err := client.profile.InitCommand()
	if err != nil {
		fmt.Printf("error getting init command: %v\n", err)
		return
	}
	client.writeString(command)
}

// WriteLine writes the given string followed by a newline to the
// ESC/POS target.
func (client *Client) WriteLine(line string) {
	client.writeString(line)
	client.writeString("\n")
}

// Write configures the client using the given Configurer then writes
// the given string to the ESC/POS target.
func (client *Client) Write(s string, fmtCfg FormatConfig) {
	fmtCfg.apply(client, client.profile)
	client.writeString(s)
	DefaultFormatConfig().apply(client, client.profile)
}

// Cut writes a command which selects the cut mode and cuts the paper.
func (client *Client) Cut() {
	command, err := client.profile.CutCommand()
	if err != nil {
		fmt.Printf("error getting cut command: %v\n", err)
		return
	}
	client.writeString(command)
}

// End signifies the printing has completed and subsequent data is
// considered separate.
func (client *Client) End() {
	command, err := client.profile.EndCommand()
	if err != nil {
		fmt.Printf("error getting end command: %v\n", err)
		return
	}
	client.writeString(command)
}

// WriteQrCode writes the given data as a QR code to the printer,
// using the given QrCodeConfig for options such as size and model.
func (client *Client) WriteQrCode(data string, cfg QrCodeConfig) {
	DefaultFormatConfig().Justify(cfg.justification).apply(client, client.profile)

	c, err := client.profile.SelectQrCodeModelCommand(&cfg)
	if err != nil {
		fmt.Printf("error getting select QR code model command: %v\n", err)
	}
	client.writeString(c)

	c, err = client.profile.SetQrCodeSizeCommand(&cfg)
	if err != nil {
		fmt.Printf("error getting set QR code size command: %v\n", err)
	}
	client.writeString(c)

	c, err = client.profile.SelectQrCodeErrorCorrectionLevelCommand(&cfg)
	if err != nil {
		fmt.Printf("error getting select QR code error correction level command: %v\n", err)
	}
	client.writeString(c)

	c, err = client.profile.StoreQrCodeDataCommand(data)
	if err != nil {
		fmt.Printf("error getting store QR code data command: %v\n", err)
	}
	client.writeString(c)

	c, err = client.profile.PrintQrCodeDataCommand()
	if err != nil {
		fmt.Printf("error getting print QR code data command: %v\n", err)
	}
	client.writeString(c)

	DefaultFormatConfig().apply(client, client.profile)
}
