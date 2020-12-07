package bridge 

import (
	"fmt"
)

var PORT := 2345
var HOSTNAME := "localhost"

var CONTROLLER_IP := "10.1.47.178"
var CONTROLLER_PORT := 4321

var MAX_DIR_DEPTH := 16

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

func RegisterJob(jobName string, numLambdas int, capacityGB int, peakMbps int, latencySensitive int) int {
	// Call Pocket's register_job function.
	// Return the job ID.
}

func DeregisterJob(jobId int) int {
	// Call Pocket's deregister_job.
	// Return err.
}

func Connect(hostname string, port int) 