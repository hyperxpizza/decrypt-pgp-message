# Decrypt-PGP-Message
A simple command line tool for decrypting armored pgp messages

## How to build?
```
make build
```

## How to use?
```
Usage:
	-key=<path to your private pgp key file>
	-message=<path to message encrypted with your public pgp key>
	-password=<password for the armored message>
All flags must be set!
Correct usage:
	decryptpgpmessage -key=<path to your private pgp key file> -message=<path to message encrypted with your public pgp key> -password=<password for the armored message>

```

## How to use inside terminal?
```
sudo cp ./bin/decryptpgpmessage /usr/bin
```