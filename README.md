# netcat messenger

in this project I will make a simple chat/messenger using netcat and golang over the command line

i will consider this project done when it meets the following expectations;

    a message can be sent and received over TCPto a virtual machine with netcat
        device a "Hello World!" => device b "Hello World!"

    a server will store these messages and the time they were delivered

    a user will connect to the server which will send the last 10 messages and the time they were sent

    the server will send new messages to every client connected

# server

listen on port

on connection send last 10 messages

when recieving new message -> push to connected clients