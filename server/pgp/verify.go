// Copyright 2021 The boltchat Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pgp

import (
	"errors"
	"strings"

	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

func getSignature(rawSig string) (*packet.Signature, error) {
	sigReader := strings.NewReader(rawSig)
	sigDecoded, decodeErr := armor.Decode(sigReader)

	if decodeErr != nil {
		return nil, decodeErr
	}

	pack, packErr := packet.Read(sigDecoded.Body)

	if packErr != nil {
		return nil, packErr
	}

	sig, ok := pack.(*packet.Signature)
	if !ok {
		return nil, errors.New("invalid signature")
	}

	return sig, nil
}

func getPublicKey(rawPubKey string) (*packet.PublicKey, error) {
	pubKeyRead := strings.NewReader(rawPubKey)
	pubKeyDecode, pubKeyDecodeErr := armor.Decode(pubKeyRead)

	if pubKeyDecodeErr != nil {
		return nil, pubKeyDecodeErr
	}

	pubKeyPack, pubKeyPackErr := packet.Read(pubKeyDecode.Body)

	if pubKeyPackErr != nil {
		return nil, pubKeyPackErr
	}

	pubKey, ok := pubKeyPack.(*packet.PublicKey)
	if !ok {
		return nil, errors.New("invalid public key")
	}

	return pubKey, nil
}

// VerifyMessageSignature verifies a message sent by a user with
// the corresponding public key and signature.
func VerifyMessageSignature(rawSig string, rawPubKey string, msg string) (*packet.PublicKey, error) {
	sig, sigErr := getSignature(rawSig)
	if sigErr != nil {
		return nil, sigErr
	}

	pubKey, pubKeyErr := getPublicKey(rawPubKey)
	if pubKeyErr != nil {
		return nil, pubKeyErr
	}

	hash := sig.Hash.New()
	_, hashErr := hash.Write([]byte(msg))

	if hashErr != nil {
		return nil, hashErr
	}

	return pubKey, pubKey.VerifySignature(hash, sig)
}
