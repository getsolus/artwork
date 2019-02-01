// This is responsible for generating our solus-wallpapers.xml.in file
// Effectively, this script will read the files in our backgrounds directory and generate an XML file so we don't have to maintain it

package main

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sort"
)

// customNames is a list of wallpapers to custom wallpaper names
var customNames = map[string]string {
	"Athlone-Shannon-1920x1080": "Athlone, Shannon",
	"HermiteCrab2560x1600": "Hermite Crab (2560x1600)",
	"IMG_9656": "IMG_9656.png",
	"IMG_9710": "IMG_9710.png",
	"LegendaryBarrels": "Legendary Barrels, Tullamore",
	"PubSide": "Pub Side, Tullamore",
	"ShannonClonmacnoise": "Clonmacnoise, Shannon",
	"TullamoreGrandCanal": "Grand Canal, Tullamore",
	"matador-rocks": "El Matador Rocks",
}

var workDir string

// GetFiles will get our files and return their file names
func GetFiles() ([]string, error) {
	var files []string
	var getFilesErr error

	if backgroundsFile, backgroundsOpenErr := os.Open(filepath.Join(workDir, "backgrounds")); backgroundsOpenErr == nil { // Open up backgrounds
		if fileInfos, readDirErr := backgroundsFile.Readdir(-1); readDirErr == nil { // Get the FileInfo objects in backgrounds
			for _, fileInfo := range fileInfos { // For each fileInfo
				if !fileInfo.IsDir() { // Not a directory
					name := fileInfo.Name()
					fileExt := strings.TrimPrefix(filepath.Ext(name), ".") // Get the extension of this file
					fileExt = strings.ToLower(fileExt)

					if (fileExt == "jpg") || (fileExt == "jpeg") || (fileExt == "png") { // JPG and PNG are the only acceptable file types
						files = append(files, name)
					}
				}
			}
		} else { // If we failed to read the directory
			getFilesErr = readDirErr
		}
	} else { // If we failed to open the directory
		getFilesErr = backgroundsOpenErr
	}

	return files, getFilesErr
}

func main() {
	var getWdErr error
	workDir, getWdErr = os.Getwd()

	if getWdErr != nil {
		fmt.Printf("Failed to get current working directory: %s\nExiting.", getWdErr)
		os.Exit(1)
	}

	if files, getFilesErr := GetFiles(); getFilesErr == nil { // Get our files
		if len(files) != 0 { // If we have files listed
			sort.Strings(files) // Ensure they are alphabetically sorted

			var wallpapers []string

			for _, file := range files { // For each file
				name := file // Set our name of this entry to file
				name = strings.Replace(name, filepath.Ext(name), "", -1) // Remove the extension

				if fileLegacyName, legacyNameExists := customNames[name]; legacyNameExists { // If we have a legacy name for this
					name = fileLegacyName // Maintain backwards-compatibility
				} else { // No exception for this
					name = strings.Replace(name, "Excl_", "", -1) // Remove Excl from name. Example: Excl_Lonely_Railway -> Lonely_Railway
					name = strings.ToTitle(name[:1]) + name[1:]// Titalize
				}

				if !strings.HasPrefix(name, "IMG_") { // If this doesn't start with IMG_, which should retain _
					name = strings.Replace(name, "_", " ", -1) // Replace _ with whitespace
				}

				fmt.Printf("Adding file: %s\n", file)
				filePath := fmt.Sprintf("@prefix@/share/backgrounds/solus/%s", file)
				
				wallpaper := Wallpaper{
					Deleted: false,
					Name: name, // Our name
					FileName: filePath, // Full path
					Options: "zoom",
					Pcolor: "#000000",
					Scolor: "#000000",
					ShadeType: "solid",
				}

				if xmlBytes, xmlMarshalErr := xml.MarshalIndent(wallpaper, "  ", "  "); xmlMarshalErr == nil { // Encode our wallpapers
					xmlString := string(xmlBytes[:])
					xmlString = strings.Replace(xmlString, "Wallpaper", "wallpaper",  -1) // Enforce lowercase
					wallpapers = append(wallpapers, xmlString) // Convert to string, add to wallpapers string array
				} else {
					fmt.Printf("Failed to generate our XML: %s\n", xmlMarshalErr)
				}
			}

			xmlFile := strings.Join(wallpapers, "\n") // Merge all our strings into a single one, separated by newlines
			xmlFile = fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE wallpapers SYSTEM "gnome-wp-list.dtd">
<wallpapers>
%s
</wallpapers>
			`, xmlFile) // Wrap in wallpapers tag
			
			ioutil.WriteFile("backgrounds/solus-wallpapers.xml.in", []byte(xmlFile), 0644)
		} else { // No files were found
			fmt.Printf("Did not find any image files in backgrounds/")
		}
	} else {
		fmt.Printf("Failed to get files in backgrounds directory: %s\n", getFilesErr)
	}
}