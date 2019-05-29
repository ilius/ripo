package ripo

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ilius/is"
)

func TestAssertError_OK(t *testing.T) {
	is := is.New(t)
	is.True(AssertError(
		t,
		NewError(Unauthenticated, "not authorized buddy", nil),
		Unauthenticated,
		"not authorized buddy",
	))
}

func TestAssertError_OK_NonRPC(t *testing.T) {
	is := is.New(t)
	is.True(AssertError(
		t,
		errors.New("some strange error"),
		Unknown,
		"some strange error",
	))
}

func TestAssertError_Fail_Nil(t *testing.T) {
	is := is.New(t)
	mtCtrl := gomock.NewController(t)
	defer mtCtrl.Finish()
	mt := NewMockSmallT(mtCtrl)
	mt.EXPECT().Helper()
	mt.EXPECT().Fatalf(
		"got err==nil, expected code=%v msg=%#v",
		InvalidArgument,
		"not authorized buddy",
	)
	is.False(AssertError(
		mt,
		nil,
		InvalidArgument,
		"not authorized buddy",
	))
}

func TestAssertError_Fail_Code(t *testing.T) {
	is := is.New(t)
	mtCtrl := gomock.NewController(t)
	defer mtCtrl.Finish()
	mt := NewMockSmallT(mtCtrl)
	mt.EXPECT().Helper()
	mt.EXPECT().Fatalf(
		"got code=%v in err==%#v, expected code=%v",
		Internal,
		"something bad happened",
		InvalidArgument,
	)
	is.False(AssertError(
		mt,
		NewError(Internal, "something bad happened", nil),
		InvalidArgument,
		"not authorized buddy",
	))
}

func TestAssertError_Fail_Msg(t *testing.T) {
	is := is.New(t)
	mtCtrl := gomock.NewController(t)
	defer mtCtrl.Finish()
	mt := NewMockSmallT(mtCtrl)
	mt.EXPECT().Helper()
	mt.EXPECT().Fatalf(
		"got err.Error()==%#v, expected %#v",
		"param2 is missing",
		"param1 is missing",
	)
	is.False(AssertError(
		mt,
		NewError(InvalidArgument, "param2 is missing", nil),
		InvalidArgument,
		"param1 is missing",
	))
}

func TestAssertErrorFail_NonRPC(t *testing.T) {
	is := is.New(t)
	mtCtrl := gomock.NewController(t)
	defer mtCtrl.Finish()
	mt := NewMockSmallT(mtCtrl)
	mt.EXPECT().Helper()
	mt.EXPECT().Fatalf(
		"got non-rpc err==%#v, expected code=%v",
		"some unknown problem",
		InvalidArgument,
	)
	is.False(AssertError(
		mt,
		errors.New("some unknown problem"),
		InvalidArgument,
		"not authorized buddy",
	))
}
