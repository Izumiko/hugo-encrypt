package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/crypto/pbkdf2"
)

func deriveKey(passphrase string, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 8)
		// http://www.ietf.org/rfc/rfc2898.txt
		// Salt.
		rand.Read(salt)
	}
	return pbkdf2.Key([]byte(passphrase), salt, 1000, 32, sha256.New), salt
}

func encrypt(passphrase, plaintext string) string {
	key, salt := deriveKey(passphrase, nil)
	iv := make([]byte, 12)
	// http://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-38d.pdf
	// Section 8.2
	rand.Read(iv)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data := aesgcm.Seal(nil, iv, []byte(plaintext), nil)
	return hex.EncodeToString(salt) + "-" + hex.EncodeToString(iv) + "-" + hex.EncodeToString(data)
}

func encryptPage(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(content))
	if err != nil {
		panic(err)
	}
	block := doc.Find("cipher-text")
	if len(block.Nodes) == 1 {
		fmt.Printf("Processing %s\n", path)

		password, _ := block.Attr("data-password")
		blockhtml, _ := block.Html()
		enchtml := encrypt(password, blockhtml)
		block.RemoveAttr("data-password")
		block.SetHtml(enchtml)
		wholehtml, _ := doc.Html()
		ioutil.WriteFile(path, []byte(wholehtml), 0644)
	}
}

func main() {
	err := filepath.Walk("./public", func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		ok := strings.HasSuffix(f.Name(), ".html")
		if ok {
			encryptPage(path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
