# Enit

Encrypt it!  
Encrypt any string in pattern of "key value", and have a key argument for encoding is blank by default.  
You can set it via `-k <your encoding_key>`, anyway, it encrypt [key, value] via aes.  

# Install

**Highly recommend:**  
copy the exec file to your system bin folder, so you can call it convenient as system commands.  

## via go
```
git clone https://github.com/hi20160616/enit.git
cd enit
go install cmd/enit.go
```
## via releases
Download from [Releases Page](https://github.com/hi20160616/enit/releases), 
decompress tarball and copy it to `/usr/bin`

# Usage

Run `enit` to read help, the fast begin case is below:
```
$ enit set <yourkey> <yourvalue>
$
$ enit get <yourkey>
// Output:
$ <yourkey> = <yourvalue>
$
$ enit get <yourkey>
$ enit ls
// Output:
$ <your storaged key-values>
$
$ enit rm <key>
// Output:
Secret <key> remove successfully.
```
Also, you can add key for encoding secret file.  
```
$ enit set -k <encoding_key> <yourkey> <yourvalue>
$
$ enit get -k <encoding_key> <yourkey>
// Output:
$ <yourkey> = <yourvalue>
```

# Trick
```
mkdir enit && cd enit
mv ~/.enit ./enit_anyname && ln -s <your prepath>/enit/enit_anyname ~/.enit
```
so, you can backup or sync the file on any plant.

# Details
The encrypted config file is located at `$HOME/.enit`.  
For windows user, it is located at `C:\Users\<username>\.enit`

Enjoy it!  

EOF
