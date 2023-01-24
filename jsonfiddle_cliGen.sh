prj=jsonfiddle
pwd=`pwd`
srcDir=$GOPATH/src/github.com/go-easygen/easygen/test
cd $srcDir || srcDir=/usr/share/gocode/src/github.com/go-easygen/easygen/test
cd $srcDir || srcDir=/usr/share/doc/easygen/examples
[ -d "$srcDir" ] || {
  echo No template file found
  exit 1
}

set -x

cd $srcDir
easygen commandlineGoFlags.header,commandlineGoFlags.ityped.tmpl,commandlineGoFlags "$pwd/$prj"_cli | gofmt > "$pwd/$prj"_cliDef.go

cd $pwd
go build -v
echo ${prj}_cliDef.go generated successful
