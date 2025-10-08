package escpos

import "io"

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
	client.writer.Write(data)
}

func (client *Client) writeString(s string) {
	client.writeRaw([]byte(s))
}

// Init clears the data in the print buffer and resets the printer
// modes to the modes that were in effect when the power was turned on.
func (client *Client) Init() {
	client.writeString("\x1B@")
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
// \x1DV is GS V, the select mode and cut command. The A0 is used to
// specify the mode.
func (client *Client) Cut() {
	client.writeString("\x1DVA0")
}

func (client *Client) End() {
	client.writeString("\xFA")
}
