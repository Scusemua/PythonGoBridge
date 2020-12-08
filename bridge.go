package bridge 

// #cgo CFLAGS: -Wall -std=c++11
// #cgo LDFLAGS: -L./libcppcrail libpocket.a -lstdc++11
// #include "pocket_dispatcher_bridge.h"
// #include <stdlib.h>
import "C"
import "fmt"
import "os/exec"
import "unsafe"

var PORT = 2345
var HOSTNAME = "localhost"

var CONTROLLER_IP = "10.1.47.178"
var CONTROLLER_PORT = 4321

var MAX_DIR_DEPTH = 16

const INT = 4
const LONG = 8
const FLOAT = 4
const SHORT = 2
const BYTE = 1

// msg_len (INT), ticket (LONG LONG), cmd (SHORT), cmd_type (SHORT), register_type (BYTE)
const REQ_STRUCT_FORMAT = "!iqhhi" 
// CMD, CMD_TYPE, IOCTL_OPCODE (note: doesn't include msg_len or ticket from NaRPC hdr)
const REQ_LEN_HDR = SHORT + SHORT + INT

// msg_len (INT), ticket (LONG LONG), cmd (SHORT), error (SHORT), register_opcode (BYTE)
const RESP_STRUCT_FORMAT = "!iqhhi"
// MSG_LEN, TICKET, CMD, ERROR, REGISTER_OPCODE 
const RESP_LEN_BYTES = INT + LONG + SHORT + SHORT + INT

const TICKET = 1000
const RPC_JOB_CMD = 14
const JOB_CMD = 14
const REGISTER_OPCODE = 0
const DEREGISTER_OPCODE = 1

type PocketDispatcher struct {
	ptr unsafe.Pointer 
}

func HelloWorld() {
	C.pocket_HelloWorld();
}

func NewPocketDispatcher() PocketDispatcher {
	var pocketDispatcher PocketDispatcher
	pocketDispatcher.ptr = C.pocket_NewPocketDispatcher()
	return pocketDispatcher
}

func (pocketDispatcher PocketDispatcher) Free() {
	C.pocket_DestroyPocketDispatcher(pocketDispatcher.ptr)
}

// Call Pocket's python function directly to register the job.
func RegisterJob(jobName string, numLambdas int, capacityGB int, peakMbps int, latencySensitive int) int {
	// Call Pocket's register_job function.
	// Return the job ID.
	//tcpAddr, err := net.ResolveTCPAddr(fmt.Sprintf("%s:%d", CONTROLLER_IP, CONTROLLER_PORT))
	//conn, err := net.DialTCP("tcp", nil, tcpAddr)

	// def register_job(jobname, num_lambdas=0, capacityGB=0, peakMbps=0, latency_sensitive=1)
	python_command := fmt.Sprintf("import registration; print registration.register_job(%s, num_lambdas=%d, capacityGB=%d, peakMbps=%d, latency_sensitive=%d",
									jobName, numLambdas, capacityGB, peakMbps, latency_sensitive)
	cmd := exec.Command("python3", "-c", python_command)
	fmt.Println(cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err);
	}

	fmt.Println(string(out))
}

// Call Pocket's python function directly to deregister the job.
func DeregisterJob(jobId int) int {
	// Call Pocket's deregister_job.
	// Return err.
	
	// def deregister_job(jobid)
	python_command := fmt.Sprintf("import registration; print registration.deregister_job(%s)", jobName)
	cmd := exec.Command("python3", "-c", python_command)
	fmt.Println(cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err);
	}

	fmt.Println(string(out))
}

func Connect(hostname string, port int) PocketDispatcher {
    fname := C.CString(hostname)
	defer C.free(unsafe.Pointer(fname))	
	
	pocketDispatcher := NewPocketDispatcher()
	_res := C.pocket_Initialize(pocketDispatcher.ptr, fname, C.int(port))

	res := int(res)

	if res != 0 {
		fmt.Printf("[ERROR] Connecting to metadata server failed!")
	}
	
	return pocketDispatcher
}

