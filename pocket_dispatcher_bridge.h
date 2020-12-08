// C header file for C++ implenetation of bridge.

#ifdef __cplusplus
extern "C" {
#endif
    void* pocket_NewPocketDispatcher();

    int pocket_Initialize(void* PocketDispatcher,  char* address,  int pocket_port);

    int pocket_MakeDir(void* PocketDispatcher, char* name);
    int pocket_Lookup(void* PocketDispatcher, char* name);
    int pocket_Enumerate(void* PocketDispatcher, char* name);
    int pocket_PutFile(void* PocketDispatcher, char* local_file, char* dst_file, int enumerable);
    int pocket_GetFile(void* PocketDispatcher, char* src_file, char* local_file);
    int pocket_PutBuffer(void* PocketDispatcher, const char buf[], int pocket_len, char* dst_file, int enumerable);
    int pocket_GetBuffer(void* PocketDispatcher, char buf[], int pocket_len, char* src_file);
    int pocket_DeleteFile(void* PocketDispatcher, char* file);
    int pocket_DeleteDir(void* PocketDispatcher, char* directory);
    int pocket_CountFiles(void* PocketDispatcher, char* directory);

    void pocket_DestroyPocketDispatcher(void* PocketDispatcher);

#ifdef __cplusplus
}
#endif