package encrypt

import (
	"github.com/monnand/dhkx"
	"github.com/pkg/errors"
)

type DHSession struct {
	Group      dhkx.DHGroup
	PrivateKey dhkx.DHKey
	PublicKey  []byte
}

func NewSession() (*DHSession, error) {
	group, err := dhkx.GetGroup(0)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get DH group")
	}

	priv, err := group.GeneratePrivateKey(nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to generate DH private key")
	}

	return &DHSession{
		Group: *group,
		PrivateKey: *priv,
		PublicKey: priv.Bytes(),
	}, nil
}

func ComputeSecret(session DHSession, foreignKey []byte) ([]byte, error) {
	foreignPublic := dhkx.NewPublicKey(foreignKey)
	secret, err := session.Group.ComputeKey(foreignPublic, &session.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to compute shared secret")
	}
	return secret.Bytes(), nil
}
