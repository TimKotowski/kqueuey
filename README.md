# Kqueuey

# highly distributed, simple message queue.

 KQueuey is a simple distributed message queue built in Go, with data replication using Raft Consensus Algorithm.
 Ability to queue messages and poll messages between software components with deadlines to ensure a message is finished or re-processed.

### High Level Overview
   1. Client distributes messages to leader node.
   2. Leader node recieves messages, and sends batch messages back to what ever client is polling messages.
   4. Clients process messages, and sends tombstones back to prevent any messages being resent, or if deadline is reached, message is re-received to any client polling messages.


![kqueuey](https://github.com/user-attachments/assets/a9a0cca6-f4f7-4431-882b-a686bae0d34c)


README Work In Progress....
