package store_test

import (
	"reflect"
	"sync"
	"testing"

	"github.com/bassforce86/trigrams/pkg/store"
)

func TestNewMapTrigramStore(t *testing.T) {
	trigramChooser := &store.RandomChooser{}
	type args struct {
		chooser store.Chooser
	}
	tests := []struct {
		name string
		args args
		want *store.TrigramMapStore
	}{
		{
			name: "new store",
			args: args{chooser: trigramChooser},
			want: store.NewMapTrigramStore(trigramChooser),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := store.NewMapTrigramStore(tt.args.chooser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMapTrigramStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrigramMapStore_AddTrigram(t *testing.T) {
	trigramChooser := &store.RandomChooser{}
	cache := store.NewMapTrigramStore(trigramChooser)

	type fields struct {
		trigrams store.TrigramMap
		mutex    *sync.Mutex
		chooser  store.Chooser
	}
	type args struct {
		trigram store.Trigram
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "add a new trigram",
			fields: fields{
				trigrams: store.TrigramMap{
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
				mutex:   &sync.Mutex{},
				chooser: trigramChooser,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache.AddTrigram(tt.args.trigram)
		})
	}
}

func TestTrigramMapStore_GenerateText(t *testing.T) {
	trigramChooser := &store.RandomChooser{}
	cache := store.NewMapTrigramStore(trigramChooser)
	cache.AddTrigram(store.Trigram{"Test", "test1", "test2"})
	cache.AddTrigram(store.Trigram{"Test", "test1", "test2"})
	cache.AddTrigram(store.Trigram{"test1", "test2", "test3"})
	cache.AddTrigram(store.Trigram{"test2", "test3", "test4"})
	cache.AddTrigram(store.Trigram{"test3", "test4", "test5"})

	type fields struct {
		trigrams store.TrigramMap
		mutex    *sync.Mutex
		chooser  store.Chooser
	}
	type args struct {
		max int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "make text using most frequent trigrams",
			fields: fields{
				trigrams: cache.Trigrams(),
				mutex:    &sync.Mutex{},
				chooser:  trigramChooser,
			},
			args: args{max: 3},
			want: "Test test1 test2 test3 test4",
		},
		{
			name: "make text using most frequent trigrams",
			fields: fields{
				trigrams: cache.Trigrams(),
				mutex:    &sync.Mutex{},
				chooser:  trigramChooser,
			},
			args: args{max: 4},
			want: "Test test1 test2 test3 test4 test5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cache.GenerateText(tt.args.max); got != tt.want {
				t.Errorf("TrigramMapStore.MakeText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrigramMapStore_Trigrams(t *testing.T) {
	trigramChooser := &store.RandomChooser{}
	cache := store.NewMapTrigramStore(trigramChooser)
	cache.AddTrigram(store.Trigram{"Test", "test1", "test2"})
	cache.AddTrigram(store.Trigram{"Story", "story1", "story2"})
	cache.AddTrigram(store.Trigram{"test1", "test2", "test3"})
	cache.AddTrigram(store.Trigram{"test2", "test3", "test4"})
	cache.AddTrigram(store.Trigram{"test3", "test4", "test5"})

	type fields struct {
		trigrams store.TrigramMap
		mutex    *sync.Mutex
		chooser  store.Chooser
	}
	tests := []struct {
		name   string
		fields fields
		want   store.TrigramMap
	}{
		{
			name: "return the cache's trigrams",
			fields: fields{
				trigrams: cache.Trigrams(),
				mutex:    &sync.Mutex{},
				chooser:  trigramChooser,
			},
			want: cache.Trigrams(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cache.Trigrams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrigramMapStore.Trigrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
