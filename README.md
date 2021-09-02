# ai1y-cli

## CLI

```
$ ai1y-cli textToSpeech hello en-US
$ ai1y-cli textToSpeech salut fr-FR
```

## Run

```
$ go run main.go

AI1Y app allow to call text-to-speech API that transform text to speech

Usage:
  ai1y-cli [command]

Available Commands:
  help         Help about any command
  textToSpeech Transform text to speech and export result in a mp3 file

Flags:
      --config string   config file (default is $HOME/.ai1y-cli.yaml)
  -h, --help            help for ai1y-cli
  -t, --toggle          Help message for toggle

Use "ai1y-cli [command] --help" for more information about a command.

$ go run main.go tts
textToSpeech called with text = Hello, World! and lang = en-US
Audio content written to file: output.mp3

$ go run main.go tts "hello scraly"
textToSpeech called with text = hello scraly and lang = en-US
Audio content written to file: output.mp3

$ go run main.go tts "salut scraly" fr-FR
textToSpeech called with text = salut scraly and lang = fr-FR
Audio content written to file: output.mp3
```

## Build program

```
$ go build -o bin/ai1y-cli main.go
```