// Send a CREATE DIRECTORY request to Pocket
// 
// :param pocket:           pocketHandle returned from connect()
// :param str src_filename: name of directory to create in Pocket 
// :param str jobid:        id unique to this job, used to separate keyspace for job
// :return: the Pocket dispatcher response 
func (pocket PocketDispatcher) CreateDir(src_filename string, jobid string) int {
	if jobid != "" {
		jobid = "/" + jobid
	}

	if src_filename != "" {
		src_filename = jobid + "/" + src_filename
	} else {
		src_filename = jobid
	}	

	return pocket.MakeDir(src_filename)
}

func (pocket PocketDispatcher) MakeDir(src_filename string) int {
	source_filename := C.CString(src_filename)
	defer C.free(unsafe.Pointer(source_filename))

	_res := C.pocket_MakeDir(pocket.ptr, source_filename)

	return int(_res)
}

// Send a LOOKUP metadata request to Pocket to see if file exists
//
// :param pocket:           	pocketHandle returned from connect()
// :param str src_filename: 	name of file/key in Pocket from which looking up
// :param str jobid:        	id unique to this job, used to separate keyspace for job
// :return: 					the Pocket dispatcher response 
func (pocket PocketDispatcher) Lookup(src_filename string, jobid string) int {
	if jobid != "" {
		jobid = "/" + jobid
	}

	get_filename := jobid + "/" + src_filename
	get_file := C.CString(get_filename)
	defer C.free(unsafe.Pointer(get_file))

	_res = C.pocket_Lookup(pocket.ptr, get_file)
	res = int(_res)

	if res != 0 {
		fmt.Printf("[ERROR] Lookup failed!\n")
	}

	return res 
}

func (pocket PocketDispatcher) Enumerate(name string) int {
	fmt.Printf("[ERROR] Enumerate has not been implemented.\n")
	return -1 
}

// Send a PUT request to Pocket to write key
// :param pocket:           	pocketHandle returned from connect()
// :param str src_filename: 	name of local file containing data to PUT
// :param str dst_filename: 	name of file/key in Pocket which writing to
// :param str jobid:        	id unique to this job, used to separate keyspace for job
// :param PERSIST_AFTER_JOB:	optional hint, if True, data written to table persisted after job done
// :return: the Pocket dispatcher response 
func (pocket PocketDispatcher) Put(src_filename string, dst_filename string, jobid string, persistAfterJob bool) int {
	if jobid != "" {
		jobid = "/" + jobid
	}

	var set_filename string 
	if persistAfterJob {
		set_filename = jobid + "-persist/" + dst_filename
	} else {
		set_filename = jobid + "/" + dst_filename
	}

	dst_file := C.CString(set_filename)
	defer C.free(unsafe.Pointer(set_filename))
	
	local_file := C.CString(src_filename)
	defer C.free(unsafe.Pointer(local_file))

	// int pocket_PutFile(void* pocketDispatcher, char* local_file, char* dst_file, bool enumerable)
	_res = C.pocket_PutFile(pocket.ptr, local_file, dst_file, false)
	return int(_res)
}

// Send a GET request to Pocket to read key
//
// :param pocket:           	pocketHandle returned from connect()
// :param str src_filename: 	name of file/key in Pocket from which reading
// :param str dst_filename: 	name of local file where want to store data from GET
// :param str jobid:        	id unique to this job, used to separate keyspace for job
// :param DELETE_AFTER_READ:	optional hint, if True, data deleted after job done
// :return: the Pocket dispatcher response 
func (pocket PocketDispatcher) Get(src_filename string, dst_filename string, jobid string, deleteAfterRead bool) int {
	if jobid != "" {
		jobid = "/" + jobid
	}

	get_filename := jobid + "/" + src_filename
	get_file := C.CString(get_filename)
	defer C.free(unsafe.Pointer(get_file))

	dst_file := C.CString(dst_filename)
	defer C.free(unsafe.Pointer(dst_file))

	_res = C.pocket_GetFile(pocket.ptr, get_file, dst_file)
	res = int(_res)

	if res != 0 {
		fmt.Printf("[ERROR] GET failed!")
		return res 
	}

	if deleteAfterRead {
		_res = pocket.Delete(src_filename, jobid)
		res = int(_res)
	}

	return res 
}

