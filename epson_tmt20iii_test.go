package escpos

import (
	"bytes"
	"testing"
)

func TestEpsonTMT20III_SimpleCommands(t *testing.T) {
	profile := EpsonTMT20III{}

	cases := []struct {
		name        string
		commandFunc func() (string, error)
		want        []byte
	}{
		{"init command returns correct value", profile.InitCommand, []byte{'\x1B', '@'}},
		{"cut command returns correct value", profile.CutCommand, []byte{'\x1D', 'V', 'A', '0'}},
		{"end command returns correct value", profile.EndCommand, []byte{'\xFA'}},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := testCase.commandFunc()

			if err != nil {
				t.Errorf("err was not nil")
			}

			gotAsBytes := []byte(got)

			if !bytes.Equal(gotAsBytes, testCase.want) {
				t.Errorf("InitCommand did not return expected bytes: wanted %v, got %v", testCase.want, gotAsBytes)
			}
		})
	}
}

func TestEpsonTMT20III_FontCommand(t *testing.T) {
	profile := EpsonTMT20III{}

	cases := []struct {
		name string
		font string
		want []byte
	}{
		{"font A returns correct value", "A", []byte{'\x1B', 'M', '0'}},
		{"font B returns correct value", "B", []byte{'\x1B', 'M', '1'}},
		{"font C returns correct value", "C", []byte{'\x1B', 'M', '2'}},
		{"font D returns correct value", "D", []byte{'\x1B', 'M', '3'}},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := profile.FontCommand(&FormatConfig{font: testCase.font})

			if err != nil {
				t.Errorf("err was not nil")
			}

			gotAsBytes := []byte(got)

			if !bytes.Equal(gotAsBytes, testCase.want) {
				t.Errorf("FontCommand did not return expected bytes: wanted %v, got %v", testCase.want, gotAsBytes)
			}
		})
	}
}

func TestEpsonTMT20III_JustificationCommand(t *testing.T) {
	profile := EpsonTMT20III{}

	cases := []struct {
		name          string
		justification string
		want          []byte
	}{
		{"justify left returns correct value", "left", []byte{'\x1B', 'a', '0'}},
		{"justify center returns correct value", "center", []byte{'\x1B', 'a', '1'}},
		{"justify right returns correct value", "right", []byte{'\x1B', 'a', '2'}},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := profile.JustificationCommand(&FormatConfig{justification: testCase.justification})

			if err != nil {
				t.Errorf("err was not nil")
			}

			gotAsBytes := []byte(got)

			if !bytes.Equal(gotAsBytes, testCase.want) {
				t.Errorf("JustificationCommand did not return expected bytes: wanted %v, got %v", testCase.want, gotAsBytes)
			}
		})
	}
}

func TestEpsonTMT20III_EmphasisCommand(t *testing.T) {
	profile := EpsonTMT20III{}

	cases := []struct {
		name     string
		emphasis bool
		want     []byte
	}{
		{"emphasis false returns correct value", false, []byte{'\x1B', 'E', '0'}},
		{"emphasis true returns correct value", true, []byte{'\x1B', 'E', '1'}},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := profile.EmphasisCommand(&FormatConfig{emphasis: testCase.emphasis})

			if err != nil {
				t.Errorf("err was not nil")
			}

			gotAsBytes := []byte(got)

			if !bytes.Equal(gotAsBytes, testCase.want) {
				t.Errorf("EmphasisCommand did not return expected bytes: wanted %v, got %v", testCase.want, gotAsBytes)
			}
		})
	}
}
