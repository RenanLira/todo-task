package utils

import "reflect"

func GetPackageName(t reflect.Type) string {
	if t.Kind().String() == "ptr" {
		t = t.Elem()
	}

	pkgPath := t.PkgPath()
	var pkgName string

	for i := len(pkgPath) - 1; i >= 0; i-- {
		if pkgPath[i] == '/' {
			pkgName = pkgPath[i+1:]
			break
		}
	}

	return pkgName
}
