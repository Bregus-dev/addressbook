package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// //////////////////////////////////////////////////////////////

var writeJsonCmd = &cobra.Command{
	Use:   "write-json",
	Short: "JSON-file generation",
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

		if pathFileWrite == FilePathIsDir {
			CmdToFilePath = filepath.Join(CmdToFilePath, "addressbook.json")
		}

		//

		obj, err := ParseYml(CmdFromFilePath)
		if err != nil {
			return err
		}

		return writeJson(obj, CmdToFilePath)
	},
}

func writeJson(obj *IronAdressbookObj, pathToFile string) error {
	file, err := os.Create(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(obj.Devices, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err == nil {
		fmt.Printf("The file was created successfully. \n\tDir: %s \n\tFile: %s\n", cyan(filepath.Dir(pathToFile)), green(filepath.Base(pathToFile)))
	}
	return err
}

// //////////////

func init() {
	writeJsonCmd.Flags().StringVar(&CmdFromFilePath, "fileFrom", "", "Path to the file")
	writeJsonCmd.Flags().StringVar(&CmdToFilePath, "fileTo", "", "Path to the file where the generated JSON-content will be written. You can specify a directory.")

	rootCmd.AddCommand(writeJsonCmd)
}
