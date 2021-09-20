package protocol

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// EncodeMsg implements msgp.Encodable
func (z *Message) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 4
	err = en.Append(0x94)
	if err != nil {
		return
	}
	err = en.WriteString(z.Tag)
	if err != nil {
		err = msgp.WrapError(err, "Tag")
		return
	}
	err = en.WriteInt64(z.Timestamp)
	if err != nil {
		err = msgp.WrapError(err, "Timestamp")
		return
	}
	err = z.Record.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Record")
		return
	}
	if z.Options == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Options.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Options")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Message) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 4
	o = append(o, 0x94)
	o = msgp.AppendString(o, z.Tag)
	o = msgp.AppendInt64(o, z.Timestamp)
	o, err = z.Record.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Record")
		return
	}
	if z.Options == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Options.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Options")
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *MessageExt) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 4
	err = en.Append(0x94)
	if err != nil {
		return
	}
	err = en.WriteString(z.Tag)
	if err != nil {
		err = msgp.WrapError(err, "Tag")
		return
	}
	err = en.WriteExtension(&z.Timestamp)
	if err != nil {
		err = msgp.WrapError(err, "Timestamp")
		return
	}
	err = z.Record.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Record")
		return
	}
	if z.Options == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Options.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Options")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MessageExt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 4
	o = append(o, 0x94)
	o = msgp.AppendString(o, z.Tag)
	o, err = msgp.AppendExtension(o, &z.Timestamp)
	if err != nil {
		err = msgp.WrapError(err, "Timestamp")
		return
	}
	o, err = z.Record.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Record")
		return
	}
	if z.Options == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Options.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Options")
			return
		}
	}
	return
}