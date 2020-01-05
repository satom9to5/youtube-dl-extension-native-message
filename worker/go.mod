module worker

go 1.13

require (
	github.com/mattn/go-sqlite3 v2.0.2+incompatible
	github.com/satom9to5/youtube-dl-queue v0.0.0-20191212161905-73fb72f540dd
)

replace github.com/satom9to5/youtube-dl-queue v0.0.0-20191212161905-73fb72f540dd => ../../youtube-dl-queue
