module github.com/kendfss/joe

go 1.17

require (
	github.com/kendfss/but v0.0.0-00010101000000-000000000000
	github.com/kendfss/namespacer v0.0.0-00010101000000-000000000000
	github.com/termie/go-shutil v0.0.0-20140729215957-bcacb06fecae
	github.com/urfave/cli v1.22.5
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
)

replace (
	github.com/kendfss/but => ../but
	github.com/kendfss/namespacer => ../namespacer
)
