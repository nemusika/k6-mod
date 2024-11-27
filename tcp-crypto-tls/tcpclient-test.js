import { TcpClient } from "k6/x/tcpclient";

export default function () {
  let client = new TcpClient();
  let conn = client.connect(true, "example.com:443", true);
  conn.write("Hello\n");
  let response = conn.readLine();
  console.log(response);
}
