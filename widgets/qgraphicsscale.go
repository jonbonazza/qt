package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsScale struct {
	QGraphicsTransform
}

type QGraphicsScale_ITF interface {
	QGraphicsTransform_ITF
	QGraphicsScale_PTR() *QGraphicsScale
}

func PointerFromQGraphicsScale(ptr QGraphicsScale_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsScale_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsScaleFromPointer(ptr unsafe.Pointer) *QGraphicsScale {
	var n = new(QGraphicsScale)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsScale_") {
		n.SetObjectName("QGraphicsScale_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsScale) QGraphicsScale_PTR() *QGraphicsScale {
	return ptr
}

func (ptr *QGraphicsScale) SetOrigin(point gui.QVector3D_ITF) {
	defer qt.Recovering("QGraphicsScale::setOrigin")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_SetOrigin(ptr.Pointer(), gui.PointerFromQVector3D(point))
	}
}

func (ptr *QGraphicsScale) SetXScale(v float64) {
	defer qt.Recovering("QGraphicsScale::setXScale")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_SetXScale(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QGraphicsScale) SetYScale(v float64) {
	defer qt.Recovering("QGraphicsScale::setYScale")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_SetYScale(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QGraphicsScale) SetZScale(v float64) {
	defer qt.Recovering("QGraphicsScale::setZScale")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_SetZScale(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QGraphicsScale) XScale() float64 {
	defer qt.Recovering("QGraphicsScale::xScale")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsScale_XScale(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsScale) YScale() float64 {
	defer qt.Recovering("QGraphicsScale::yScale")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsScale_YScale(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsScale) ZScale() float64 {
	defer qt.Recovering("QGraphicsScale::zScale")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsScale_ZScale(ptr.Pointer()))
	}
	return 0
}

func NewQGraphicsScale(parent core.QObject_ITF) *QGraphicsScale {
	defer qt.Recovering("QGraphicsScale::QGraphicsScale")

	return NewQGraphicsScaleFromPointer(C.QGraphicsScale_NewQGraphicsScale(core.PointerFromQObject(parent)))
}

func (ptr *QGraphicsScale) ConnectApplyTo(f func(matrix *gui.QMatrix4x4)) {
	defer qt.Recovering("connect QGraphicsScale::applyTo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "applyTo", f)
	}
}

func (ptr *QGraphicsScale) DisconnectApplyTo() {
	defer qt.Recovering("disconnect QGraphicsScale::applyTo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "applyTo")
	}
}

//export callbackQGraphicsScaleApplyTo
func callbackQGraphicsScaleApplyTo(ptr unsafe.Pointer, ptrName *C.char, matrix unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsScale::applyTo")

	if signal := qt.GetSignal(C.GoString(ptrName), "applyTo"); signal != nil {
		signal.(func(*gui.QMatrix4x4))(gui.NewQMatrix4x4FromPointer(matrix))
	}

}

func (ptr *QGraphicsScale) ApplyTo(matrix gui.QMatrix4x4_ITF) {
	defer qt.Recovering("QGraphicsScale::applyTo")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_ApplyTo(ptr.Pointer(), gui.PointerFromQMatrix4x4(matrix))
	}
}

func (ptr *QGraphicsScale) ApplyToDefault(matrix gui.QMatrix4x4_ITF) {
	defer qt.Recovering("QGraphicsScale::applyTo")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_ApplyToDefault(ptr.Pointer(), gui.PointerFromQMatrix4x4(matrix))
	}
}

func (ptr *QGraphicsScale) ConnectOriginChanged(f func()) {
	defer qt.Recovering("connect QGraphicsScale::originChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectOriginChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "originChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectOriginChanged() {
	defer qt.Recovering("disconnect QGraphicsScale::originChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectOriginChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "originChanged")
	}
}

//export callbackQGraphicsScaleOriginChanged
func callbackQGraphicsScaleOriginChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsScale::originChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "originChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsScale) OriginChanged() {
	defer qt.Recovering("QGraphicsScale::originChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_OriginChanged(ptr.Pointer())
	}
}

func (ptr *QGraphicsScale) ConnectScaleChanged(f func()) {
	defer qt.Recovering("connect QGraphicsScale::scaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectScaleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "scaleChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectScaleChanged() {
	defer qt.Recovering("disconnect QGraphicsScale::scaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectScaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "scaleChanged")
	}
}

//export callbackQGraphicsScaleScaleChanged
func callbackQGraphicsScaleScaleChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsScale::scaleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "scaleChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsScale) ScaleChanged() {
	defer qt.Recovering("QGraphicsScale::scaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_ScaleChanged(ptr.Pointer())
	}
}

func (ptr *QGraphicsScale) ConnectXScaleChanged(f func()) {
	defer qt.Recovering("connect QGraphicsScale::xScaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectXScaleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "xScaleChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectXScaleChanged() {
	defer qt.Recovering("disconnect QGraphicsScale::xScaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectXScaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "xScaleChanged")
	}
}

//export callbackQGraphicsScaleXScaleChanged
func callbackQGraphicsScaleXScaleChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsScale::xScaleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "xScaleChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsScale) XScaleChanged() {
	defer qt.Recovering("QGraphicsScale::xScaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_XScaleChanged(ptr.Pointer())
	}
}

func (ptr *QGraphicsScale) ConnectYScaleChanged(f func()) {
	defer qt.Recovering("connect QGraphicsScale::yScaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectYScaleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "yScaleChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectYScaleChanged() {
	defer qt.Recovering("disconnect QGraphicsScale::yScaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectYScaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "yScaleChanged")
	}
}

//export callbackQGraphicsScaleYScaleChanged
func callbackQGraphicsScaleYScaleChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsScale::yScaleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "yScaleChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsScale) YScaleChanged() {
	defer qt.Recovering("QGraphicsScale::yScaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_YScaleChanged(ptr.Pointer())
	}
}

func (ptr *QGraphicsScale) ConnectZScaleChanged(f func()) {
	defer qt.Recovering("connect QGraphicsScale::zScaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectZScaleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "zScaleChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectZScaleChanged() {
	defer qt.Recovering("disconnect QGraphicsScale::zScaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectZScaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "zScaleChanged")
	}
}

//export callbackQGraphicsScaleZScaleChanged
func callbackQGraphicsScaleZScaleChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsScale::zScaleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "zScaleChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsScale) ZScaleChanged() {
	defer qt.Recovering("QGraphicsScale::zScaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_ZScaleChanged(ptr.Pointer())
	}
}

func (ptr *QGraphicsScale) DestroyQGraphicsScale() {
	defer qt.Recovering("QGraphicsScale::~QGraphicsScale")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_DestroyQGraphicsScale(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsScale) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGraphicsScale::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGraphicsScale) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGraphicsScale::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGraphicsScaleTimerEvent
func callbackQGraphicsScaleTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsScale::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGraphicsScaleFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGraphicsScale) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsScale::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGraphicsScale) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsScale::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGraphicsScale) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGraphicsScale::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGraphicsScale) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGraphicsScale::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGraphicsScaleChildEvent
func callbackQGraphicsScaleChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsScale::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGraphicsScaleFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGraphicsScale) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsScale::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGraphicsScale) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsScale::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGraphicsScale) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsScale::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGraphicsScale) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGraphicsScale::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGraphicsScaleCustomEvent
func callbackQGraphicsScaleCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsScale::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsScaleFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsScale) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsScale::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsScale) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsScale::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}