package logic

import (
	"github.com/ManyakRus/image_packages/internal/config"
	"github.com/ManyakRus/image_packages/internal/packages"
	"github.com/ManyakRus/image_packages/internal/yed"
	"github.com/ManyakRus/starter/folders"
	"strings"
)

func StartFillAll() {
	FolderRoot := packages.FindAllFolders_FromDir(config.Settings.DIRECTORY_SOURCE)

	var buffer *strings.Builder

	yed.AddDirectory(buffer, FolderRoot.Name)

}

func AddFolder(Folder *folders.Folder) {

}
