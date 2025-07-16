package config

import (
	"encoding/json"
	"log"
	"os"
)

const configFileLocation string = "/Users/stu/.config/macos_keymapper/config.json"

type Parser struct {
	hotkeyItems []HotkeyItem
}

func NewParser() *Parser {
	return &Parser{
		hotkeyItems: []HotkeyItem{},
	}
}

func (p *Parser) ParseConfigFile() {
	contents, err := os.ReadFile(configFileLocation)

	if err != nil {
		log.Fatal(err)
	}

	var hotkeyItems []HotkeyItem

	err = json.Unmarshal(contents, &hotkeyItems)

	if err != nil {
		log.Fatal(err)
	}

	p.hotkeyItems = hotkeyItems
}

func (p *Parser) RegisterHotkeys() {
	for _, hotkeyItem := range p.hotkeyItems {
		_, err := hotkeyItem.RegisterHotkey()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (p *Parser) PrintHotkeys() {
	for _, hotkeyItem := range p.hotkeyItems {
		log.Println(hotkeyItem)
	}
}
