module example

go 1.18

require (
	gee v0.0.0
)
//使用 replace 将 gee 指向 ./gee
replace (
	gee => ./gee
)