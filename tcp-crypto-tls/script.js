import tcp from "k6/x/tcp-crypto-tls"; // デフォルトエクスポート
// import { TCP } from 'k6/x/tcp-crypto-tls'; // 名前付きエクスポート

export default function () {
  const address = "google.com:443"; // HTTPSサーバーのアドレス
  const conn = tcp.ConnectTLS(address, true); // TLS接続
  console.log(`Connected to ${address}`);

  tcp.WriteLn(conn, "Hello, server!");
  const response = tcp.Read(conn, 1024);
  console.log(`Response: ${String.fromCharCode(...response)}`);

  tcp.Close(conn);
}
