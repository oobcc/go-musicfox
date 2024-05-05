// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package frames

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

const SignatureMediaFrameSourceInfo string = "rc(Windows.Media.Capture.Frames.MediaFrameSourceInfo;{87bdc9cd-4601-408f-91cf-038318cd0af3})"

type MediaFrameSourceInfo struct {
	ole.IUnknown
}

const GUIDiMediaFrameSourceInfo string = "87bdc9cd-4601-408f-91cf-038318cd0af3"
const SignatureiMediaFrameSourceInfo string = "{87bdc9cd-4601-408f-91cf-038318cd0af3}"

type iMediaFrameSourceInfo struct {
	ole.IInspectable
}

type iMediaFrameSourceInfoVtbl struct {
	ole.IInspectableVtbl

	GetId                uintptr
	GetMediaStreamType   uintptr
	GetSourceKind        uintptr
	GetSourceGroup       uintptr
	GetDeviceInformation uintptr
	GetProperties        uintptr
	GetCoordinateSystem  uintptr
}

func (v *iMediaFrameSourceInfo) VTable() *iMediaFrameSourceInfoVtbl {
	return (*iMediaFrameSourceInfoVtbl)(unsafe.Pointer(v.RawVTable))
}

const GUIDiMediaFrameSourceInfo2 string = "195a7855-6457-42c6-a769-19b65bd32e6e"
const SignatureiMediaFrameSourceInfo2 string = "{195a7855-6457-42c6-a769-19b65bd32e6e}"

type iMediaFrameSourceInfo2 struct {
	ole.IInspectable
}

type iMediaFrameSourceInfo2Vtbl struct {
	ole.IInspectableVtbl

	GetProfileId                    uintptr
	GetVideoProfileMediaDescription uintptr
}

func (v *iMediaFrameSourceInfo2) VTable() *iMediaFrameSourceInfo2Vtbl {
	return (*iMediaFrameSourceInfo2Vtbl)(unsafe.Pointer(v.RawVTable))
}

const GUIDiMediaFrameSourceInfo3 string = "ca824ab6-66ea-5885-a2b6-26c0eeec3c7b"
const SignatureiMediaFrameSourceInfo3 string = "{ca824ab6-66ea-5885-a2b6-26c0eeec3c7b}"

type iMediaFrameSourceInfo3 struct {
	ole.IInspectable
}

type iMediaFrameSourceInfo3Vtbl struct {
	ole.IInspectableVtbl

	GetRelativePanel uintptr
}

func (v *iMediaFrameSourceInfo3) VTable() *iMediaFrameSourceInfo3Vtbl {
	return (*iMediaFrameSourceInfo3Vtbl)(unsafe.Pointer(v.RawVTable))
}