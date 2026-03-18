package gopiano

import "testing"

func TestEncrypt(t *testing.T) {
	t.Parallel()

	t.Run("encrypts plaintext to expected ciphertext", func(t *testing.T) {
		t.Parallel()

		client, err := NewClient(AndroidClient)
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		input := "foobar"
		expected := "3c739d4e29b5d6c6"
		got := client.encrypt(input)
		if got != expected {
			t.Errorf("encrypt(%q) = %q, want %q", input, got, expected)
		}
	})
}

func TestDecrypt(t *testing.T) {
	t.Parallel()

	t.Run("decrypts ciphertext to expected plaintext", func(t *testing.T) {
		t.Parallel()

		client, err := NewClient(AndroidClient)
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		input := "95b6027f2d427dc0"
		expected := "foobar"
		got, err := client.decrypt(input)
		if err != nil {
			t.Fatalf("decrypt(%q) returned unexpected error: %v", input, err)
		}
		if got != expected {
			t.Errorf("decrypt(%q) = %q, want %q", input, got, expected)
		}
	})
}
