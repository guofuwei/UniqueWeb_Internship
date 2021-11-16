package LRU

import (
	"container/list"
	"time"
)

type LRU struct {
	MaxCount int
	CurCount int
	TheList  *list.List
	Cache    map[string]*list.Element
}

type Node struct {
	key        string
	filename   string
	expiretime int64
	code       string
}

var LruCache *LRU

func init() {
	LruCache = CreateLRU(10)
}

func CreateLRU(maxSize int) *LRU {
	return &LRU{
		MaxCount: maxSize,
		CurCount: 0,
		TheList:  list.New(),
		Cache:    make(map[string]*list.Element),
	}
}

func (lru *LRU) Get(key string) (code string) {
	if ele, ok := lru.Cache[key]; ok {
		lru.TheList.MoveToFront(ele)
		if ele.Value.(*Node).expiretime < time.Now().Unix() {
			//若已过期则直接删除
			lru.TheList.Remove(ele)
			delete(lru.Cache, ele.Value.(*Node).key)
			lru.CurCount--
			return ""
		}
		return ele.Value.(*Node).code
	}
	return ""
}

func (lru *LRU) Set(key string, filename string, expiretime int64, code string) (status bool) {
	//尝试查询
	if ele, ok := lru.Cache[key]; ok {
		//查询成功就更新
		lru.TheList.MoveToFront(ele)
		node := ele.Value.(*Node)
		node.filename = filename
		node.expiretime = expiretime
		node.code = code
	} else {
		//查询不成功时进行新建节点
		newnode := &Node{key, filename, expiretime, code}
		lru.Cache[key] = lru.TheList.PushFront(newnode)
		lru.CurCount++
	}
	//开始进行lru淘汰算法
	if lru.CurCount > lru.MaxCount {
		lru.DeleteOldest()
	}
	return true
}

func (lru *LRU) DeleteOldest() {
	ele := lru.TheList.Back()
	lru.TheList.Remove(ele)
	delete(lru.Cache, ele.Value.(*Node).key)
	lru.CurCount--
}
