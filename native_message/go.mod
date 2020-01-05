module native_message

go 1.13

require (
	github.com/mattn/go-sqlite3 v2.0.2+incompatible
	github.com/mitchellh/mapstructure v1.1.2
	github.com/satom9to5/pidfile v0.0.0-20190604150648-b4983bc136e3
	github.com/satom9to5/webext v0.0.0-20190226145152-1cd0537acbbc
	github.com/satom9to5/youtube-dl-queue v0.0.0-20191212161905-73fb72f540dd
)

replace (
	github.com/satom9to5/pidfile v0.0.0-20190604150648-b4983bc136e3 => ../../pidfile
	github.com/satom9to5/webext v0.0.0-20190226145152-1cd0537acbbc => ../../webext
	github.com/satom9to5/youtube-dl-queue v0.0.0-20191212161905-73fb72f540dd => ../../youtube-dl-queue
)
