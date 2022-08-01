package store_test

import (
	"reflect"
	"testing"

	"github.com/bassforce86/trigrams/pkg/store"
)

func TestRandomChooser_ChooseInitialTrigram(t *testing.T) {
	c := &store.RandomChooser{}
	type args struct {
		trigramMap store.TrigramMap
	}
	tests := []struct {
		name string
		c    *store.RandomChooser
		args args
		want store.Trigram
	}{
		{
			name: "trigrams are empty",
			c:    c,
			args: args{
				trigramMap: map[string]map[string]map[string]int{},
			},
			want: store.Trigram{"", "", ""},
		},
		{
			name: "word begins with Uppercase",
			c:    c,
			args: args{
				trigramMap: map[string]map[string]map[string]int{
					"test": {
						"test1": {
							"test2": 1,
						},
					},
					"Story": {
						"story1": {
							"story2": 2,
						},
					},
				},
			},
			want: store.Trigram{"Story", "story1", "story2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := c.ChooseInitialTrigram(tt.args.trigramMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RandomChooser.ChooseInitialTrigram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomChooser_ChooseNextWord(t *testing.T) {
	c := &store.RandomChooser{}
	type args struct {
		wordFreqs map[string]int
	}
	tests := []struct {
		name string
		c    *store.RandomChooser
		args args
		want string
	}{
		{
			name: "invalid next word",
			c:    c,
			args: args{
				wordFreqs: map[string]int{},
			},
			want: "",
		},
		{
			name: "valid next word",
			c:    c,
			args: args{
				wordFreqs: map[string]int{"word": 1},
			},
			want: "word",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := c.ChooseNextWord(tt.args.wordFreqs); got != tt.want {
				t.Errorf("RandomChooser.ChooseNextWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
