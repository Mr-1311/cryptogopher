package main

import (
	"crypto/md5"
	"reflect"
	"testing"
)

var TestDatas = []struct {
	OriginalData string
	Passphrase   string
}{
	{"abcdefghijklmnoprjstuvwxyz", "hello_world"},
	{"ABCDEFGHIJKLMNOPRJSTUVWXYZ", "world_hello"},
	{"ığüşöç", "küçük_harf_türkçe_karakter_test"},
	{"İĞÜŞÖÇ", "BÜYÜK_HARF_TÜRKÇE_KARAKTER_TEST"},
	{"É!'^+%&/()=?_", "special_characters_shift"},
	{"<>£#$½¾|", "special_characters_alt"},
}

func TestEncryptDecrypt(t *testing.T) {
	hasher := md5.New()
	for _, data := range TestDatas {
		encrypted := string(Encrypt([]byte(data.OriginalData), []byte(data.Passphrase)))
		decrypted := string(Decrypt([]byte(encrypted), []byte(data.Passphrase)))

		if !reflect.DeepEqual(hasher.Sum([]byte(decrypted)), hasher.Sum([]byte(data.OriginalData))) {
			t.Errorf("Encrypt: Expected: %s, Actual: %s", data.OriginalData, decrypted)
		} else {
			t.Logf("[PASSED]: Original: %s, Encrypted: %s, Decrypted: %s", data.OriginalData, encrypted, decrypted)

		}
	}
}
