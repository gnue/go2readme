package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func bindata_read(data []byte, name string) ([]byte, error) {
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

type bindata_file_info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _assets_readme_md = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\x31\x4f\xc3\x30\x10\x85\x77\xff\x8a\x53\xcc\xc4\x10\x24\xfe\x42\x61\xe8\x02\x5d\xd8\x73\xd4\x87\x1b\xc9\xf6\x45\xb1\x2b\x81\xaa\xfc\x77\xce\x76\x4c\x53\x3a\xde\xbb\xef\x9e\xdf\xb3\x86\xcb\x05\xfa\x37\xf4\x04\xcb\xa2\x54\x1e\x5e\x28\x1e\xe7\x71\x4a\x23\x87\xa2\x69\x0d\xfb\x10\x13\x3a\x87\x59\x53\x6a\x18\x86\x78\x52\x0f\x60\x19\x2c\xa5\x62\xb0\xf7\x13\xcf\xe9\x80\xe9\x94\x4f\x04\x28\x67\x1f\x11\x2d\x89\xe7\xf8\x25\x44\xdc\xb1\xf7\x18\x4c\x03\xf2\x59\x01\xae\x02\xb9\x48\x75\xb0\xac\xc6\x62\x09\xdd\x9d\x7d\xd7\xe8\x60\x04\xae\xee\xaf\xdf\xe8\x27\x47\x51\x04\x79\xb7\x4d\xb2\x9c\x31\xc8\x0b\xdb\x7d\x3d\xc8\x8d\x97\x45\xeb\xdb\xfe\xab\x69\x45\x0e\x0e\x7f\xe0\x2f\x4e\xc6\x36\xca\x6a\xf3\x7e\x4e\xd3\x39\x6d\x2b\xfd\x57\x5a\xca\x9b\x6e\x19\xdc\xb1\xa1\x7b\x4c\xbe\x9f\x8f\x90\x23\xac\x49\xda\xa2\x16\xb1\x8e\x3f\xa1\xeb\x2d\x3f\xcf\x84\xc6\xd3\xd3\x63\xef\x4d\x57\x12\xd7\xef\xea\xaf\x2d\xd4\x6f\x00\x00\x00\xff\xff\xc4\xfc\xb0\x52\xdc\x01\x00\x00")

func assets_readme_md_bytes() ([]byte, error) {
	return bindata_read(
		_assets_readme_md,
		"assets/README.md",
	)
}

func assets_readme_md() (*asset, error) {
	bytes, err := assets_readme_md_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "assets/README.md", size: 476, mode: os.FileMode(420), modTime: time.Unix(1462774918, 0)}
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
	"assets/README.md": assets_readme_md,
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

type _bintree_t struct {
	Func     func() (*asset, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"assets": &_bintree_t{nil, map[string]*_bintree_t{
		"README.md": &_bintree_t{assets_readme_md, map[string]*_bintree_t{}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	if err != nil { // File
		return RestoreAsset(dir, name)
	} else { // Dir
		for _, child := range children {
			err = RestoreAssets(dir, path.Join(name, child))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
