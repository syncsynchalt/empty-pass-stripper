# empty-pass-stripper

Some applications (such as
[micromdm/scep](https://github.com/micromdm/scep)) can create an RSA
key and apply an empty string as password.  This can be difficult to
read using tools such as openssl since they won't recognize an empty
string as a valid password.  This utility converts an encrypted RSA key
(in PEM format) to an unencrypted RSA key so it can be more easily worked
with using command line tools.

## Usage

To use this tool:

```
go install github.com/syncsynchalt/empty-pass-stripper@latest
~/go/bin/empty-pass-stripper /path/to/encrypted-key.pem > unencrypted-key.pem
```
