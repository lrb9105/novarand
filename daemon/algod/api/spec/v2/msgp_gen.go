package v2

// Code generated by github.com/algorand/msgp DO NOT EDIT.

import (
	"github.com/algorand/msgp/msgp"

	"github.com/Orca18/go-novarand/data/basics"
)

// The following msgp objects are implemented in this file:
// AccountApplicationModel
//            |-----> (*) MarshalMsg
//            |-----> (*) CanMarshalMsg
//            |-----> (*) UnmarshalMsg
//            |-----> (*) CanUnmarshalMsg
//            |-----> (*) Msgsize
//            |-----> (*) MsgIsZero
//
// AccountAssetModel
//         |-----> (*) MarshalMsg
//         |-----> (*) CanMarshalMsg
//         |-----> (*) UnmarshalMsg
//         |-----> (*) CanUnmarshalMsg
//         |-----> (*) Msgsize
//         |-----> (*) MsgIsZero
//

// MarshalMsg implements msgp.Marshaler
func (z *AccountApplicationModel) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(2)
	var zb0001Mask uint8 /* 3 bits */
	if (*z).AppLocalState == nil {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).AppParams == nil {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "app-local-state"
			o = append(o, 0xaf, 0x61, 0x70, 0x70, 0x2d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x65)
			if (*z).AppLocalState == nil {
				o = msgp.AppendNil(o)
			} else {
				o = (*z).AppLocalState.MarshalMsg(o)
			}
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "app-params"
			o = append(o, 0xaa, 0x61, 0x70, 0x70, 0x2d, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73)
			if (*z).AppParams == nil {
				o = msgp.AppendNil(o)
			} else {
				o = (*z).AppParams.MarshalMsg(o)
			}
		}
	}
	return
}

func (_ *AccountApplicationModel) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*AccountApplicationModel)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AccountApplicationModel) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 int
	var zb0002 bool
	zb0001, zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0001, zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0001 > 0 {
			zb0001--
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				(*z).AppLocalState = nil
			} else {
				if (*z).AppLocalState == nil {
					(*z).AppLocalState = new(basics.AppLocalState)
				}
				bts, err = (*z).AppLocalState.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "AppLocalState")
					return
				}
			}
		}
		if zb0001 > 0 {
			zb0001--
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				(*z).AppParams = nil
			} else {
				if (*z).AppParams == nil {
					(*z).AppParams = new(basics.AppParams)
				}
				bts, err = (*z).AppParams.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "AppParams")
					return
				}
			}
		}
		if zb0001 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0001)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0002 {
			(*z) = AccountApplicationModel{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "app-local-state":
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					(*z).AppLocalState = nil
				} else {
					if (*z).AppLocalState == nil {
						(*z).AppLocalState = new(basics.AppLocalState)
					}
					bts, err = (*z).AppLocalState.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "AppLocalState")
						return
					}
				}
			case "app-params":
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					(*z).AppParams = nil
				} else {
					if (*z).AppParams == nil {
						(*z).AppParams = new(basics.AppParams)
					}
					bts, err = (*z).AppParams.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "AppParams")
						return
					}
				}
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
		}
	}
	o = bts
	return
}

func (_ *AccountApplicationModel) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*AccountApplicationModel)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AccountApplicationModel) Msgsize() (s int) {
	s = 1 + 16
	if (*z).AppLocalState == nil {
		s += msgp.NilSize
	} else {
		s += (*z).AppLocalState.Msgsize()
	}
	s += 11
	if (*z).AppParams == nil {
		s += msgp.NilSize
	} else {
		s += (*z).AppParams.Msgsize()
	}
	return
}

// MsgIsZero returns whether this is a zero value
func (z *AccountApplicationModel) MsgIsZero() bool {
	return ((*z).AppLocalState == nil) && ((*z).AppParams == nil)
}

// MarshalMsg implements msgp.Marshaler
func (z *AccountAssetModel) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(2)
	var zb0001Mask uint8 /* 3 bits */
	if (*z).AssetHolding == nil {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).AssetParams == nil {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "asset-holding"
			o = append(o, 0xad, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2d, 0x68, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67)
			if (*z).AssetHolding == nil {
				o = msgp.AppendNil(o)
			} else {
				o = (*z).AssetHolding.MarshalMsg(o)
			}
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "asset-params"
			o = append(o, 0xac, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2d, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73)
			if (*z).AssetParams == nil {
				o = msgp.AppendNil(o)
			} else {
				o = (*z).AssetParams.MarshalMsg(o)
			}
		}
	}
	return
}

func (_ *AccountAssetModel) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*AccountAssetModel)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AccountAssetModel) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 int
	var zb0002 bool
	zb0001, zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0001, zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0001 > 0 {
			zb0001--
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				(*z).AssetParams = nil
			} else {
				if (*z).AssetParams == nil {
					(*z).AssetParams = new(basics.AssetParams)
				}
				bts, err = (*z).AssetParams.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "AssetParams")
					return
				}
			}
		}
		if zb0001 > 0 {
			zb0001--
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				(*z).AssetHolding = nil
			} else {
				if (*z).AssetHolding == nil {
					(*z).AssetHolding = new(basics.AssetHolding)
				}
				bts, err = (*z).AssetHolding.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "AssetHolding")
					return
				}
			}
		}
		if zb0001 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0001)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0002 {
			(*z) = AccountAssetModel{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "asset-params":
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					(*z).AssetParams = nil
				} else {
					if (*z).AssetParams == nil {
						(*z).AssetParams = new(basics.AssetParams)
					}
					bts, err = (*z).AssetParams.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "AssetParams")
						return
					}
				}
			case "asset-holding":
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					(*z).AssetHolding = nil
				} else {
					if (*z).AssetHolding == nil {
						(*z).AssetHolding = new(basics.AssetHolding)
					}
					bts, err = (*z).AssetHolding.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "AssetHolding")
						return
					}
				}
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
		}
	}
	o = bts
	return
}

func (_ *AccountAssetModel) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*AccountAssetModel)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AccountAssetModel) Msgsize() (s int) {
	s = 1 + 13
	if (*z).AssetParams == nil {
		s += msgp.NilSize
	} else {
		s += (*z).AssetParams.Msgsize()
	}
	s += 14
	if (*z).AssetHolding == nil {
		s += msgp.NilSize
	} else {
		s += (*z).AssetHolding.Msgsize()
	}
	return
}

// MsgIsZero returns whether this is a zero value
func (z *AccountAssetModel) MsgIsZero() bool {
	return ((*z).AssetParams == nil) && ((*z).AssetHolding == nil)
}
