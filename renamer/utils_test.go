package renamer

import (
	"os"
	"testing"
)

func TestOsRename(t *testing.T) {
	tmpPath := "/tmp"
	sourcePath := tmpPath + "/testFile"
	//destPath := "/Onedrive/owen/testFile"
	destPath := tmpPath + "/testFile1"
	if f, err := os.Create(sourcePath); err != nil {
		t.Fatal(err.Error())
	} else {
		f.Write([]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'})
	}
	if err := OsRename(sourcePath, destPath, ""); err != nil {
		t.Fatal(err.Error())
	}
	defer func() {
		if err := os.Remove(destPath); err != nil {
			println(err.Error())
		}
	}()
}
