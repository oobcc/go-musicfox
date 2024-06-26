// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package core

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
)

const SignatureMseStreamSource string = "rc(Windows.Media.Core.MseStreamSource;{b0b4198d-02f4-4923-88dd-81bc3f360ffa})"

type MseStreamSource struct {
	ole.IUnknown
}

func NewMseStreamSource() (*MseStreamSource, error) {
	inspectable, err := ole.RoActivateInstance("Windows.Media.Core.MseStreamSource")
	if err != nil {
		return nil, err
	}
	return (*MseStreamSource)(unsafe.Pointer(inspectable)), nil
}

func (impl *MseStreamSource) AddOpened(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseStreamSource))
	defer itf.Release()
	v := (*iMseStreamSource)(unsafe.Pointer(itf))
	return v.AddOpened(handler)
}

func (impl *MseStreamSource) RemoveOpened(token foundation.EventRegistrationToken) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseStreamSource))
	defer itf.Release()
	v := (*iMseStreamSource)(unsafe.Pointer(itf))
	return v.RemoveOpened(token)
}

func (impl *MseStreamSource) AddEnded(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseStreamSource))
	defer itf.Release()
	v := (*iMseStreamSource)(unsafe.Pointer(itf))
	return v.AddEnded(handler)
}

func (impl *MseStreamSource) RemoveEnded(token foundation.EventRegistrationToken) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseStreamSource))
	defer itf.Release()
	v := (*iMseStreamSource)(unsafe.Pointer(itf))
	return v.RemoveEnded(token)
}

func (impl *MseStreamSource) AddClosed(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseStreamSource))
	defer itf.Release()
	v := (*iMseStreamSource)(unsafe.Pointer(itf))
	return v.AddClosed(handler)
}

func (impl *MseStreamSource) RemoveClosed(token foundation.EventRegistrationToken) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseStreamSource))
	defer itf.Release()
	v := (*iMseStreamSource)(unsafe.Pointer(itf))
	return v.RemoveClosed(token)
}

func (impl *MseStreamSource) GetSourceBuffers() (*MseSourceBufferList, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseStreamSource))
	defer itf.Release()
	v := (*iMseStreamSource)(unsafe.Pointer(itf))
	return v.GetSourceBuffers()
}

func (impl *MseStreamSource) GetActiveSourceBuffers() (*MseSourceBufferList, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseStreamSource))
	defer itf.Release()
	v := (*iMseStreamSource)(unsafe.Pointer(itf))
	return v.GetActiveSourceBuffers()
}

func (impl *MseStreamSource) GetReadyState() (MseReadyState, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseStreamSource))
	defer itf.Release()
	v := (*iMseStreamSource)(unsafe.Pointer(itf))
	return v.GetReadyState()
}

func (impl *MseStreamSource) GetDuration() (*foundation.IReference, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseStreamSource))
	defer itf.Release()
	v := (*iMseStreamSource)(unsafe.Pointer(itf))
	return v.GetDuration()
}

func (impl *MseStreamSource) SetDuration(value *foundation.IReference) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseStreamSource))
	defer itf.Release()
	v := (*iMseStreamSource)(unsafe.Pointer(itf))
	return v.SetDuration(value)
}

func (impl *MseStreamSource) AddSourceBuffer(mimeType string) (*MseSourceBuffer, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseStreamSource))
	defer itf.Release()
	v := (*iMseStreamSource)(unsafe.Pointer(itf))
	return v.AddSourceBuffer(mimeType)
}

func (impl *MseStreamSource) RemoveSourceBuffer(buffer *MseSourceBuffer) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseStreamSource))
	defer itf.Release()
	v := (*iMseStreamSource)(unsafe.Pointer(itf))
	return v.RemoveSourceBuffer(buffer)
}

func (impl *MseStreamSource) EndOfStream(status MseEndOfStreamStatus) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseStreamSource))
	defer itf.Release()
	v := (*iMseStreamSource)(unsafe.Pointer(itf))
	return v.EndOfStream(status)
}

func (impl *MseStreamSource) GetLiveSeekableRange() (*foundation.IReference, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseStreamSource2))
	defer itf.Release()
	v := (*iMseStreamSource2)(unsafe.Pointer(itf))
	return v.GetLiveSeekableRange()
}

func (impl *MseStreamSource) SetLiveSeekableRange(value *foundation.IReference) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseStreamSource2))
	defer itf.Release()
	v := (*iMseStreamSource2)(unsafe.Pointer(itf))
	return v.SetLiveSeekableRange(value)
}

const GUIDiMseStreamSource string = "b0b4198d-02f4-4923-88dd-81bc3f360ffa"
const SignatureiMseStreamSource string = "{b0b4198d-02f4-4923-88dd-81bc3f360ffa}"

type iMseStreamSource struct {
	ole.IInspectable
}

type iMseStreamSourceVtbl struct {
	ole.IInspectableVtbl

	AddOpened              uintptr
	RemoveOpened           uintptr
	AddEnded               uintptr
	RemoveEnded            uintptr
	AddClosed              uintptr
	RemoveClosed           uintptr
	GetSourceBuffers       uintptr
	GetActiveSourceBuffers uintptr
	GetReadyState          uintptr
	GetDuration            uintptr
	SetDuration            uintptr
	AddSourceBuffer        uintptr
	RemoveSourceBuffer     uintptr
	EndOfStream            uintptr
}

