# netcat messenger

in this project I will make a simple chat/messenger using netcat and golang over the command line

i will consider this project done when it meets the following expectations;

    a message can be sent and received over TCP to a virtual machine with netcat
        device a "Hello World!" => device b "Hello World!"

    the server will send new messages to every client connected

## planning

### server

#### listen for messages from connected clients

    net package
        packet, err := net.Listen(network, addr)

   conn, err := packet.Accept()

   n, err := conn.Read(buffer)
        needs a buffer (slice of type byte) to write to
        buff := make([]byte, 256)
        n is the number of bytes read?

    *client would convert bytes to string*

#### pass message to all other connected clients

    track connected clients
        FIFO arr
        sender := conn.RemoteAddr()

            let q = [];
            // Adds elements {0, 1, 2, 3, 4} to queue
            for (let i = 0; i < 5; i++)
                q.push(i);
            
            // Display contents of the queue.
            document.write("Elements of queue-[" + q.join(", ")+"]<br>");
            
            // To remove the head of queue.
            // In this the oldest element '0' will be removed
            let removedele = q.shift();
            document.write("removed element-" + removedele+"<br>");
            
            document.write("["+q.join(", ")+"]<br>");
            
            // To view the head of queue
            let head = q[0];
            document.write("head of queue-" + head+"<br>");
            
            // Rest all methods of collection interface,
            // Like size and contains can be used with this
            // implementation.
            let size = q.length;
            document.write("Size of queue-" + size+"<br>");
            
            
            // This code is contributed by avanitrachhadiya2155`

    for each connected client
        for addr, _ := range FIFOarr
            if addr != sender
                func sendMessage
                return
    
    func sendMessage(messageBuf []byte)
        conn, err := net.Dial(network, addr)

        n, err := conn.Write(messageBuf)

        conn.close()

### client

#### send message to server at port

#### recieve messages to display