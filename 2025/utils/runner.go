package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
)

// Run the solutions for quest.
func Run(
	solvePart1 func(string) string,
	solvePart2 func(string) string,
	solvePart3 func(string) string,
	submit bool,
) {
	_, callingFilePath, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("Failed to get caller info for quest ID detection")
	}

	eventId := "2025"
	questId := parseQuestId(callingFilePath)

	sessionCookie := os.Getenv("SESSION_COOKIE")
	if sessionCookie == "" {
		log.Fatal("SESSION_COOKIE environment variable is not set")
	}

	seed := fetchSeed(sessionCookie)
	encryptedInputs := fetchEncryptedInputs(eventId, questId, seed, sessionCookie)
	aesKeys := fetchAesKeys(eventId, questId, sessionCookie)

	runPart := func(part string, solve func(string) string) {
		key, ok := aesKeys[fmt.Sprintf("key%s", part)]
		if !ok {
			fmt.Printf("Input not yet available for Part %s\n", part)
			return
		}

		decryptedInput := decryptInput(encryptedInputs[part], key)

		answer := solve(decryptedInput)
		if answer == "" {
			fmt.Printf("No answer for Part %s\n", part)
			return
		}

		fmt.Printf("Part %s Answer: %s\n", part, answer)

		if !submit {
			return
		}

		submitAnswer(eventId, questId, part, answer, sessionCookie)
	}

	runPart("1", solvePart1)
	runPart("2", solvePart2)
	runPart("3", solvePart3)
}

// Parse quest ID from caller file path.
func parseQuestId(callerFilePath string) string {
	matches := regexp.MustCompile(`quest(\d{2})`).FindStringSubmatch(filepath.Base(callerFilePath))
	if len(matches) != 2 {
		log.Fatalf("Failed to find quest ID in calling file path: %s", callerFilePath)
	}

	questId, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Fatalf("Failed to extract numeric quest ID: %s (%s)", callerFilePath, matches[1])
	}

	return strconv.Itoa(questId)
}

// Fetch seed from website.
func fetchSeed(sessionCookie string) string {
	url := "https://everybody.codes/api/user/me"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request for user seed: %v", err)
	}

	request.AddCookie(&http.Cookie{Name: "everybody-codes", Value: sessionCookie})

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("Failed to send request for user seed: %v", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read response body for user seed: %v", err)
	}

	if response.StatusCode != 200 {
		log.Fatalf("Unexpected status %d when fetching user seed: %s", response.StatusCode, string(responseBody))
	}

	var responseData struct {
		Seed int `json:"seed"`
	}

	if err := json.Unmarshal(responseBody, &responseData); err != nil {
		log.Fatalf("Failed to decode JSON response for user seed: %v\nResponse body: %s", err, string(responseBody))
	}

	if responseData.Seed == 0 {
		log.Fatal("User seed is 0 - verify that the session cookie is correct")
	}

	return strconv.Itoa(responseData.Seed)
}

// Fetch encrypted inputs for all available quest parts.
func fetchEncryptedInputs(eventID, questID, seed, sessionCookie string) map[string]string {
	url := fmt.Sprintf("https://everybody-codes.b-cdn.net/assets/%s/%s/input/%s.json", eventID, questID, seed)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request for encrypted inputs: %v", err)
	}

	request.AddCookie(&http.Cookie{Name: "everybody-codes", Value: sessionCookie})

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("Failed to send request for encrypted inputs: %v", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read response body for encrypted inputs: %v", err)
	}

	if response.StatusCode != 200 {
		log.Fatalf("Unexpected status %d when fetching encrypted inputs: %s", response.StatusCode, string(responseBody))
	}

	var responseData map[string]string
	if err := json.Unmarshal(responseBody, &responseData); err != nil {
		log.Fatalf("Failed to decode JSON response for encrypted inputs: %v\nResponse body: %s", err, string(responseBody))
	}

	return responseData
}

// Fetch AES keys needed to decrypt inputs.
func fetchAesKeys(eventID, questID, sessionCookie string) map[string]string {
	url := fmt.Sprintf("https://everybody.codes/api/event/%s/quest/%s", eventID, questID)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request for AES keys: %v", err)
	}

	request.AddCookie(&http.Cookie{Name: "everybody-codes", Value: sessionCookie})

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("Failed to send request for AES keys: %v", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read response body for AES keys: %v", err)
	}

	if response.StatusCode != 200 {
		log.Fatalf("Unexpected status %d when fetching AES keys:\n%s", response.StatusCode, string(responseBody))
	}

	var responseData map[string]string
	if err := json.Unmarshal(responseBody, &responseData); err != nil {
		log.Fatalf("Failed to decode JSON response for AES keys: %v\nResponse body: %s", err, string(responseBody))
	}

	return responseData
}

// Decrypt an encrypted hex input using the given AES key.
func decryptInput(encryptedInput, key string) string {
	keyBytes := []byte(key)

	encryptedInputBytes, err := hex.DecodeString(encryptedInput)
	if err != nil {
		log.Fatalf("Failed to decode encrypted input bytes: %v", err)
	}

	decryptedInputBytes := decryptBytesAes(encryptedInputBytes, keyBytes)

	return string(decryptedInputBytes)
}

// Decrypt a ciphertext using AES CBC with PKCS7 padding.
func decryptBytesAes(data, key []byte) []byte {
	if len(data) < aes.BlockSize {
		log.Fatalf("Ciphertext is too short: %v", data)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Failed to create a cipher block for key %s: %v", key, err)
	}

	iv := key[:aes.BlockSize]
	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(data, data)

	byteCount := len(data)
	if byteCount == 0 || byteCount%aes.BlockSize != 0 {
		log.Fatalf("Invalid bytes count: %d", byteCount)
	}

	padByteCount := int(data[byteCount-1])
	if padByteCount == 0 || padByteCount > aes.BlockSize {
		log.Fatalf("Invalid padding: %d", padByteCount)
	}

	for i := byteCount - padByteCount; i < byteCount; i++ {
		if data[i] != byte(padByteCount) {
			log.Fatalf("Invalid padding: %v", data[byteCount-padByteCount:])
		}
	}

	return data[:byteCount-padByteCount]
}

// Submit an answer to the website.
func submitAnswer(eventID, questID, part, answer, sessionCookie string) {
	url := fmt.Sprintf("https://everybody.codes/api/event/%s/quest/%s/part/%s/answer", eventID, questID, part)

	payload := map[string]string{"answer": answer}
	requestBody, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Failed to conver answer payload to JSON: %v", err)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Failed to create request to submit answer: %v", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.AddCookie(&http.Cookie{Name: "everybody-codes", Value: sessionCookie})

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("Failed to send request to submit answer: %v", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read response body for answer submission: %v", err)
	}

	switch response.StatusCode {
	case 200:
		var data map[string]any
		if err := json.Unmarshal(responseBody, &data); err != nil {
			log.Fatalf("Failed to decode JSON response for answer submission: %v\nResponse body: %s", err, string(responseBody))
		}

		if data["correct"].(bool) {
			fmt.Printf("  Submitted correct answer for Part %s!\n", part)
		} else {
			log.Printf("  Submitted incorrect answer for Part %s, try again\n", part)
		}
	case 409:
		fmt.Printf("  Already submitted a correct answer for Part %s\n", part)
	case 423:
		log.Printf("  Submitted answer too recently, try again after one minute\n")
	default:
		log.Fatalf("Failed to submit answer with status %d: %s", response.StatusCode, string(responseBody))
	}
}
