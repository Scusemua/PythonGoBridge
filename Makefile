pocketlibrary.so: pocket_dispatcher.cc pocket_dispatcher_bridge.cpp
	clang++ -o pocketlibrary.so pocket_dispatcher.cc pocket_dispatcher_bridge.cpp \
	-std=c++14 -O3 -Wall -Wextra -fPIC -shared -L/libpocket.so
