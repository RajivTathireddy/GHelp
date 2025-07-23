package structure

import (
	"github.com/google/go-cmp/cmp"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateDirPath(t *testing.T) {
	t.Parallel()
	t.Run("creates Project directory1", func(t *testing.T) {
		temppath := t.TempDir()
		userpath := "test1"
		gotdirpath, gotcmdpath := createPath(temppath, userpath, true)
		wantdirpath, wantcmdpath := filepath.Join(temppath, userpath), filepath.Join(temppath, userpath, "cmd")
		assertCorrectPath(t, wantdirpath, gotdirpath)
		assertCorrectPath(t, wantcmdpath, gotcmdpath)
	})
	t.Run("creates Project directory2", func(t *testing.T) {
		temppath := t.TempDir()
		userpath := "test2"
		gotdirpath, gotcmdpath := createPath(temppath, userpath, false)
		wantdirpath, wantcmdpath := filepath.Join(temppath, userpath), filepath.Join(temppath, userpath, "pkg")
		assertCorrectPath(t, wantdirpath, gotdirpath)
		assertCorrectPath(t, wantcmdpath, gotcmdpath)
	})

}
func TestCreateProjectDir(t *testing.T) {
	t.Run("test dir creation1", func(t *testing.T) {
		tmpDir := t.TempDir()
		newDir := filepath.Join(tmpDir, "testdir")

		err := createProjectDir(newDir)
		if err != nil {
			t.Fatal(err)
		}

		info, err := os.Stat(newDir)
		if err != nil {
			t.Fatal(err)
		}
		expected := os.FileMode(0751)
		actual := info.Mode().Perm()
		if diff := cmp.Diff(expected, actual); diff != "" {
			t.Fatal(diff)
		}
	})
}

func TestCreateProjectFiles(t *testing.T){
	t.Run("test . files", func(t *testing.T) {
		tmpDir := t.TempDir()

		err := createProjectFiles(tmpDir)
		if err != nil {
			t.Fatal(err)
		}

		_, err = os.Stat(filepath.Join(tmpDir,".gitignore"))
		if err != nil {
			t.Fatal(err)
		}
		_, err = os.Stat(filepath.Join(tmpDir,".env"))
		if err != nil {
			t.Fatal(err)
		}
	})
}

func assertCorrectPath(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
