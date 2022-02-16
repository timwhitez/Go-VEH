package main

/*
#include<Windows.h>
*/
import "C"
import (
	"fmt"
)

const(
	_EXCEPTION_CONTINUE_EXECUTION = -0x1
	_EXCEPTION_CONTINUE_SEARCH    = 0x0
	_EXCEPTION_CONTINUE_HANDLE    = 0x1
)

//export CustomFunc
func CustomFunc(pExcepInfo C.PEXCEPTION_POINTERS)int32 {
	println("VEH Hooked!")
	vehHook(pExcepInfo)
	return _EXCEPTION_CONTINUE_SEARCH

}

func vehHook(pExcepInfo C.PEXCEPTION_POINTERS){
	fmt.Printf("ExceptionCode = 0x%x\n", pExcepInfo.ExceptionRecord.ExceptionCode)
	fmt.Printf("ExceptionAddress = 0x%x\n",pExcepInfo.ExceptionRecord.ExceptionAddress)
	fmt.Printf("Rsp = 0x%x\n",pExcepInfo.ContextRecord.Rsp)
	fmt.Printf("Rip = 0x%x\n",pExcepInfo.ContextRecord.Rip)
}