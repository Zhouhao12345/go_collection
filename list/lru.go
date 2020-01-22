package list

import (
	"sync"
)



type LRUNode struct {
	prev *LRUNode
	me map[string]interface{}
	next *LRUNode
}

type LRU struct {
	lock *sync.Mutex
	count int
	dmap map[string]*LRUNode
	first *LRUNode
	last *LRUNode
}

func NewLRU(count int, pairs map[string]interface{},) (*LRU, error) {
	locker := new(sync.Mutex)
	newLRU := &LRU{
		lock: locker,
		count: count,
		dmap: make(map[string]*LRUNode),
	}
	for k,v := range pairs {
		if err := newLRU.Insert(k, v); err != nil {
			return nil, err
		}
	}
	return newLRU, nil
}


func (l *LRU) Get(token string) (interface{}, error)  {
	defer l.lock.Unlock()
	l.lock.Lock()
	return l.get(token)
}
func (l *LRU) get(token string) (interface{}, error) {
	var (
		node *LRUNode
		ok bool
		value interface{}
		err error
	)
	if node, ok = l.dmap[token]; !ok {
		return node, nil
	}
	token = node.me["token"].(string)
	value = node.me["value"]
	if err = l.insert(token, value); err != nil {
		return nil, err
	}
	return value, nil
}

func (l *LRU) Delete(token string) error {
	defer l.lock.Unlock()
	l.lock.Lock()
	return l.delete(token)
}

func (l *LRU) delete(token string) error {
	var (
		node *LRUNode
		ok bool
	)
	if node, ok = l.dmap[token]; !ok {
		return nil
	}
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.first = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		l.last = node.prev
	}
	delete(l.dmap, token)
	return nil
}

func (l *LRU) Insert(token string, value interface{}) error {
	defer l.lock.Unlock()
	l.lock.Lock()
	return l.insert(token, value)
}

func (l *LRU) insert(token string, value interface{}) error{
	if l.contain(token) {
		if err := l.delete(token); err != nil {
			return err
		}
	}
	newNode := &LRUNode{
		prev: l.last,
		me: map[string]interface{}{
			"token": token,
			"value": value,
		},
		next: nil,
	}
	if l.first == nil {
		l.first = newNode
	}
	if l.last != nil {
		l.last.next = newNode
	}
	l.last = newNode
	l.dmap[token] = newNode
	if len(l.dmap) > l.count {
		if l.first == l.last {
			l.first = nil
			l.last = nil
			return nil
		}
		a := l.first
		a.next.prev = nil
		l.first = a.next
		a.next = nil
		delete(l.dmap, a.me["token"].(string))
		a = nil
	}
	return nil
}

func (l *LRU) Contain(token string) bool{
	defer l.lock.Unlock()
	l.lock.Lock()
	return l.contain(token)
}

func (l *LRU) contain(token string) bool {
	_, ok := l.dmap[token]
	return ok
}
