package scriptfinder

import (
	"errors"
	"github.com/mcabezas/migrations/commons"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type fileBackend struct {
	path string
}

func NewFileBackend(path string) commons.ScriptFinderBackend {
	return &fileBackend{path: path}
}

func (f *fileBackend) GetRequiredDeltas() ([]*commons.Delta, error) {
	ups := map[string][]byte{}
	downs := map[string][]byte{}
	err := filepath.Walk(f.path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		splits := strings.Split(file.Name(), "_")
		versionID := splits[0]
		if strings.HasSuffix(file.Name(), "up.sql") {
			ups[versionID] = bytes
			return nil
		}
		if strings.HasSuffix(file.Name(), "down.sql") {
			downs[versionID] = bytes
			return nil
		}
		return nil
	})

	if len(ups) != len(downs) {
		return nil, errors.New("THERE are not the same quantity of ups and down scripts")
	}

	var results []*commons.Delta
	for id, up := range ups {
		down, found := downs[id]
		if !found {
			return nil, errors.New("VERSION " + id + " does not have down script")
		}
		results = append(results, &commons.Delta{
			ID:         id,
			UpScript:   up,
			DownScript: down,
		})
	}
	return results, err
}