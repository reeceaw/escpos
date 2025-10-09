package escpos

import (
	"bytes"
	"testing"
)

func TestNewClient(t *testing.T) {
	var writer bytes.Buffer
	NewClient(&writer, EpsonTMT20III{})

	got := writer.Bytes()
	want := []byte{'\x1B', '@'}

	if !bytes.Equal(got, want) {
		t.Errorf("NewClient did not perform init, buffer got %v, wanted %v", got, want)
	}
}

func TestClient_WriteLine(t *testing.T) {
	var writer bytes.Buffer
	client := NewClient(&writer, EpsonTMT20III{})
	writer.Reset()

	client.WriteLine("Hello!")

	got := writer.Bytes()
	want := []byte{'H', 'e', 'l', 'l', 'o', '!', '\n'}

	if !bytes.Equal(got, want) {
		t.Errorf("WriteLine did not write expected bytes, buffer got %s, wanted %s", got, want)
	}
}

func TestClient_Cut(t *testing.T) {
	var writer bytes.Buffer
	client := NewClient(&writer, EpsonTMT20III{})
	writer.Reset()

	client.Cut()

	got := writer.Bytes()
	want := []byte{'\x1D', 'V', 'A', '0'}

	if !bytes.Equal(got, want) {
		t.Errorf("Cut did not write expected bytes, buffer got %s, wanted %s", got, want)
	}
}

func TestClient_End(t *testing.T) {
	var writer bytes.Buffer
	client := NewClient(&writer, EpsonTMT20III{})
	writer.Reset()

	client.End()

	got := writer.Bytes()
	want := []byte{'\xFA'}

	if !bytes.Equal(got, want) {
		t.Errorf("End did not write expected bytes, buffer got %s, wanted %s", got, want)
	}
}
