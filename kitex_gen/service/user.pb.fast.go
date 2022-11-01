// Code generated by FastPB v0.0.1. DO NOT EDIT.

package service

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

func (x *GetUserListReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserListReq[number], err)
}

func (x *GetUserListReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Page, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetUserListReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.PageSize, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetUserListReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetUserListResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserListResp[number], err)
}

func (x *GetUserListResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v UserData
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Data = append(x.Data, &v)
	return offset, nil
}

func (x *UserData) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UserData[number], err)
}

func (x *UserData) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UserData) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserData) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Sex, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetUserByIdReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserByIdReq[number], err)
}

func (x *GetUserByIdReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetUserListReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *GetUserListReq) fastWriteField1(buf []byte) (offset int) {
	if x.Page == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.Page)
	return offset
}

func (x *GetUserListReq) fastWriteField2(buf []byte) (offset int) {
	if x.PageSize == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.PageSize)
	return offset
}

func (x *GetUserListReq) fastWriteField3(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.Username)
	return offset
}

func (x *GetUserListResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetUserListResp) fastWriteField1(buf []byte) (offset int) {
	if x.Data == nil {
		return offset
	}
	for i := range x.Data {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.Data[i])
	}
	return offset
}

func (x *UserData) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *UserData) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.Id)
	return offset
}

func (x *UserData) fastWriteField2(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.Username)
	return offset
}

func (x *UserData) fastWriteField3(buf []byte) (offset int) {
	if x.Sex == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.Sex)
	return offset
}

func (x *GetUserByIdReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetUserByIdReq) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.Id)
	return offset
}

func (x *GetUserListReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *GetUserListReq) sizeField1() (n int) {
	if x.Page == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.Page)
	return n
}

func (x *GetUserListReq) sizeField2() (n int) {
	if x.PageSize == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.PageSize)
	return n
}

func (x *GetUserListReq) sizeField3() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(3, x.Username)
	return n
}

func (x *GetUserListResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetUserListResp) sizeField1() (n int) {
	if x.Data == nil {
		return n
	}
	for i := range x.Data {
		n += fastpb.SizeMessage(1, x.Data[i])
	}
	return n
}

func (x *UserData) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *UserData) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.Id)
	return n
}

func (x *UserData) sizeField2() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(2, x.Username)
	return n
}

func (x *UserData) sizeField3() (n int) {
	if x.Sex == "" {
		return n
	}
	n += fastpb.SizeString(3, x.Sex)
	return n
}

func (x *GetUserByIdReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetUserByIdReq) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.Id)
	return n
}

var fieldIDToName_GetUserListReq = map[int32]string{
	1: "Page",
	2: "PageSize",
	3: "Username",
}

var fieldIDToName_GetUserListResp = map[int32]string{
	1: "Data",
}

var fieldIDToName_UserData = map[int32]string{
	1: "Id",
	2: "Username",
	3: "Sex",
}

var fieldIDToName_GetUserByIdReq = map[int32]string{
	1: "Id",
}
