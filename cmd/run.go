package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/bassforce86/trigrams/pkg/store"
	"github.com/spf13/cobra"
)

var cache = store.NewMapTrigramStore(&store.RandomChooser{})
var file string
var length int

var run = &cobra.Command{
	Use:   "run",
	Short: "Run the trigram generator against the given file",
	Run: func(cmd *cobra.Command, args []string) {
		if file == "" {
			fmt.Println("Cannot locate file to learn from")
			fmt.Println("")
			cmd.Help()
		}
		if bytes, err := os.ReadFile(file); err != nil {
			log.Fatal(err)
			os.Exit(2)
		} else {
			ingest(string(bytes))
		}
		fmt.Println(cache.GenerateText(length))
	},
}

func init() {
	run.Flags().StringVarP(&file, "from-file", "f", "", "File to learn")
	run.Flags().IntVarP(&length, "length", "l", 100, "How many words should the result be? (approx)")
	run.MarkFlagRequired("file")
	base.AddCommand(run)
}

func ingest(text string) error {
	trigrams, err := parseTrigrams(text)

	if err != nil {
		return err
	}

	for _, trigram := range trigrams {
		cache.AddTrigram(trigram)
	}

	return nil
}

func parseTrigrams(text string) ([]store.Trigram, error) {
	isSpace := func(char rune) bool { return unicode.IsSpace(char) }

	stripped := regexp.MustCompile(`\_|;|!|\?`).ReplaceAllString(text, "")
	words := strings.FieldsFunc(stripped, isSpace)

	if len(words) < 3 {
		return nil, errors.New("text to learn needs to have more than 3 words")
	}

	var trigrams []store.Trigram

	for i := 0; i < len(words)-2; i++ {
		trigram := store.Trigram{words[i], words[i+1], words[i+2]}
		trigrams = append(trigrams, trigram)
	}

	return trigrams, nil
}
