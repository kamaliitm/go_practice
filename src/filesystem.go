package main

import (
	"errors"
	"strings"
)

const PATH_DELIMITER = "/"

type FSNode struct {
	name     string
	path     string
	data     string
	children map[string]*FSNode
	isFile   bool
}

// type File struct {
// 	FSObject
// 	data string
// }

// type Directory struct {
// 	FSObject
// 	files []File
// }

type FileSystem interface {
	mkdir(path string) error
	// createFile(path string) error
	// list(path string) ([]FSNode, error)
	// remove(path string) error
	// move(path string) error
	// cd(path string)
}

type UserSession struct {
	userId      string
	fileSystem  *FileSystem
	currentPath string
}

func NewUserSession(userId string, fileSystem *FileSystem) *UserSession {
	return &UserSession{
		userId:      userId,
		fileSystem:  fileSystem,
		currentPath: PATH_DELIMITER,
	}
}

type InMemoryFileSystem struct {
	root           *FSNode
	userCurrentDir map[string]string
}

func NewInMemoryFileSystem() FileSystem {
	return &InMemoryFileSystem{
		root:           &FSNode{name: "", path: PATH_DELIMITER},
		userCurrentDir: map[string]string{},
	}
}

func (fs *InMemoryFileSystem) mkdir(path string) error {

	pathTokens := strings.Split(path, PATH_DELIMITER)
	pathLen := len(pathTokens)
	if pathLen == 0 || pathLen == 1 {
		return errors.New("root directory already exists")
	}

	pathTokens = pathTokens[1:]

	curr := fs.root
	for _, currDir := range pathTokens {
		currDirNode, ok := curr.children[currDir]
		if !ok {
			currDirNode = &FSNode{
				name:     currDir,
				children: map[string]*FSNode{},
			}

			curr.children[currDir] = currDirNode
		}

		curr = currDirNode
	}

	return nil
}
