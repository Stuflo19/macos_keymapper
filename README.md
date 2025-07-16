## What is this for?
This is a tool for quickly creating hotkeys for launching app's on MacOS.

## How to set up a hotkey?
To set up a hotkey you simply need to add the following structure into `~/.config/macos_keymapper/config.json`

```
[
    {
        "name": "Safari",
        "key": "ctrl+shift+s",
        "application": "Safari"
    }
]
```

By defeault this will search for the file `Safari.app` inside of `/Applications`. However you can add an `applicationPath` if you want to launch the app from a specific location. This for example could be the path `/System/Applications`

## How to setup as a background process?
To set the app as a background process simply add the following into a file named `com.stuflo.macoskeymapper.plist` with the following content:
```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
  <dict>
    <key>Label</key>
    <string>com.stuflo.macoskeymapper</string>
    <key>ProgramArguments</key>
    <array>
      <string>/path/to/macos_keymapper</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
    <key>KeepAlive</key>
    <true/>
  </dict>
</plist>
```
