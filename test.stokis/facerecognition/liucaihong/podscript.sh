#!/bin/sh
mkdir /result
chmod a+x /user/lib/golang/src/github.com/liuchjlu/stokis/test/start.sh
/bin/sh /user/lib/golang/src/github.com/liuchjlu/stokis/test/start.sh

ftp -n -i <<-EOF
open  192.168.11.51
user  stokis 111111
binary
hash
mkdir facerecognition
cd facerecognition
mkdir liucaihong
cd liucaihong
mkdir v1
cd v1
lcd /result
mput *
bye
EOF

cd /result
size=0
for f in `ls`
do
let size=`stat -c %s $f`
if [ $size -gt 52428800 ]
then
echo The size of $f is bigger than 50Mb.
rm -rf $f
fi
done

ftp -n -i <<-EOF
open  192.168.11.51
user  stokis 111111
binary
hash
mkdir facerecognition
cd facerecognition
mkdir liucaihong
cd liucaihong
mkdir v1
cd v1
lcd /result
mput *
bye
EOF
