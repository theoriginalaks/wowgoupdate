﻿package main

import "regexp"

const (
	buildAddonErr         string = "Addon could not be built"
	containerInitErr      string = "Container could not be initialized"
	containerGetAddonsErr string = "Container could not get addon"
	containerLoadFileErr  string = "Container could not load data file"
	fileutilIterErr       string = "No match."
)

const (
	critical = iota
	warning
	info
)

const (
	userAgent string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36"
	curseURL  string = "https://www.curseforge.com/wow/addons/"
)

const (
	productFile = `C:\ProgramData\Battle.net\Agent\product.db`
	saveFile    = "data.json"
)

var reDownloadURL = regexp.MustCompile(`download\/([0-9]+)\"`)
var reVersionNum = regexp.MustCompile(`\<span\sclass\=\"table\_\_content file\_\_name full\"\>([A-z0-9.-]+)\<`)
var reAddonFolderPath = regexp.MustCompile(`([A-z]\:\/)(Program\sFiles\s\(x86\)[\/])?(([A-z\s]+[\/]))*(World of Warcraft)[.]*`)
var reAddonName = regexp.MustCompile(`Title:(?:(?:\|[A-z0-9]{9}|[r])?([A-z0-9\s\.]+)(?:\|[A-z0-9]{9})?([A-z0-9\s\.]+)?)+`)
var reAlphaNum = regexp.MustCompile(`[a-zA-Z0-9]+`)
var reNotAlphaNum = regexp.MustCompile(`[^A-z^0-9]+`)
var reNotASCII = regexp.MustCompile(`[^[:ascii:]]`)
var reBrTag = regexp.MustCompile(`\<br\>|\<br\/\>`)
var reIsExt = regexp.MustCompile(`\.[A-z]+`)
