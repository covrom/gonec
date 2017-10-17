package core

import (
	"bytes"
	"compress/flate"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/tls"
	"io"
	"sync"
)

var TLSKeyGonec = []byte(`-----BEGIN EC PARAMETERS-----
BgUrgQQAIg==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MIGkAgEBBDCmKNdI+QpNPKaTTkCc+S09XJ+xHPaZwoBCnKP60MjfNf/nYk9rPDKv
AabSoKbY/sWgBwYFK4EEACKhZANiAAR5j/stywkcTihb06Ye+cVFLFCvYCooe1CC
Z8I1BFp+/F4x9zas1NuKVqR3o/CdRBt3pqPtji2NMq5Bq6uKj92yppHhVXiriT19
d2AI06NcTbToyEMa4ookqgIVV3c29/M=
-----END EC PRIVATE KEY-----
`)

var TLSCertGonec = []byte(`-----BEGIN CERTIFICATE-----
MIICDTCCAZSgAwIBAgIJAI9YNQ6VRzGRMAoGCCqGSM49BAMCMEUxCzAJBgNVBAYT
AlJVMRMwEQYDVQQIDApTb21lLVN0YXRlMREwDwYDVQQKDAh0c292LnBybzEOMAwG
A1UEAwwFZ29uZWMwHhcNMTcxMDE3MDYxMjQ2WhcNNDUwMzA0MDYxMjQ2WjBFMQsw
CQYDVQQGEwJSVTETMBEGA1UECAwKU29tZS1TdGF0ZTERMA8GA1UECgwIdHNvdi5w
cm8xDjAMBgNVBAMMBWdvbmVjMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEeY/7LcsJ
HE4oW9OmHvnFRSxQr2AqKHtQgmfCNQRafvxeMfc2rNTbilakd6PwnUQbd6aj7Y4t
jTKuQaurio/dsqaR4VV4q4k9fXdgCNOjXE206MhDGuKKJKoCFVd3Nvfzo1AwTjAd
BgNVHQ4EFgQU6eLQKM4SXqvZQu9zG34zpL9hdEkwHwYDVR0jBBgwFoAU6eLQKM4S
XqvZQu9zG34zpL9hdEkwDAYDVR0TBAUwAwEB/zAKBggqhkjOPQQDAgNnADBkAjAH
T6xYE6AGYbyF2SOt/X+pVo/zI88wzYlFYvx92ozVviLCLzlDZOTFdJkxv7eeetwC
MFWsEfrik+vBTLviWgGqu/y8ESQSyfOnakWE/cbSNnJptU4+iyWcrAKozssX4jEH
qw==
-----END CERTIFICATE-----
`)

var TLSKeyPair, _ = tls.X509KeyPair(TLSCertGonec, TLSKeyGonec)

var aesKey = []byte("oUwhsPdfj439pfoi")

// TODO: перенести в core-функции языка

func EncryptAES128(plaintext []byte) ([]byte, error) {
	c, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func DecryptAES128(ciphertext []byte) ([]byte, error) {
	c, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, VMErrorSmallDecodeBuffer
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

var zipPool sync.Pool

const zipSizeBuf = 32 * 1024

func getZipBuf() []byte {
	sl := zipPool.Get()
	if sl != nil {
		return sl.([]byte)
	}
	return make([]byte, zipSizeBuf)
}

func putZipBuf(sl []byte) {
	zipPool.Put(sl)
}

func GZip(src []byte) ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, len(src)*2))
	r := bytes.NewReader(src)

	w, err := flate.NewWriter(b, flate.DefaultCompression)
	if err != nil {
		return nil, err
	}

	buf := getZipBuf()
	_, err = io.CopyBuffer(w, r, buf)
	putZipBuf(buf)
	if err != nil {
		return nil, err
	}
	w.Close()
	return b.Bytes(), nil
}

func UnGZip(src []byte) ([]byte, error) {
	b := bytes.NewReader(src)

	r := flate.NewReader(b)
	defer r.Close()

	bo := bytes.NewBuffer(make([]byte, 0, len(src)*2))
	buf := getZipBuf()
	_, err := io.CopyBuffer(bo, r, buf)
	putZipBuf(buf)
	if err != nil {
		return nil, err
	}
	return bo.Bytes(), nil
}
