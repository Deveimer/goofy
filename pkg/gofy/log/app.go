package log

import "sync"

type appInfo struct {
	Name      string                 `json:"name"`
	Version   string                 `json:"version"`
	Framework string                 `json:"framework"`
	Data      map[string]interface{} `json:"data,omitempty"`
	syncData  *sync.Map
}

func (a *appInfo) getAppData() appInfo {
	res := appInfo{}

	res.Name = a.Name
	res.Version = a.Version
	res.Framework = a.Framework
	res.Data = make(map[string]interface{})

	if a.syncData != nil {
		a.syncData.Range(func(key, value interface{}) bool {
			res.Data[key.(string)] = value
			return true
		})
	}

	return res
}

func (k *logger) AddData(key string, value interface{}) {
	k.app.syncData.Store(key, value)
}

func (k *logger) removeData(key string) {
	k.app.syncData.Delete(key)
}
