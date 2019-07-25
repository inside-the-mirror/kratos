package dotenv

import (
	"os"
	"path"
	"sync"

	"github.com/inside-the-mirror/kratos/pkg/log"
	godotenv "github.com/joho/godotenv"
)

//EnvMap .
type EnvMap struct {
	lock   *sync.RWMutex
	values map[string]string
}

var (
	keys []string
	kv   EnvMap
)

func init() {
	kv = EnvMap{
		lock:   new(sync.RWMutex),
		values: map[string]string{},
	}
	keys = []string{}
}

func load() {
	kv.lock.Lock()
	defer kv.lock.Unlock()

	env := os.Getenv("ENV")

	runPath := dir(".")
	envFile := path.Join(runPath, ".env")

	log.Info("About loading env from: %s\n", envFile)

	if exists(envFile) {
		m, err := godotenv.Read(envFile)
		if err == nil {
			kv.values = m
		}
		for k := range kv.values {
			keys = append(keys, k)
		}
		if env != "dev" && env != "" {
			// [todo] sync from remote
		}
	}
}

func (em EnvMap) set(key string, value string) {
	em.values[key] = value
}

//Get get value by key string
func Get(key string) string {
	if v, ok := kv.values[key]; ok {
		return v
	}
	return ""
}

//Gets get values by list of key
func Gets(keys []string) map[string]string {
	m := map[string]string{}
	for _, key := range keys {
		if v, ok := kv.values[key]; ok {
			m[key] = v
		}
	}
	return m
}
