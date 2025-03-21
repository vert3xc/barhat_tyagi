package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/gob"
	"errors"
	"log"
	"os"
	"strings"
	"time"
)

var secretKey = []byte(os.Getenv("secretKey"))

func CreateSession(session SessionData) (string, error) {
	return SignSession(session)
}

func SignSession(data SessionData) (string, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}
	mac := hmac.New(sha256.New, secretKey)
	mac.Write(buf.Bytes())
	sig := mac.Sum(nil)
	return base64.URLEncoding.EncodeToString(buf.Bytes()) + "." + base64.URLEncoding.EncodeToString(sig), nil
}

func DecodeSession(session string) (SessionData, error) {
	parts := strings.Split(session, ".")
	if len(parts) != 2 {
		log.Println("unexpected")
		return SessionData{}, errors.New("unexpected format")
	}
	payload, err := base64.URLEncoding.DecodeString(parts[0])
	if err != nil {
		log.Println("not b64")
		return SessionData{}, err
	}
	mac := hmac.New(sha256.New, secretKey)
	mac.Write(payload)
	expectedSig := mac.Sum(nil)
	receivedSig, err := base64.URLEncoding.DecodeString(parts[1])
	if err != nil {
		log.Println("not b64")
		return SessionData{}, err
	}
	if !hmac.Equal(expectedSig, receivedSig) {
		log.Println("sigfault")
		return SessionData{}, errors.New("invalid signature")
	}
	var data SessionData
	decoder := gob.NewDecoder(bytes.NewReader(payload))
	if err := decoder.Decode(&data); err != nil {
		log.Println("error decoding")
		return SessionData{}, err
	}
	t := time.Now()
	if t.After(data.Expiry) {
		log.Println("expired")
		return SessionData{}, errors.New("session expired")
	}
	return data, nil
}
