package main

//func main() {
//	http.HandleFunc("/test", handle)
//	http.ListenAndServe(":8080", nil)
//}
//
//func handle(res http.ResponseWriter, req *http.Request) {
//	switch req.Method {
//	case "GET":
//		res.Write([]byte("我收到了你的get请求"))
//		break
//	case "POST":
//		b, _ := io.ReadAll(req.Body)
//		header := res.Header()
//		header["test"] = []string{"test1", "test2"}
//		res.WriteHeader(http.StatusOK)
//		res.Write(b)
//		break
//	default:
//		res.Write([]byte("我没有收到"))
//	}
//
//}
