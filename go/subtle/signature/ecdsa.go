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
//
////////////////////////////////////////////////////////////////////////////////

// Package signature provides subtle implementations of the Signer and Verifier primitives.
package signature

import (
	"fmt"
	"math/big"
)

var errUnsupportedEncoding = fmt.Errorf("ecdsa: unsupported encoding")

// ECDSASignerature is a struct holding r and s values of an ECDSA signature.
type ECDSASignerature struct {
	R, S *big.Int
}

// NewECDSASignerature creates a new ecdsaSignature object.
func NewECDSASignerature(r, s *big.Int) *ECDSASignerature {
	return &ECDSASignerature{R: r, S: s}
}

// EncodeECDSASignature converts the signature to the given encoding format.
// Only DER encoding is supported now.
func (sig *ECDSASignerature) EncodeECDSASignature(encoding string) ([]byte, error) {
	switch encoding {
	case "DER":
		return asn1encode(sig)
	default:
		return nil, errUnsupportedEncoding
	}
}

// DecodeECDSASignerature creates a new ECDSA signature using the given byte slice.
// The function assumes that the byte slice is the concatenation of the BigEndian
// representation of two big integer r and s.
func DecodeECDSASignerature(encodedBytes []byte,
	encoding string) (*ECDSASignerature, error) {
	switch encoding {
	case "DER":
		return asn1decode(encodedBytes)
	default:
		return nil, errUnsupportedEncoding
	}
}

// ValidateECDSAParams validates ECDSA parameters.
// The hash's strength must not be weaker than the curve's strength.
// Only DER encoding is supported now.
func ValidateECDSAParams(hashAlg string, curve string, encoding string) error {
	switch encoding {
	case "DER":
		break
	default:
		return errUnsupportedEncoding
	}
	switch curve {
	case "NIST_P256":
		if hashAlg != "SHA256" {
			return fmt.Errorf("invalid hash type, expect SHA-256")
		}
	case "NIST_P384", "NIST_P521":
		if hashAlg != "SHA512" {
			return fmt.Errorf("invalid hash type, expect SHA-512")
		}
	default:
		return fmt.Errorf("unsupported curve: %s", curve)
	}
	return nil
}
