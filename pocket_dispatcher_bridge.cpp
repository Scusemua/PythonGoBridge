#include <iostream>
#include <string>

#include "pocket_dispatcher_bridge.h"
#include "pocket_dispatcher.h"

PocketDispatcher* return AsPocketDispatcher(void* pocketDispatcher) {
    reinterpret_cast<PocketDispatcher*>(pocketDispatcher);
}

void* pocket_NewPocketDispatcher() {
    auto dispatcher = new PocketDispatcher();

    return dispatcher;
}

int pocket_Initialize(void* pocketDispatcher, char* address, int pocket_port) {
    std::cout << "[C++ Bridge] pocket_Initialize" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->Initialize(str(address), pocket_port);
}

int pocket_MakeDir(void* pocketDispatcher, char* name) {
    std::cout << "[C++ Bridge] pocket_MakeDir" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->MakeDir(str(name));
}

int pocket_Lookup(void* pocketDispatcher, char* name) {
    std::cout << "[C++ Bridge] pocket_Lookup" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->Lookup(str(name));
}

int pocket_Enumerate(void* pocketDispatcher, char* name) {
    std::cout << "[C++ Bridge] pocket_Enumerate" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->Enumerate(str(name));
}

int pocket_PutFile(void* pocketDispatcher, char* local_file, char* dst_file, bool enumerable) {
    std::cout << "[C++ Bridge] pocket_PutFile" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->PutFile(str(local_file), str(dst_file), enumerable);
}

int pocket_GetFile(void* pocketDispatcher, char* src_file, char* local_file) {
    std::cout << "[C++ Bridge] pocket_GetFile" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->GetFile(str(src_file), str(local_file));
}

int pocket_PutBuffer(void* pocketDispatcher, const char buf[], int pocket_len, char* dst_file, bool enumerable) {
    std::cout << "[C++ Bridge] pocket_PutBuffer" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->PutBuffer(buf, pocket_len, str(dst_file), enumerable);
}

int pocket_GetBuffer(void* pocketDispatcher, char buf[], int pocket_len, char* src_file) {
    std::cout << "[C++ Bridge] pocket_GetBuffer" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->GetBuffer(buf, pocket_len, str(src_file));
}

int pocket_DeleteFile(void* pocketDispatcher, char* file) {
    std::cout << "[C++ Bridge] pocket_DeleteFile" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->DeleteFile(str(file));
}

int pocket_DeleteDir(void* pocketDispatcher, char* directory) {
    std::cout << "[C++ Bridge] pocket_DeleteDir" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->DeleteDir(str(directory));
}

int pocket_CountFiles(void* pocketDispatcher, char* directory) {
    std::cout << "[C++ Bridge] pocket_CountFiles" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->CountFiles(str(directory));
}

void pocket_DestroyPocketDispatcher(void* pocketDispatcher) {
    std::cout << "[C++ Bridge] pocket_DestroyPocketDispatcher" << std::endl;
    return AsPocketDispatcher(pocketDispatcher)->~PocketDispatcher();
}