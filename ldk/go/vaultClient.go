package ldk

import (
	"context"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/proto"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/server"
)

// VaultClient is used by loops to make vault vault service requests
type VaultClient struct {
	VaultClient proto.VaultClient
	Session     *server.Session
}

// Delete is used by loops to delete a vault entry
func (s *VaultClient) Delete(ctx context.Context, key string) error {
	_, err := s.VaultClient.VaultDelete(ctx, &proto.VaultDeleteRequest{
		Key:     key,
		Session: s.Session.ToProto(),
	})
	return err
}

// Exists is used by loops to check if a key exists in the vault
func (s *VaultClient) Exists(ctx context.Context, key string) (bool, error) {
	resp, err := s.VaultClient.VaultExists(ctx, &proto.VaultExistsRequest{
		Key:     key,
		Session: s.Session.ToProto(),
	})
	if err != nil {
		return false, err
	}

	return resp.Exists, nil
}

// Read is used by loops to get the value of an entry
func (s *VaultClient) Read(ctx context.Context, key string) (string, error) {
	resp, err := s.VaultClient.VaultRead(ctx, &proto.VaultReadRequest{
		Key:     key,
		Session: s.Session.ToProto(),
	})
	if err != nil {
		return "", err
	}

	return resp.Value, nil
}

// Write is used by plugins to set a vault entry value
func (s *VaultClient) Write(ctx context.Context, key, value string) error {
	_, err := s.VaultClient.VaultWrite(ctx, &proto.VaultWriteRequest{
		Key:     key,
		Value:   value,
		Session: s.Session.ToProto(),
	})
	return err
}
