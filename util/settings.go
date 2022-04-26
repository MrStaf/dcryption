package util

import (
	"encoding/json"
	"io/ioutil"
)

type Parameters struct {
	Language      string `json:"language"`
	Theme         string `json:"theme"`
	FontSize      int    `json:"fontSize"`
	FontFamily    string `json:"fontFamily"`
	EncryptedPath string `json:"encryptedPath"`
	DecryptedPath string `json:"decryptedPath"`
	Key           string
}

type Settings struct {
	P *Parameters `json:"parameters"`
}

func (s *Settings) SetSettings(language string, theme string, fontSize int, fontFamily string, encryptedPath string, decryptedPath string, key string) {
	s.P.Language = language
	s.P.Theme = theme
	s.P.FontSize = fontSize
	s.P.FontFamily = fontFamily
	s.P.EncryptedPath = encryptedPath
	s.P.DecryptedPath = decryptedPath
	s.P.Key = key
}

func (s *Settings) GetLanguage() string {
	return s.P.Language
}

func (s *Settings) GetTheme() string {
	return s.P.Theme
}

func (s *Settings) GetFontSize() int {
	return s.P.FontSize
}

func (s *Settings) GetFontFamily() string {
	return s.P.FontFamily
}

func (s *Settings) GetEncryptedPath() string {
	return s.P.EncryptedPath
}

func (s *Settings) GetDecryptedPath() string {
	return s.P.DecryptedPath
}

func (s *Settings) GetKey() string {
	return s.P.Key
}

func (s *Settings) SetLanguage(language string) {
	s.P.Language = language
}

func (s *Settings) SetTheme(theme string) {
	s.P.Theme = theme
}

func (s *Settings) SetFontSize(fontSize int) {
	s.P.FontSize = fontSize
}

func (s *Settings) SetFontFamily(fontFamily string) {
	s.P.FontFamily = fontFamily
}

func (s *Settings) SetEncryptedPath(encryptedPath string) {
	s.P.EncryptedPath = encryptedPath
}

func (s *Settings) SetDecryptedPath(decryptedPath string) {
	s.P.DecryptedPath = decryptedPath
}

func (s *Settings) SetKey(key string) {
	s.P.Key = key
}

func LoadSettingsFile() *Settings {
	// Load the settings file
	rawJson, err := ioutil.ReadFile("./settings.json")
	if err != nil {
		panic(err)
	}
	// Parse the settings file
	var s Settings
	err = json.Unmarshal(rawJson, &s)
	if err != nil {
		panic(err)
	}
	// Return the settings
	return &s
}

func SaveSettingsFile(s *Settings) {
	// Save the settings file
	rawJson, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./settings.json", rawJson, 0644)
	if err != nil {
		panic(err)
	}
}
