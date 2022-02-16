#include<Windows.h>
#include "_cgo_export.h"

LONG NTAPI Firstcontinuehandler(PEXCEPTION_POINTERS pExcepInfo)
{
    return CustomFunc(pExcepInfo);
}

void break0()
{
   __debugbreak();
}

LPVOID setptr(){
	return &Firstcontinuehandler;
}