// Send a DEL request to Pocket to delete key
//
// :param pocket:           pocketHandle returned from connect()
// :param str src_filename: name of file/key in Pocket which deleting
// :param str jobid:        id unique to this job, used to separate keyspace for job
// :return: the Pocket dispatcher response 
func (pocket PocketDispatcher) Delete(src_filename string, jobid string) {
	if jobid != "" {
		jobid = "/" + jobid
	}

	if src_filename != "" {
		src_filename = jobid + "/" + src_filename
	} else {
		src_filename = jobid
	}

	_res = pocket.DeleteDir(src_filename) // Recursive delete.
	return int(res)
}

func (pocket PocketDispatcher) DeleteDir(file string) int {
	filename := C.CString(file)
	defer C.free(unsafe.Pointer(filename))

	_res = C.pocket_DeleteDir(pocket.ptr, filename)
	return int(_res)
}

func (pocket PocketDispatcher) CountFiles(directory string) int {
	directoryString := C.CString(directory)
	defer C.free(unsafe.Pointer(directory))

	// int pocket_CountFiles(void* pocketDispatcher, char* directory)
	_res = C.pocket_CountFiles(pocket.ptr, directoryString)
	return int(_res)
}

// Send a PUT request to Pocket to write key
//
// :param pocket:           	pocketHandle returned from connect()
// :param str src: 	   			name of local object containing data to PUT
// :param str dst_filename: 	name of file/key in Pocket which writing to
// :param str jobid:        	id unique to this job, used to separate keyspace for job
// :param PERSIST_AFTER_JOB:	optional hint, if True, data written to table persisted after job done
// :return: the Pocket dispatcher response 
func (pocket PocketDispatcher) PutBuffer(data []byte, len int, dst_filename string, jobid int, persistAfterJob bool) int {
	if jobid != "" {
		jobid = "/" + jobid
	}

	var set_filename string 
	if persistAfterJob {
		set_filename = jobid + "-persist/" + dst_filename
	} else {
		set_filename = jobid + "/" + dst_filename
	}

	dst_file := C.CString(set_filename)
	defer C.free(unsafe.Pointer(set_filename))

	// int pocket_PutBuffer(void* pocketDispatcher, const char buf[], int pocket_len, char* dst_file, bool enumerable)
	_res = C.pocket_PutBuffer(pocket.ptr, (*C.char)(unsafe.Pointer(&dst[0])), C.int(len), set_filename, false)
	return _res 
}

// Send a GET request to Pocket to read key
//
// :param pocket:           	pocketHandle returned from connect()
// :param str src_filename: 	name of file/key in Pocket from which reading
// :param str dst: 				name of local object  where want to store data from GET
// :param str jobid:        	id unique to this job, used to separate keyspace for job
// :param delete_after_read:	optional hint, if True, data deleted after job done
// :return: the Pocket dispatcher response 
func (pocket PocketDispatcher) GetBuffer(src_filename string, dst []byte, len int, jobid string, delete_after_read bool) int {
	if jobid != "" {
		jobid = "/" + jobid
	}

	get_filename := jobid + "/" + src_filename
	c_get_filename := C.CString(get_filename)
	defer C.free(unsafe.Pointer(c_get_filename))

	// int pocket_GetBuffer(void* PocketDispatcher, char buf[], int pocket_len, char* src_file);
	_res := pocket_GetBuffer(pocket.ptr, (*C.char)(unsafe.Pointer(&dst[0])), C.int(pocket_len), c_get_filename)
	res := int(_res)

	if res != 0 {
		fmt.Printf("[ERROR] GET BUFFER failed!\n")
		return res 
	}

	if delete_after_read {
		_res = pocket.Delete(src_filename, jobid)
		res = int(_res)
	}

	return res 
}
