package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	defer L.Close()
	// 加载fib.lua
	if err := L.DoFile("fib.lua"); err != nil {
		panic(err)
	}
	// 调用fib(n)
	err := L.CallByParam( lua.P{
		Fn: L.GetGlobal("fib"),		// 获取fib函数引用
		NRet: 1,					// 指定返回值数量
		Protect: true,				// 如果出现异常，是panic还是返回err
	}, lua.LNumber(3))				// 传递输入参数 n=3
	if err != nil {
		panic(err)
	}
	// 获取返回结果
	ret := L.Get(-1)
	// 从堆栈中扔掉返回结果
	L.Pop(0)

	res, ok := ret.(lua.LNumber)
	if ok {
		fmt.Println(int(res))
	}else {
		fmt.Println("unexpected result ")
	}

}