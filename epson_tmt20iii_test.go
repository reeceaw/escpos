package escpos

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEpsonTMT20III_SimpleCommands(t *testing.T) {
	var profile Profile = EpsonTMT20III{}

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
	var profile Profile = EpsonTMT20III{}

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

	t.Run("font with unknown value returns error", func(t *testing.T) {
		got, err := profile.FontCommand(&FormatConfig{font: "blabla"})

		if got != "" || err == nil {
			t.Errorf("command was not nil, expected empty string and error")
		}

		if err.Error() != "invalid font option in FormatConfig" {
			t.Errorf("FontCommand did not return expected error, got %s", err.Error())
		}
	})
}

func TestEpsonTMT20III_JustificationCommand(t *testing.T) {
	var profile Profile = EpsonTMT20III{}

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

	t.Run("justification with unknown value returns error", func(t *testing.T) {
		got, err := profile.JustificationCommand(&FormatConfig{justification: "blabla"})

		if got != "" || err == nil {
			t.Errorf("command was not nil, expected empty string and error")
		}

		if err.Error() != "invalid justification option in FormatConfig" {
			t.Errorf("JustificationCommand did not return expected error, got %s", err.Error())
		}
	})
}

func TestEpsonTMT20III_EmphasisCommand(t *testing.T) {
	var profile Profile = EpsonTMT20III{}

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

func TestEpsonTMT20III_UnderlineCommand(t *testing.T) {
	var profile Profile = EpsonTMT20III{}

	cases := []struct {
		name      string
		underline string
		want      []byte
	}{
		{"underline off returns correct value", "off", []byte{'\x1B', '-', '0'}},
		{"underline 1-dot returns correct value", "1-dot", []byte{'\x1B', '-', '1'}},
		{"underline 2-dots returns correct value", "2-dots", []byte{'\x1B', '-', '2'}},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := profile.UnderlineCommand(&FormatConfig{underline: testCase.underline})

			if err != nil {
				t.Errorf("err was not nil")
			}

			gotAsBytes := []byte(got)

			if !bytes.Equal(gotAsBytes, testCase.want) {
				t.Errorf("UnderlineCommand did not return expected bytes: wanted %v, got %v", testCase.want, gotAsBytes)
			}
		})
	}

	t.Run("underline with unknown value returns error", func(t *testing.T) {
		got, err := profile.UnderlineCommand(&FormatConfig{underline: "blabla"})

		if got != "" || err == nil {
			t.Errorf("returned command was not nil, expected empty string and error")
		}

		if err.Error() != "invalid underline option in FormatConfig" {
			t.Errorf("UnderlineCommand did not return expected error, got %s", err.Error())
		}
	})
}

func TestEpsonTMT20III_CharSizeCommand(t *testing.T) {
	var profile Profile = EpsonTMT20III{}

	cases := []struct {
		name   string
		width  uint8
		height uint8
		want   []byte
	}{
		{"charsize with width 1, height 1 returns correct value", 1, 1, []byte{'\x1D', '!', '\x00'}},
		{"charsize with width 1, height 2 returns correct value", 1, 2, []byte{'\x1D', '!', '\x01'}},
		{"charsize with width 2, height 1 returns correct value", 2, 1, []byte{'\x1D', '!', '\x10'}},
		{"charsize with width 8, height 8 returns correct value", 8, 8, []byte{'\x1D', '!', '\x77'}},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := profile.CharSizeCommand(&FormatConfig{charWidth: testCase.width, charHeight: testCase.height})

			if err != nil {
				t.Errorf("err was not nil")
			}

			gotAsBytes := []byte(got)

			if !bytes.Equal(gotAsBytes, testCase.want) {
				t.Errorf("CharSizeCommand did not return expected bytes: wanted %v, got %v", testCase.want, gotAsBytes)
			}
		})
	}

	negativeCases := []struct {
		name   string
		width  uint8
		height uint8
		want   []byte
	}{
		{"charsize with width 0, height 0 returns error", 0, 0, []byte{'\x1D', '!', '0'}},
		{"charsize with width 0, height 1 returns error", 0, 1, []byte{'\x1D', '!', '0'}},
		{"charsize with width 1, height 0 returns error", 1, 0, []byte{'\x1D', '!', '0'}},
		{"charsize with width 9, height 8 returns error", 9, 8, []byte{'\x1D', '!', '0'}},
		{"charsize with width 8, height 9 returns error", 8, 9, []byte{'\x1D', '!', '0'}},
		{"charsize with width 9, height 9 returns error", 9, 9, []byte{'\x1D', '!', '0'}},
	}

	for _, testCase := range negativeCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := profile.CharSizeCommand(&FormatConfig{charWidth: testCase.width, charHeight: testCase.height})

			if got != "" || err == nil {
				t.Errorf("returned command was not nil, expected empty string and error")
			}

			expectedError := fmt.Sprintf("invalid charsize options in FormatConfig: width %v, height %v\n", testCase.width, testCase.height)

			if err.Error() != expectedError {
				t.Errorf("CharSizeCommand did not return expected error, got %s, wanted %s", err.Error(), expectedError)
			}
		})
	}
}
