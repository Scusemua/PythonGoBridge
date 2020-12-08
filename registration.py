import time
import sys
import os
import socket
import struct
import errno
import libpocket
from subprocess import call, Popen

def register_job(jobname, num_lambdas=0, capacityGB=0, peakMbps=0, latency_sensitive=1):
  # connect to controller
  sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
  sock.connect((CONTROLLER_IP, CONTROLLER_PORT))

  # send register request to controller
  msg_packer = struct.Struct(REQ_STRUCT_FORMAT + "i" + str(len(jobname)) + "s" + "iiih") 
  msgLen = REQ_LEN_HDR + INT + len(jobname) + 3*INT + SHORT
  sampleMsg = (msgLen, TICKET, RPC_JOB_CMD, JOB_CMD, REGISTER_OPCODE, len(jobname), jobname, \
                 num_lambdas, int(capacityGB), int(peakMbps), latency_sensitive)
  pkt = msg_packer.pack(*sampleMsg)
  sock.sendall(pkt)

  # get jobid response
  data = sock.recv(RESP_LEN_BYTES + INT)
  resp_packer = struct.Struct(RESP_STRUCT_FORMAT + "i")
  [length, ticket, type_, err, opcode, jobIdNum] = resp_packer.unpack(data)
  if err != 0:
    jobid = None
    print("Error registering job: ", err)
  else:
    jobid = jobname + "-" + str(jobIdNum)
    print("Registered jobid ", jobid)
  sock.close()
  return jobid
 

def deregister_job(jobid):
  # connect to controller
  sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
  sock.connect((CONTROLLER_IP, CONTROLLER_PORT))
  
  # send register request to controller
  msg_packer = struct.Struct(REQ_STRUCT_FORMAT + "i" + str(len(jobid)) + "s") # len(jobname) (INT) + jobname (STRING)
  msgLen = REQ_LEN_HDR + INT + len(jobid)
  sampleMsg = (msgLen, TICKET, RPC_JOB_CMD, JOB_CMD, DEREGISTER_OPCODE, len(jobid), jobid)
  pkt = msg_packer.pack(*sampleMsg)
  sock.sendall(pkt)

  # get jobid response
  data = sock.recv(RESP_LEN_BYTES)
  resp_packer = struct.Struct(RESP_STRUCT_FORMAT)
  [length, ticket, type_, err, opcode] = resp_packer.unpack(data)
  if err != 0:
    print("Error deregistering job: ", err)
  else:
    print("Successfully deregistered jobid ", jobid)
  sock.close() 
  return err