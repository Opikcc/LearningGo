module gitlab.com/opikcc

go 1.16

replace gitlab.com/opikcc/calc => ../calc

replace gitlab.com/opikcc/strcon => ../strcon

replace github.com/gorilla/mux => ../mux-master

replace github.com/gorilla/handlers => ../handlers-master

replace github.com/felixge/httpsnoop => ../httpsnoop-master

require (
	github.com/gorilla/handlers v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v0.0.0-00010101000000-000000000000
	gitlab.com/opikcc/calc v0.0.0-00010101000000-000000000000
	gitlab.com/opikcc/strcon v0.0.0-00010101000000-000000000000
)
