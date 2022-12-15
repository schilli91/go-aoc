package day_7

import (
	"reflect"
	"testing"

	"github.com/schilli91/go-aoc/pkg/common"
)

func TestNewFolder(t *testing.T) {
	given := "dir a"
	got := NewFolder(given)
	want := &Folder{name: "a", folders: make(map[string]*Folder)}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong folder, \ngot \t%v, \nwant \t%v\n", got, want)
	}
}

func TestNewFile(t *testing.T) {
	given := "14848514 b.txt"
	got := NewFile(given)
	want := &File{name: "b.txt", size: 14848514}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong file, \ngot \t%v, \nwant \t%v\n", got, want)
	}
}

func TestPopulateFileTree(t *testing.T) {
	t.Run("Partial test data", func(t *testing.T) {
		given := []string{
			"$ cd /",
			"$ ls",
			"dir a",
			"14848514 b.txt",
			"8504156 c.dat",
			"dir d",
		}
		want := NewFolder("dir /")
		a := NewFolder("dir a")
		want.AddFolder(a)
		b := NewFile("14848514 b.txt")
		want.AddFile(b)
		c := NewFile("8504156 c.dat")
		want.AddFile(c)
		d := NewFolder("dir d")
		want.AddFolder(d)
		got := NewFolder("dir /")
		got.populateFileTree(given)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wronge file tree, \ngot \t%v \nwant \t%v", got, want)
		}
	})
	t.Run("Full test data", func(t *testing.T) {
		want := NewFolder("dir /")
		a := NewFolder("dir a")
		want.AddFolder(a)
		b := NewFile("14848514 b.txt")
		want.AddFile(b)
		c := NewFile("8504156 c.dat")
		want.AddFile(c)
		d := NewFolder("dir d")
		want.AddFolder(d)
		e := NewFolder("dir e")
		a.AddFolder(e)
		f := NewFile("29116 f")
		a.AddFile(f)
		g := NewFile("2557 g")
		a.AddFile(g)
		h := NewFile("62596 h.lst")
		a.AddFile(h)
		i := NewFile("584 i")
		e.AddFile(i)
		j := NewFile("4060174 j")
		d.AddFile(j)
		dl := NewFile("8033020 d.log")
		d.AddFile(dl)
		de := NewFile("5626152 d.ext")
		d.AddFile(de)
		k := NewFile("7214296 k")
		d.AddFile(k)

		input := common.OpenInputFile("test_data.txt")
		defer input.Close()
		given := common.ReadLinesString(input)
		got := NewFolder("dir /")
		got.populateFileTree(given)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wrong file tree, \ngot \t%v \nwant \t%v", got, want)
		}
	})
}

func TestPopulateFileTreeSize(t *testing.T) {
	input := common.OpenInputFile("test_data.txt")
	defer input.Close()
	given := common.ReadLinesString(input)
	got := NewFolder("dir /")
	got.populateFileTree(given)
	want := 48381165
	if got.size != want {
		t.Errorf("wrong size for /, \ngot \t%d \nwant \t%d", got.size, want)
	}
}

func TestSumFolderSizes(t *testing.T) {
	input := common.OpenInputFile("test_data.txt")
	defer input.Close()
	lines := common.ReadLinesString(input)
	given := NewFolder("dir /")
	given.populateFileTree(lines)
	got := given.sumFolderSizes()
	want := 95437
	if got != want {
		t.Errorf("wrong sum of sizes, \ngot \t%d \nwant \t%d", got, want)
	}
}
