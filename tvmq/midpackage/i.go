package midpackage

import (
	"code.tvmining.com/tvplay/tvmq/allmap"
)

type MapReader interface {
	Reader() map[string]allmap.AMap
}

func GetMap(reader MapReader) map[string]allmap.AMap {
	return reader.Reader()
}
