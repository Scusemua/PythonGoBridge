#include <iostream>

#include "pocket_dispatcher_bridge.h"
#include "pocket_dispatcher.h"

PocketDispatcher* return AsPocketDispatcher(void* pocketDispatcher) {
    reinterpret_cast<PocketDispatcher*>(pocketDispatcher);
}

void* pocket_NewPocketDispatcher() {
    auto dispatcher = new PocketDispatcher();

    return dispatcher;
}

int pocket_Initialize(void* pocketDispatcher, string address, int pocket_port) {
    std::cout << "[C++ Bridge] pocket_Initialize" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->Initialize(address, pocket_port);
}

int pocket_MakeDir(void* pocketDispatcher, string name) {
    std::cout << "[C++ Bridge] pocket_MakeDir" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->MakeDir(name);
}

int pocket_Lookup(void* pocketDispatcher, string name) {
    std::cout << "[C++ Bridge] pocket_Lookup" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->Lookup(name);
}

int pocket_Enumerate(void* pocketDispatcher, string name) {
    std::cout << "[C++ Bridge] pocket_Enumerate" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->Enumerate(name);
}

int pocket_PutFile(void* pocketDispatcher, string local_file, string dst_file, bool enumerable) {
    std::cout << "[C++ Bridge] pocket_PutFile" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->PutFile(local_file, dst_file, enumerable);
}

int pocket_GetFile(void* pocketDispatcher, string src_file, string local_file) {
    std::cout << "[C++ Bridge] pocket_GetFile" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->GetFile(src_file, local_file);
}

int pocket_PutBuffer(void* pocketDispatcher, const char buf[], int pocket_len, string dst_file, bool enumerable) {
    std::cout << "[C++ Bridge] pocket_PutBuffer" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->PutBuffer(buf, pocket_len, dst_file, enumerable);
}

int pocket_GetBuffer(void* pocketDispatcher, char buf[], int pocket_len, string src_file) {
    std::cout << "[C++ Bridge] pocket_GetBuffer" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->GetBuffer(buf, pocket_len, src_file);
}

int pocket_DeleteFile(void* pocketDispatcher, string file) {
    std::cout << "[C++ Bridge] pocket_DeleteFile" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->DeleteFile(file);
}

int pocket_DeleteDir(void* pocketDispatcher, string directory) {
    std::cout << "[C++ Bridge] pocket_DeleteDir" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->DeleteDir(directory);
}

int pocket_CountFiles(void* pocketDispatcher, string directory) {
    std::cout << "[C++ Bridge] pocket_CountFiles" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->CountFiles(directory);
}

void pocket_DestroyPocketDispatcher(void* pocketDispatcher) {
    std::cout << "[C++ Bridge] pocket_DestroyPocketDispatcher" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->~PocketDispatcher();
}