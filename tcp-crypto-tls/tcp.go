package tcp-crypto-tls

import (
	"crypto/tls"
	"net"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/tcp", new(TCP))
}

type TCP struct{}

// TCP接続
func (tcp *TCP) Connect(addr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// TLS接続
func (tcp *TCP) ConnectTLS(addr string, insecureSkipVerify bool) (net.Conn, error) {
	// TLS設定を構築
	tlsConfig := &tls.Config{
		InsecureSkipVerify: insecureSkipVerify, // 証明書検証をスキップするか
	}
	
	// TLS接続を確立
	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// 書き込み
func (tcp *TCP) Write(conn net.Conn, data []byte) error {
	_, err := conn.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// 読み込み
func (tcp *TCP) Read(conn net.Conn, size int) ([]byte, error) {
	buf := make([]byte, size)
	_, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// 改行付きで書き込み
func (tcp *TCP) WriteLn(conn net.Conn, data []byte) error {
	return tcp.Write(conn, append(data, []byte("\n")...))
}

// 接続を閉じる
func (tcp *TCP) Close(conn net.Conn) error {
	err := conn.Close()
	if err != nil {
		return err
	}
	return nil
}
