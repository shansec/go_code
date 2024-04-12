package main

//func main() {
//	tcpErrpr, _ := net.ResolveTCPAddr("tcp", ":8888")
//	lister, _ := net.ListenTCP("tcp", tcpErrpr)
//	for {
//		conn, err := lister.AcceptTCP()
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//
//		go handleConnection(conn)
//	}
//}
//
//func handleConnection(conn *net.TCPConn) {
//	for {
//		bufer := make([]byte, 1024)
//		n, err := conn.Read(bufer)
//		if err != nil {
//			fmt.Println(err)
//			break
//		}
//		fmt.Println(conn.RemoteAddr().String() + "发送：" + string(bufer[0:n]))
//	}
//
//}
