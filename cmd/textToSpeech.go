/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"context"
    "io/ioutil"
	"log"

    texttospeech "cloud.google.com/go/texttospeech/apiv1"
    texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

// textToSpeechCmd represents the textToSpeech command
var textToSpeechCmd = &cobra.Command{
	Use:   "textToSpeech",
	Aliases: []string{"tts"},
	Short: "Transform text to speech and export result in a mp3 file",
	Long: `Usage: textToSpeech <text> <lang>`,
	Run: func(cmd *cobra.Command, args []string) {
		var text = "Hello, World!"
		var lang = "en-US"

		if len(args) >= 1 && args[0] != "" {
			text = args[0]
		}

		if len(args) == 2 && args[1] != "" {
			lang = args[1]
		}
		fmt.Println("textToSpeech called with text =",text,"and lang =",lang)

		textToSpeach(text, lang)
	},
}

func init() {
	rootCmd.AddCommand(textToSpeechCmd)
	// textToSpeechCmd.Flags().String("text", "", "Text to transform")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	textToSpeechCmd.PersistentFlags().String("text", "", "Text to transform")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// textToSpeechCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func textToSpeach(text, lang string) {
        ctx := context.Background()

        client, err := texttospeech.NewClient(ctx)
        if err != nil {
                log.Fatal(err)
        }

        // Perform the text-to-speech request on the text input with the selected
        // voice parameters and audio file type.
        req := texttospeechpb.SynthesizeSpeechRequest{
                // Set the text input to be synthesized.
                Input: &texttospeechpb.SynthesisInput{
                        InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
                },
                // Build the voice request, select the language code ("en-US") and the SSML
                // voice gender ("neutral").
                Voice: &texttospeechpb.VoiceSelectionParams{
                        LanguageCode: lang,
                        SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
                },
                // Select the type of audio file you want returned.
                AudioConfig: &texttospeechpb.AudioConfig{
                        AudioEncoding: texttospeechpb.AudioEncoding_MP3,
                },
        }

        resp, err := client.SynthesizeSpeech(ctx, &req)
        if err != nil {
                log.Fatal(err)
        }

        // The resp's AudioContent is binary.
        filename := "output.mp3"
        err = ioutil.WriteFile(filename, resp.AudioContent, 0644)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("Audio content written to file: %v\n", filename)
}

