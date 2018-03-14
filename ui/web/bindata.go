// Code generated by go-bindata.
// sources:
// admin_http_assets/app.css
// admin_http_assets/app.js
// admin_http_assets/index.html
// DO NOT EDIT!

package web

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _admin_http_assetsAppCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xce\x4b\x8f\xb1\x4a\xce\xc9\x4f\xcc\x8e\xd5\x51\x88\xce\x4b\xd7\x85\xb3\x53\x12\x4b\x12\x75\x91\x05\x2a\x90\x79\x7a\x30\xb6\x8e\x82\x1e\x42\x02\x22\x9e\x91\x99\x92\xaa\x50\xcd\xa5\xa0\xa0\xa0\x90\x92\x59\x5c\x90\x93\x58\x69\xa5\x90\x97\x9f\x97\xaa\xa0\x98\x99\x5b\x90\x5f\x54\x92\x98\x57\x62\xcd\x55\x0b\x08\x00\x00\xff\xff\x9c\x9e\x21\xb0\x7a\x00\x00\x00")

func admin_http_assetsAppCssBytes() ([]byte, error) {
	return bindataRead(
		_admin_http_assetsAppCss,
		"admin_http_assets/app.css",
	)
}

func admin_http_assetsAppCss() (*asset, error) {
	bytes, err := admin_http_assetsAppCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "admin_http_assets/app.css", size: 122, mode: os.FileMode(420), modTime: time.Unix(1498236484, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _admin_http_assetsAppJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x58\x4d\x8f\xdb\x36\x13\xbe\xfb\x57\x0c\x84\x05\x2c\x21\x8a\x9c\xf7\x3d\x14\xa8\x16\x46\xea\xa6\x0d\x9a\xc3\x02\xc5\x36\x6d\x0f\x5b\x07\xe0\x4a\x63\x2d\x61\x9a\x14\x48\xca\x6b\xc3\xd1\x7f\x2f\x48\x7d\x51\xb4\xb3\xd9\x78\xdd\xd6\x17\x5b\xd4\x7c\x3c\xf3\x70\x66\x38\xf4\x96\x48\x20\x65\x09\x73\xe0\xf8\x08\x84\x17\x15\x23\x32\xd9\x88\xbc\x62\x18\x06\x19\x91\xf7\x82\xbf\x96\xc8\xc8\xfe\x35\x2f\x82\x18\xee\x02\x5e\xdc\xa2\x12\x95\xcc\x30\x88\x21\xa8\x68\x72\x2f\x84\x56\x5a\x92\x32\x58\x46\xd7\x93\x09\x29\xcb\x24\x13\x5c\x4b\xc1\x18\xca\x30\xb8\x21\x94\xbf\xd3\xcc\xea\x5e\xa9\x4c\x94\x56\xef\x4a\x3a\x46\xae\x36\x22\x27\x46\x62\x55\xf1\x4c\x53\xc1\xc3\x46\x30\x86\x5e\x2c\x86\x46\x28\x3a\x4c\x00\x9a\xb7\x09\x61\x28\xb5\x82\x39\xdc\x2d\xaf\x27\x00\x26\x94\x77\x82\xaf\x68\x01\xf3\x41\x31\x0c\x66\x99\x5d\x9c\x05\x51\x27\xf5\x91\xdc\x33\xf4\x84\xb4\x59\x73\x64\x6e\xf1\x51\x52\x8d\xd2\x13\x93\xed\xb2\x9a\xa5\x94\xe7\xb8\x1b\x14\x7e\x64\x24\x5b\x33\xaa\xb4\xa7\x71\xdf\xad\x1f\xab\x2c\x8a\x42\x62\x41\xb4\xf0\xbd\x90\xfe\xc5\xb1\xd2\xad\xa8\xb4\x0f\x5e\x9a\x35\x35\x4b\xd7\xb8\x0f\x62\x38\xac\x71\x9f\xc2\xf4\x87\x35\xee\xa7\x75\x0c\x87\xba\xd7\xfd\x09\x95\xa6\x9c\x18\x8a\xbf\x6c\x61\x96\x0f\x52\xae\xfb\xc9\xc0\xfc\x96\x30\x9a\x2f\xf2\x5c\xa2\x32\xfc\xcf\x3e\xdd\x7d\x4a\x97\xaf\xfe\x4a\xef\xde\xbc\xfe\x7e\xf9\x2a\x4c\xed\x63\xf4\xf6\x6a\x76\xed\xe9\x58\xf0\x1f\xf7\x25\x5a\xad\x50\x21\xcf\xc3\x05\x63\x9f\xdf\x53\xa9\x74\x74\x43\x74\xf6\x10\x7d\x0e\x33\xc1\x15\x55\x1a\xb9\xfe\x85\xa8\x07\xca\x8b\x68\xe6\xdb\xc1\x02\x77\x30\x87\xb0\xcf\x98\x08\x4c\x66\x98\x8f\x44\x5d\x49\xde\x3f\x9a\x8f\x46\xa5\xd3\x21\xbb\xb6\x84\x55\x18\x8d\x24\xcc\xc7\x30\x44\xd5\x1f\xc6\x3e\xcc\x41\xcb\x0a\xaf\x3d\x09\x2d\xf7\x47\x5a\x60\x2b\xe7\x16\x8b\x9f\x77\x65\x6b\xd9\x57\xab\x21\x33\x91\x85\x27\x7c\x82\xe3\x71\x45\x98\x3a\x72\x59\x7b\xcf\x6d\x74\xad\x96\x2b\xdd\x49\xd6\x66\xb1\x8e\x42\xb3\x65\xde\x8e\x15\xc5\xfb\x8a\x67\x97\xe6\x8d\xae\xa0\x79\x05\xf3\x39\x04\x64\x5b\x04\xa7\xe2\xec\x1d\x9c\x22\xd6\x8f\x72\x6c\x32\x47\xa6\xc9\xe5\x8d\x4a\xba\xc5\x4b\x5b\x65\x44\xe9\x4b\xdb\xdc\x90\xdd\xc5\x4d\x52\x7e\x69\x93\x4a\xe7\xb8\xbd\xb8\xd1\x6a\x73\x8e\xc9\xd3\x35\x73\x54\x5f\x4f\x56\x4c\xdb\xca\xfb\xe4\xa7\xf9\x2e\x6a\x80\xd8\xe3\x23\x29\x50\x0f\x55\x94\x13\x4d\xa2\x0e\x66\x6b\x40\xb7\xa7\x8c\x79\xd7\x78\x6d\x1a\x71\xed\x7b\x09\xed\x72\x73\x74\x8d\xcd\x66\xab\xa2\x8b\xbe\x95\xcf\xba\x03\x2e\x5b\x15\x0d\x6a\xd7\x1a\xc7\x47\xe7\xdc\x6a\xfa\x52\xf3\xd8\xf8\xe8\x0e\xce\x3c\x77\xc4\x8e\x5a\xc1\xe9\xe3\xf5\x94\x93\xe4\x4a\x91\x2d\x86\x51\xa2\x1f\x90\x0f\xb0\x25\xaa\x72\xd8\xb5\x67\x62\x73\x44\x07\x52\x00\xea\xb8\x37\x8b\x52\x1e\x59\x1d\x40\x1e\x36\xaa\x48\x01\xa5\x4c\x0c\xe1\x09\x4a\x29\x64\xbd\xf4\x88\xef\xf5\x24\x6e\xc4\x16\x4f\xb1\x30\xec\xb3\x49\x45\xcb\xb7\xdc\x84\xd3\x85\x44\xd8\x8b\x0a\x54\xd5\xfe\x78\x24\x5c\x83\x16\x90\x23\x43\x8d\xd0\x0d\x06\x80\xdc\x9c\x12\x5c\x24\x30\x85\x57\x60\xac\x0d\xa0\x7b\xde\x1a\xa5\xf0\x30\xb5\xa7\xeb\x34\xa5\xf9\xae\x7e\x92\x05\x3f\x6f\x38\x3e\x2e\x8a\xa2\x25\x72\x98\x23\xc2\x83\x39\x5c\x53\x08\x48\x61\xa6\xb5\x0f\x5c\xa3\xdc\x12\x96\x7e\xf7\x26\x86\x3f\x09\xd5\x29\xfc\xef\xff\x6f\x6a\x2f\x17\x46\x63\xc8\x19\xd9\xb0\x28\x8a\x6f\x4c\x84\x45\x9e\xbf\x04\xfa\x7f\x91\x2a\xa7\x49\x3a\x33\x59\x86\xf9\xee\xc9\x74\x19\x7c\x5e\x24\x61\xba\xc9\xd1\xd6\x9e\xf9\x1d\x1e\x7e\xa5\xd9\x9a\x61\xda\xb4\xc6\x18\x7e\x2b\x85\x60\xa9\x6d\xaa\x31\x78\xef\xda\xdd\x31\x23\xdb\x82\x31\x3b\xaa\x05\x7e\x26\x75\x1e\xce\x69\x29\x46\xf5\x5b\xfb\xc9\x3f\x14\xd0\xbf\x92\x5f\x03\x71\x26\x66\x97\x34\x3b\x8b\x7f\x8d\xb9\x86\x2f\x4b\xd7\xa1\x8e\xc1\xea\xc4\x23\xe2\xc7\x11\x40\x1d\x37\x70\xbd\x18\x9e\x8d\xfe\x14\xf2\xa6\x34\xdc\x9b\xcf\x4b\x2b\xa3\xbf\x2d\x3d\x59\x18\xbd\xc7\x97\xd7\x45\x7b\x10\xf8\x89\xbb\xc6\xfd\xd7\xe0\xbf\x9d\xba\x9d\xdd\x6e\x47\x8f\xc6\x5c\xbd\xcc\x5d\xea\x1c\x2c\xe3\x4b\x9a\x8b\x28\x86\x67\xb0\x3a\x82\xe5\xde\xf8\x8e\xd1\xc5\x70\x36\x6d\xa2\x44\xee\x93\x06\x76\xd3\x5b\xdf\xe6\x32\x65\x6f\xec\x1f\xb8\xd2\x84\x67\xf6\xca\x6a\x17\xac\x6e\xd8\x21\xd4\xb8\x29\x19\xd1\xf8\xbb\x64\x29\x4c\xab\x32\x27\xba\xd9\x8d\x1b\x2b\xfb\xa0\x37\x6c\xda\x66\x2e\xac\x71\x7f\x2f\x88\xcc\xbb\x12\x6e\x97\x87\x7f\x1b\x52\x07\x4b\xff\x17\xc2\x08\x45\x5b\x29\xee\x4c\xd9\x91\xdf\x46\x63\xbf\xaf\xfd\xb7\x62\x3d\x0a\x74\x3c\x93\x8e\x5d\x24\x19\x13\x0a\x43\xd7\xac\x73\x23\xac\x8f\x4c\x67\x46\x89\x3d\xdf\x7c\x4e\xd5\x86\x2a\x15\x4e\x1b\xc5\xe9\x29\xe3\x5d\xb1\x83\x75\x91\xb6\xae\xe2\xfe\xa6\xa7\x04\xdb\x62\xea\x78\xb1\x30\xd3\x2f\x42\x68\x27\xe8\xee\xff\xa1\x4c\x94\xfb\x51\x80\xea\x8e\xe6\xbb\xa5\x8b\x64\xe2\x7e\xb7\x53\x2a\x8c\x33\x22\x91\xa8\x2a\xa6\xc7\xcd\x1e\x42\x6f\x83\x9c\x26\x19\xba\x64\x0e\xed\xa8\x36\x8e\xff\x0e\x00\x00\xff\xff\x76\x2a\x2d\xe7\xc9\x12\x00\x00")

func admin_http_assetsAppJsBytes() ([]byte, error) {
	return bindataRead(
		_admin_http_assetsAppJs,
		"admin_http_assets/app.js",
	)
}

func admin_http_assetsAppJs() (*asset, error) {
	bytes, err := admin_http_assetsAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "admin_http_assets/app.js", size: 4809, mode: os.FileMode(420), modTime: time.Unix(1508436520, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _admin_http_assetsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x3b\x7b\x6f\xdb\xb6\xf6\xff\xf7\x53\x30\xc2\x80\xb4\xf8\x45\x52\xe2\x6e\xd8\x16\xd8\x06\xb2\xac\xfd\xad\xf7\xae\x6b\x91\x6e\xf7\x81\x61\x77\xa0\xc5\x63\x89\x09\x45\xaa\x24\xe5\xc4\x37\xd8\x77\xbf\x20\x25\x5a\x92\xad\x97\x9b\x57\x9b\x3f\x62\x51\x3c\x3c\x3c\x4f\x9e\x73\x48\x71\x7a\xf0\xe3\xbb\xf3\x5f\xff\xfd\xfe\x15\x4a\x74\xca\xe6\xcf\xa6\xe6\x07\x31\xcc\xe3\x99\x07\xdc\x43\x3c\xf6\x71\x96\xcd\xbc\x08\xcb\x85\xe0\xbe\x04\x86\xd7\x3e\x8f\x3d\x03\x09\x98\xcc\x9f\x21\x34\x4d\x41\x63\x14\x25\x58\x2a\xd0\x33\x2f\xd7\x4b\xff\x3b\xaf\xea\x48\xb4\xce\x7c\xf8\x98\xd3\xd5\xcc\xfb\x97\xff\xdb\x99\x7f\x2e\xd2\x0c\x6b\xba\x60\xe0\xa1\x48\x70\x0d\x5c\xcf\xbc\x37\xaf\x66\x40\x62\xa8\x8d\xe3\x38\x85\x99\xb7\xa2\x70\x9d\x09\xa9\x6b\xa0\xd7\x94\xe8\x64\x46\x60\x45\x23\xf0\x6d\xe3\x08\x51\x4e\x35\xc5\xcc\x57\x11\x66\x30\x3b\xd9\x41\x43\x40\x45\x92\x66\x9a\x0a\x5e\xc3\xb4\x03\x86\x73\x9d\x08\xb9\x03\xc1\x28\xbf\x42\x12\xd8\xcc\xa3\x91\x41\x90\x48\x58\xce\xbc\x20\x08\x83\x20\x5c\xe2\x95\x79\x19\xd0\x48\x78\xf3\x67\x06\x5a\x53\xcd\x60\x7e\x5e\x08\xec\xc2\x0a\xec\x97\xff\x47\x98\xa4\x94\x4f\xc3\xa2\xd3\xc2\x1d\xf8\x3e\xfa\x19\x6b\x50\x1a\x45\x22\xcd\x28\x03\x82\x30\x27\x28\xa5\x9c\x2e\x29\x10\x74\xfe\xe1\x03\xf2\xfd\x2d\x0a\x94\x5e\x33\x50\x09\x80\x76\x74\x84\x61\x8a\x6f\x22\xc2\x83\x85\x10\x5a\x69\x89\x33\xd3\x88\x44\x1a\x6e\x5e\x84\x2f\x83\x49\x70\x1c\x46\x4a\x55\xef\x82\x94\xf2\x20\x52\xca\xab\xa8\x79\x67\x05\x84\x19\xd2\x09\xa4\xf0\x80\x73\xfb\x76\x82\x2d\x0a\x7a\xe7\xc1\x59\xb6\x45\xec\x4f\xbf\xbe\xfd\xf9\x1b\xa4\x12\x9a\x5a\xa9\x5d\x80\xca\x04\x27\xc1\xa5\x42\x6f\x5e\x7d\x87\x54\x9e\x19\xb3\x41\x62\x59\x02\x02\x83\x14\xb8\x56\x85\x88\x81\x50\x8c\x3e\xe6\x20\x29\x28\xc7\xe7\x81\xef\xff\x4e\x97\x88\x69\xf4\xe6\x15\xfa\xfe\x0f\xfb\xae\xb0\x1a\xa4\x64\x34\xf3\x8c\x21\xab\xd3\x30\x14\x4a\x05\x25\xd7\x86\x51\xe3\x30\xdf\xa8\x84\xae\xc2\x97\xc1\xb7\xc1\xa4\x6a\x5b\xf6\x2e\x95\x37\x9f\x86\x05\x9a\xb1\x18\x65\xc1\x4a\x78\x12\x7c\x1d\x4c\x5c\xab\x1d\xdb\xc1\xef\xc0\x09\x5d\xfe\x61\x58\x98\x86\x85\x47\x3e\x9b\x2e\x04\x59\x23\x29\x98\x31\x7c\x11\xe5\x86\x6f\x0f\x55\x92\x7b\x4d\x6f\x80\x20\x8e\x57\x0b\x2c\x1d\xf3\x84\xae\x50\xc4\xb0\x52\x33\xaf\xec\x28\x7e\x7c\xca\x57\x20\x15\x78\x25\x3e\x8e\x57\x34\xc6\xd6\x8f\xcc\xb8\xe6\x48\xe3\x36\x98\x72\x90\x65\x5f\x1b\x5e\xdf\x10\x59\x83\x40\x68\x8a\xb7\x20\x16\x12\x73\xb2\xd1\xbc\xb7\xed\x4a\xd3\x10\x6f\xd0\x87\x84\xae\x7a\xe6\x8a\x04\x63\x38\x53\x80\xdc\x43\x7d\xda\x9c\xd5\xa0\x1d\xbb\x1c\xaf\x6a\x30\xd6\x2a\x1d\x14\x8e\x34\x5d\x41\xa3\xd7\x12\xbf\xa1\xf3\x27\x91\x42\x8d\xb8\x82\x40\x46\xb7\xd0\x75\x8d\x5f\x60\xf2\x16\xb4\xa4\x91\x0a\x27\x5f\x27\xc1\xa5\x32\x22\xfe\x01\x1b\x63\xb5\x6f\x7b\x31\x4f\xc3\x9c\xb5\x0b\xe5\xc0\xf7\xc3\x80\xe3\x55\x25\x0b\xdf\x9f\x57\x30\xe5\xc3\xb3\x2e\x45\x96\x6a\x4f\x31\x2d\x82\x81\xe9\x91\x82\x31\x90\x33\xef\x2d\xa6\xfc\x5c\xb3\x9a\x21\xb4\xa9\x42\x8a\xeb\x62\x24\x13\xf8\xaa\xae\x75\x06\x52\x9b\x0e\x09\x19\x60\x3d\xf3\x8a\x17\x94\x23\xfb\xa0\xbc\xf9\xed\xad\x7d\x0a\x52\x15\xff\xf5\xd7\x34\xb4\x8d\x1a\x82\x06\xbd\xcc\x4f\x89\x7f\x32\x69\xea\x2e\x99\xcc\xff\x81\x19\x25\xd6\x5e\xa7\x61\x32\xd9\x92\xbd\xc6\x0b\x06\x0e\x47\xd1\xb0\xff\x0d\x97\x04\xb8\x02\xb2\xa5\x6d\x33\xc6\x85\xbd\xed\xf7\x72\xf7\xa5\x05\x2f\xcd\x77\x1a\xea\xa4\x0b\xa2\xd4\x3b\x9a\x04\xc7\x7d\x60\x25\x2f\x80\x84\x24\x20\xdb\x21\xa7\xe1\x2e\x21\x06\xb2\x85\xe8\xa9\x36\xeb\xc4\x1e\xac\x90\xf9\xed\x6d\x24\xf8\x92\xc6\x41\x25\xd6\x3f\x19\xac\x80\xfd\xc9\x20\xc6\xd1\xda\xa8\x49\xb7\x48\x67\x68\x74\x3a\x39\xde\x6f\x28\xfc\x69\x65\xd0\x35\xa8\x43\x0a\xdb\xfc\x4e\x43\xab\xee\xba\x1b\xd5\x8c\x78\x8c\x89\x15\x46\xf6\x03\xc3\xd1\x15\xa3\x4a\x3f\x99\x8d\xbd\xc5\x3a\x4a\x50\x26\x61\x49\x6f\x7a\x2d\xcd\xc2\xa9\x7c\xa1\xb4\xa4\x3c\x1e\x06\x95\x10\x43\x2f\xc6\xb3\xc8\xe8\x51\xdd\x93\x35\xd6\xd7\x83\x85\x59\x0b\xac\xac\x82\x85\x13\xf0\x8e\xac\x3a\xa5\xe2\xec\x66\x11\x14\x52\xe9\xb6\xaf\x0a\x72\x23\x97\x31\xc0\x56\x32\x03\x80\x1b\xbd\xc3\x8d\xf6\x23\xe0\xda\x04\xbe\x29\x2e\x96\x43\x1a\x5d\xcd\x3c\x09\xa9\x58\xc1\xc6\x80\x9e\x7f\x45\x39\x81\x9b\x17\xde\x7c\xba\x89\x3b\x31\x5b\x67\x89\xc9\x31\xd1\xe6\xc9\x2f\x86\xf9\x11\x95\x11\x03\x2f\x9c\x9b\xe0\x70\x27\x57\x68\x71\x06\x54\x5a\xf7\x05\x5c\x4b\xaa\x41\xaa\x16\xeb\x5e\x0a\x99\x96\xb9\xb3\x2c\xc1\x5e\x0b\x99\x7a\x8e\x76\xd3\xef\x63\x42\x90\x7d\x88\xa5\xc8\x33\x94\x60\xe5\x2f\x01\xc8\x02\x47\x57\x2e\xbe\x2c\xed\x20\x1e\xfb\x2a\x5f\xa4\xd4\x04\x03\x42\xdc\xbc\xcf\x5f\x78\x88\x8b\x55\xe9\xf8\x4f\xe2\x5e\xef\x18\xe9\xf3\x81\x5f\xe0\xba\xdf\x93\x9e\xc8\x83\x64\xe5\x41\x4e\x3b\x6a\x7f\x0f\x92\x81\x60\x64\xd8\x23\x64\xc0\xe1\x7a\x0c\x58\x8a\xef\xcb\x6d\x36\x16\xf2\x74\x5e\x83\xf6\x0f\xa4\x0d\x16\xbb\xfc\xa2\x7d\xa0\x1d\x4c\x79\x96\xdb\xc4\x29\x15\xc4\x54\x4b\x1c\xae\x9d\x20\x82\x77\x8c\x78\xa5\x3f\x0a\xf3\x58\x9f\xa5\x4c\xdd\x3c\x94\x31\x1c\x41\x22\x18\x31\x49\x9c\x05\x93\xa6\x3e\x97\xd0\xa1\x11\xe4\x42\xa1\xf1\xcf\x44\x5c\x37\x7d\xdd\x58\x47\xf0\x15\xe5\xd6\x41\x7b\xe8\xb6\x58\x54\x86\x79\x1f\x1a\x90\x52\xc8\x20\xc3\x5a\x83\xe4\xde\xfc\xd5\x4d\x06\x91\x36\xc5\x8a\xe0\x3e\xa4\x99\x5e\x23\x17\xb8\x0c\xa6\x1e\x72\x9b\x81\xbc\xd9\x35\xc2\xf2\xee\x59\x2d\xbf\xc0\xb5\x53\x0b\x37\x8f\xc3\x6a\x31\x60\x4f\x41\xbf\x5e\x67\x86\xc8\x3c\x5d\x98\xe4\xbf\x9d\x9b\xb7\xf8\xc6\x71\x93\x9a\xc7\x56\x6e\xee\x64\x52\x29\xbe\xb9\x0f\x93\xb2\x68\xba\x4c\x0a\xa3\x82\x4b\xa4\x05\x22\xc0\x85\x06\x94\xe2\x1b\x24\xc1\xea\xa1\xd8\x2b\x78\x8e\x51\x26\x14\x35\x75\x5f\x09\x7d\x84\x84\x44\xfe\x89\x89\x67\x88\x0b\xc4\x68\x4a\xf5\x8b\x87\xb2\xc6\x8e\x8e\x45\xae\xb5\xe0\x4e\xec\x0b\xcd\xd1\x42\x73\x5f\xa5\xf6\x27\x93\x34\xc5\x72\x6d\x9f\x17\x4c\x98\x10\x5b\xe8\xb4\x88\xac\x56\xa7\x84\x2a\x13\x14\xc8\x96\xb8\x2a\x89\x9f\x11\x32\x0d\x8b\x69\xf6\x21\xbb\x6d\xd5\xdc\x2b\xdb\x08\x8d\x0d\xed\x66\x20\x67\x71\x2c\x21\xc6\x5a\x0c\xe5\x20\x38\x8e\xef\x2d\xfd\xa8\x26\xfd\x0c\x12\x90\xd7\x39\x8f\x8a\x22\xb6\x3b\x8f\xb8\x18\x4a\xd5\xdf\xe5\xda\xb8\xb8\x61\x16\xeb\x3e\xc0\x73\x1c\x25\xd0\x07\xf0\xc6\x84\xe4\x15\x66\xe8\xb9\x82\xe8\x45\x1f\xe4\x3f\x31\xd5\xc3\x50\x0f\x97\x02\xe1\x2a\x05\xc2\x95\x15\xed\x9f\x04\xe1\x60\x99\xf3\xe1\xec\x06\x8f\x2a\x0b\x2c\xe0\xbb\x5c\xbf\x4e\xf5\x18\x48\xab\x8e\x31\x80\x4e\x2d\x63\x60\x8d\x62\xee\x29\x0d\xab\x79\xca\x9d\x13\xb1\xbb\xac\x29\xe8\xb3\xc8\xc5\xce\xe2\x38\x78\x9d\x73\x17\x21\x97\xe6\x71\x38\xde\x3b\x07\xaf\x02\xa7\xc1\x5a\x46\xad\x99\x67\x17\x9f\xb3\x38\x36\x60\x7d\xf4\x34\x42\x6a\xb9\x1a\x1a\xcb\xfd\xc4\x68\xda\xc0\xd0\x19\x48\x4b\xf5\x53\xc1\xd1\xb2\x64\xe3\x14\xe1\x55\x7c\x84\x08\x30\x8d\xcd\x8f\xa4\x2b\x38\x42\x0c\x2b\x7d\x64\xa2\xec\x11\x4a\x29\x3f\x42\x4a\x13\x58\x99\x80\xaa\xf2\xf4\xcb\xc8\xe8\x8c\x72\xed\x3a\xeb\x6d\xaa\x5e\xdb\x18\x56\x70\x09\xb8\xa3\xd4\x02\xdb\xbe\x2a\xb5\xd8\xee\xa8\xd4\x12\x47\x97\x5a\x2d\x66\x24\x21\xce\x19\x96\x08\x6e\x32\x09\x4a\xd9\x10\xf4\xa5\x28\xaa\x58\x61\x37\xf5\x50\xae\x97\xa6\x35\xa2\x24\xca\xf5\x50\xfe\xfa\x20\x9c\x34\x52\xf0\x28\x81\xe8\x6a\x21\x6e\xbc\x5d\xbe\x6c\x3c\x70\x6c\x95\x8d\x36\xae\x3e\xb3\xfa\xc1\x90\xee\x22\x94\xa3\x9e\x6e\xda\xad\x6a\xa9\x7b\x4b\xf8\x9f\xdf\x8f\xfd\xef\xff\xf8\xbf\xaf\xc2\xbd\xbd\xc5\xcd\x72\x47\x87\xa9\xd0\x0c\xd6\x14\xcf\x29\x47\x0a\x4c\x0a\xa8\x5e\xd4\x0a\x8c\x8f\x39\xe6\x9a\xfe\x97\xf2\x18\x39\x64\x9f\xad\x3b\x0d\x68\xd2\xe4\x0f\x4e\x8b\xd7\xf6\xf9\x21\x35\x68\x66\xb8\xa3\xf6\x0a\x14\x9f\xaa\x39\x53\x1a\x1a\x0c\x5f\x6e\xa9\xe7\xe4\xf0\xd9\x56\x79\x17\x22\xd7\x76\x6b\xa7\x77\x97\x59\xe4\x1a\xee\x6f\x8b\xd9\x60\x7b\xe2\xf2\xae\xa8\x80\x6c\xe9\x53\x1c\xcf\x5f\xc3\x21\x63\xc8\xb8\xa0\xc9\x95\x15\x4a\x40\xba\x6f\x31\xda\x46\x5a\x1e\xd0\x15\xac\x7b\x6b\x43\x0b\x64\xac\x63\xf8\xc8\xe7\x49\x4e\x91\x08\x31\xe9\x45\x1f\xc8\x87\x4c\x08\xd6\x07\xf0\x9e\x46\x57\xac\x8f\x3f\x14\x09\x66\xbc\x77\x36\x41\x8f\xb4\xe9\x6e\xa4\xbe\x47\xb1\x69\xd6\x81\x62\xf5\xea\xa9\x9d\x32\x86\xd7\x1e\xc2\x92\x62\x3f\xa1\x84\x00\x9f\x79\x5a\xe6\x60\x3f\x0c\x31\x4b\x53\xcf\x31\xaa\x43\x4b\xf9\x52\x78\x76\x2f\xfe\x0a\xfa\xcf\x6c\x77\x47\x18\x1b\xda\x73\x48\x6a\x6c\x00\xe4\xe0\x41\x5c\xef\xe0\x11\x67\x73\xbd\xe3\x07\xea\xf2\xed\xb1\x1b\x5b\xf1\x5e\x7a\xa3\x45\x5a\x0d\x9a\xb4\x9f\x57\xd8\xe5\xc6\x4a\xfd\x81\xcf\x2a\xac\x89\xd5\xad\x92\x18\xab\x94\x01\x01\xa5\x29\xaf\x7f\xbf\xd3\xc2\xd3\xdc\x2c\x43\x83\x76\x18\x25\xb0\x92\x86\x50\x1a\x27\xba\xcf\x20\x7d\xbf\x53\x80\xe5\x74\x9f\xde\x5b\x88\xd8\xd2\x78\x7b\x48\x30\x8f\x41\x1e\x9e\xa2\x03\x12\x08\xce\x28\x87\x23\x74\x68\x14\x73\x78\x8a\xdc\x9b\xbf\x8c\x55\x90\xd1\x26\x79\x2f\x93\x8c\x3d\x56\xbe\xeb\x3c\xa3\x4e\xa4\x3f\x71\x0e\x5c\x2c\xd0\xf7\x8f\xbd\x75\x93\xa9\x1d\xbf\x35\xc0\x4d\x5e\x47\x02\x65\x22\x82\xd7\x67\xa3\x09\x21\x5e\xd8\xea\xb9\x3d\x1e\x5d\xe3\x01\x39\x26\xd0\xc3\x71\x91\xd9\xb0\xd5\xcb\x86\xb8\xf2\x15\x8d\xf9\x63\xb1\xd2\xbf\x88\xfd\x58\x2d\x21\xc5\x52\x76\xf4\x94\xa7\xaf\x9f\xf0\x11\xd3\x63\x44\xd8\xfd\xaa\xae\xd6\xb3\x43\x13\x2b\x82\xbf\xc3\xda\x15\x59\xbf\xae\xb3\x8e\x3a\xbf\xda\x32\x6c\xee\x38\xb9\xc4\xb0\x75\xd6\x47\x64\xa2\xa0\x7c\x90\x8b\x06\xf1\x0a\x38\x39\x63\xcc\xa6\x8f\x7d\x7b\xa2\x76\x06\x8b\xb4\x83\xaa\xe6\x21\xa3\xab\x1e\x2c\x4d\xc3\x05\xe5\xf6\xe1\xe2\xd6\xf0\xae\x62\xb2\x4e\xfc\x91\x6d\xbd\xa6\x52\xe9\xb2\x2d\x24\x32\x79\x3d\x55\x1a\xb8\xfe\x09\xab\x64\xe0\x3c\xbb\xb3\xa4\x7c\x44\x15\xbe\xb7\x71\xd2\x29\xd1\xb5\x86\xd5\x58\x42\x3e\x31\xf5\x1f\x5c\x00\x76\x0c\xd4\x5e\x0c\xf3\x50\x01\x3f\x31\x1b\x8d\x8d\xe7\x8b\xb1\x1b\xcf\x17\x7b\x6f\x3c\x3f\x22\x4b\x65\xdd\xe7\x98\xda\x34\xf7\x58\xe6\x2a\x14\x3b\xe7\x25\x65\xcf\x63\xf2\xd8\xbb\x79\x5b\x5a\x63\x91\xb6\x94\x96\xd8\xc8\x61\x06\xb7\x6f\x9f\x90\xec\xf7\x65\x9e\x52\x2e\x01\xcd\xac\x65\x24\xe1\x1d\x11\xf9\xc1\xf7\xbe\xaa\x85\x7b\xf4\xee\xd7\xfd\x7f\x4d\xb9\xbd\xef\xd5\xfc\x4e\xbf\x6a\xd4\x3e\xc9\xdf\x7c\xa5\x7f\xe0\xfb\x28\xdc\x7c\x95\x6f\x77\x83\xdc\xeb\x1f\xdc\xa5\x1a\x14\x09\x09\xe8\x6f\x78\x85\x3f\xd8\x0b\x22\x16\xd9\x6c\xef\xbf\xda\x75\x18\xf4\xde\x38\x1a\x41\x58\x23\x9d\x00\x02\x4e\x90\x58\xda\x47\x77\xad\x04\x29\x61\xdb\x19\x8e\x41\x21\x26\x30\x41\x4b\xac\x34\x6c\xee\x95\xb4\x5d\x77\xc1\x97\xf8\x26\x88\x85\x88\x19\xe0\x8c\x2a\x7b\xe7\xc5\xbc\x0b\x19\x5d\xa8\xf0\xf2\x63\x0e\x72\x1d\x9e\x04\x27\x27\xc1\x49\xd9\x1a\xbe\x4a\x33\xfe\x22\xd2\xe5\xf6\x1d\xa8\x26\xde\x2e\xa2\x23\x41\x20\xc0\xdc\x9e\x86\x5d\xaa\x40\xc8\x38\x3c\x09\x26\xc1\xc9\x71\x58\xbe\x1c\x7f\xdd\x67\x10\x95\x2f\x41\x89\x5c\x46\x30\x86\xef\x88\xf0\x4b\x15\x44\x4c\xe4\x64\xc9\xb0\x84\x2d\x71\x3a\x94\x39\xf5\x2b\x49\x1c\x1b\xe1\x1e\x87\xf5\x77\xbe\xce\x98\x1a\x90\x85\xbd\x85\xd5\x45\x4e\xe1\x7e\xa6\x00\x0a\x79\xec\x6b\x48\x33\x86\x35\x78\x88\x92\x99\x97\x67\x04\xeb\x62\xfb\xe3\xad\x20\x98\x05\x89\x4e\x59\xcb\x05\xa2\xd4\x74\x6e\xdf\x10\x9a\x26\x2f\x9b\xfd\xf6\x0e\x9d\x37\xff\xcd\x22\x45\xd6\xb7\x4f\xd1\xed\xad\x7d\x70\x3b\x5a\xc9\xcb\x86\x33\xa1\xad\x6d\xe4\xe5\x1d\x76\x90\x77\xb6\x8a\x77\x39\x30\xeb\x41\xfd\xaa\x51\x0d\xa0\x9a\x63\xeb\x9e\x11\x5e\x00\x33\x14\xcc\x3c\x13\xc5\xbc\xf9\xfb\x22\x96\x4d\x43\xdb\xd3\x80\xdd\x0e\xa7\x05\xe7\x66\x80\x5b\x9e\x2d\x8a\x31\xf9\x59\x99\xbb\x8e\x49\x0f\x9a\xa9\xb4\x41\x6a\x53\xdf\xae\x2c\x7a\x2b\x77\xae\xc1\xf7\x9f\x38\xf7\x9d\x34\x6f\xdf\xcf\xe8\xbe\xae\x31\x42\xcc\x98\x10\xe9\x55\x3b\xcf\x63\xc5\x6c\x06\x38\x31\x5b\x14\x23\xc4\x7c\x32\xf9\x36\x38\x0e\x8e\x83\x93\xd3\xc9\xf1\xf1\xd7\xbd\xdf\x78\xb4\xe4\x2c\x2d\x82\x37\x13\xef\x23\xf8\x02\xbe\x43\xf0\xa7\xe8\x30\x11\x4a\x9f\x66\x42\xea\xc3\xfd\x84\xde\x79\x01\xaf\xf0\x83\xa5\x10\xcd\x9d\x90\xb6\xa0\x6e\x42\x38\x81\x25\xce\x99\xf6\x6a\xbb\x0d\x11\xe6\x11\xb0\xe7\x2f\xbc\xf9\x39\x13\x0a\x76\x63\x75\x47\x82\x50\x66\x06\x5b\x19\xc0\xb2\x11\xfc\x6b\xd3\x88\x2b\x33\x45\xb1\x96\x6c\xcf\xd1\x08\xca\x2e\x78\x57\x6b\xdf\x34\x2c\xe2\xfe\x34\x2c\xee\x4e\xff\x2f\x00\x00\xff\xff\x15\x89\xd7\xa8\x4c\x3d\x00\x00")

func admin_http_assetsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_admin_http_assetsIndexHtml,
		"admin_http_assets/index.html",
	)
}

func admin_http_assetsIndexHtml() (*asset, error) {
	bytes, err := admin_http_assetsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "admin_http_assets/index.html", size: 15692, mode: os.FileMode(420), modTime: time.Unix(1508436520, 0)}
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
	"admin_http_assets/app.css":    admin_http_assetsAppCss,
	"admin_http_assets/app.js":     admin_http_assetsAppJs,
	"admin_http_assets/index.html": admin_http_assetsIndexHtml,
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
	"admin_http_assets": {nil, map[string]*bintree{
		"app.css":    {admin_http_assetsAppCss, map[string]*bintree{}},
		"app.js":     {admin_http_assetsAppJs, map[string]*bintree{}},
		"index.html": {admin_http_assetsIndexHtml, map[string]*bintree{}},
	}},
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
