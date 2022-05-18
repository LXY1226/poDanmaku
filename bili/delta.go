package bili

// Package bili GENERATED CODE by copilot - DO NOT EDIT

const (
	RoomId = iota
	ShortId
	LiveStatus
	Attention
	Online
	Description
	AreaId
	AreaName
	ParentAreaId
	ParentAreaName
	Background
	Title
	UserCover
	IsStrictRoom
	Tags
	IsAnchor
)

func (info *RoomInfo) GenDelta(from *RoomInfo) []byte {

	builder := Encoder{}
	if info.RoomId != from.RoomId {
		builder.AppendI32(RoomId, info.RoomId)
	}
	if info.ShortId != from.ShortId {
		builder.AppendI32(ShortId, info.ShortId)
	}
	if info.LiveStatus != from.LiveStatus {
		builder.AppendI8(LiveStatus, info.LiveStatus)
	}
	if info.Attention != from.Attention {
		builder.AppendI64(Attention, info.Attention)
	}
	if info.Online != from.Online {
		builder.AppendI64(Online, info.Online)
	}
	if info.Description != from.Description {
		builder.AppendString(Description, info.Description)
	}
	if info.AreaId != from.AreaId {
		builder.AppendI16(AreaId, info.AreaId)
	}
	if info.AreaName != from.AreaName {
		builder.AppendString(AreaName, info.AreaName)
	}
	if info.ParentAreaId != from.ParentAreaId {
		builder.AppendI16(ParentAreaId, info.ParentAreaId)
	}
	if info.ParentAreaName != from.ParentAreaName {
		builder.AppendString(ParentAreaName, info.ParentAreaName)
	}
	if info.Background != from.Background {
		builder.AppendString(Background, info.Background)
	}
	if info.Title != from.Title {
		builder.AppendString(Title, info.Title)
	}
	if info.UserCover != from.UserCover {
		builder.AppendString(UserCover, info.UserCover)
	}
	if info.IsStrictRoom != from.IsStrictRoom {
		builder.AppendBool(IsStrictRoom, info.IsStrictRoom)
	}
	if info.Tags != from.Tags {
		builder.AppendString(Tags, info.Tags)
	}
	if info.IsAnchor != from.IsAnchor {
		builder.AppendI8(IsAnchor, info.IsAnchor)
	}
	return builder.EncodeAppend(nil)
}
func (info *RoomInfo) UpdateDelta(delta []byte) {
	decoder := Decoder(delta)
	tag, typ := decoder.Next()
	if typ == TypeEndOfData {
		return
	}
	if tag == RoomId {
		info.RoomId = decoder.AsI32()
		tag, typ = decoder.Next()
		if typ == TypeEndOfData {
			return
		}
	}
	if tag == ShortId {
		info.ShortId = decoder.AsI32()
		tag, typ = decoder.Next()
		if typ == TypeEndOfData {
			return
		}
	}
	if tag == LiveStatus {
		info.LiveStatus = decoder.AsI8()
		tag, typ = decoder.Next()
		if typ == TypeEndOfData {
			return
		}
	}
	if tag == Attention {
		info.Attention = decoder.AsI64()
		tag, typ = decoder.Next()
		if typ == TypeEndOfData {
			return
		}
	}
	if tag == Online {
		info.Online = decoder.AsI64()
		tag, typ = decoder.Next()
		if typ == TypeEndOfData {
			return
		}
	}
	if tag == Description {
		info.Description = decoder.AsString(typ)
		tag, typ = decoder.Next()
		if typ == TypeEndOfData {
			return
		}
	}
	if tag == AreaId {
		info.AreaId = decoder.AsI16()
		tag, typ = decoder.Next()
		if typ == TypeEndOfData {
			return
		}
	}
	if tag == AreaName {
		info.AreaName = decoder.AsString(typ)
		tag, typ = decoder.Next()
		if typ == TypeEndOfData {
			return
		}
	}
	if tag == ParentAreaId {
		info.ParentAreaId = decoder.AsI16()
		tag, typ = decoder.Next()
		if typ == TypeEndOfData {
			return
		}
	}
	if tag == ParentAreaName {
		info.ParentAreaName = decoder.AsString(typ)
		tag, typ = decoder.Next()
		if typ == TypeEndOfData {
			return
		}
	}
	if tag == Background {
		info.Background = decoder.AsString(typ)
		tag, typ = decoder.Next()
		if typ == TypeEndOfData {
			return
		}
	}
	if tag == Title {
		info.Title = decoder.AsString(typ)
		tag, typ = decoder.Next()
		if typ == TypeEndOfData {
			return
		}
	}
	if tag == UserCover {
		info.UserCover = decoder.AsString(typ)
		tag, typ = decoder.Next()
		if typ == TypeEndOfData {
			return
		}
	}
	if tag == IsStrictRoom {
		info.IsStrictRoom = decoder.AsBool(typ)
		tag, typ = decoder.Next()
		if typ == TypeEndOfData {
			return
		}
	}
	if tag == Tags {
		info.Tags = decoder.AsString(typ)
		tag, typ = decoder.Next()
		if typ == TypeEndOfData {
			return
		}
	}
	if tag == IsAnchor {
		info.IsAnchor = decoder.AsI8()
		tag, typ = decoder.Next()
		if typ == TypeEndOfData {
			return
		}
	}
}

const (
	UID = iota
	Uname
	Face
)

func (info *UserInfo) GenDelta(from *UserInfo) []byte {
	builder := Encoder{}
	if info.UID != from.UID {
		builder.AppendI32(UID, info.UID)
	}
	if info.Uname != from.Uname {
		builder.AppendString(Uname, info.Uname)
	}
	if info.Face != from.Face {
		builder.AppendString(Face, info.Face)
	}
	return builder.EncodeAppend(nil)
}

func (info *UserInfo) UpdateDelta(delta []byte) {
	decoder := Decoder(delta)
	tag, typ := decoder.Next()
	if typ == TypeEndOfData {
		return
	}
	if tag == UID {
		info.UID = decoder.AsI32()
		tag, typ = decoder.Next()
		if typ == TypeEndOfData {
			return
		}
	}
	if tag == Uname {
		info.Uname = decoder.AsString(typ)
		tag, typ = decoder.Next()
		if typ == TypeEndOfData {
			return
		}
	}
	if tag == Face {
		info.Face = decoder.AsString(typ)
		tag, typ = decoder.Next()
		if typ == TypeEndOfData {
			return
		}
	}
}