func (v *iMseStreamSource) VTable() *iMseStreamSourceVtbl {
	return (*iMseStreamSourceVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iMseStreamSource) AddOpened(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	var out foundation.EventRegistrationToken
	hr, _, _ := syscall.SyscallN(
		v.VTable().AddOpened,
		uintptr(unsafe.Pointer(v)),       // this
		uintptr(unsafe.Pointer(handler)), // in foundation.TypedEventHandler
		uintptr(unsafe.Pointer(&out)),    // out foundation.EventRegistrationToken
	)

	if hr != 0 {
		return foundation.EventRegistrationToken{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMseStreamSource) RemoveOpened(token foundation.EventRegistrationToken) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().RemoveOpened,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&token)), // in foundation.EventRegistrationToken
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iMseStreamSource) AddEnded(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	var out foundation.EventRegistrationToken
	hr, _, _ := syscall.SyscallN(
		v.VTable().AddEnded,
		uintptr(unsafe.Pointer(v)),       // this
		uintptr(unsafe.Pointer(handler)), // in foundation.TypedEventHandler
		uintptr(unsafe.Pointer(&out)),    // out foundation.EventRegistrationToken
	)

	if hr != 0 {
		return foundation.EventRegistrationToken{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMseStreamSource) RemoveEnded(token foundation.EventRegistrationToken) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().RemoveEnded,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&token)), // in foundation.EventRegistrationToken
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iMseStreamSource) AddClosed(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	var out foundation.EventRegistrationToken
	hr, _, _ := syscall.SyscallN(
		v.VTable().AddClosed,
		uintptr(unsafe.Pointer(v)),       // this
		uintptr(unsafe.Pointer(handler)), // in foundation.TypedEventHandler
		uintptr(unsafe.Pointer(&out)),    // out foundation.EventRegistrationToken
	)

	if hr != 0 {
		return foundation.EventRegistrationToken{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMseStreamSource) RemoveClosed(token foundation.EventRegistrationToken) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().RemoveClosed,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&token)), // in foundation.EventRegistrationToken
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iMseStreamSource) GetSourceBuffers() (*MseSourceBufferList, error) {
	var out *MseSourceBufferList
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetSourceBuffers,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out MseSourceBufferList
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMseStreamSource) GetActiveSourceBuffers() (*MseSourceBufferList, error) {
	var out *MseSourceBufferList
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetActiveSourceBuffers,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out MseSourceBufferList
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMseStreamSource) GetReadyState() (MseReadyState, error) {
	var out MseReadyState
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetReadyState,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out MseReadyState
	)

	if hr != 0 {
		return MseReadyStateClosed, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMseStreamSource) GetDuration() (*foundation.IReference, error) {
	var out *foundation.IReference
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetDuration,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.IReference
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMseStreamSource) SetDuration(value *foundation.IReference) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetDuration,
		uintptr(unsafe.Pointer(v)),     // this
		uintptr(unsafe.Pointer(value)), // in foundation.IReference
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iMseStreamSource) AddSourceBuffer(mimeType string) (*MseSourceBuffer, error) {
	var out *MseSourceBuffer
	mimeTypeHStr, err := ole.NewHString(mimeType)
	if err != nil {
		return nil, err
	}
	hr, _, _ := syscall.SyscallN(
		v.VTable().AddSourceBuffer,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(mimeTypeHStr),         // in string
		uintptr(unsafe.Pointer(&out)), // out MseSourceBuffer
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMseStreamSource) RemoveSourceBuffer(buffer *MseSourceBuffer) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().RemoveSourceBuffer,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(buffer)), // in MseSourceBuffer
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iMseStreamSource) EndOfStream(status MseEndOfStreamStatus) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().EndOfStream,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(status),            // in MseEndOfStreamStatus
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

const GUIDiMseStreamSource2 string = "66f57d37-f9e7-418a-9cde-a020e956552b"
const SignatureiMseStreamSource2 string = "{66f57d37-f9e7-418a-9cde-a020e956552b}"

type iMseStreamSource2 struct {
	ole.IInspectable
}

type iMseStreamSource2Vtbl struct {
	ole.IInspectableVtbl

	GetLiveSeekableRange uintptr
	SetLiveSeekableRange uintptr
}

func (v *iMseStreamSource2) VTable() *iMseStreamSource2Vtbl {
	return (*iMseStreamSource2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iMseStreamSource2) GetLiveSeekableRange() (*foundation.IReference, error) {
	var out *foundation.IReference
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetLiveSeekableRange,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.IReference
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMseStreamSource2) SetLiveSeekableRange(value *foundation.IReference) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetLiveSeekableRange,
		uintptr(unsafe.Pointer(v)),     // this
		uintptr(unsafe.Pointer(value)), // in foundation.IReference
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

const GUIDiMseStreamSourceStatics string = "465c679d-d570-43ce-ba21-0bff5f3fbd0a"
const SignatureiMseStreamSourceStatics string = "{465c679d-d570-43ce-ba21-0bff5f3fbd0a}"

type iMseStreamSourceStatics struct {
	ole.IInspectable
}

type iMseStreamSourceStaticsVtbl struct {
	ole.IInspectableVtbl

	MseStreamSourceIsContentTypeSupported uintptr
}

func (v *iMseStreamSourceStatics) VTable() *iMseStreamSourceStaticsVtbl {
	return (*iMseStreamSourceStaticsVtbl)(unsafe.Pointer(v.RawVTable))
}

func MseStreamSourceIsContentTypeSupported(contentType string) (bool, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Media.Core.MseStreamSource", ole.NewGUID(GUIDiMseStreamSourceStatics))
	if err != nil {
		return false, err
	}
	v := (*iMseStreamSourceStatics)(unsafe.Pointer(inspectable))

	var out bool
	contentTypeHStr, err := ole.NewHString(contentType)
	if err != nil {
		return false, err
	}
	hr, _, _ := syscall.SyscallN(
		v.VTable().MseStreamSourceIsContentTypeSupported,
		0,                             // this is a static func, so there's no this
		uintptr(contentTypeHStr),      // in string
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}
