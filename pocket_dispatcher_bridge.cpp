// C++ implementation of C header file.

#include <iostream>
#include <string>

#include "pocket_dispatcher_bridge.h"
#include "pocket_dispatcher.h"

PocketDispatcher* AsPocketDispatcher(void* pocketDispatcher) {
    return reinterpret_cast<PocketDispatcher*>(pocketDispatcher);
}

void* pocket_NewPocketDispatcher() {
    auto dispatcher = new PocketDispatcher();

    return dispatcher;
}

void pocket_HelloWorld() {
    std::cout << "Hello, world!" << std::endl;
}

int pocket_Initialize(void* pocketDispatcher, char* address, int pocket_port) {
    std::cout << "[C++ Bridge] pocket_Initialize" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->Initialize(string(address), pocket_port);
}

int pocket_MakeDir(void* pocketDispatcher, char* name) {
    std::cout << "[C++ Bridge] pocket_MakeDir" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->MakeDir(string(name));
}

int pocket_Lookup(void* pocketDispatcher, char* name) {
    std::cout << "[C++ Bridge] pocket_Lookup" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->Lookup(string(name));
}

int pocket_Enumerate(void* pocketDispatcher, char* name) {
    std::cout << "[C++ Bridge] pocket_Enumerate" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->Enumerate(string(name));
}

int pocket_PutFile(void* pocketDispatcher, char* local_file, char* dst_file, int enumerable) {
    std::cout << "[C++ Bridge] pocket_PutFile" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->PutFile(string(local_file), string(dst_file), enumerable);
}

int pocket_GetFile(void* pocketDispatcher, char* src_file, char* local_file) {
    std::cout << "[C++ Bridge] pocket_GetFile" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->GetFile(string(src_file), string(local_file));
}

int pocket_PutBuffer(void* pocketDispatcher, const char buf[], int pocket_len, char* dst_file, int enumerable) {
    std::cout << "[C++ Bridge] pocket_PutBuffer" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->PutBuffer(buf, pocket_len, string(dst_file), enumerable);
}

int pocket_GetBuffer(void* pocketDispatcher, char buf[], int pocket_len, char* src_file) {
    std::cout << "[C++ Bridge] pocket_GetBuffer" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->GetBuffer(buf, pocket_len, string(src_file));
}

int pocket_DeleteFile(void* pocketDispatcher, char* file) {
    std::cout << "[C++ Bridge] pocket_DeleteFile" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->DeleteFile(string(file));
}

int pocket_DeleteDir(void* pocketDispatcher, char* directory) {
    std::cout << "[C++ Bridge] pocket_DeleteDir" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->DeleteDir(string(directory));
}

int pocket_CountFiles(void* pocketDispatcher, char* directory) {
    std::cout << "[C++ Bridge] pocket_CountFiles" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->CountFiles(string(directory));
}

void pocket_DestroyPocketDispatcher(void* pocketDispatcher) {
    std::cout << "[C++ Bridge] pocket_DestroyPocketDispatcher" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->~PocketDispatcher();
}