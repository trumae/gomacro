// this file was generated by gomacro command: import "encoding/binary"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "encoding/binary"
	. "reflect"
)

func Package_encoding_binary() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"BigEndian":      ValueOf(&pkg.BigEndian).Elem(),
			"LittleEndian":   ValueOf(&pkg.LittleEndian).Elem(),
			"MaxVarintLen16": ValueOf(pkg.MaxVarintLen16),
			"MaxVarintLen32": ValueOf(pkg.MaxVarintLen32),
			"MaxVarintLen64": ValueOf(pkg.MaxVarintLen64),
			"PutUvarint":     ValueOf(pkg.PutUvarint),
			"PutVarint":      ValueOf(pkg.PutVarint),
			"Read":           ValueOf(pkg.Read),
			"ReadUvarint":    ValueOf(pkg.ReadUvarint),
			"ReadVarint":     ValueOf(pkg.ReadVarint),
			"Size":           ValueOf(pkg.Size),
			"Uvarint":        ValueOf(pkg.Uvarint),
			"Varint":         ValueOf(pkg.Varint),
			"Write":          ValueOf(pkg.Write),
		}, map[string]Type{
			"ByteOrder": TypeOf((*pkg.ByteOrder)(nil)).Elem(),
		}
}

func init() {
	binds, types := Package_encoding_binary()
	Binds["encoding/binary"] = binds
	Types["encoding/binary"] = types
}