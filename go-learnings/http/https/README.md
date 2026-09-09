# gohttps

This repository contains examples of HTTPS clients and servers. It includes simple server with minimal TLS configuration and a more advanced server that covers additional TLS configuration options. They demonstrate a range of behavorios between TLS clients and servers. See [Create Secure Clients and Servers in Golang Using HTTPS](https://youngkin.github.io/post/gohttpsclientserver/) for more information regarding this project.


-------------------------------- FLOW --------------------------------------
- client → public, private
- server → public, private
- CA (certificate Authority)

server’s public cert is signed by CA private key 

when client requests for https://:[server.com](http://server.com) then in the browser (usually has all the popular CA public key) so server will send its public key that is signed by CA private key which can be unlocked by CA public key present in browser 

- FLOW
    - CA
        - public
        - private
    - server
        - public > this is sent to CA > CA signs it using its private key
        - private
    - client
        - when clients send requst to [server.com](http://server.com) > server.com sends its public key in the response (which is signed by CA private key)
        - generally all browsers have popular CA's public key -  hence client is able to decrypt it by CA public key from the browser
        - now clients send their data by encrypting using [server.com](http://server.com) public key
        - which can only decrypted by [server.com](http://server.com) bcz now one has its private key