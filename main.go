package main

/*
extern uintptr_t setptr();
extern void break0();
*/
import "C"
import (
	"syscall"
)

func VEH_proc(){
	//add VEH
	vehA := syscall.NewLazyDLL("kernel32.dll").NewProc("AddVectoredExceptionHandler")
	vehA.Call(1,uintptr(C.setptr()))
}



func main() {
	VEH_proc()
	C.break0()
}
