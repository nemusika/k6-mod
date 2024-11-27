# k6-mod

- build k6
```
xk6 build --with github.com/nemusika/k6-mod/tcp-crypto-tls@latest
```
- local
```
xk6 build --with github.com/nemusika/k6-mod/tcp-crypto-tls=C:\<dist-path>
```
- check version
```
C:\Users\nemusika\src\k6-mod>k6 version
k6.exe v0.55.0 (go1.23.3, windows/amd64)
Extensions:
  github.com/nemusika/k6-mod/tcp-crypto-tls (devel), k6/x/tcp [js]

```
