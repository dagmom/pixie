// Package schema Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// schema.graphql
package schema

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x4f\x6f\xe3\xba\x11\xbf\xeb\x53\x4c\x90\xc3\x4b\x80\xc4\x87\xa2\x7d\x07\xdd\x54\xdb\xfb\x56\xdd\xc4\x49\x63\x27\xdb\xd7\x87\x20\xa0\xc5\xb1\x45\x44\x22\xb5\xe4\xc8\x89\x5b\xec\x77\x2f\x86\xa4\x64\xc9\x49\xba\xd8\x02\x3d\x59\x26\x39\x33\xbf\xf9\xf3\x1b\x0e\x4f\x61\x55\x2a\x07\x1b\x55\x21\x48\x74\x85\x55\x6b\x74\x40\x25\x82\x2b\x4a\xac\x05\x6c\xac\xa9\xfd\xff\xec\x36\x07\x87\x76\xa7\x0a\x9c\x24\xa7\xc9\x29\xe4\xf4\x8b\x03\x6d\x08\x94\x44\x51\x5d\xc0\xba\x25\x78\x41\xd0\x88\x12\xc8\x40\x2d\x74\x2b\xaa\x6a\x0f\x5b\xd4\x68\x05\x21\xd0\xbe\x41\x07\x1b\x63\xbd\xbe\xd5\xbe\xc1\x65\x61\x55\x43\x70\x9f\x27\xa7\xf0\x52\xa2\x06\xea\xc1\x28\x07\x6d\x23\x05\xa1\x9c\x04\x88\x85\xd0\xb0\x46\x90\x46\x23\xac\xf7\x60\x5b\xad\x95\xde\xa6\xc9\x29\xc0\xd6\x8a\xa6\xfc\x56\x5d\x06\xc8\x97\xde\x4e\xd0\xdc\xd9\xbe\x24\x17\x1d\x9a\xc4\xc3\x70\x79\x69\x5a\x6a\x5a\xea\xd6\xe5\x84\x9c\x87\xa1\x8a\x12\x5e\x54\x55\x0d\x80\x97\x08\xf1\x30\xeb\x0e\x00\xa9\x14\x14\xce\xad\x11\x1a\x55\x3c\xa3\x84\xb6\x61\x68\x7c\xfc\x3e\x9f\x24\x31\xb6\x03\xfd\x5e\xd2\x81\x2b\x4d\x5b\x49\xc0\x57\xe5\x08\x94\x0e\xe1\x16\x35\x82\x54\x16\x0b\x32\x76\x0f\x62\x98\x84\x1e\x33\x8b\x4f\x92\x24\xa6\xe6\xdf\x09\xc0\xb7\x16\xed\x3e\x85\xbf\xf3\x4f\x02\x50\xb7\x24\x48\x19\x9d\xc2\x75\xfc\x4a\xbe\x27\x89\x07\x7d\xef\xd0\xe6\x7a\x63\xbc\x98\x92\x29\xe4\xb3\x93\x04\x40\x8b\x1a\x53\x58\x92\x55\x7a\xcb\xff\xb1\x16\xaa\x1a\x2e\x34\xaa\xa0\xd6\x8e\xce\x18\xbb\x5d\x8c\xc4\xbe\x27\x09\xea\xb6\x86\xcc\x92\xda\x88\x82\x38\xb7\xde\x0e\x40\xb6\x7a\xba\x5f\x7c\x59\xdc\x7c\x5d\x74\x7f\xaf\xf2\xc5\xfd\x3f\x9e\xb2\xeb\xd9\xaf\x7f\xee\x96\x66\xd9\xdd\xd7\x7c\x31\x5e\x9b\xde\x2c\x56\x59\xbe\x98\xdf\x3d\x2d\xe7\xab\xa7\xdf\xb3\xeb\xab\xe5\xfb\x5b\x43\x7d\x3d\x90\x96\x4c\x61\xea\xa6\x42\xc2\xac\xe0\x38\xf4\x90\xb2\x11\xa2\x53\xc8\x34\xa0\x54\x04\xc2\x1f\x03\x53\x14\xad\x75\xa0\x36\x20\xa0\x75\x68\xa1\x14\x0e\x6a\x23\xd5\x46\x71\x5d\x97\x08\x4a\xfb\x42\xc0\x57\xe2\x64\x2b\xed\xd0\x92\xd2\x5b\x30\x16\x24\x56\xe8\xbf\x8b\x52\x58\x51\x10\x5a\x37\xf1\x46\x7c\x21\x28\x5d\x54\xad\x64\x7a\xed\x1b\x2f\x10\x32\xff\x8c\xfb\xb5\x11\x56\x82\xd0\x12\x1a\xe1\x82\x02\x53\xd7\x42\x4b\x2f\xce\x88\xe7\xb3\x7c\x15\xe0\x82\xc3\x0a\x8b\x03\x5e\x5d\xed\xdf\x07\x5d\x94\xc6\xa1\x06\xa1\x41\x0c\xa2\x01\xae\xdd\x6e\xd1\xb1\xec\xa4\x83\x25\x55\x21\x88\x71\x19\x6f\x82\x41\x8d\x44\x7c\xa9\x2b\xea\xea\xb6\x36\xbb\xc0\x09\x36\xf5\x8b\x03\xb6\xcd\xa4\x36\x7e\x51\x73\x60\x44\xd3\x58\xd3\x58\xe5\xd9\x23\xd6\x9d\x17\xcb\xf9\xd5\x7c\xba\x7a\x37\x4b\x73\x4d\x8a\xf6\x5f\x94\x96\x21\x4b\xf3\x2f\x83\x2c\xf1\xbf\xdb\x9b\x59\xfc\x5a\x3e\x4c\xbb\xaf\xe9\x5d\x7e\xbb\x8a\x7f\x16\xd9\xf5\x7c\x79\x9b\x4d\xe7\x7d\xc9\x7b\x56\x78\x75\x8c\x34\xed\x29\x70\x12\x72\x72\x33\xbb\x39\xe3\xec\x59\xa5\xcd\x79\x0a\xd7\xe2\x19\x21\x9f\x81\xc5\x6f\xad\xb2\x28\xc1\xe8\x82\x89\xec\xfd\x75\x60\x76\xe8\x7d\xac\xdb\x8a\xd4\x65\x51\xb5\x8e\xd0\x82\x6b\x9b\xc6\x58\x62\x07\xe3\xd2\x59\xe0\xd6\x79\x0a\xd3\xb0\xd0\x59\x8c\xfb\x2e\x85\x3f\x86\x3b\x8f\xff\x57\x34\x53\xa3\x35\xfa\x4a\x79\x83\xeb\xb0\x75\x40\xa8\x3a\x06\x9f\x89\x01\x95\xd3\x11\xb1\x59\xc3\x55\xde\xad\xb0\x5c\x77\xd6\xf5\x52\xc3\xf6\x70\x7e\x10\x77\x9d\xa5\x61\x79\x9d\x79\x42\x75\xa7\x2f\x62\x39\xdd\x1a\x97\x42\xae\xe9\x22\x16\x7a\xfa\x01\xa7\xcf\xc7\x1b\x77\xe8\xda\x8a\x8e\x4d\x7c\x52\x58\xc9\x63\x3b\x1b\x5e\x8c\xee\xbd\x5b\x88\x17\xbe\xdf\x74\x09\xc8\xec\x96\x0f\x73\xfa\xde\x3f\xfe\x78\x9e\x8e\x76\x96\x3d\xd1\x1e\x13\x9f\xe2\x70\xdb\xd5\x5b\x0b\xa8\x65\x63\x94\x26\x77\x01\x16\x37\x21\x93\xd2\x14\xcc\x45\x28\x2a\xd3\x4a\xd1\xa8\x49\x63\x8d\x27\x64\xa5\x76\xf8\xa0\xf0\x85\x2d\x5f\xc5\xef\x6b\x24\x21\x05\x89\x50\x3d\xdd\x89\xa9\xd1\x84\x9a\x5c\x4c\xf5\xc9\x79\x0a\x57\x47\x5b\x7c\x3c\xdc\x8d\xac\x2e\x20\x1a\x2b\x0b\xbb\xef\xa8\x5a\x8e\x36\x4e\x7a\x96\xbd\x0d\xbf\xa7\x1c\x37\x3d\xe4\xdb\xbe\x16\x44\x28\x63\xdb\x54\x6e\xd0\x43\x5d\xcc\x44\xb8\x73\xb9\x67\xad\x11\x35\x34\xc2\x3a\x94\xdd\x4d\x3a\xee\x44\xa6\x6f\x57\xa1\x55\x89\xf5\x92\x4c\x03\x8d\x71\x8a\x23\xed\xfb\x65\x6f\x33\x1f\x26\xdc\x9f\xff\x5a\x22\x95\x68\xdf\x60\x60\x5c\x02\x76\xa2\x52\xf2\x02\xf0\x15\x8b\x96\xc4\xba\xc2\xae\x0d\xb3\x56\xe5\xe6\xfd\x7a\x0a\x7f\x35\xa6\x42\xa1\x43\x4b\xae\xaa\x41\x57\x0d\x13\x0e\x8a\xa2\x04\xb3\xf1\x86\x22\x48\x8f\x8d\xbf\x0f\x47\x53\xf8\x63\x35\x5c\x78\xec\x83\x3a\x5a\x1e\xc4\x53\x69\x89\xaf\x03\xc5\xa1\x37\x53\x89\x0e\x47\x18\x84\xf5\xb1\x8f\x26\x73\x96\xf2\x74\x1a\x45\x21\xdc\x24\xec\xbe\x18\x08\xc7\x09\x8d\x33\x25\xd6\xd1\xa0\x9f\x73\x6a\x6e\x49\x6c\x37\x46\x65\x10\x28\xb6\x73\xf8\x97\x6d\x08\xed\xd2\x2b\x1f\x46\xca\x8d\x1c\xff\x88\x2a\xef\x95\xd5\x51\x28\x9e\x95\x96\x1f\x91\xf6\x68\xa4\x49\x20\x8e\xb5\x4d\x68\x22\xfd\x6a\x2d\xa8\x28\xb9\x44\x24\xbe\x7a\x52\xe7\x9a\x1e\xfb\x9b\x29\xb6\xc8\x25\x09\x6a\xdd\x20\xfc\x12\x37\x82\x0b\xdc\x11\xdf\x6c\x6a\xc3\xf3\x6f\x19\xeb\xe7\x59\x9b\x17\xcd\x81\x78\xf8\xe7\xd3\x72\x3c\x63\xb0\x68\x14\x71\x50\xa2\xa8\xa8\xdc\xb3\x74\x89\xc2\xd2\x1a\x05\x85\x84\x59\x2c\x50\xed\x7c\xaf\x07\x8b\xdb\xb6\x12\x16\x94\x26\xb4\x3b\x51\x39\x3f\x1e\x50\x19\xea\xbe\x6b\xf8\xca\x81\x45\xd7\x18\x2d\x19\x04\x19\xdf\xad\xd0\x91\x3b\xe0\xf8\x3c\xcf\xae\x56\x9f\x7f\x3f\xc2\x11\x06\x5c\xe3\x1b\x8f\x72\x45\xb8\x0a\x98\xa5\xa1\xb2\x7e\xbb\xbb\x9d\x42\xd1\x5f\x10\xb0\xb6\x28\x9e\xdd\xc4\x2b\x28\x4d\x83\x81\xc7\x82\xfa\x79\xa1\x03\xe4\xf5\x16\xa6\x46\x58\x8b\xe2\x99\xa7\x13\xa5\xd1\x43\xb7\xe8\xda\x9a\x0b\x18\x22\xa2\x80\xe4\x00\x74\x96\x2f\xa7\x37\x8b\xc5\x7c\xba\x9a\xcf\xde\x46\xcd\x3f\x06\xd8\xc9\xf8\x4e\xc0\x61\x0c\xe2\x0c\xdd\x58\x53\xa0\x73\x4c\x8f\xee\xf8\x20\x1f\xb7\xb3\x6c\x95\x2f\x7e\xeb\x55\xef\xd4\xbf\x54\x37\x2a\x75\xfe\x87\x57\x0c\x2f\xf1\xc3\xc6\xa1\x26\x10\x7a\x0f\xc6\xd3\x65\xd3\xda\x40\x9b\x50\x15\xe1\x79\xe2\x40\xac\x4d\x1b\x02\xf1\x12\x79\xa5\x68\x98\x67\x63\xdf\x41\xf3\xd6\xd3\x08\xe7\x85\x27\x7f\xbb\x8f\xe9\x0c\x36\x02\xaa\x8d\x50\x15\x86\x29\x51\x31\xbe\x17\x76\x5b\xc0\x5a\xc8\xe3\x48\x7a\x57\xe7\x4f\x9f\xb2\xfc\x6a\x3e\xeb\x09\xf5\xe0\x0d\x4c\x8d\xde\xa8\xad\x2f\xe9\x46\x38\x47\xa5\x35\xed\xb6\x9c\x6b\xe6\xad\x3c\xb0\xb5\x13\x1a\x4c\x2a\x47\x0f\x87\x10\x85\x74\x4c\x15\x7f\x1b\x09\x47\x9f\xbb\xc2\xbe\x76\x29\x7c\xaa\x8c\xf0\x57\xf2\x6e\x80\x20\x1d\xe1\x39\xec\x3e\xa0\x75\x63\xa6\xc6\x3c\x7f\xb8\xb1\x18\xd3\x3d\xae\xde\xe7\xb3\x7e\xf1\xc8\x99\xf1\xe0\x13\xdc\x6a\x32\x29\x2d\x3a\x37\x7c\xe6\x90\x79\x46\x3d\x7a\xe4\x78\x2d\xdd\xcb\xca\x0b\x4e\x2d\x0a\xc2\xa8\x78\x34\xf2\x25\x00\xf7\x3e\x7b\x43\x3f\xcf\x22\x3a\x06\x97\xcf\x4e\x2e\xfe\x5b\x0e\xce\xfb\xaf\x83\xed\xc1\xe0\x15\xdf\x57\xad\x1d\xbd\xd6\x00\x5c\x29\xfe\xf4\x97\x5f\xdf\xc2\x1e\xcd\x60\xc1\x69\xc2\xda\x77\xe2\xb8\xf3\xf8\xe6\xac\x3f\xb6\x1b\x07\xde\x0f\x8a\xa5\xd0\x5b\xac\xcc\x76\x14\x2e\x55\xa3\x23\x51\x37\x83\x9c\x7f\x4f\x92\x53\xb8\xfb\xc1\x88\xe3\x4d\x1e\x4f\x36\x3f\x78\xa6\x72\x53\x1f\xf9\xf8\x93\x66\xba\x31\xc6\x9b\xa9\xa3\xcd\xf4\x0d\x0a\xff\x00\x7e\xad\xba\xd3\x43\x04\x3b\xe5\xfe\xb6\xbc\x59\xfc\x2f\x20\xc6\x63\xd7\x4f\x79\x0a\xdc\x9c\x3a\x94\xe3\x02\xf9\x29\xe3\x1f\xf8\x7f\x34\x10\x72\xaa\xdf\xb8\xfe\x3d\xf9\x4f\x00\x00\x00\xff\xff\xe4\xf4\xf2\xd5\x35\x12\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 4661, mode: os.FileMode(436), modTime: time.Unix(1590205897, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"schema.graphql": schemaGraphql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
