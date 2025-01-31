// Code generated by kratos tool genbts. DO NOT EDIT.

/*
  Package testdata is a generated cache proxy package.
  It is generated from:
  type _bts interface {
		// bts: -batch=2 -max_group=20 -batch_err=break -nullcache=&Demo{ID:-1} -check_null_code=$.ID==-1
		Demos(c context.Context, keys []int64) (map[int64]*Demo, error)
	    // bts: -batch=2 -max_group=20 -batch_err=continue -nullcache=&Demo{ID:-1} -check_null_code=$.ID==-1
	    Demos1(c context.Context, keys []int64) (map[int64]*Demo, error)
		// bts: -sync=true -nullcache=&Demo{ID:-1} -check_null_code=$.ID==-1
		Demo(c context.Context, key int64) (*Demo, error)
		// bts: -paging=true
		Demo1(c context.Context, key int64, pn int, ps int) (*Demo, error)
		// bts: -nullcache=&Demo{ID:-1} -check_null_code=$.ID==-1
		None(c context.Context) (*Demo, error)
	}
*/

package testdata

import (
	"context"
	"sync"

	"github.com/inside-the-mirror/kratos/pkg/stat/prom"
	"github.com/inside-the-mirror/kratos/pkg/sync/errgroup"
)

var _ _bts

// Demos get data from cache if miss will call source method, then add to cache.
func (d *Dao) Demos(c context.Context, keys []int64) (res map[int64]*Demo, err error) {
	if len(keys) == 0 {
		return
	}
	addCache := true
	if res, err = d.CacheDemos(c, keys); err != nil {
		addCache = false
		res = nil
		err = nil
	}
	var miss []int64
	for _, key := range keys {
		if (res == nil) || (res[key] == nil) {
			miss = append(miss, key)
		}
	}
	prom.CacheHit.Add("Demos", int64(len(keys)-len(miss)))
	for k, v := range res {
		if v.ID == -1 {
			delete(res, k)
		}
	}
	missLen := len(miss)
	if missLen == 0 {
		return
	}
	missData := make(map[int64]*Demo, missLen)
	prom.CacheMiss.Add("Demos", int64(missLen))
	var mutex sync.Mutex
	group := errgroup.WithCancel(c)
	if missLen > 20 {
		group.GOMAXPROCS(20)
	}
	var run = func(ms []int64) {
		group.Go(func(ctx context.Context) (err error) {
			data, err := d.RawDemos(ctx, ms)
			mutex.Lock()
			for k, v := range data {
				missData[k] = v
			}
			mutex.Unlock()
			return
		})
	}
	var (
		i int
		n = missLen / 2
	)
	for i = 0; i < n; i++ {
		run(miss[i*2 : (i+1)*2])
	}
	if len(miss[i*2:]) > 0 {
		run(miss[i*2:])
	}
	err = group.Wait()
	if res == nil {
		res = make(map[int64]*Demo, len(keys))
	}
	for k, v := range missData {
		res[k] = v
	}
	if err != nil {
		return
	}
	for _, key := range miss {
		if res[key] == nil {
			missData[key] = &Demo{ID: -1}
		}
	}
	if !addCache {
		return
	}
	d.cache.Do(c, func(c context.Context) {
		d.AddCacheDemos(c, missData)
	})
	return
}

// Demos1 get data from cache if miss will call source method, then add to cache.
func (d *Dao) Demos1(c context.Context, keys []int64) (res map[int64]*Demo, err error) {
	if len(keys) == 0 {
		return
	}
	addCache := true
	if res, err = d.CacheDemos1(c, keys); err != nil {
		addCache = false
		res = nil
		err = nil
	}
	var miss []int64
	for _, key := range keys {
		if (res == nil) || (res[key] == nil) {
			miss = append(miss, key)
		}
	}
	prom.CacheHit.Add("Demos1", int64(len(keys)-len(miss)))
	for k, v := range res {
		if v.ID == -1 {
			delete(res, k)
		}
	}
	missLen := len(miss)
	if missLen == 0 {
		return
	}
	missData := make(map[int64]*Demo, missLen)
	prom.CacheMiss.Add("Demos1", int64(missLen))
	var mutex sync.Mutex
	group := errgroup.WithContext(c)
	if missLen > 20 {
		group.GOMAXPROCS(20)
	}
	var run = func(ms []int64) {
		group.Go(func(ctx context.Context) (err error) {
			data, err := d.RawDemos1(ctx, ms)
			mutex.Lock()
			for k, v := range data {
				missData[k] = v
			}
			mutex.Unlock()
			return
		})
	}
	var (
		i int
		n = missLen / 2
	)
	for i = 0; i < n; i++ {
		run(miss[i*2 : (i+1)*2])
	}
	if len(miss[i*2:]) > 0 {
		run(miss[i*2:])
	}
	err = group.Wait()
	if res == nil {
		res = make(map[int64]*Demo, len(keys))
	}
	for k, v := range missData {
		res[k] = v
	}
	if err != nil {
		return
	}
	for _, key := range miss {
		if res[key] == nil {
			missData[key] = &Demo{ID: -1}
		}
	}
	if !addCache {
		return
	}
	d.cache.Do(c, func(c context.Context) {
		d.AddCacheDemos1(c, missData)
	})
	return
}

// Demo get data from cache if miss will call source method, then add to cache.
func (d *Dao) Demo(c context.Context, key int64) (res *Demo, err error) {
	addCache := true
	res, err = d.CacheDemo(c, key)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res.ID == -1 {
			res = nil
		}
	}()
	if res != nil {
		prom.CacheHit.Incr("Demo")
		return
	}
	prom.CacheMiss.Incr("Demo")
	res, err = d.RawDemo(c, key)
	if err != nil {
		return
	}
	miss := res
	if miss == nil {
		miss = &Demo{ID: -1}
	}
	if !addCache {
		return
	}
	d.AddCacheDemo(c, key, miss)
	return
}

// Demo1 get data from cache if miss will call source method, then add to cache.
func (d *Dao) Demo1(c context.Context, key int64, pn int, ps int) (res *Demo, err error) {
	addCache := true
	res, err = d.CacheDemo1(c, key, pn, ps)
	if err != nil {
		addCache = false
		err = nil
	}
	if res != nil {
		prom.CacheHit.Incr("Demo1")
		return
	}
	var miss *Demo
	prom.CacheMiss.Incr("Demo1")
	res, miss, err = d.RawDemo1(c, key, pn, ps)
	if err != nil {
		return
	}
	if !addCache {
		return
	}
	d.cache.Do(c, func(c context.Context) {
		d.AddCacheDemo1(c, key, miss, pn, ps)
	})
	return
}

// None get data from cache if miss will call source method, then add to cache.
func (d *Dao) None(c context.Context) (res *Demo, err error) {
	addCache := true
	res, err = d.CacheNone(c)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res.ID == -1 {
			res = nil
		}
	}()
	if res != nil {
		prom.CacheHit.Incr("None")
		return
	}
	prom.CacheMiss.Incr("None")
	res, err = d.RawNone(c)
	if err != nil {
		return
	}
	var miss = res
	if miss == nil {
		miss = &Demo{ID: -1}
	}
	if !addCache {
		return
	}
	d.cache.Do(c, func(c context.Context) {
		d.AddCacheNone(c, miss)
	})
	return
}
