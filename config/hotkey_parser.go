package config

import (
	"errors"
	"strings"

	"golang.design/x/hotkey"
)

var modifierMap = map[string]hotkey.Modifier{
	"cmd":   hotkey.ModCmd,
	"ctrl":  hotkey.ModCtrl,
	"shift": hotkey.ModShift,
	"opt":   hotkey.ModOption,
}

var keyMap = map[string]hotkey.Key{
	"up":     hotkey.KeyUp,
	"down":   hotkey.KeyDown,
	"left":   hotkey.KeyLeft,
	"right":  hotkey.KeyRight,
	"return": hotkey.KeyReturn,
	"tab":    hotkey.KeyTab,
	"space":  hotkey.KeySpace,
	"delete": hotkey.KeyDelete,
	"escape": hotkey.KeyEscape,
	"0":      hotkey.Key0,
	"1":      hotkey.Key1,
	"2":      hotkey.Key2,
	"3":      hotkey.Key3,
	"4":      hotkey.Key4,
	"5":      hotkey.Key5,
	"6":      hotkey.Key6,
	"7":      hotkey.Key7,
	"8":      hotkey.Key8,
	"9":      hotkey.Key9,
	"f1":     hotkey.KeyF1,
	"f2":     hotkey.KeyF2,
	"f3":     hotkey.KeyF3,
	"f4":     hotkey.KeyF4,
	"f5":     hotkey.KeyF5,
	"f6":     hotkey.KeyF6,
	"f7":     hotkey.KeyF7,
	"f8":     hotkey.KeyF8,
	"f9":     hotkey.KeyF9,
	"f10":    hotkey.KeyF10,
	"f11":    hotkey.KeyF11,
	"f12":    hotkey.KeyF12,
	"f13":    hotkey.KeyF13,
	"f14":    hotkey.KeyF14,
	"f15":    hotkey.KeyF15,
	"f16":    hotkey.KeyF16,
	"f17":    hotkey.KeyF17,
	"f18":    hotkey.KeyF18,
	"f19":    hotkey.KeyF19,
	"f20":    hotkey.KeyF20,
	"a":      hotkey.KeyA,
	"b":      hotkey.KeyB,
	"c":      hotkey.KeyC,
	"d":      hotkey.KeyD,
	"e":      hotkey.KeyE,
	"f":      hotkey.KeyF,
	"g":      hotkey.KeyG,
	"h":      hotkey.KeyH,
	"i":      hotkey.KeyI,
	"j":      hotkey.KeyJ,
	"k":      hotkey.KeyK,
	"l":      hotkey.KeyL,
	"m":      hotkey.KeyM,
	"n":      hotkey.KeyN,
	"o":      hotkey.KeyO,
	"p":      hotkey.KeyP,
	"q":      hotkey.KeyQ,
	"r":      hotkey.KeyR,
	"s":      hotkey.KeyS,
	"t":      hotkey.KeyT,
	"u":      hotkey.KeyU,
	"v":      hotkey.KeyV,
	"w":      hotkey.KeyW,
	"x":      hotkey.KeyX,
	"y":      hotkey.KeyY,
	"z":      hotkey.KeyZ,
}

func getModifier(modifier string) (hotkey.Modifier, error) {
	val, ok := modifierMap[strings.ToLower(modifier)]

	if !ok {
		return 0, errors.New("unknown modifier")
	}

	return val, nil
}

func getKey(key string) (hotkey.Key, error) {
	val, ok := keyMap[strings.ToLower(key)]

	if !ok {
		return 0, errors.New("unknown key")
	}

	return val, nil
}
