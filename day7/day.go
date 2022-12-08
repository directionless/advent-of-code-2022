package day7

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	maxDepth = 100 // Sorry.
)

type dayHandler struct {
	// Note that one of these is an array, and the other a slice. Fun times
	fileSizes             map[string]int
	directorySizesDoubled map[string]int
	cwd                   []string
	dirs                  []string
}

func New() *dayHandler {
	h := &dayHandler{
		directorySizesDoubled: make(map[string]int),
		fileSizes:             make(map[string]int),
		cwd:                   []string{},
	}

	return h

}

func (h *dayHandler) Consume(lineBytes []byte) error {
	if len(lineBytes) == 0 {
		return nil
	}

	// Go maps cannot use slices as keys, so we're just going to use filepath to marshall to strings

	line := string(lineBytes)

	switch {
	case strings.HasPrefix(line, "$ cd /"):
		h.cwd = []string{"/"}
		return nil

	case strings.HasPrefix(line, "$ cd .."):
		if len(h.cwd) < 1 {
			return fmt.Errorf("cannot cd .. from: %v", h.cwd)
		}
		h.cwd = h.cwd[:len(h.cwd)-1]

	case strings.HasPrefix(line, "$ cd "):
		dir := strings.TrimPrefix(line, "$ cd ")
		h.cwd = append(h.cwd, dir)

		//fmt.Printf("now in: %v", h.cwd)

	case strings.HasPrefix(line, "$ ls"):
		// Just an FYI, we don't need to handle it
		// could add error handling, but :shrug:
		return nil

	case strings.HasPrefix(line, "dir "):
		// This is ls output showing we're looking at a dir
		return nil

	default: // must be a file!
		components := strings.SplitN(line, " ", 2)
		size, err := strconv.Atoi(components[0])
		if err != nil {
			return fmt.Errorf("cannot parse size %s: %w", components[0], err)
		}

		filename := filepath.Join(append(h.cwd, components[1])...)
		h.fileSizes[filename] = size
		h.addSizeToDirs(filename, size, 1)

	}
	return nil
}

func (h *dayHandler) AnswerPart1() int {
	// The elves want:
	//
	// Find all of the directories with a total size of at most 100000. What is the sum
	// of the total sizes of those directories?
	//
	// Kind of a weird thing to want. But those elves...

	total := 0
	for _, size := range h.directorySizesDoubled {
		if size > 100000 {
			continue
		}
		//fmt.Printf("Adding %d for %s\n", size, pathJoin(dir))
		total += size

	}

	return total
}

func (h *dayHandler) AnswerPart2() int {
	return 0
}

func (h *dayHandler) DirSize(path string) int {
	//h.addUpSubdirs()

	return h.directorySizesDoubled[path]
}

func (h *dayHandler) addSizeToDirs(p string, size int, reverseDepth int) {
	dir := filepath.Dir(p)
	h.directorySizesDoubled[dir] += size * reverseDepth

	// If dir is _not_ `/`, do it again
	if dir != "/" && dir != "" {
		h.addSizeToDirs(dir, size, reverseDepth)
	}
}

/*
func (h *dayHandler) sortDirs() {
	if len(h.dirs) > 0 {
		return
	}
	for dir := range h.directorySizes {
		h.dirs = append(h.dirs, dir)
	}

	sort.Slice(h.dirs, func(i, j int) bool {
		return pathLen(h.dirs[i]) > pathLen(h.dirs[j])
	})
}
*/

/*
func (h *dayHandler) addUpFiles() {
	if len(h.directorySizes) > 0 {
		return
	}

	// We have files. Let's add them to the directories!

	for f, size := range h.fileSizes {
		dir := filepath.Dir(f)
		h.directorySizes[dir] += size

		for {

		}

		size := h.directorySizes[dir]
		h.fullDirectorySizes[dir] += size

		if pathLen(dir) == 0 {
			fmt.Println("found /")
			continue
		}

		parent := parentname(dir)
		h.fullDirectorySizes[parent] += size

		fmt.Printf("Add %d to %s. Now %d (For %s)\n", size, pathJoin(parent), h.fullDirectorySizes[parent], pathJoin(dir))

	}

	h.subdirsAdded = true
}

*/

func (h *dayHandler) DumpTree() {
	//h.sortDirs()
	//h.addUpSubdirs()

	for f, sz := range h.fileSizes {
		fmt.Printf("file: %s: %d\n", f, sz)
	}

	for d, sz := range h.directorySizesDoubled {
		fmt.Printf("dir:  %s: %d\n", d, sz)
	}

}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}

func pathLen(path [maxDepth]string) int {
	size := 0
	for _, dir := range path {
		if dir == "" {
			return size
		}

		size++
	}
	return size
}

func pathJoin(path [maxDepth]string) string {
	var sb strings.Builder
	for _, dir := range path {
		sb.WriteString("/")

		if dir == "" {
			break
		}
		sb.WriteString(dir)
	}

	return sb.String()
}

func parentname(path [maxDepth]string) [maxDepth]string {
	var parent [maxDepth]string
	for i, dir := range path {
		if dir == "" {
			if i > 0 {
				parent[i-1] = ""
			}
			return parent
		}
		parent[i] = dir
	}
	return parent
}
