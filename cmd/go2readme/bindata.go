package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _assets_readme_md = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\x41\x4f\xc3\x30\x0c\x85\xef\xf9\x15\x56\xcb\x89\x43\x91\xf8\x0b\x83\xc3\x2e\xb0\x0b\xf7\x9a\xc5\x64\x95\x92\xb8\x6a\x32\x09\x34\xed\xbf\xe3\x38\x64\x74\xdb\xa9\xf2\xf3\xf7\x9e\x5f\xda\xc3\xe9\x04\xc3\x1b\x06\x82\xf3\xd9\x98\x32\xbc\x50\xda\x2f\xd3\x9c\x27\x8e\xaa\xf5\x3d\x6c\x63\xca\xe8\x3d\x16\xcd\x98\x71\x1c\xd3\xc1\x3c\x80\x63\x70\x94\x35\x60\x1b\x66\x5e\xf2\x0e\xf3\xa1\x58\x04\x50\xdb\x47\x42\x47\x92\x39\x7d\x09\x91\x36\x1c\x02\x46\xdb\x80\x62\x53\xe0\x5f\x20\x9f\xa8\x0e\x8e\xcd\xa4\x91\xd0\xdd\xc5\x77\x8d\x8e\x56\xe0\x9a\xfe\xfa\x8d\x61\xf6\x94\x44\x90\xbb\x6d\x92\xe5\x82\x51\x2e\xac\xf7\xfa\x44\xde\x4b\x50\xb5\xee\x3c\xfe\xc0\xe5\x6a\xd9\xae\x94\xbf\xf8\xf7\x63\x9e\x8f\x79\xdd\xfc\x56\x69\x65\xae\x9e\x50\xc0\x0d\x5b\xba\xc1\xae\x3f\x97\x92\xce\xf3\x27\x74\x83\xe3\xe7\x85\xd0\x06\x7a\x7a\x1c\x82\xed\xb4\x66\xfd\x15\x43\xbd\xa0\xa6\xdf\x00\x00\x00\xff\xff\x0b\xed\x59\xc7\xb8\x01\x00\x00")

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

	info := bindata_file_info{name: "assets/README.md", size: 440, mode: os.FileMode(420), modTime: time.Unix(1449886187, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"assets": &_bintree_t{nil, map[string]*_bintree_t{
		"README.md": &_bintree_t{assets_readme_md, map[string]*_bintree_t{
		}},
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

