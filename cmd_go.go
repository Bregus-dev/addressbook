package main

import (
	"errors"
	"fmt"
	generator "github.com/Bookshelf-Writer/SimpleGenerator"
	"github.com/spf13/cobra"
	"path/filepath"
)

// //////////////////////////////////////////////////////////////

var writeGoCmd = &cobra.Command{
	Use:   "write-go",
	Short: "Go-file generation",
	RunE: func(cmd *cobra.Command, args []string) error {
		if CmdFromFilePath == "" {
			return errors.New("you must specify a file using " + cyan("--fileFrom"))
		}

		pathFileRead := CheckFilePath(CmdFromFilePath)
		if pathFileRead != FilePathValid {
			return errors.New(cyan("--fileFrom") + ":\t" + pathFileRead.String())
		}

		CmdFromFilePath, err := filepath.Abs(CmdFromFilePath)
		if err != nil {
			return errors.New(cyan("--fileFrom") + ":\t" + err.Error())
		}

		//

		if CmdToFilePath == "" {
			return errors.New("you must specify a file using " + cyan("--fileTo"))
		}

		pathFileWrite := CheckFilePath(CmdToFilePath)
		if pathFileWrite != FilePathValid && pathFileWrite != FilePathIsDir && pathFileWrite != FilePathValidDir {
			return errors.New(cyan("--fileTo") + ":\t" + pathFileWrite.String())
		}

		CmdToFilePath, err := filepath.Abs(CmdToFilePath)
		if err != nil {
			return errors.New(cyan("--fileTo") + ":\t" + err.Error())
		}

		//

		if CmdPackage == "" {
			if pathFileWrite == FilePathIsDir {
				CmdPackage = CmdToFilePath
			} else {
				CmdPackage = filepath.Dir(CmdToFilePath)
			}
			CmdPackage = filepath.Base(CmdPackage)
		}

		if pathFileWrite == FilePathIsDir {
			CmdToFilePath = filepath.Join(CmdToFilePath, "addressbook.go")
		}

		//

		obj, err := ParseYml(CmdFromFilePath)
		if err != nil {
			return err
		}

		return writeGo(obj, CmdToFilePath, CmdPackage)
	},
}

func writeGo(obj *IronAdressbookObj, pathToFile, packageName string) error {
	goGen := generator.NewFilePathName(filepath.Dir(pathToFile), packageName)

	goGen.AddObjStruct(obj.Devices).SeparatorX6().LN()
	goGen.AddObjValue("BregusDev", obj.Devices)

	err := goGen.Save(filepath.Base(pathToFile))
	if err == nil {
		fmt.Printf("The file was created successfully. \n\tDir: %s \n\tFile: %s\n", cyan(filepath.Dir(pathToFile)), green(filepath.Base(pathToFile)))
	}
	return err
}

// //////////////

func init() {
	writeGoCmd.Flags().StringVar(&CmdFromFilePath, "fileFrom", "", "Path to the file")
	writeGoCmd.Flags().StringVar(&CmdToFilePath, "fileTo", "", "Path to the file where the generated Go-code will be written. You can specify a directory.")

	writeGoCmd.Flags().StringVar(&CmdPackage, "package", "", "Package name. The default is the directory name.")

	rootCmd.AddCommand(writeGoCmd)
}
