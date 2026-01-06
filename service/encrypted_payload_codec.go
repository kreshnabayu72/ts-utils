package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/gogo/protobuf/proto"
	commonpb "go.temporal.io/api/common/v1"
)

const (
	// MetadataEncodingEncrypted is the metadata key for encrypted payloads
	MetadataEncodingEncrypted = "binary/encrypted"
	// MetadataEncryptionKeyID is the metadata key for encryption key identifier
	MetadataEncryptionKeyID = "encryption-key-id"
)

// EncryptionCodec implements payload encryption using AES-256-GCM.
// It encrypts workflow payloads before they are stored in Temporal,
// ensuring sensitive data is not visible in Temporal UI or database.
type EncryptionCodec struct {
	keyID         string
	encryptionKey []byte
}

// NewEncryptionCodec creates a new encryption codec with the provided key.
// The key must be exactly 32 bytes for AES-256 encryption.
// keyID is an identifier for the key (not secret, used for key rotation).
func NewEncryptionCodec(keyID string, encryptionKey []byte) (*EncryptionCodec, error) {
	if len(encryptionKey) != 32 {
		return nil, errors.New("encryption key must be 32 bytes for AES-256")
	}
	if keyID == "" {
		return nil, errors.New("keyID cannot be empty")
	}

	return &EncryptionCodec{
		keyID:         keyID,
		encryptionKey: encryptionKey,
	}, nil
}

// Encode encrypts all payloads using AES-256-GCM.
// The entire payload (including metadata) is serialized and encrypted.
func (ec *EncryptionCodec) Encode(payloads []*commonpb.Payload) ([]*commonpb.Payload, error) {
	result := make([]*commonpb.Payload, len(payloads))

	for i, payload := range payloads {
		// Skip if already encrypted
		if ec.isEncrypted(payload) {
			result[i] = payload
			continue
		}

		// Serialize the entire payload (including metadata) to bytes
		serialized, err := proto.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to serialize payload %d: %w", i, err)
		}

		// Encrypt the serialized payload
		encryptedData, err := ec.encrypt(serialized)
		if err != nil {
			return nil, fmt.Errorf("failed to encrypt payload %d: %w", i, err)
		}

		// Create new payload with encrypted data and metadata
		result[i] = &commonpb.Payload{
			Metadata: map[string][]byte{
				"encoding":              []byte(MetadataEncodingEncrypted),
				MetadataEncryptionKeyID: []byte(ec.keyID),
			},
			Data: encryptedData,
		}
	}

	return result, nil
}

// Decode decrypts all encrypted payloads.
// The decrypted data is deserialized to restore the original payload.
func (ec *EncryptionCodec) Decode(payloads []*commonpb.Payload) ([]*commonpb.Payload, error) {
	result := make([]*commonpb.Payload, len(payloads))

	for i, payload := range payloads {
		// Only decrypt if encrypted
		if !ec.isEncrypted(payload) {
			result[i] = payload
			continue
		}

		// Verify key ID matches
		keyID := string(payload.Metadata[MetadataEncryptionKeyID])
		if keyID != ec.keyID {
			return nil, fmt.Errorf("payload encrypted with key '%s', but codec has key '%s'", keyID, ec.keyID)
		}

		// Decrypt the payload data
		decryptedData, err := ec.decrypt(payload.Data)
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt payload %d: %w", i, err)
		}

		// Deserialize to restore the original payload (with original metadata)
		originalPayload := &commonpb.Payload{}
		if err := proto.Unmarshal(decryptedData, originalPayload); err != nil {
			return nil, fmt.Errorf("failed to deserialize payload %d: %w", i, err)
		}

		result[i] = originalPayload
	}

	return result, nil
}

// isEncrypted checks if a payload is encrypted
func (ec *EncryptionCodec) isEncrypted(p *commonpb.Payload) bool {
	if p.Metadata == nil {
		return false
	}
	encoding := string(p.Metadata["encoding"])
	return encoding == MetadataEncodingEncrypted
}

// encrypt encrypts data using AES-256-GCM
func (ec *EncryptionCodec) encrypt(plaintext []byte) ([]byte, error) {
	// Create AES cipher
	block, err := aes.NewCipher(ec.encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	// Generate random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}

	// Encrypt and prepend nonce
	// Format: [nonce][ciphertext+tag]
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// decrypt decrypts data using AES-256-GCM
func (ec *EncryptionCodec) decrypt(ciphertext []byte) ([]byte, error) {
	// Create AES cipher
	block, err := aes.NewCipher(ec.encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	// Verify ciphertext is long enough
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	// Split nonce and ciphertext
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt: %w", err)
	}

	return plaintext, nil
}
