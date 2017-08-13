package set

import "sync"

//golang no set
//using map implementation 
//lock sync.RWMutex for multi-ops
//flattened to store the items in a list

var pool = sync.Pool{}

type Set struct{
	items	map[interface{}]struct{}
	lock	sync.RWMutex
	flattened	[]interface{}
}

func (set *Set) Add(items ...interface{}){
	set.lock.Lock()
	defer set.lock.Unlock()

	set.flattened = nil
	for _,item := range items{
		// init the struct ?
		set.items[item] = struct{}{}
	}
}

func (set *Set) Remove(items ...interface{}){
	set.lock.Lock()
	defer set.lock.Unlock()

	set.flattened = nil
	for _,item := range items{
		//set.items=append(set.items,item)
		delete(set.items,item)
	}
}

func (set *Set) Exists(item interface{}) bool{
	set.lock.RLock()
	// exists true , else false
	_,ok := set.items[item]
	set.lock.RUnlock()
	return ok
}


func (set *Set) Flatten() []interface{}{
	set.lock.Lock()
	defer set.lock.Unlock()

	if set.flattened != nil{
		return set.flattened
	}
	set.flattened = make([]interface{},0,len(set.items))
	for item := range set.items{
		set.flattened = append(set.flattened,item)
	}
	return set.flattened
}

func (set *Set) Len() int64{
	set.lock.RLock()
	size := int64(len(set.items))
	set.lock.RUnlock()
	return size
}

func (set *Set) Clear(){
	set.lock.Lock()
	set.items = map[interface{}]struct{}{}
	set.lock.Unlock()
}

func (set *Set) All(items ...interface{}) bool{
	set.lock.RLock()
	defer set.lock.RUnlock()

	for _,item := range set.items{
		if _,ok := set.items[item];  !ok{
			return false
		}
	}
	return true
}

func (set *Set) Dispose(){
	set.lock.Lock()
	defer set.lock.Unlock()

	for k:= range set.items{
		delete(set.items,k)
	}
	for i:=0;i<len(set.flattened);i++{
		set.flattened[i]=nil
	}
	set.flattened = set.flattened[:0]
	pool.Put(set)
}

func New(items ...interface{}) *Set{
	set := pool.Get.(*Set)
	for _,item := range items{
		set.items[item] = struct{}{}
	}
	if len(items) > 0 {
		set.flattened =nil
	}
	return set
}

func init(){
	pool.New = func()interface{}{
		return &Set{
			items:make(map[interface{}]struct{},10),
		}
	}
}