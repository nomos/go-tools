package tools

import (
	"errors"
	"github.com/nomos/go-lokas/util"
	"os"
	"path"
	"regexp"
	"strings"
)

func FileNameReplace(dir string,pattern interface{},replacer string)(map[string]string,error){
	ret:=make(map[string]string)
	patternRegexp,isRegexp:=pattern.(*regexp.Regexp)
	patternString:=""
	if !isRegexp {
		ok:=false
		patternString,ok=pattern.(string)
		if !ok {
			return ret,errors.New("param error")
		}
	}
	_,err:=util.WalkDirFilesWithFunc(dir, func(filePath string, file os.FileInfo) bool {
		fileName:=path.Base(filePath)
		fileDir:=path.Dir(filePath)
		exportStr:=""
		if isRegexp {
			exportStr = patternRegexp.ReplaceAllString(fileName,replacer)
		} else {
			exportStr = strings.Replace(fileName,patternString,replacer,-1)
		}
		exportStr = strings.ToLower(exportStr)
		ret[filePath] = path.Join(fileDir,exportStr)
		return false
	},true)
	if err != nil {
		return ret,err
	}
	return ret,nil
}

