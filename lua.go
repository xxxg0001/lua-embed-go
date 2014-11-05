package lua

/*
#cgo LDFLAGS: -L ../ -llua -ldl -lm 
#include <stdlib.h>
#include "../lualib/src/lua.h"
#include "../lualib/src/lualib.h"
#include "../lualib/src/lauxlib.h"
*/
import "C"
import "fmt"
import "unsafe"

func Run(file string) {
	L := C.lua_open(500) /* opens Lua */
	C.lua_baselibopen(L)
	C.lua_iolibopen(L)
	C.lua_strlibopen(L)
	C.lua_mathlibopen(L)
	defer C.lua_close(L)
	cs := C.CString(file)
	defer C.free(unsafe.Pointer(cs))
	if C.int(C.lua_dofile(L, cs)) != 0 {
		fmt.Println("error!\n")
		return
	}
	return
}
