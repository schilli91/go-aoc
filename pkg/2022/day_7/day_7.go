package day_7

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/schilli91/go-aoc/pkg/common"
)

func Run() {
	day := 7
	fmt.Printf("Day %d of AOC 2022.\n", day)

	input := common.OpenInputFile(fmt.Sprintf("pkg/2022/day_%d/puzzle_input.txt", day))
	defer input.Close()

	lines := common.ReadLinesString(input)
	given := NewFolder("dir /")
	given.populateFileTree(lines)
	v := given.sumFolderSizes()
	fmt.Printf("The sum of folders* sizes is %d\n", v)
}

type Folder struct {
	name    string
	parent  *Folder
	folders map[string]*Folder
	files   []*File
	size    int
}

type File struct {
	name string
	size int
}

func NewFolder(line string) *Folder {
	if line[:3] != "dir" {
		log.Fatalf("invalid folder line, got %s", line)
	}
	f := &Folder{
		name: line[4:],
	}
	f.folders = make(map[string]*Folder)
	return f
}

func (f *Folder) AddFolder(child *Folder) {
	f.folders[child.name] = child
	child.parent = f
	//f.size += child.size
}

func (f *Folder) AddFile(file *File) {
	f.files = append(f.files, file)
	f.size += file.size
}

func NewFile(line string) *File {
	p := strings.Split(line, " ")
	s, err := strconv.Atoi(p[0])
	if err != nil {
		log.Fatalf("invalid file size (can't convert), got line: %s", line)
	}
	f := &File{
		name: p[1],
		size: s,
	}
	return f
}

func (f *Folder) populateFileTree(terminalLines []string) int {
	defer func() {
		// calc size once folder is populated
		for _, subFolder := range f.folders {
			f.size += subFolder.size
		}
	}()

	c := 0
	for i := 0; i < len(terminalLines); i++ {
		l := terminalLines[i]
		if l == "$ cd .." {
			c++
			return c
		}
		if l == "$ cd /" {
			c++
			continue
		}
		if l == "$ ls" {
			c++
			continue
		}
		if l[:4] == "$ cd" {
			numFiles := f.folders[l[5:]].populateFileTree(terminalLines[i+1:])

			i += numFiles
			c += numFiles + 1 // don't forget to count the current line
			continue
		}
		if l[:3] == "dir" {
			f.AddFolder(NewFolder(l))
			c++
			continue
		}
		// Expecting file line now
		n := NewFile(l)
		f.AddFile(n)
		c++
	}
	return c
}

func (f Folder) sumFolderSizes() int {
	sum := 0
	if f.size <= 100000 {
		sum += f.size
	}

	for _, i := range f.folders {
		sum += i.sumFolderSizes()
	}

	return sum
}
