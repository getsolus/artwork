package main

// Wallpaper is our wallpaper struct
type Wallpaper struct {
	Deleted bool `xml:"deleted,attr"`
	Name string `xml:"name"`
	FileName string `xml:"filename"`
	Options string `xml:"options"`
	Pcolor string `xml:"pcolor"`
	Scolor string `xml:"scolor"`
	ShadeType string `xml:"shade_type"`
}