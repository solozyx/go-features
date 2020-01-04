package main

import (
	"errors"
	"fmt"
	"time"
)

// 对象池
func main() {
	pool := NewObjPool(10)
	for i := 0; i < 11; i++ {
		if obj, err := pool.GetObj(time.Second * 1); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("%T\n", obj)
			if err := pool.ReleaseObj(obj); err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

type ReusableObj struct{}

type ObjPool struct {
	bufChan chan *ReusableObj
}

func (pool *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case obj := <-pool.bufChan:
		return obj, nil
	case <-time.After(timeout):
		return nil, errors.New("get ObjPool obj timeout")
	}
}

// 释放对象
func (pool *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case pool.bufChan <- obj:
		return nil
	default:
		return errors.New("chan of ObjPool is full")
	}
}

// 初始化对象池
func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}
