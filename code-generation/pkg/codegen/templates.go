// Code generated by go-bindata.
// sources:
// assets/aws-service-operator.yaml.templ
// assets/base.go.templ
// assets/cft.go.templ
// assets/operator.go.templ
// assets/template_functions.go.templ
// assets/types.go.templ
// DO NOT EDIT!

package codegen

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

var _awsServiceOperatorYamlTempl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x4d\x8f\xdb\x36\x10\xbd\xeb\x57\x0c\x76\x03\x14\x28\x2a\x39\xee\xa9\x10\x76\x17\xd8\xda\x6e\x6a\x64\xeb\x5d\xd8\x4e\xaf\xc9\x98\x1a\xcb\xac\x29\x92\x20\x29\x3b\x8e\xe0\xff\x5e\x50\x5f\xfe\xf6\x2e\xda\x44\x27\x71\xf4\x38\xef\xcd\xe3\x70\x74\x0b\x0f\xd5\x03\xfd\x67\x18\x3d\x4f\x61\xd0\x1f\x4e\x61\xfa\xe7\x70\x02\x7f\x0c\x9f\x06\x70\xd7\x3e\xc1\x2d\x4c\x17\xdc\xc2\x9c\x0b\x02\x6e\x01\x73\xa7\x52\x92\x64\xd0\x51\x02\x2b\x8e\xf0\x05\xd7\x36\x54\xda\x47\x94\x81\xe6\xdb\x97\xe0\x16\x86\x73\xd8\xa8\xfc\xa7\x04\x04\x5f\x12\xb8\x05\x01\x5b\xa0\x4c\x09\x50\x6e\xdc\x82\xcb\x14\x70\xa6\x72\x07\xae\x25\xc8\x70\x49\x40\x09\x77\x16\x9c\x2a\x77\x44\x8e\x32\x2d\x82\xdb\x5a\x80\x2c\x83\x7a\x99\x76\x98\x4a\x28\x25\xd9\x41\x6b\xc9\x59\x48\xb8\x21\xe6\x94\xd9\x44\x41\x80\x9a\xff\x4d\xc6\x72\x25\x63\x58\x75\x83\x25\x97\x49\x0c\x4f\xdc\xba\x80\x3b\xca\x6c\x1c\x84\x50\xc5\x46\x98\x91\xd5\xc8\x28\x00\x38\xda\x04\x90\x91\xc3\x04\x1d\xc6\x01\x00\x80\xc4\x8c\x62\xf0\xa5\x5a\x32\x2b\xce\xa8\x2d\x39\x28\x0a\x53\x56\xf5\x8e\xcb\x84\xbe\xfe\x02\xef\x48\x50\x46\xd2\x41\x7c\x0f\xd1\xd0\x33\x6e\xb7\x41\x78\x40\x80\x9a\xd3\x57\x47\xd2\xaf\x6c\xb4\xfc\xcd\x46\x5c\x75\x56\xdd\x19\x39\xf4\xd4\x95\xba\x5e\x6e\x9d\xca\xc6\x64\x55\x6e\x18\xf5\x69\xce\x25\x77\x5c\xc9\x0b\xda\x8a\xa2\x21\x8e\x26\x9a\x58\xd4\x6c\x8c\x5e\x44\x6e\x50\x6c\xb7\xd1\xb1\xf2\x08\xd7\x36\x00\xb0\x9a\x58\x95\x28\x35\x2a\xd7\x31\x5c\xc0\x55\x44\xb6\x82\x36\x22\x8f\x59\x3f\x72\x99\x6c\xb7\x35\x44\x70\xeb\x3e\x5e\x81\x95\x67\x52\x41\x75\x29\xf2\x0d\x55\xd4\x78\xbb\x50\xc6\x8d\xf6\xf5\x14\x45\x08\x67\x0e\xc2\x9f\xc2\x85\x94\x13\x9f\xa3\xac\xa9\x4d\x1b\x36\x02\x22\x9f\xbb\x0d\xfb\xdc\xb4\x57\x98\xe5\x32\xcd\x05\x9a\xb7\xea\xb5\x4c\xe9\x6b\x47\x34\xf1\xdf\x6b\xec\x6a\xd7\x85\x28\xf4\x02\xbb\x41\x51\x54\xdc\x4d\xdb\xf6\x44\x6e\x1d\x99\xb1\x12\xc7\x8d\x6b\x66\xc8\x22\xcc\xdd\x42\x19\xfe\x0d\x7d\xb3\x9c\x36\xd7\x9b\xfb\x1a\xc0\xe4\xa2\xf2\xb7\xec\xde\x0f\xbe\x3b\x6a\xbb\x43\xb8\xb9\x29\x5f\x4c\x5d\x42\x1b\xb7\xc4\x0c\x39\x5b\xaf\xb4\x4a\x9a\x57\xa6\xe4\x9c\xa7\x19\x6a\xdb\x22\x4b\xc2\x66\x49\x2b\x92\xf5\xbe\x15\x99\x59\x9b\x30\x25\x57\xbf\x89\xa6\x5d\x42\x58\xa3\x63\x8b\x26\xb1\x21\x74\x54\x2f\x12\x12\xd4\x2e\x72\x9d\xec\xbe\xe8\x7a\xcb\x99\x5a\xce\x5d\xc7\xf3\xd5\xb1\xf2\x4e\x36\xe1\xa4\xbd\x93\xdf\x4f\xf7\x19\x79\x17\xef\xe3\x89\xbc\x9b\x9f\x6f\x4e\x85\xf8\x60\xdb\x3b\x93\x2a\xd7\x23\x63\x2a\x97\xee\x7f\xcd\xbd\x76\x22\xf8\x09\x7a\x01\x73\xae\x67\x7f\xe7\x32\xe1\x32\xfd\xc1\xad\xab\x04\x8d\x69\x5e\x21\x1b\x47\xaf\xb0\x04\xbb\x91\x76\x78\xbd\x5e\xe1\xb1\xf9\xec\x1f\x62\xae\xbe\x25\x17\x3c\xfe\x7e\x36\xf6\x49\x0b\xb5\xf1\x13\xe4\xc8\x3e\xd4\xda\xfe\x37\xa7\x5e\x67\xdf\xff\x45\x18\xd2\x82\x33\xb4\x31\x74\xcb\x75\xf9\x7b\x46\x47\xcd\x14\x3e\x24\x2e\xcd\x97\x52\xb9\xd2\x6a\xbb\x0b\x02\x70\xcc\x22\xcc\xf0\x9b\x92\xb8\xb6\x11\x53\x59\xc7\x1f\xd9\x15\x95\xe5\x0f\x05\x67\x24\x0e\xd2\xa0\xd6\x57\xf7\xec\x94\x97\xab\x83\x93\x19\x5d\xb7\xc5\x3f\x4c\x49\x87\x5c\x92\xd9\x23\x0d\x5f\xf3\xb3\x2e\x30\xc3\xb4\x42\xd5\xa0\x06\xd3\x39\xb7\x31\x5e\xbd\x8f\xde\x47\xdd\xb0\x1c\xf9\xbf\x1e\xa7\x79\xc9\x85\x78\x51\x82\xb3\x4d\x0c\x8f\x62\x8d\x1b\xbb\x6f\x81\x49\x0f\x2c\x69\x46\x06\x99\xa3\x60\x18\xb2\xaa\xb5\x43\x5f\xc0\xfd\x5d\xef\xe9\xd3\x64\x3a\x18\x7f\x1e\x3d\xfe\x35\x78\x38\xc1\x1a\x4a\xb9\x92\xf7\x77\xe3\xc1\x87\xe1\xf3\xe8\xf4\x3b\x56\x26\x86\x3c\xb9\xbf\x7b\xec\xf5\x9e\x3f\x8d\xa6\x9f\x87\xfd\x87\x7f\x03\x00\x00\xff\xff\x1b\x6a\x36\x65\x5b\x0a\x00\x00")

func awsServiceOperatorYamlTemplBytes() ([]byte, error) {
	return bindataRead(
		_awsServiceOperatorYamlTempl,
		"aws-service-operator.yaml.templ",
	)
}

