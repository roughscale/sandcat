package util

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"os"
	"strings"
	"time"
	"unicode"
)

// Encode base64 encodes bytes
func Encode(b []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(b))
}

// Decode base64 decodes a string
func Decode(s string) []byte {
	raw, _ := base64.StdEncoding.DecodeString(s)
	return raw
}

// Unpack converts bytes into JSON
func Unpack(b []byte) (out map[string]interface{}) {
	_ = json.Unmarshal(b, &out)
	return
}

// Exists checks for a file
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// CheckErrorMessage Check for error message
func CheckErrorMessage(err error) bool {
	if err != nil && err.Error() != "The operation completed successfully." {
		println(err.Error())
		return true
	}
	return false
}

// StringToByteArrayString transforms an input string to a byte string
func StringToByteArrayString(input string) ([]byte, error) {
	temp := removeWhiteSpace(input)
	temp = strings.Replace(temp, "0x", "", -1)
	temp = strings.Replace(temp, "\\x", "", -1)
	temp = strings.Replace(temp, ",", "", -1)
	return hex.DecodeString(temp)
}

// Sleep sleeps for a desired interval
func Sleep(interval float64) {
	time.Sleep(time.Duration(interval) * time.Second)
}

func removeWhiteSpace(input string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, input)
}

func StopProcess(pid int) {
	proc, _ := os.FindProcess(pid)
	_ = proc.Kill()
}

// TimeoutWatchdog executes a timer that signals to a channel if a command broke it's timeout duration
func TimeoutWatchdog(timeoutChan chan bool, timeout time.Duration) {
	time.Sleep(timeout * time.Second)
	timeoutChan <- true
}