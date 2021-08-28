package markdown

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type NoteTitle struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type NoteTitles []NoteTitle

type NotesList struct {
	Name         string    `json:"name"`
	ModifiedTime time.Time `json:"modified_time"`
}

type NotesLists []NotesList

func ListNotes() (string, error) {
	notesDir := os.Getenv("NOTES_DIR")

	files, err := fetchIndex(notesDir)
	if err != nil {
		return files, err
	}

	return files, err
}

func fetchIndex(root string) (string, error) {
	filesInfo, err := IOReadDir(root)
	if err != nil {
		return "", err
	}

	j, err := json.Marshal(filesInfo)
	if err != nil {
		return "", err
	}

	return string(j), err
}

func IOReadDir(root string) ([]NotesList, error) {
	var files []NotesList
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		file_info := NotesList{
			Name:         file.Name(),
			ModifiedTime: file.ModTime(),
		}

		files = append(files, file_info)
	}

	return files, nil
}
