package test_extension_file

import "path"

func getExtensionFromFile(filename string) string {
	return path.Ext(filename)
}