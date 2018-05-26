package tests

import "path"

func getExtensionFromFile(filename string) string {
	return path.Ext(filename)
}