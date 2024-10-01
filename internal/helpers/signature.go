package helpers

import (
	"encoding/base64"
	"errors"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
)

func VerifySignature(addressString, pubKeyString, signatureString, nonce string) error {
	// Verify the username is a valid cosmos address
	address, err := sdk.AccAddressFromBech32(addressString)
	if err != nil {
		return errors.New("provided address is not valid")
	}

	// Unmarshal the provided signdoc and
	signDoc := &tx.SignDoc{
		AccountNumber: 0,
		AuthInfoBytes: []byte{},
		BodyBytes:     []byte(nonce),
		ChainId:       "",
	}

	// Marshal the signature in raw bytes
	signBytes, err := signDoc.Marshal()
	if err != nil {
		return errors.New("failed to serialize the signature")
	}

	// Decode the public key
	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKeyString)
	if err != nil {
		return errors.New("failed to decode pubkey")
	}
	pubKey := &secp256k1.PubKey{Key: pubKeyBytes}

	// Verify the address used corresponds to the public key used
	if addressFromPK := sdk.AccAddress(pubKey.Address().Bytes()); addressFromPK.String() != address.String() {
		return errors.New("address in provided signature is not related to the public key used to sign")
	}

	// Decode the signature
	signature, err := base64.StdEncoding.DecodeString(signatureString)
	if err != nil {
		return errors.New("failed to decode signature")
	}

	// Check signature validity
	if ok := pubKey.VerifySignature(signBytes, signature); !ok {
		return errors.New("invalid signature")
	}

	return nil
}
