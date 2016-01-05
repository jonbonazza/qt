package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QFileSystemWatcher struct {
	QObject
}

type QFileSystemWatcher_ITF interface {
	QObject_ITF
	QFileSystemWatcher_PTR() *QFileSystemWatcher
}

func PointerFromQFileSystemWatcher(ptr QFileSystemWatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileSystemWatcher_PTR().Pointer()
	}
	return nil
}

func NewQFileSystemWatcherFromPointer(ptr unsafe.Pointer) *QFileSystemWatcher {
	var n = new(QFileSystemWatcher)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QFileSystemWatcher_") {
		n.SetObjectName("QFileSystemWatcher_" + qt.Identifier())
	}
	return n
}

func (ptr *QFileSystemWatcher) QFileSystemWatcher_PTR() *QFileSystemWatcher {
	return ptr
}

func (ptr *QFileSystemWatcher) Directories() []string {
	defer qt.Recovering("QFileSystemWatcher::directories")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSystemWatcher_Directories(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QFileSystemWatcher) Files() []string {
	defer qt.Recovering("QFileSystemWatcher::files")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSystemWatcher_Files(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func NewQFileSystemWatcher(parent QObject_ITF) *QFileSystemWatcher {
	defer qt.Recovering("QFileSystemWatcher::QFileSystemWatcher")

	return NewQFileSystemWatcherFromPointer(C.QFileSystemWatcher_NewQFileSystemWatcher(PointerFromQObject(parent)))
}

func NewQFileSystemWatcher2(paths []string, parent QObject_ITF) *QFileSystemWatcher {
	defer qt.Recovering("QFileSystemWatcher::QFileSystemWatcher")

	return NewQFileSystemWatcherFromPointer(C.QFileSystemWatcher_NewQFileSystemWatcher2(C.CString(strings.Join(paths, ",,,")), PointerFromQObject(parent)))
}

func (ptr *QFileSystemWatcher) AddPath(path string) bool {
	defer qt.Recovering("QFileSystemWatcher::addPath")

	if ptr.Pointer() != nil {
		return C.QFileSystemWatcher_AddPath(ptr.Pointer(), C.CString(path)) != 0
	}
	return false
}

func (ptr *QFileSystemWatcher) AddPaths(paths []string) []string {
	defer qt.Recovering("QFileSystemWatcher::addPaths")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSystemWatcher_AddPaths(ptr.Pointer(), C.CString(strings.Join(paths, ",,,")))), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QFileSystemWatcher) ConnectDirectoryChanged(f func(path string)) {
	defer qt.Recovering("connect QFileSystemWatcher::directoryChanged")

	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_ConnectDirectoryChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "directoryChanged", f)
	}
}

func (ptr *QFileSystemWatcher) DisconnectDirectoryChanged() {
	defer qt.Recovering("disconnect QFileSystemWatcher::directoryChanged")

	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_DisconnectDirectoryChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "directoryChanged")
	}
}

//export callbackQFileSystemWatcherDirectoryChanged
func callbackQFileSystemWatcherDirectoryChanged(ptr unsafe.Pointer, ptrName *C.char, path *C.char) {
	defer qt.Recovering("callback QFileSystemWatcher::directoryChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "directoryChanged"); signal != nil {
		signal.(func(string))(C.GoString(path))
	}

}

func (ptr *QFileSystemWatcher) ConnectFileChanged(f func(path string)) {
	defer qt.Recovering("connect QFileSystemWatcher::fileChanged")

	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_ConnectFileChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fileChanged", f)
	}
}

func (ptr *QFileSystemWatcher) DisconnectFileChanged() {
	defer qt.Recovering("disconnect QFileSystemWatcher::fileChanged")

	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_DisconnectFileChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fileChanged")
	}
}

//export callbackQFileSystemWatcherFileChanged
func callbackQFileSystemWatcherFileChanged(ptr unsafe.Pointer, ptrName *C.char, path *C.char) {
	defer qt.Recovering("callback QFileSystemWatcher::fileChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "fileChanged"); signal != nil {
		signal.(func(string))(C.GoString(path))
	}

}

func (ptr *QFileSystemWatcher) RemovePath(path string) bool {
	defer qt.Recovering("QFileSystemWatcher::removePath")

	if ptr.Pointer() != nil {
		return C.QFileSystemWatcher_RemovePath(ptr.Pointer(), C.CString(path)) != 0
	}
	return false
}

func (ptr *QFileSystemWatcher) RemovePaths(paths []string) []string {
	defer qt.Recovering("QFileSystemWatcher::removePaths")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSystemWatcher_RemovePaths(ptr.Pointer(), C.CString(strings.Join(paths, ",,,")))), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QFileSystemWatcher) DestroyQFileSystemWatcher() {
	defer qt.Recovering("QFileSystemWatcher::~QFileSystemWatcher")

	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_DestroyQFileSystemWatcher(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QFileSystemWatcher) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QFileSystemWatcher::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QFileSystemWatcher) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QFileSystemWatcher::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQFileSystemWatcherTimerEvent
func callbackQFileSystemWatcherTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFileSystemWatcher::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQFileSystemWatcherFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QFileSystemWatcher) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QFileSystemWatcher::timerEvent")

	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QFileSystemWatcher) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QFileSystemWatcher::timerEvent")

	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QFileSystemWatcher) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QFileSystemWatcher::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QFileSystemWatcher) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QFileSystemWatcher::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQFileSystemWatcherChildEvent
func callbackQFileSystemWatcherChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFileSystemWatcher::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQFileSystemWatcherFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QFileSystemWatcher) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QFileSystemWatcher::childEvent")

	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QFileSystemWatcher) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QFileSystemWatcher::childEvent")

	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QFileSystemWatcher) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QFileSystemWatcher::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QFileSystemWatcher) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QFileSystemWatcher::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQFileSystemWatcherCustomEvent
func callbackQFileSystemWatcherCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFileSystemWatcher::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQFileSystemWatcherFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QFileSystemWatcher) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QFileSystemWatcher::customEvent")

	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QFileSystemWatcher) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QFileSystemWatcher::customEvent")

	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}