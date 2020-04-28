// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

// Mode return file modify time
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

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x56\x4f\x6f\xe3\xba\x11\xbf\xeb\x53\x4c\x90\xc3\x4b\x80\xc4\x87\xa2\x7d\x07\xdd\x54\xdb\x7d\xab\x6e\xe2\xa4\xb1\x93\xed\xeb\xc3\x22\x18\x93\x63\x8b\xb0\x44\x6a\xc9\x91\x13\xb7\xc8\x77\x2f\x48\x4a\xb2\xe4\x64\xbb\xd8\x9e\x44\x91\xf3\xe7\x37\xff\xe7\x1c\x56\x85\x72\xb0\x51\x25\x81\x24\x27\xac\x5a\x93\x03\x2e\x08\x9c\x28\xa8\x42\xd8\x58\x53\x85\xff\xec\x3e\x07\x47\x76\xaf\x04\x4d\x92\xf3\xe4\x1c\x72\xfe\xc5\x81\x36\x0c\x4a\x12\x96\x57\xb0\x6e\x18\x5e\x08\x34\x91\x04\x36\x50\xa1\x6e\xb0\x2c\x0f\xb0\x25\x4d\x16\x99\x80\x0f\x35\x39\xd8\x18\x1b\xe4\xad\x0e\x35\x2d\x85\x55\x35\xc3\x63\x9e\x9c\xc3\x4b\x41\x1a\xb8\x07\xa3\x1c\x34\xb5\x44\x26\x39\x89\x10\x05\x6a\x58\x13\x48\xa3\x09\xd6\x07\xb0\x8d\xd6\x4a\x6f\xd3\xe4\x1c\x60\x6b\xb1\x2e\xbe\x95\xd7\x11\xf2\x75\xd0\x13\x25\x77\xba\xaf\xd9\xb5\x06\x4d\x5a\x62\xb8\xbe\x36\x0d\xd7\x0d\x77\xf7\x72\xc2\x2e\xc0\x50\xa2\x80\x17\x55\x96\x03\xe0\x05\x41\x4b\xec\x65\x47\x80\x5c\x20\x47\xba\x35\x41\xad\xc4\x8e\x24\x34\xb5\x87\xe6\xc9\x1f\xf3\x49\xd2\xfa\x76\x20\x3f\x70\x3a\x70\x85\x69\x4a\x09\xf4\xaa\x1c\x83\xd2\xd1\xdd\x58\x11\x48\x65\x49\xb0\xb1\x07\xc0\x61\x10\x7a\xcc\x9e\x7d\x92\x24\x6d\x68\xfe\x93\x00\x7c\x6b\xc8\x1e\x52\xf8\x87\xff\x24\x00\x55\xc3\xc8\xca\xe8\x14\x6e\xdb\x53\xf2\x96\x24\x01\xf4\xa3\x23\x9b\xeb\x8d\x09\x6c\x4a\xa6\x90\xcf\xce\x12\x00\x8d\x15\xa5\xb0\x64\xab\xf4\xd6\xff\x53\x85\xaa\x1c\x5e\xd4\x4a\x70\x63\x47\x34\xc6\x6e\x17\x23\xb6\xb7\x24\x21\xdd\x54\x90\x59\x56\x1b\x14\xec\x63\x1b\xf4\x00\x64\xab\xe7\xc7\xc5\xe7\xc5\xdd\x97\x45\xf7\x7b\x93\x2f\x1e\xff\xf9\x9c\xdd\xce\x7e\xfd\x73\x77\x35\xcb\x1e\xbe\xe4\x8b\xf1\xdd\xf4\x6e\xb1\xca\xf2\xc5\xfc\xe1\x79\x39\x5f\x3d\xff\x9e\xdd\xde\x2c\x3f\x7e\x1a\xca\xeb\x81\x34\x6c\x84\xa9\xea\x92\x98\x32\xe1\xfd\xd0\x43\xca\x46\x88\xce\x21\xd3\x40\x52\x31\x60\x20\x03\x23\x44\x63\x1d\xa8\x0d\x20\x34\x8e\x2c\x14\xe8\xa0\x32\x52\x6d\x94\xcf\xeb\x82\x40\xe9\x90\x08\xf4\xca\x3e\xd8\x4a\x3b\xb2\xac\xf4\x16\x8c\x05\x49\x25\x85\xb3\x28\xd0\xa2\x60\xb2\x6e\x12\x94\x84\x44\x50\x5a\x94\x8d\xf4\xe5\x75\xa8\x03\x43\x8c\xfc\x8e\x0e\x6b\x83\x56\x02\x6a\x09\x35\xba\x28\xc0\x54\x15\x6a\x19\xd8\x3d\xe2\xf9\x2c\x5f\x45\xb8\xe0\xa8\x24\x71\xc4\xab\xcb\xc3\xc7\xa0\x45\x61\x1c\x69\x40\x0d\x38\xf0\x06\xb8\x66\xbb\x25\xe7\x79\x27\x1d\x2c\xa9\x04\xb2\xc7\x65\x82\x0a\x0f\x6a\xc4\x12\x52\x5d\x71\x97\xb7\x95\xd9\xc7\x9a\xf0\xaa\x7e\x71\xe0\x75\xfb\xa2\x36\xe1\x52\x7b\xc7\x60\x5d\x5b\x53\x5b\x15\xaa\x07\xd7\x9d\x15\xcb\xf9\xcd\x7c\xba\xfa\x30\x4a\x73\xcd\x8a\x0f\x9f\x95\x96\x31\x4a\xf3\xcf\x83\x28\xf9\xbf\xfb\xbb\x59\x7b\x5a\x3e\x4d\xbb\xd3\xf4\x21\xbf\x5f\xb5\x3f\x8b\xec\x76\xbe\xbc\xcf\xa6\xf3\x3e\xe5\x43\x55\x04\x71\x1e\x69\xda\x97\x80\xcf\x61\x51\x36\x8e\xfd\xe5\x34\x1e\x4e\xee\xa7\x46\x6b\x12\xb1\x94\xa6\xa7\x57\x47\x5a\xd5\x65\xfc\x05\x0e\x52\x3f\x1d\x15\xc2\x65\x0a\xd3\x9b\xbc\xbb\xf1\x7c\x1d\xad\xeb\xb9\x86\xe5\x74\x79\x64\x77\x9d\xa6\x61\x38\x2e\x42\x02\x76\xd4\x57\xad\xfb\xef\x8d\x4b\x21\xd7\x7c\xd5\x26\x46\xfa\x9d\x1a\xb8\x1c\x3f\x3c\x90\x6b\x4a\x3e\x4b\x42\xe0\x63\x3b\xae\xb6\x16\x48\xcb\xda\x28\xcd\xee\x0a\x2c\x6d\x28\x04\x57\x1a\xe1\x93\x05\x44\x69\x1a\x89\xb5\x9a\xd4\xd6\x84\x8c\x29\xd5\x9e\x9e\x14\xbd\xb8\x14\xfe\xb8\x69\xcf\xb7\xc4\x28\x91\xf1\xec\xeb\xd9\x80\x62\x6a\x34\x93\x66\x77\xd1\xf6\x9f\xcb\x14\x6e\x4e\x9e\x3c\x79\x6c\xde\x5e\x5c\x44\x34\x16\x16\x5f\x3f\x10\xb5\x1c\x3d\x9c\xf5\x69\xf0\xde\xde\x90\x13\xbe\x2a\xc9\x8f\xa3\x0a\x99\x49\xb6\x75\xad\xdc\xa0\xc8\x5d\xeb\xe2\x38\x14\x7c\x51\xad\x89\x34\xd4\x68\x1d\xc9\xae\xd5\x8f\x4b\xc5\xf4\xf5\x14\x6b\x09\xd7\x4b\x36\x35\xd4\xc6\x29\x1f\x82\x50\xd0\xbd\xce\x7c\x18\xc9\x40\xff\xa5\x20\x2e\xc8\xbe\xc3\xe0\x71\x21\xec\xb1\x54\xf2\x0a\xe8\x95\x44\xc3\xb8\x2e\xa9\xeb\x13\x5e\xaa\x72\xf3\xfe\x3e\x85\xbf\x1a\x53\x12\xea\xd8\x33\xca\x72\x50\xf6\x71\x04\x13\x8a\x02\xcc\x26\x28\x6a\x41\x06\x6c\xfe\x7c\x24\x4d\xe1\x8f\xd5\xf0\xe2\x6b\xef\xd4\xd1\xf5\xc0\x9f\x4a\x4b\x7a\x1d\x08\x8e\xcd\x83\x0b\x72\x34\xc2\x80\x36\xf8\xbe\x55\x99\x7b\xae\x90\xbf\x23\x2f\xc4\x56\xe7\xcd\xc7\x01\x73\xbb\x42\xf8\x48\xe1\xba\x55\x18\x06\x71\x85\xbb\xd8\x98\x5a\xaf\x0c\x1c\xe5\xf5\x1c\xff\xb2\x0d\x93\x5d\x06\xe1\x43\x4f\xb9\x91\xe1\xc3\xb4\xf9\xc8\x03\x1f\xbf\x07\x57\xec\x94\x96\xe9\x77\xda\xdb\xc9\xcc\x4d\xa0\xdd\xbb\xea\x58\xb5\xed\x6d\xd7\x22\xdb\xde\xb3\x64\xe4\xc6\x0d\xdc\x2c\x69\x83\x3e\x91\x1d\xfb\x16\xab\x36\x7e\x11\x2b\xda\x3c\xd9\x69\xf3\xa2\xbd\xc1\x4f\xff\x7a\x5e\x8e\x87\x9d\x67\x6d\x59\x1c\x14\x84\x25\x17\x07\xcf\x5d\x10\x5a\x5e\x13\x72\x0c\x8c\x25\x41\x6a\x4f\xd2\x8f\x28\x4b\xdb\xa6\x44\x0b\x4a\x33\xd9\x3d\x96\x2e\xcc\x29\x2e\x62\x7e\xb7\xed\xd2\x8b\xb3\xe4\x6a\xa3\xa5\x07\xc1\x06\x2c\x7d\x6b\xc8\xb1\x3b\xe2\xf8\x34\xcf\x6e\x56\x9f\x7e\x3f\xc1\x11\x37\x2d\x13\x1a\x8c\x72\x22\xf6\x58\x5f\x8d\x31\x83\x7e\x7b\xb8\x9f\x82\xe8\x3b\x2f\xac\x2d\xe1\xce\x4d\x82\x80\xc2\xd4\x14\xeb\x15\xb9\x1f\x5c\x1d\xa0\x20\x57\x98\x8a\x60\x8d\x62\xe7\xc7\xa4\xd2\x14\xa0\x5b\x72\x4d\xe5\x13\x15\x5a\x44\x11\xc9\x11\xe8\x2c\x5f\x4e\xef\x16\x8b\xf9\x74\x35\x9f\x1d\xbd\x78\x3f\xcb\x56\xf9\xe2\xb7\x3e\xfc\x4f\xea\xdf\x2a\x0c\x85\x8d\xda\x86\xc0\xd4\xe8\x1c\x17\xd6\x34\xdb\x62\xae\x7d\x96\xc9\x63\x6e\x75\x4c\x83\x61\x73\xb2\x87\xb9\x10\xe1\x74\x1c\xf0\xd0\x3b\xd1\xf1\xa7\x2e\x3c\xb7\x2e\x85\xbf\x95\x06\xc3\x18\xd9\x0f\x10\xa4\x23\x3c\x67\xa7\x0a\xc7\xb3\x2b\xaa\xae\x33\x29\x2d\x39\x37\xdc\xec\xd8\xec\x48\x8f\xf6\xba\x20\xa5\x5b\x26\x03\xe3\xd4\x12\x32\x4d\x3f\x98\x9f\x09\xc0\x63\xd8\xd9\x87\x58\x2e\xda\x88\xe4\xb3\x60\xea\xd5\xff\xf2\xd3\x65\x7f\x1a\x58\x70\x9c\x9d\xed\x4a\xd9\xd8\xd1\x82\x0a\xe0\x0a\xfc\xd3\x5f\x7e\x7d\x0f\x7b\x34\x46\xa3\xd1\x4c\x55\xa8\xed\xf6\xe5\xeb\x3b\xda\x40\xb6\x27\xeb\x06\xd5\x18\x66\x7d\x81\x7a\x4b\xa5\xd9\x8e\xdc\xa5\x2a\x72\x8c\x55\x3d\x88\xcb\x5b\x92\x9c\xc3\xc3\x0f\x86\x66\x50\x79\x3a\x2b\x7f\xb0\x99\xfb\x36\x31\xb2\xf1\x27\xd5\x74\x83\x31\xa8\xa9\x5a\x9d\xe9\x3b\x14\x61\xe7\x7f\x2d\x3b\xea\x21\x82\xbd\x72\x7f\x5f\xde\x2d\xfe\x1f\x10\xe3\x41\xfe\x53\x96\x82\x1f\xbc\x1d\xca\x71\x82\xfc\x94\xf2\xef\xd8\x7f\xb2\x62\xf8\x50\xbf\x33\xfd\x2d\xf9\x6f\x00\x00\x00\xff\xff\xe0\x01\x96\xca\x28\x0f\x00\x00")

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

	info := bindataFileInfo{name: "schema.graphql", size: 3880, mode: os.FileMode(436), modTime: time.Unix(1588048302, 0)}
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
