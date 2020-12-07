#ifdef __cplusplus
extern "C" {
#endif
    void* pocket_NewPocketDispatcher();

    int pocket_Initialize(void* PocketDispatcher, string address, int pocket_port);

    int pocket_MakeDir(void* PocketDispatcher,string name);
    int pocket_Lookup(void* PocketDispatcher,string name);
    int pocket_Enumerate(void* PocketDispatcher,string name);
    int pocket_PutFile(void* PocketDispatcher,string local_file, string dst_file, bool enumerable);
    int pocket_GetFile(void* PocketDispatcher,string src_file, string local_file);
    int pocket_PutBuffer(void* PocketDispatcher,const char buf[], int pocket_len, string dst_file, bool enumerable);
    int pocket_GetBuffer(void* PocketDispatcher,char buf[], int pocket_len, string src_file);
    int pocket_DeleteFile(void* PocketDispatcher,string file);
    int pocket_DeleteDir(void* PocketDispatcher,string directory);
    int pocket_CountFiles(void* PocketDispatcher,string directory);

    void pocket_DestroyPocketDispatcher(void* PocketDispatcher);

#ifdef __cplusplus
}
#endif