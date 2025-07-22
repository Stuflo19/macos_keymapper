package config

import (
	"log"
	"os/exec"
	"strings"

	"golang.design/x/hotkey"
)

type HotkeyItem struct {
	Name            string `json:"name"`
	Key             string `json:"key"`
	Application     string `json:"application"`
	ApplicationPath string `json:"applicationPath",omitempty`
}

func (h HotkeyItem) String() string {
	return h.Name
}

func (h HotkeyItem) getHotkey() *hotkey.Hotkey {
	split := strings.Split(h.Key, "+")

	modifiers := make([]hotkey.Modifier, 0)
	var key hotkey.Key

	for _, k := range split {
		modifier, err := getModifier(k)
		if err == nil {
			modifiers = append(modifiers, modifier)
			continue
		}

		k, err := getKey(k)
		if err == nil {
			key = k
			continue
		}
	}

	return hotkey.New(modifiers, key)
}

func (h HotkeyItem) SetupChannel(hk *hotkey.Hotkey) {
	for {
		select {
		case <-hk.Keydown():
			h.LaunchApplication()
		}
	}
}

func (h HotkeyItem) RegisterHotkey() (*hotkey.Hotkey, error) {
	hk := h.getHotkey()

	err := hk.Register()
	if err != nil {
		log.Fatalf("hotkey: failed to register hotkey: %v", err)
		return nil, err
	}

	go h.SetupChannel(hk)

	return hk, nil
}

func (h HotkeyItem) LaunchApplication() {
	if h.ApplicationPath == "" {
		h.ApplicationPath = "/Applications/"
	}
	if !strings.HasSuffix(h.ApplicationPath, "/") {
		h.ApplicationPath += "/"
	}

	cmd := exec.Command("open", h.ApplicationPath+h.Application+".app")
	if err := cmd.Start(); err != nil {
		log.Printf("hotkey: failed to start cmd: %v", err)
	}
}
