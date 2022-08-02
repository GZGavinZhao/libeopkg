package shared

type Delta struct {
	ReleaseFrom int `xml:"releaseFrom,attr"`
	PackageURI  string
	PackageSize int64
	PackageHash string
}
