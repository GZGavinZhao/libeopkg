package shared

type Delta struct {
	ReleaseFrom int `xml:"releaseFrom,attr"`
	PackageURI  string
	PackageSize string
	PackageHash string
}
