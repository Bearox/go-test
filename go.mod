module github.com/Bearox/go-test

go 1.16

require (
	github.com/Bearox/go-test/chan-test v0.0.0-00010101000000-000000000000
	github.com/Bearox/go-test/greetings v0.0.0-20210403143038-aeee0f13f128
	github.com/Bearox/go-test/util-test v0.0.0-00010101000000-000000000000
)

replace (
	github.com/Bearox/go-test/chan-test => ./chan-test
	github.com/Bearox/go-test/greetings => ./greetings
	github.com/Bearox/go-test/util-test => ./util-test
)