func awsServiceOperatorYamlTempl() (*asset, error) {
	bytes, err := awsServiceOperatorYamlTemplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "aws-service-operator.yaml.templ", size: 2651, mode: os.FileMode(420), modTime: time.Unix(1543602919, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _baseGoTempl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\xbf\xae\x9b\x30\x14\x87\xe7\xf8\x29\x8e\xa2\xab\x0a\x10\x98\x1d\xb5\x53\xbb\x74\x68\xaa\xf6\x0e\x1d\xaa\x0e\xc6\xf7\xc4\x17\x25\xd8\xd4\x3e\x34\x54\xc8\xef\x5e\xd9\x40\x43\x86\x2b\x92\xe9\xc4\xff\x7f\xdf\x97\x43\x27\xe4\x49\x28\x84\x5a\x38\x64\xac\x69\x3b\x63\x09\x12\x06\xb0\x97\x46\x13\x0e\xb4\x0f\xbf\x55\x43\xaf\x7d\xcd\xa5\x69\x4b\x71\x71\x67\x51\xbb\x50\x0b\x87\xf6\x4f\x23\xb1\x30\x1d\x5a\x41\xc6\x96\xdd\x49\x95\xd2\xe8\x63\xa3\x1e\x3e\xf6\xbb\xc7\x1e\x5b\xa1\x85\x42\x1b\x0e\x8f\x63\x01\x56\x68\x85\xf0\xd4\xe8\x17\x1c\x72\x78\xc2\x33\xb6\xa8\x09\xaa\x0f\xc0\x3f\x13\xb6\xce\xfb\x47\x5f\x59\x06\xae\x1c\xc7\xe5\x3e\xfe\xdc\xa1\xe4\xdf\xd1\x99\xde\x4a\xe4\x07\xd1\xa2\xf7\x4b\x04\xd4\x2f\xde\xb3\x94\x31\xfa\xdb\x4d\x96\xc0\x91\xed\x25\xc1\xc8\x00\x26\x56\xc8\xa6\xca\x3f\xc6\xc2\x00\x22\xcc\x97\x09\x06\xb2\x35\x1a\xff\xb6\x5a\x7a\x04\x73\x23\x2f\x64\x1b\x1b\xf8\xd7\x19\xfd\x06\xcc\x33\x76\xec\xb5\x84\x03\x5e\x92\xb7\x78\xf2\xfb\x81\x72\x96\x42\x16\x25\x05\x3b\x16\xa9\xb7\x1a\xde\x85\x89\x30\x5e\xee\xaf\xe6\x9a\xc7\xb9\xf5\xd5\xd5\xcd\x68\x5a\xbf\xdb\xd0\xa6\xa3\x6a\x6b\x03\x3f\xe0\x65\xd1\x94\xcc\x19\x6f\x12\xa5\xd7\x48\x93\x3f\x80\xab\xc3\xa4\x9e\xd8\x53\xf8\x21\x48\xbe\x26\x92\x06\x98\x3f\xa2\x60\x32\xd4\x1c\xb4\x68\xd1\x75\x42\xc6\x36\x6a\xb4\x4a\xa3\xa9\xbb\x19\x9b\x23\x84\x4e\x8f\xff\xce\x12\xde\xfd\xdc\x6f\x35\xf3\xaf\xf8\xca\x4e\x19\xa8\xf9\x96\x83\x67\x12\x96\xfe\x13\xac\x12\xa7\x91\x76\x77\x85\x67\xbb\xf7\x85\xa4\x81\x7f\x32\x1a\x93\x94\xf9\x7f\x01\x00\x00\xff\xff\x51\xa9\xcf\x48\x4c\x04\x00\x00")

func baseGoTemplBytes() ([]byte, error) {
	return bindataRead(
		_baseGoTempl,
		"base.go.templ",
	)
}

func baseGoTempl() (*asset, error) {
	bytes, err := baseGoTemplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "base.go.templ", size: 1100, mode: os.FileMode(420), modTime: time.Unix(1543602919, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cftGoTempl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x6d\x6f\x9b\xc8\x13\x7f\x0d\x9f\x62\xfe\x56\xd5\x3f\x44\x0e\xbe\xbe\xb5\x2e\x27\xf5\x9a\x3e\x44\xad\x92\x2a\x76\xdb\x17\x51\xd4\xae\x61\x8c\xb7\xc1\x2c\xc7\x2e\xf1\xf9\x2c\xbe\xfb\x69\x1f\x30\x0b\x36\x76\xdc\xf6\xda\xeb\xc9\xa8\x95\x61\x99\xf9\xcd\xe3\xce\xec\x84\xc1\x00\x7e\xd3\x17\x9c\x5f\xc1\xe5\xd5\x18\x9e\x9f\x5f\x8c\x61\xfc\xea\x62\x04\x2f\x2e\xde\x3c\x87\x5f\xd7\x97\x3b\x18\xc0\x78\x46\x39\x4c\x69\x82\x40\x39\x90\x42\xb0\x18\x53\xcc\x89\xc0\x08\xee\x29\x81\x4f\x64\xc1\x4f\x59\x26\x57\x58\x7e\x1a\xb2\x08\x63\x4c\x21\xcb\x59\x88\x9c\x7f\x92\x00\x17\x53\x58\xb2\xe2\xff\x11\x24\xf4\x0e\x41\xcc\x10\xc2\x19\x49\x63\x04\x92\x2e\xc5\x8c\xa6\x31\x90\x09\x2b\x04\x88\xb5\xa0\x39\xb9\x43\xc0\x88\x0a\x0e\x82\x29\x8e\x40\xe0\x3c\x4b\x24\x9a\xd6\x24\x55\xab\xd9\x5d\x3c\x30\x12\x07\x84\x73\x14\x1c\x22\x9a\x63\x28\x58\xbe\x0c\x5c\x37\x23\xe1\x1d\x89\x11\x56\xab\x60\x94\x61\x18\x5c\x23\x67\x45\x1e\x62\x70\x49\xe6\x58\x96\xae\x4b\xe7\x19\xcb\x05\x78\xae\xd3\xc3\x3c\x67\x39\xef\xb9\x4e\x2f\xa6\x62\x56\x4c\x82\x90\xcd\x07\x64\xc1\xe5\xff\x53\x1e\xdd\x9d\xc6\x4c\xde\xee\x26\xe0\x98\xdf\xd3\x10\x07\x61\xc2\x8a\x68\xca\xf2\x39\x11\x94\xa5\x3d\xd7\x21\x0b\xfe\xfe\x09\x49\xb2\x19\x79\x02\x2d\xfe\x84\x4c\x0c\x86\x66\x5e\xfb\x72\x20\xad\x23\x19\xe5\x83\xf6\x9b\x40\x8a\xbd\x37\x78\x9b\x1a\xed\x46\x0c\x59\x3a\xa5\xf1\xa1\x5c\x33\x4c\x32\x94\xfe\xf1\x5d\x19\x84\x4b\x5c\x40\x95\x06\x1c\x08\xa4\xb8\x00\x36\xf9\x8c\xa1\x70\xa7\x45\x1a\xca\xf7\x9e\x16\x04\x27\xfa\x37\x78\xa6\x7e\xfa\x5d\xc1\x80\x13\xcb\x47\x41\x45\xf4\x9a\xa6\x51\x59\xf6\x41\xb0\x8c\x86\x4f\xaf\x2f\x81\x8b\x9c\xa6\xb1\x0f\x27\xcf\x1a\x2e\x86\x95\xeb\xe4\x28\x8a\x3c\x85\xc7\xcd\x37\x2b\xd7\x71\x9a\x68\xc3\x2e\x15\xfa\xae\xe3\x68\x65\x87\x8e\xbc\xf4\x7d\xdf\x05\x80\xb5\x02\x43\xd0\x57\xf5\xdc\x77\x9d\xd2\x2d\x95\x4f\x5a\x1a\x45\x38\xa5\x29\x72\x95\xa7\x5d\x36\x87\x53\xc1\x5d\xb1\xcc\xb0\xcd\xcc\x45\x5e\x84\x42\x5a\x65\xbc\x58\x5d\x4d\x6f\xba\x2d\xd3\x76\x39\xd1\xad\x95\x36\x58\xda\x97\x46\xfb\x91\x20\xe1\x9d\x54\x0b\xb4\x1f\xb5\xe2\xa9\x5c\x60\x53\x75\xcf\x25\x05\x4c\x08\xc7\x08\x98\xde\x7f\xad\xad\xaf\x54\x52\xf1\xf7\x78\x3b\x42\x7e\x2d\xc1\xf3\x8d\x68\x2b\x6a\x26\xbf\x82\x9a\x88\x07\x95\xa9\x49\xc1\x05\xe6\x72\xb1\x0f\xbd\x0e\x57\xf6\xfa\xc0\x5b\x06\x07\x9a\x63\xeb\x32\xcf\x48\x88\xbe\x31\xfd\x25\x8a\xab\x42\x64\x85\xe0\xc6\x76\xcb\x5c\x66\x5e\x4c\x73\x36\x57\xcb\xe7\xc8\xc3\x9c\x4e\x50\x29\xca\x21\x24\x49\xd2\x6d\x72\x8d\xec\xf9\xe0\xcd\x49\x76\xa3\x0d\xbf\xd5\x3f\x7d\x50\x55\xc7\x97\x7e\xa8\x04\x0d\xcf\x60\x83\x6e\x55\xba\x0e\x47\xae\x5e\xae\xdd\xf2\xf4\xc3\x68\x84\x9c\x53\x96\xba\x0e\xbf\x0f\xe5\xbb\x66\xd9\x09\xe4\x26\x94\x6c\xbe\xeb\x3a\xca\x9a\x8b\xb4\x12\xd1\xa2\x6c\x1a\xa5\xc8\xe4\xbe\x59\x07\x43\x66\x3d\x59\xc8\xe8\x48\x75\x3c\x3b\x4c\xbe\xaf\xb6\x40\x65\x80\x32\x49\xe9\x79\x1f\xb6\x60\xbd\xc7\x96\x12\xbe\xeb\xd0\xa9\xa2\xfd\xdf\x19\xa4\x34\x91\x2e\xa8\x72\x21\xa5\x89\x82\x91\xb8\x8e\x2c\x36\x4c\x00\x2f\x72\x04\x3a\xd5\x3d\x82\x72\xc0\x7b\xd9\x64\x18\xe7\x74\x92\xa0\xc2\x4a\x30\xf5\xb4\x0e\x5a\x3b\xee\x4b\xe8\x27\xdb\x80\x59\xce\x95\x73\x7a\x29\xd3\x61\xae\x02\x8f\x11\x2c\xa8\x98\x81\x98\x11\x61\x12\x40\xee\x80\x9e\xaf\x4d\x9c\xb2\x1c\x3e\xf6\x65\x4e\x48\x0b\x73\xd5\xc3\x1a\x22\x6f\x7e\xb9\x0d\xaa\x54\x92\x72\x4d\x50\x6f\x4e\x58\x21\xcc\x8b\xd7\xb8\xbc\x85\x33\xb0\x56\xde\x93\xa4\x40\x2d\xc0\xe8\x69\xd8\xb4\x13\x4c\x71\xc9\x91\x08\xed\x48\x58\xd0\x24\x81\x50\x2d\x58\x99\x6a\x14\x47\xe0\x45\x96\x25\x14\x23\xc8\x48\x4e\xe6\xbc\x3b\x3b\x2d\x4c\x99\x9e\x5a\x2a\x9c\xb4\x72\xc3\xa2\xba\xb2\x42\x5c\x67\xee\x57\x66\x66\x38\x55\xad\x5d\x1a\x33\x3c\x5b\xd7\x81\x97\x28\x94\xba\x2f\x2a\x9e\xb1\x21\x5a\xd7\x85\xc3\x6a\x41\xf5\x38\x5e\x66\x58\x96\xc1\x76\xec\x8e\x82\xf1\x60\x5e\x53\x55\xf6\xed\x36\xcb\xa3\x87\x6f\x35\xa7\x12\xf8\xee\xfa\xcd\xd0\xa6\xac\xfd\xa8\xc8\x2e\x99\xa0\x53\x1a\x2a\x89\x4f\xaf\x2f\xf9\x10\x6e\x6e\x4f\x4c\x39\x71\x1d\xc7\x69\xc8\xa8\x1a\x83\xe2\x2c\xfb\x55\x32\x6a\xaf\xaa\xa6\x60\x85\x46\xab\xff\x56\x26\x97\xd7\xbb\xb6\x88\xba\x6a\xb0\x5f\x63\xbd\xc7\x5c\xe6\xc5\x5e\x38\x43\xb7\x0d\xb1\x45\xe2\xbb\x4e\x5a\xb9\xbe\x13\x76\x1d\x9c\x2e\x15\x4d\xe4\x9c\xb0\xee\x35\x9d\x60\x56\x3f\x52\x70\x9b\x6d\xca\x77\x57\xab\x53\x53\x1f\x1e\x09\x96\x49\x28\x2d\xf2\x77\x16\x2d\x83\x51\x38\xc3\x39\x09\xde\xe6\xb2\x75\x0a\x8a\xbc\x2c\x65\x23\x3f\x95\xf5\x6d\xcb\xaa\xc1\x49\x91\xcb\x23\xb7\x84\x6a\xd0\xd4\xac\x26\x33\xe8\x5f\x18\x69\x5e\x29\x3b\x78\x8d\xcb\xb2\x5c\xad\x0c\xbf\x7e\x94\x94\x12\xa9\x57\xaf\x57\x69\x55\x96\x3d\x8d\x88\x09\xc7\x87\xc3\xc8\x7f\x65\x19\x5c\x4d\x3e\x07\xab\xd5\xa3\xc6\x7e\x31\x00\x23\x75\xa4\x51\x7c\x2a\x02\xf5\xa3\x62\xed\xd5\x82\xd3\x7d\xea\xab\x8a\xb9\xee\x34\x55\x90\x6a\xf3\xbd\x3d\x3a\xf7\xd7\x3c\xe7\x44\x90\xd5\xd5\xe4\xf3\x50\xa5\xc5\xa3\xe6\xa1\x53\x9f\xb1\x86\x50\x57\x9c\x57\x9a\x6d\xb8\xe6\x97\x85\xcc\x2f\x77\x35\x33\xab\x2d\xca\x6d\x05\xb0\x43\xb7\xce\x94\xdb\x0c\x93\xa2\xef\xf5\xad\x73\x93\xdc\xca\x74\xba\xdc\x65\xbb\xf2\x9b\xef\xb7\xdc\x6c\x87\x5a\x25\x92\x87\x7f\xa8\xb4\x55\x01\x84\x9e\x5d\x07\x7a\x7e\x3b\x34\xbb\x74\x56\x18\x7b\x15\xde\xf0\xbc\x2e\x1a\x7e\x43\xb7\x07\x65\xb9\xc9\x47\xa7\x2d\x7b\x67\x4e\x5b\x4c\x07\x25\xf1\xfe\xac\x3d\x28\x4d\x7f\x70\x5e\x7e\xe3\x98\x6e\xb8\x61\x5b\xd6\xa9\xdb\xfa\xce\x75\xd4\x81\x05\x05\xe6\xaa\x6b\xde\xdc\xb6\x8f\x22\x6f\xab\xf7\xf2\x40\x6c\x11\x9f\x01\xc9\x32\x4c\x23\xaf\x5e\xeb\x83\x9d\xb6\xfe\xc3\xc9\xeb\xb6\xb2\x8f\x23\xb5\xfa\xc6\x3e\xda\xf0\x07\x35\x87\xbd\x8a\xed\xa8\x16\xdd\x75\xe2\x10\x50\xbf\x2b\xd8\xd6\x19\x29\x18\xa1\x58\xc7\x96\x5b\x50\x7e\xeb\x18\x32\x26\xf1\x66\x6e\x8e\x49\xfc\xa5\xe7\x90\xfd\x78\x5f\x76\x10\xe9\xc6\xfd\x92\x93\x48\x37\xda\x43\x8e\x22\xae\x23\x48\xdc\xb1\xa1\xc6\x44\xcd\x96\x8a\x60\x1d\x49\xf9\xd4\xdc\x3e\x63\x12\xfb\xbb\xa9\x6a\x7f\x76\x10\xda\xae\xe9\x20\x69\x5a\xec\x6f\xa6\xc8\x98\xc4\x5c\xd1\xfa\xae\x0b\x76\x01\x03\x3d\x6e\xda\xf3\x4c\x6b\xd6\xd4\x45\xcf\xcc\x52\xef\xb2\xa8\x39\x4b\x15\x6a\x41\x0d\x4f\xf8\x27\xe5\x82\xa6\xb1\x1e\xaa\xba\x67\x27\x0b\xc3\xd3\xec\xd1\xae\x3f\xbe\x74\x8f\x57\x16\xd0\xcf\x33\x5e\x19\x8b\x0f\x99\xaa\x0e\x66\x79\xe0\x30\x65\xf9\xef\x38\x4c\x1d\x87\xa9\xe3\x30\xf5\x6f\x18\xa6\xcc\x66\xff\x0e\x47\xd5\xe3\x04\xf5\x4f\x4c\x50\xf0\x5f\x9a\xa0\xbe\x5f\x32\x1e\xe7\xa6\xe3\xdc\x74\x9c\x9b\x8e\x73\xd3\x4f\x36\x37\xd9\xb3\xcc\xae\xb9\xe9\x1c\x13\x6c\xcc\x4d\x91\x5a\xa8\xbf\x41\x75\x8f\x4b\x16\xab\xe7\x83\xf7\x4d\x27\x9c\x7d\x9f\x36\xd7\x92\xf5\x7c\x50\x6e\xf8\xc7\xfe\xe6\x6c\x0f\x08\xd2\x5b\x1f\x6d\x47\xd9\x56\xec\x72\xd4\x07\x42\xc5\xbb\x54\xd0\x44\x91\x6a\xae\xe8\x50\x97\x6d\x05\xf9\xee\xce\xfb\xea\xef\xc2\x60\xf9\x6f\x9b\x49\xcf\xd8\x3c\x93\xbf\x9d\x0e\xfd\x3b\x00\x00\xff\xff\xdd\xfa\xa4\x1b\xae\x24\x00\x00")

func cftGoTemplBytes() ([]byte, error) {
	return bindataRead(
		_cftGoTempl,
		"cft.go.templ",
	)
}

func cftGoTempl() (*asset, error) {
	bytes, err := cftGoTemplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cft.go.templ", size: 9390, mode: os.FileMode(420), modTime: time.Unix(1543602919, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _operatorGoTempl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5f\x73\xdb\xb6\xb2\x7f\x26\x3f\x05\xca\xdb\xa6\x64\xaa\x90\xcd\xdb\x1d\xcf\xf5\xed\xb8\xb6\xda\x7a\xea\xd8\x3e\xb6\xd3\x3c\x64\x32\x29\x4c\xae\x24\xc6\x24\xc0\x02\x90\x15\x57\xd5\x77\x3f\xb3\x58\x80\x22\x65\x49\x96\x9a\x4c\xcf\x99\x73\xea\x17\xf1\x0f\xb0\xd8\xdd\xdf\x6f\x17\xbb\xa0\xb3\x8c\xfd\x3f\xfd\xb1\x93\x0b\x76\x7e\x71\xc3\x86\x27\xa7\x37\xec\xe6\xa7\xd3\x6b\xf6\xc3\xe9\xd9\x90\xfd\x5f\xfb\x17\x66\x19\xbb\x99\x94\x9a\x8d\xca\x0a\x58\xa9\x19\x9f\x1a\x39\x06\x01\x8a\x1b\x28\xd8\x7d\xc9\xd9\xaf\x7c\xa6\x5f\xc8\x06\x9f\x48\xc5\xfc\xbb\x5f\x71\xe6\xe9\x88\x3d\xc8\xe9\xd7\x05\xab\xca\x3b\x60\x66\x02\x2c\x9f\x70\x31\x06\xc6\xc5\x83\x99\x94\x62\xcc\xf8\xad\x9c\x1a\x66\xda\x15\x6a\x7e\x07\x0c\x8a\xd2\x68\x66\xa4\x9d\x91\x1a\xa8\x9b\x0a\xa5\x91\x0a\xc2\x3e\x6d\xee\xc6\x59\x2e\x0b\x18\x83\xc8\xb8\xd6\x60\x34\x2b\x4a\x05\xb9\x91\xea\x21\x0d\xc3\x86\xe7\x77\x7c\x0c\x6c\x3e\x4f\xaf\x1b\xc8\xd3\x2b\xd0\x72\xaa\x72\x48\xcf\x79\x0d\x8b\x45\x18\x96\x75\x23\x95\x61\x71\x18\x44\xb9\x14\x06\x3e\x9a\x28\x0c\xe6\xf3\x17\xac\x1c\x31\x9a\x72\xaa\x8f\xa7\xda\xc8\xba\xfc\x1d\x8a\xc5\x22\x0c\x58\x0d\x86\xdf\xbf\x64\xd1\xdd\xff\xea\xb4\x94\x19\x6f\xca\x9a\xe7\x93\x52\x80\x7a\xc8\x50\x1d\xde\x94\x3a\xc3\x41\xd9\xfd\xcb\x28\x0c\xa2\x71\x69\x26\xd3\xdb\x34\x97\x75\xc6\x67\xba\xe2\xb7\x1a\x7f\x5f\x68\x50\xf7\x65\x0e\xad\xc7\xec\xdc\x09\x54\x0d\x28\x8d\xd3\x14\x8c\x2a\xc8\xbd\x3a\x20\x70\xf1\x3d\xa5\xe5\x52\x8c\xca\x71\x14\x32\xb6\xd7\xb4\xdf\xa6\x30\x85\x9a\x0b\x3e\x06\x85\x93\x7b\xee\xf8\x07\xbe\x5c\x2c\x42\xc6\x72\xa9\xa0\xef\x87\x0c\x1f\x59\xab\xfb\x2b\x96\x5c\xe4\xb2\x42\x91\x99\x36\x2a\xe7\x1a\xec\x08\x6d\x54\x29\xc6\x68\x2c\x9f\xe9\xbc\x2a\x41\x98\xfd\x14\xa5\x39\xee\x47\x83\xc9\xee\x41\xe9\x52\x0a\x28\x32\xf3\xd0\x40\x91\xad\xce\x4a\xf9\x4c\x67\xf7\x2f\x79\xd5\x4c\xf8\xde\xd8\x58\xb7\xf4\xe0\xd8\x6f\xbe\xbf\xc1\x75\x9d\xcf\x48\xf3\x17\x63\x99\x19\x29\x2b\x9d\xe5\x3c\x9f\x40\x14\x5a\x87\xfc\xe2\xd4\xdc\xcf\x25\x96\x7c\x4f\x99\xdd\x03\xd4\xb3\x9b\x9b\x52\x0a\x9d\x5e\x52\xc4\x58\x84\x23\x1f\x36\x9b\xc6\xf4\xbc\x91\x84\x18\x9c\x17\x3e\x01\x28\x68\x14\x68\x04\x86\x71\x86\xb1\xa5\x64\x55\x81\x62\xf2\xf6\x03\xe4\x86\x8d\x64\x7b\xa9\x8d\x54\xc0\x72\xbb\x06\x53\x2e\x42\x75\x88\x10\x2e\xc5\x69\xa3\xa6\xb9\x61\xf3\x30\x20\x56\x33\xfa\x7b\x4e\x77\xe9\xb1\xfd\x09\x03\x23\x9b\x32\x3f\xba\x3a\xb7\x2f\x89\x60\x61\x60\x91\x7b\x45\x84\x66\xcf\xbb\xf4\x26\x3a\xbb\x57\xe1\xc2\x5a\x70\x0e\xb3\x76\xd5\x5c\x01\x37\xd0\x55\x1f\xf5\x9e\x71\x93\xdb\xac\xb5\xd5\x00\x37\xb9\x08\x47\x53\x91\x77\xa5\xc6\xce\x82\xbe\xee\x03\xb6\xa3\x9a\x09\x7b\xde\xea\x37\x5f\x01\xd3\x47\x27\x99\x9c\x1b\x55\xb1\x83\x43\x12\x9c\x9e\xc3\xcc\xad\x3c\x60\x6e\xe5\xa3\x37\xd7\xc7\x3e\x78\x06\xec\xe5\xb7\xc9\xd2\x81\x03\xf6\xbe\x9d\x8a\x72\xd2\x2b\x18\x97\xda\x80\x8a\xa3\x0d\xc9\x34\x4a\xfa\x9e\x4e\x8f\x8a\x22\x5e\x8a\xeb\xd9\xf3\x13\x17\x45\x05\xea\x87\xa9\xc8\x63\xab\xf3\xeb\xa6\xe0\x06\x54\x92\xf4\xd3\x9d\x02\x33\x55\x82\x3d\xf3\x06\xcf\xc3\xc0\xe1\x7f\xe0\x08\xe0\x2c\x0a\x83\x0d\x8e\x68\x0d\x3a\x60\xad\x2e\x6e\xb0\x0b\xe2\x9e\xd2\x07\x3d\x14\x06\x61\xb0\x70\xac\xb8\x36\x5c\x99\x37\x08\x3c\xc1\x0f\xda\x52\xa1\x14\xda\x70\x81\x68\xcb\x11\xbb\x20\x3a\x5c\xaf\xa7\x03\x17\x05\xe3\xb9\xd1\x4c\xda\xad\xab\x26\x5e\xc4\xf9\x12\xce\xa4\xb3\x4a\x9c\x9b\x8f\xcc\x6d\x49\xc8\x10\xfc\x1d\x30\xc1\x6b\xd0\x0d\xcf\xc1\x71\x3b\x41\x06\xf8\x25\x9c\x53\x35\xe2\x66\x33\x49\x8b\xd0\xf0\x1e\x84\xe9\xf8\x5c\xa3\x1f\x8f\x8a\x02\xaf\xad\x23\xf3\x54\x8a\xa3\xa2\x40\xcf\x10\x12\xf4\x06\x1f\xd3\x3d\xbe\x39\x81\x0a\xba\x6f\xe8\xde\xfa\x28\x0c\x30\xd7\xe0\xc2\x6d\xce\x41\xba\x3d\xa6\xca\x65\x35\x55\xbc\x5a\x2c\xa2\x8e\x2d\x03\xb6\x6a\xc1\x80\xe5\xe9\x1a\x8a\xa6\x57\xc3\xeb\x1b\xba\x8b\x91\x29\xb8\x56\x4a\xde\x7a\xd6\xc9\x98\xa9\x5f\xf5\xe7\x12\x11\x9e\x2f\x06\x2c\x37\x1f\xd3\x13\x29\x00\xa7\x2d\xc2\x70\x2d\x55\xb2\x8c\x75\xa9\xc8\x66\x65\x55\x31\xc3\x5d\xc1\x52\x83\xd6\x7c\x8c\xa8\x2b\x59\xdb\x27\x96\x28\x16\xd5\x46\xc9\x1c\xb4\xee\xa0\xda\x15\xb4\x29\xdc\x6b\x3d\x5e\x89\xf2\x57\xb4\xc6\xf7\xb2\x78\x48\x18\x28\x45\x01\x5e\xc9\xf1\x98\x7c\xeb\x04\x9c\xd9\x07\x61\x70\xcf\x95\x75\xe2\x63\x5a\x84\x41\x39\x42\xf9\xa9\x55\x81\xdf\x56\x80\x82\x02\x1c\xc6\x0e\xed\x0b\x0f\x08\x86\xae\x7b\x43\x02\xe8\xf5\xb9\xbf\x0f\x83\x05\x83\x4a\xd3\x7c\xda\xae\xae\x31\x57\xd8\xe4\xd0\x6e\xdb\x08\xf6\x0f\x52\x91\x65\xce\x5e\x02\xcb\x5e\x26\x61\xd0\x92\x54\x0f\xd0\x32\x6b\x8d\x97\xd6\xe2\x45\xe4\xa0\x6c\x12\x47\x51\x92\x9e\x95\xda\xc4\x54\x6f\xd9\xeb\x8b\xc6\x6e\x40\xf3\x05\x4a\x2c\x47\x56\xd2\x17\x87\x4c\x94\x95\xd5\xcf\xb9\x2a\x7d\x53\x9a\xc9\x10\xdd\x17\x83\x52\x49\x4a\x97\x11\x39\x74\x0c\xc6\x60\xea\xde\xcc\x4c\x94\xed\x13\x0f\x28\x15\x06\x01\x26\x09\x0c\xf6\xf7\x4b\xa6\xa2\x01\xca\x96\xb0\xad\x61\xe9\xa9\x81\x5a\x93\x22\xe5\xa8\x7d\x9e\x5e\x1b\x6e\xa6\x1a\x7f\xf2\xbb\xd3\x13\x76\x48\x1e\xbe\xe4\x4a\x43\xe1\x10\x7f\x1b\xd1\xdb\x22\x7a\x47\xf3\x3d\x54\xaa\x9b\x62\xdb\x17\x1e\xa9\xde\x5b\x07\x17\x69\xbb\xa0\xa0\x2c\x47\x96\x1a\xe8\xa3\x28\x62\xcf\x9e\x75\x88\x42\x8f\x70\x31\x2e\x84\x34\xb4\xb3\xa3\x55\x35\x6f\xde\x12\x8b\xde\xd1\x8f\x55\xc8\x29\x78\x12\x1d\x6c\xd5\x7e\xb0\x1c\x8b\x3a\x6d\x19\x6d\x5f\xbb\xf1\x1e\x83\x9b\x87\x66\xc3\x94\xde\x08\x3b\x6b\x41\x0c\xd8\x32\x96\xfc\x1e\xbd\x43\x8f\x47\x57\x17\x67\x67\xdf\x1f\x1d\xff\xfc\xfe\xf8\xe2\xd5\xe5\xd9\xf0\x66\x48\xb6\x07\xf2\xf6\x43\x4b\xc8\xc2\xe6\x33\xab\x5d\xbb\x51\xae\xc4\xd7\x60\xab\xf5\x89\x83\x7e\x95\x96\x7d\x36\x59\xd5\x03\x1f\x24\x90\x4b\x55\xe0\x66\x49\x28\x40\x61\x73\xf5\x28\xb6\x8a\x75\xa0\x19\xb8\x7a\x3b\xb5\xef\xd1\x0f\x6f\xb8\x12\xa5\x18\x0f\x98\xab\xa9\xd3\x1b\x79\xcc\x6b\xa8\x62\x57\x57\xa7\x37\xf2\x4c\xce\x40\xc5\x3b\xf8\x28\x49\xd6\x5a\xd6\x1f\x76\x05\x5c\x4b\x41\x66\xba\xac\xb0\x07\x02\x27\x43\x74\xfb\x36\xff\x4f\x6d\xbe\xa4\x49\x7f\x0e\x80\x1d\x8c\xd8\x69\x50\xd7\xd2\xff\x12\x40\x03\xb4\xf1\x90\x95\x22\x57\x50\x83\x30\x57\xb2\xaa\x6e\x79\x7e\x77\x2c\xa7\xc2\x6c\x42\x63\x0f\xff\x74\xb6\x91\xbf\x41\xef\x80\x7e\x2e\x55\xcd\xab\x7f\x51\x10\x87\xb4\x4d\x38\xc3\x44\x59\x85\x8b\xb0\x53\x7e\xaf\x29\x50\x6d\x99\x88\x66\xb1\x52\x18\x50\x23\x9e\xc3\x7c\x61\xab\x50\xbb\x79\xc8\xdb\x0f\x69\xfc\x7c\x73\x2d\x96\xa4\x27\x00\xcd\xb1\x6c\x1e\xe2\x24\x6c\x7b\x17\x21\xcd\xfa\x66\xf4\xa8\x28\x6c\x23\x5a\x8e\x98\xf6\xfb\x68\xdf\x1c\x9b\x5a\x22\xf6\xc7\x1f\x5b\x07\xac\xc9\x3d\xb6\xea\x1d\x19\x54\xda\x36\x46\xa9\xe7\x9e\x2d\x3e\x7d\xab\x90\xd8\x71\x72\x6a\x9a\xa9\x59\x16\x2e\x23\x93\x1e\xdb\xde\x8e\xf6\x0a\x1a\xb4\xca\x18\xd7\x9e\xa4\xbd\xa2\x6d\x6d\x69\x32\xf2\xb5\x89\x6d\x18\xd7\x16\x27\x54\x13\xb1\xaf\xbf\xd2\x5f\x47\x03\xa6\xed\x7d\xe2\x96\x20\xf4\xec\xcd\x22\x5c\xb7\xe6\xa9\x18\xc9\x51\x1c\xf1\xa2\x80\x62\xab\x68\x36\x2b\xcd\x84\x69\x57\xa8\xf4\x16\x1b\xb8\xca\x32\x7e\x4e\xce\x70\xe5\x4c\x91\x24\xdb\x96\xbc\x2f\x61\xc6\xb8\x61\x13\x63\x1a\x7d\x90\x65\xb9\x14\x5a\x56\x90\xf2\x99\x4e\x79\xcd\x7f\x97\xc2\x9e\x6a\xe4\x95\x9c\x16\x23\x0c\x04\x84\x3d\x9b\xc8\x1a\xbe\xfb\x9f\xcc\xea\x91\x15\x60\x78\x59\x7d\x47\x4a\x15\x87\x5f\xe9\x68\x8b\x2a\x61\x10\xbc\x27\x98\x56\x13\xcb\x12\x5e\x6f\xce\xb2\x68\xda\x28\x70\xc0\xa2\xe3\xab\xe1\xd1\xcd\xf0\xfd\xe9\xf9\xfb\xcb\xab\x8b\x1f\xaf\x86\xd7\xd7\xd1\x80\x45\xd1\x86\xfa\x73\x17\xb4\x3d\xd8\x56\x3f\x04\x5b\x53\xf8\x52\x40\x22\x86\xd4\x9c\x56\xda\xf6\xaf\x1b\x8e\x60\x6c\x64\x74\xcd\xea\x35\xce\x9b\x62\x97\x3a\x92\x58\x56\xc5\x05\x26\x26\x01\xb3\x8b\x35\x91\x2c\xa5\x0d\x65\x3b\x68\xf7\x68\x0e\x84\x9d\x46\x32\xf7\x49\x02\x36\x01\xec\x17\xbc\xe8\x69\xca\x44\x42\x26\xd6\x67\x4f\xe7\x12\xb2\xdd\xa7\x13\x77\xd0\x9a\x9e\x6a\x0b\xf6\xb1\xac\x1b\x2c\xfe\x62\xb9\x41\x93\x01\x1b\xf1\x4a\x43\x82\x45\xf4\x17\xee\x68\xd6\x9a\x31\xfc\x6d\xca\x2b\x9c\xe6\x6d\xc4\xa4\xbe\x58\x0c\xd0\xa6\xfe\xa3\x64\x5b\xd6\x91\x72\xa7\xb4\xf3\xda\x93\x3a\xbf\x73\xa6\x7f\x96\xc4\xd3\x72\xf1\xe9\xec\x20\x60\xc6\x1a\xae\x78\xad\xd9\x57\xdf\xdc\xdb\xf6\x57\x56\x05\x5e\x47\xd6\x66\x8a\x2e\x21\xd1\xa4\x7d\x33\x14\x85\xec\x2e\x39\xca\x69\xe0\x52\x54\xbb\xec\x7f\x48\x8e\x92\xde\x1e\x77\xf1\x64\x96\x7a\x7d\x79\xf2\x97\x67\x29\xfa\x2e\xf0\x44\x9a\xf2\x41\xb7\xc2\x74\xa4\xee\x0e\xe9\x8a\x0e\x98\xfe\xb2\x6a\x83\x96\x43\x6b\x76\x29\x0c\x82\x4e\x58\x9e\x74\x3a\xc7\x24\x5c\xe3\xf3\xbd\xa2\xd1\xf6\xa1\x7b\x94\x01\xbe\x32\xa5\x62\x6e\x03\xcd\xa9\xb9\xdd\x1e\x5b\x5d\x99\x3b\x80\xeb\xfd\xb5\x6d\x1b\x0a\xb7\x7d\x4b\xb3\xa0\x6f\xed\x38\x1e\x1d\x9a\xd9\xe3\x0c\x0a\x84\x75\xa7\xa1\x4f\x1f\x9a\x7d\xc2\x19\x96\x3f\x6e\xd9\xf5\x04\x6b\xd9\x26\xa5\x3f\x82\x89\xa9\xa9\x71\xa7\x59\x3f\x42\xef\x30\xeb\x31\x63\x3e\xdb\x51\x56\xb7\x6b\xa1\x62\x9f\x46\x61\x58\xd8\x53\x2c\x3f\xab\xbb\x97\x77\x07\x39\xe1\x5d\x74\x3a\x67\x4f\x1b\x46\x7c\xf3\x32\x0c\x97\x79\x6e\x77\x47\xb9\x0a\xa5\x2b\xfd\xcf\xba\xa7\x4d\x5d\x5e\xd8\x1a\x77\xf4\x5b\x1f\x22\xe4\x9a\x8e\x74\x3f\x1a\x0e\xda\x12\xba\x73\x8f\xd5\x8c\xbf\x55\xb6\x07\x6b\x49\xbb\x2d\x85\x0d\x88\xd2\xc9\xdf\x9c\x5e\xc3\x69\x51\x56\x83\x4f\x22\xf6\xfa\x9a\xd3\xc1\xb5\xcb\x60\xea\xa6\x6d\x38\xe0\xc5\xfa\x29\xed\xc9\xaf\xe7\x05\x15\xbd\x9b\x4a\x50\xdd\x2f\x38\xed\xf6\xd1\xd9\x8f\x1c\x03\xbb\x0b\xf9\xbd\x9e\x0a\x03\xdd\xab\x18\x11\x14\x7a\x1c\x7f\xfa\xa1\xb9\x5b\x20\xf2\x87\x07\x36\xd7\xd3\x31\xf8\x97\x8d\x92\x0d\x2e\x4a\xb8\x5d\xb8\x1a\x25\x9f\x40\xcd\xd3\x4b\x25\x1b\x50\xa6\x04\xcd\xec\x17\x37\xb7\x29\xd8\x39\xe9\x0d\xd4\x4d\xc5\x8d\xff\xff\x8a\x60\x3e\xa7\xe7\x3f\xc3\x03\xf2\xdf\x32\xda\xfb\x6a\x39\x36\x8e\xfc\x38\xf7\x0c\xec\x47\x25\x3f\xf0\x84\x1b\x3e\xbf\xb8\xfd\x70\xb0\xe2\xa8\x63\xf7\xf1\x90\xdc\xb8\xe8\x7e\x99\xb0\x90\x39\xbd\xbd\xec\x6b\xfb\xd5\xd9\x6a\xc2\x0e\x59\x4f\x33\xff\x11\xd1\xed\x91\x7b\x88\x71\x5e\x7c\xfb\xc8\x02\x3b\x20\x7a\xb7\xf2\x79\xb2\x73\xbd\xf8\xf7\x4e\xaa\xfd\x78\xdc\x8f\xe2\x64\x93\x7e\x10\xf9\x51\x51\x94\x98\x45\x78\xe5\x63\x4d\xaf\x65\xfd\x5e\x74\xc6\x2a\xc8\xdb\x80\x8b\x74\x4d\xf0\x7c\xee\xec\x07\x7d\xd2\x74\x77\x87\xc7\x1f\x09\x3e\x6d\x73\xf8\x3b\xfd\x7f\xa6\xf4\xbf\x25\x47\xba\xfc\xe8\xc2\x66\x97\x7a\x7d\xed\x0a\xcb\xf9\x6f\x78\x69\x5e\x0b\x53\x56\x56\x04\x49\x2b\x68\x7b\xe9\xd1\x87\x26\x7b\xea\x6c\x27\xf7\x23\x1a\x69\xb6\xad\xb5\x61\xe8\xbd\x0e\x3b\xfe\x42\x12\x74\x8e\xce\x88\x06\xfe\x4c\x6d\x57\x22\x3c\xaa\xc1\xdc\x37\xd5\xf5\x7b\x36\x76\x14\xcb\x16\x62\x8d\x03\xd3\x6b\xfa\x7f\x27\xdc\x5c\xb4\xbf\x3c\x38\x64\x6f\xfd\x07\xcc\xc5\x7c\xee\x12\x68\x67\xbf\xea\x7e\xc9\xdd\x49\x74\xdb\x40\xc2\x6f\xcb\xd9\xf6\x6c\x87\x45\xc3\x8f\x06\x94\x20\x0f\x45\x34\xf4\x4b\xb5\xd2\x61\x75\xb6\x31\x3a\x3d\xee\x4e\x72\xeb\xc4\x9d\x66\x73\xed\x49\x65\xd4\x15\xdc\x15\x60\x77\xbf\xf9\x3c\x2e\x45\x01\x1f\x3b\xfa\x5d\x4a\x65\x34\xfb\x36\xb1\x17\x0b\xc4\xa3\x75\xd1\x21\xe3\x4d\x03\xa2\x88\xfd\x13\x14\xb0\xaa\x76\xaf\xa1\xeb\x5f\xee\x8a\x49\x2b\x71\xeb\xb8\x43\xe6\xd5\xf0\x68\x3d\x0d\x3c\x51\xf8\x15\x6f\x70\x19\x72\x5d\x8d\x37\x9f\x09\xfc\x9e\xf8\x75\x98\x62\x95\xb1\xf1\x9b\x79\x67\xb9\xd2\x40\x8d\xe3\x96\xf3\xed\x4c\xbb\xbd\x23\xa2\xf8\x9e\xaa\x8a\xe8\x80\xb5\x0f\x7e\xe1\xd5\xd4\xe2\xba\x5a\x04\xec\xc4\xae\x56\xf9\x27\x39\xb5\xde\xae\xc4\xff\xa7\x9d\xf5\x68\xcb\x95\xe5\xb3\x5d\xd8\xb2\x3b\x7a\x5b\x49\xd2\x19\xe9\x77\x3f\xd4\xa0\x25\xca\xee\x55\x51\x2f\x77\xad\xd4\x45\x3b\xe4\x2a\x7f\x96\xf9\xe8\x5b\xd9\x3f\x03\x00\x00\xff\xff\xdb\xc3\xc9\xb8\x4b\x2d\x00\x00")

func operatorGoTemplBytes() ([]byte, error) {
	return bindataRead(
		_operatorGoTempl,
		"operator.go.templ",
	)
}

func operatorGoTempl() (*asset, error) {
	bytes, err := operatorGoTemplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "operator.go.templ", size: 11595, mode: os.FileMode(420), modTime: time.Unix(1543602919, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _template_functionsGoTempl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\xc1\x8e\xdb\x36\x10\x3d\x8b\x5f\x31\x15\x02\xd4\x2a\xbc\x22\xf6\x56\x04\x4d\x0f\x6d\x36\x89\x91\x60\xb7\xc8\x1a\xc8\xb1\xa1\xa5\xb1\x34\x30\x45\x12\xe4\xc8\x8e\x21\xe8\xdf\x0b\x92\xb2\x03\x6c\x91\xee\x56\x17\x4a\x23\xce\x7b\xf3\x1e\x67\x28\x25\xfc\x9e\x1f\x78\xfb\x00\xf7\x0f\x5b\xb8\x7b\xbb\xd9\xc2\xf6\xc3\xe6\x11\xde\x6d\x3e\xdd\xc1\x6f\xd7\x47\x48\x09\xdb\x9e\x02\xec\x49\x23\x50\x00\x35\xb2\xed\xd0\xa0\x57\x8c\x2d\x1c\x49\xc1\x57\x75\x0a\x37\xd6\xc5\x88\xf5\x37\x8d\x6d\xb1\x43\x03\xce\xdb\x06\x43\xf8\x1a\x01\x36\x7b\x38\xdb\xf1\xe7\x16\x34\x1d\x10\xb8\x47\x68\x7a\x65\x3a\x04\x65\xce\xdc\x93\xe9\x40\xed\xec\xc8\xc0\x57\xa2\x41\x1d\x10\xb0\x25\x0e\xc0\x36\x65\xd4\x8c\x83\xd3\x11\x2d\x57\x62\x52\xd4\x1d\x3a\xb9\x30\x4a\x15\x02\x72\x80\x96\x3c\x36\x6c\xfd\xb9\x16\xc2\xa9\xe6\xa0\x3a\x84\x1e\xb5\x43\x1f\x84\xa0\xc1\x59\xcf\xb0\x12\x85\x3a\x85\x46\x13\x1a\x86\xb2\x23\xee\xc7\x5d\xdd\xd8\x41\xaa\x53\xd0\x6a\x17\xe2\x7a\x13\xd0\x1f\xa9\xc1\xab\x34\x99\xc8\x52\xce\xb2\x04\x64\x79\x44\x1f\xc8\x1a\x6c\x25\x9f\x1d\xb6\xf2\x69\x56\xad\x4e\x41\x1e\x6f\x95\x76\xbd\xba\x2d\x45\xf1\xff\xd8\xac\xd9\x53\x57\x8a\x62\x40\x56\xc7\x5b\x28\x0f\xbf\x86\x9a\xac\x54\x8e\x06\xd5\xf4\x64\xd0\x9f\xd3\x46\xe5\x28\xc8\xb8\x49\x1e\x6f\x4b\x51\x89\xe8\xd3\x3d\x9e\xe0\x0b\x69\x0d\x1e\x79\xf4\xe6\x62\x42\x34\x74\x87\xe0\xa2\x5d\x2d\x90\x59\xfc\x4d\xf6\x2a\xc6\x20\xf6\xa3\x69\x62\xf2\xaa\x82\x0f\x4b\xca\x24\x8a\x05\x64\x89\x4c\xa2\x00\xf8\x38\xee\xd0\x1b\x64\x0c\x9f\x31\xd8\xd1\x37\x78\xaf\x06\x7c\xfd\x83\xf8\x5a\x14\xc5\x34\xdd\x80\x4f\x27\xff\x8a\x4c\x8b\xdf\xd6\xf0\x0a\x35\x0e\xf1\x14\x5e\xbf\x81\x7a\xc3\x38\x04\x98\x67\x51\x14\xef\x91\xa7\xe9\xf2\xb3\x7e\x74\xd8\xd4\x1f\xc9\xb4\xf3\xfc\xc7\x39\x93\x3c\xb3\xe1\xc2\x86\x31\x24\x8a\x59\xcc\xc9\x93\x8b\xa0\x16\xf7\x64\x30\x80\xd2\x3a\xa9\xcf\x71\x88\xd2\x99\xac\x09\xc0\xbd\x62\x50\x1e\x01\xbf\x39\x1b\x8d\xfa\x97\x4d\xf1\xb8\xaf\x78\x81\xfd\xd8\x30\x4c\xe2\x47\xae\x24\xe8\x55\x60\x4f\xa6\xab\xe0\xf2\x22\x5e\xee\xc8\x33\x7a\x33\xfe\x2f\xb9\x61\xea\x3f\xd3\xb2\x86\x4c\x73\x59\x2b\x58\x91\x61\xf4\x7b\xd5\xe0\x34\xaf\x01\xbd\xb7\x7e\xa9\x21\xfb\x34\x0b\xf1\xd2\x82\xa4\x7c\xee\x08\xe0\x14\x9b\x6f\x4f\xa6\x4d\xce\xf9\xc5\x0d\xd8\x9d\xc1\xa8\x01\x73\x9f\x3d\x83\xb1\xca\x82\xe0\xa9\xb0\x08\x70\x55\x17\x3f\x82\x53\x0d\xfe\xa7\xce\xd8\xc4\xda\x76\x1d\xfa\x28\x64\xc1\xfb\x94\x02\xa2\xc8\xf3\xfc\x88\xbc\x86\xbf\xe3\xef\xeb\xed\x50\xdf\xe3\xe9\x9d\xf5\x99\x77\xa9\xa6\xfe\x7c\xf7\xb8\xcd\x91\x2a\x0e\x46\xd6\x95\x78\x12\xf4\x05\xab\x7e\xaa\xec\x2f\x3d\x7a\xa5\xa3\xae\x79\x5e\x5d\xab\xae\xea\xf7\xc8\xe9\x73\x0d\x79\xce\x63\xe0\xc1\xa5\x46\x9c\xe6\x4a\x14\xb4\x4f\xd8\x3f\xbd\x01\x43\x3a\xea\x58\x84\xd4\x5f\x88\xfb\xbb\xa8\x6e\x85\xde\x57\x75\x7e\x2d\x93\x5e\xe8\x90\x39\x5e\xad\x21\x76\xb3\x75\xd4\x94\x95\x28\x2e\x63\x5c\x96\xa9\xdc\x38\x18\xd7\xd1\xfe\x2e\xc4\x90\x16\xb3\xf8\xde\x16\xff\x04\x00\x00\xff\xff\xe5\x13\x3f\x32\x30\x06\x00\x00")

func template_functionsGoTemplBytes() ([]byte, error) {
	return bindataRead(
		_template_functionsGoTempl,
		"template_functions.go.templ",
	)
}

func template_functionsGoTempl() (*asset, error) {
	bytes, err := template_functionsGoTemplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template_functions.go.templ", size: 1584, mode: os.FileMode(420), modTime: time.Unix(1543602919, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typesGoTempl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x56\x4d\x6f\xe3\x36\x10\x3d\x4b\xbf\x62\x60\x04\x85\xd5\x26\x36\x7c\x5b\x18\xd8\xc3\x6e\x16\x2d\xb6\x69\x36\x45\x9c\xf6\xb2\x58\x20\x0c\x35\x92\xb9\x96\x48\x96\xa4\xd2\x1a\x86\xfe\x7b\xc1\x0f\xc9\xa6\x25\x6d\xd3\x43\x0e\xab\x93\x34\xf3\xe6\x91\x7c\xf3\x21\x4a\x42\x77\xa4\x44\x78\x5e\x91\x4a\x6e\xc9\x2a\x4d\x59\x2d\x85\x32\x30\x4f\x93\x1a\x0d\x79\x5e\xc1\x6c\xf7\x46\x2f\x98\x58\x12\xc9\x6a\x42\xb7\x8c\xa3\xda\x2f\xe5\xae\xb4\x06\xbd\xb4\xa0\xe5\xf3\x6a\x96\x26\x93\x38\xd5\x70\xc3\x6a\x9c\xa5\x59\x9a\x2e\x97\xf0\x53\x89\x9c\x56\x0c\xb9\x89\xbf\xd6\x5c\x6c\x0c\x31\x8d\x76\xe6\xdd\x1b\xbd\xce\x11\x25\x15\x72\x7f\x55\x22\x5f\x33\x6e\x50\x15\x84\xa2\x7e\xfb\x1f\x0b\x2d\xee\x9e\xbe\x22\x35\x6e\xb1\xc3\x61\xb1\x91\x48\x17\x37\x8c\xe7\x6d\x0b\x39\x16\x8c\xa3\x06\xb3\x45\x78\x22\x1a\x41\xa1\x16\x8d\xa2\x98\x9a\xbd\xc4\x73\xb4\x36\xaa\xa1\x06\x0e\x9d\x14\x8b\x87\xbd\xc4\x5b\x34\x04\x00\x1e\xbf\x6a\xc1\xd7\xb3\x4b\xc6\x2b\xc6\x71\xf6\xd8\x63\xfc\xe2\x0e\x15\x30\xd6\x91\x13\x43\x2c\xa8\x5b\xc1\x32\xb5\x2d\x00\x9c\xad\x39\xf0\x07\x8e\x53\xfb\xaf\x9b\xbb\x4f\x6d\x6b\xd9\xbc\x5e\x70\xf2\xc4\x6c\x03\x7f\x60\xd3\xce\x6e\x19\xee\x1a\x23\x1b\x33\xcd\x30\xf0\x07\x06\xe1\xec\x96\xe1\x5d\x9e\x33\xc3\x04\x27\xd5\x7d\xd0\x52\x9f\x71\x8c\x21\x02\x0d\x19\xba\x66\x8f\x69\x9b\xa6\x87\xc3\x15\xb0\x02\xe6\x25\xc2\xbc\x42\x0e\x9e\xef\xbd\xc8\xf7\x8b\x0d\xdd\x62\x4d\x16\x37\xb8\xbf\x25\x52\x32\x5e\x66\xb0\xca\xda\xd6\x85\x28\xc2\x4b\x84\x0b\xc6\x73\xfc\xe7\x12\x2e\xb0\xc2\x1a\xb9\x81\xf5\xdb\x6f\x12\xb4\xad\x2f\x95\x8b\x38\x13\x9e\xe6\xac\x6a\x4e\xcc\x5d\xed\x40\x21\xd4\x79\x74\x57\x4f\x53\x94\xc7\xd2\xea\x4e\x2a\x14\xcc\xf1\xaf\xb0\x77\x98\xd9\xb0\x59\x16\x99\x3e\xd8\x1a\xca\xec\x51\x93\xeb\x4a\x34\xf9\xcf\x42\xd5\xc4\xaa\xf7\x80\xb5\xac\x88\xc1\x4f\xa4\x46\xcb\xcc\x78\xd9\x09\x4c\x27\x81\x36\x77\xd3\x34\x5a\x12\xfa\x72\x2e\x87\xb6\x84\xf7\xa2\xaa\x9e\x08\xdd\x5d\x8b\x86\x1b\x60\xdc\x74\xb1\xea\xd4\xe1\x1b\xe1\x0a\xd0\x29\x65\x4b\xae\xcf\x9c\x54\x42\xda\x74\x75\xa9\xb3\x87\xed\x14\xe2\xc1\xef\x9a\x00\x66\x9d\xfc\xee\x30\x59\x20\xf2\x80\x8d\x53\xf7\x06\xf7\x6d\xdb\xdb\x6c\xd0\x5d\xe1\xbb\x13\xa2\xb4\x1c\x7b\xcc\x23\x5d\xdc\xd8\x1e\xfd\xbb\xad\x34\xff\xe6\x6c\x95\xc6\xc9\x74\x47\xbd\x7c\x4c\xf9\x77\x9f\xbd\x70\x74\x27\xc2\x70\xca\x86\x99\x71\xda\x35\x7e\x5c\x9c\x77\xcc\x58\xc3\x8c\x10\xf5\xc2\x8d\x56\x8a\x0f\xf0\xd0\xae\xb5\x7f\x57\x42\xa2\x32\x0c\x35\xbc\x42\x5d\xc4\xc5\x30\x3c\x7f\x98\xba\x5b\x51\xe5\xfe\xf4\x7e\xdc\x82\x28\xdc\x97\x4b\x5b\xd1\x25\x02\x4c\xc8\xc4\x98\x00\x81\xe9\x44\x80\x6e\x4c\x46\x93\x3d\x4e\xb4\x8a\x20\x6e\xbf\x71\xd4\x3d\x12\x2d\xf8\x37\xa3\x3c\xa4\x3f\x2b\x2b\x82\xcc\x7f\x68\x8c\xab\xce\xc9\xbb\x31\x84\xee\x3e\x7e\x80\xe8\x89\xe9\xb5\x87\xbc\x40\xbd\xb1\xbf\xc5\x51\xca\xe3\x0f\xa3\x2f\x26\x3d\xa6\xdc\x18\xcb\x60\xe6\xfa\x80\x11\xe8\x62\x83\xea\x99\x51\xd4\x76\xfc\x74\xef\x49\xf2\xf9\x8b\x3f\x54\xd2\x1d\x2a\x78\xe2\x59\xf1\x12\x76\xaa\xd0\x04\x72\xf7\x9a\x8c\x91\x3b\xcf\xff\xe5\xbe\x16\xbc\x60\xe5\x2d\x91\x8e\xfe\xf8\x35\xe0\xa7\xbd\x2b\x5e\xc2\xe7\xe4\x75\xee\x5f\xbf\x31\x1d\xcf\x85\xca\x1a\x88\x31\x8a\x3d\x35\xc6\x8f\x05\xff\x93\x8d\xfa\xd0\xe6\x77\x2c\xc9\x8e\x6e\xfa\x92\x36\x79\x45\xb3\x71\x53\x17\xb4\x8f\x06\xeb\xd3\x1b\xd3\xe7\x2f\x67\x9b\x09\x31\xcc\xe2\xfc\x5d\xa5\x68\x38\x05\xc6\x99\x99\x67\x76\x1b\x95\xa0\xa4\x72\x83\x08\xdf\x37\xac\xca\x51\x2d\xee\xb1\x64\xda\xa0\x9a\x93\x3c\x8f\xe9\xec\x66\x75\xd6\xb3\x8c\xfb\xe7\xda\xb1\xc1\x8f\x9d\xc0\x9e\x3d\x03\x54\xca\x0e\xd2\x34\xf1\x00\x5b\x10\x37\x5c\xfc\xcd\x7d\x94\x47\xfd\xa2\x44\x23\xff\x44\xa5\x99\xe0\x97\x69\x92\xfc\x70\x76\xe1\x6c\x47\x8c\x56\x20\xe7\xc8\x7a\xcd\xde\xe5\xf9\x83\x38\xe5\x0a\x9b\xba\x84\xe1\x32\x59\x9a\x28\x34\x8d\xe2\xc0\x59\x95\xb6\xe9\xbf\x01\x00\x00\xff\xff\x10\x11\xf5\x38\x65\x0c\x00\x00")

func typesGoTemplBytes() ([]byte, error) {
	return bindataRead(
		_typesGoTempl,
		"types.go.templ",
	)
}

func typesGoTempl() (*asset, error) {
	bytes, err := typesGoTemplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "types.go.templ", size: 3173, mode: os.FileMode(420), modTime: time.Unix(1543602919, 0)}
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
	"aws-service-operator.yaml.templ": awsServiceOperatorYamlTempl,
	"base.go.templ":                   baseGoTempl,
	"cft.go.templ":                    cftGoTempl,
	"operator.go.templ":               operatorGoTempl,
	"template_functions.go.templ":     template_functionsGoTempl,
	"types.go.templ":                  typesGoTempl,
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
	"aws-service-operator.yaml.templ": &bintree{awsServiceOperatorYamlTempl, map[string]*bintree{}},
	"base.go.templ":                   &bintree{baseGoTempl, map[string]*bintree{}},
	"cft.go.templ":                    &bintree{cftGoTempl, map[string]*bintree{}},
	"operator.go.templ":               &bintree{operatorGoTempl, map[string]*bintree{}},
	"template_functions.go.templ":     &bintree{template_functionsGoTempl, map[string]*bintree{}},
	"types.go.templ":                  &bintree{typesGoTempl, map[string]*bintree{}},
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
