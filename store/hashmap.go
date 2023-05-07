package store

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"sync"
)

type hashmap struct {
	sync.RWMutex
	m map[string]interface{}
}

func (hm *hashmap) Get(key string) (string, bool) {
	hm.RLock()
	defer hm.RUnlock()
	val, ok := hm.m[key]
	return val.(string), ok
}

func (hm *hashmap) Set(key string, value string) {
	hm.Lock()
	defer hm.Unlock()
	hm.m[key] = value
}

func (hm *hashmap) Remove(key string) {
	hm.Lock()
	defer hm.Unlock()
	delete(hm.m, key)
}

func (hm *hashmap) Len() int {
	hm.RLock()
	defer hm.RUnlock()
	return len(hm.m)
}

func (hm *hashmap) Keys() []string {
	hm.RLock()
	defer hm.RUnlock()
	keys := make([]string, 0, len(hm.m))
	for k := range hm.m {
		keys = append(keys, k)
	}
	return keys
}

// set any structure.
func (hm *hashmap) SetAny(key string, v any) {
	hm.Lock()
	defer hm.Unlock()
	hm.m[key] = v
}

// get any structure. Decode. pass the structure reference to this function.
func (hm *hashmap) GetAny(key string, v any) error {
	hm.RLock()
	defer hm.RUnlock()
	receivedInterface := hm.m[key]
	rawBytes, err := json.Marshal(receivedInterface)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawBytes, v)
	if err != nil {
		return err
	}
	return nil
}

func (hm *hashmap) SetCompressed(key string, v any) error {
	hm.Lock()
	defer hm.Unlock()
	compressed, err := CompressGzip(v)
	if err != nil {
		return err
	}
	hm.m[key] = compressed
	return nil
}

func (hm *hashmap) GetCompressed(key string, v any) error {
	hm.RLock()
	defer hm.RUnlock()
	compressed := hm.m[key]
	err := DecompressGzip(compressed.(string), v)
	if err != nil {
		return err
	}
	return nil
}

// CompressGzip compresses the given data using gzip and returns a base64-encoded string
func CompressGzip(data interface{}) (string, error) {
	// 1. convert struct to json string
	// 2. compress into byte sclice
	// 3. return base64 encoded version of the byte slice
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(jsonData); err != nil {
		return "", err
	}

	if err := gz.Close(); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

// DecompressGzip takes a base64-encoded gzip string and returns a new instance of the provided struct
func DecompressGzip(data string, v any) error {
	// 1. first decode by base64. This gives the compressed byte string
	// 2. Uncompress the data
	// 3. load the data into user type object
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err
	}

	reader, err := gzip.NewReader(bytes.NewReader(decoded))
	if err != nil {
		return err
	}
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(v)
	if err != nil {
		return err
	}

	return nil
}
