pocketlibrary.a: pocket_dispatcher.cc pocket_dispatcher_bridge.cpp
	clang++ -o pocketlibrary.a pocket_dispatcher.cc pocket_dispatcher_bridge.cpp libcppcrail.a \
	-std=c++14 -O3 -Wall -Wextra -fPIC -static -nostartfiles
