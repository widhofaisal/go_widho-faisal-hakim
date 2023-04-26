#!/bin/bash

# Mendapatkan argumen yang dimasukkan pada command line
arg1=$1
arg2=$2
arg3=$3

# mengambil nama folder dari tanggal dan waktu saat ini
mainfolder=$(date "+${arg1} at %a %b %e %T %Z %Y" | tr '[:upper:]' '[:lower:]')

# membuat folder utama
mkdir "$mainfolder"

# masuk ke folder utama
cd "$mainfolder"

# membuat folder about me di dalam folder utama
mkdir "about me"
mkdir "my_friends"
mkdir "my_system_info"

cd "about me"
mkdir "personal"
mkdir "professional"

cd "personal"
echo "https://www.facebook.com/${arg2}" > facebook.txt
cd ..

cd "professional"
echo "https://www.linkedin.com/in/${arg3}" > linkedin.txt
cd ..
cd ..
pwd

cd "my_friends"
cat <<EOT >>list_of_my_friends.txt
1) Achmad Miftahul - amulum
2) Achmad Rafiq - achmadrafiq97
3) Adiststi - adiststi
4) Agung - agungajin19
5) Azzahra - al7262
6) Charisma - fadzrichar
7) Farida - ulfarida
8 ) Garry - garryarielcussoy
9) Hamdi - hamdiranu
10) Hedy Gading - Gading09
11) Ilham - aamfatur
12) Lelianto - Lelianto
13) Mohammad - daffa99
14) Muhammad Fadhil - fabdulkarim
15) Ofbimon - bimon-alta
16) Purnama Syafitri - pipitmnr
17) Setyo - setyoyo07
18) Wildan - wiflash
19) Willy - sumarnowilly94
20) Woka - woka20
EOT
cd ..

cd "my_system_info"
echo "My username: tegarimansyah" > about_this_laptop.txt
echo "With host: $(uname -n) $(uname -v) $(uname -s) $(uname -r) $(uname -m)" >> about_this_laptop.txt

echo "Connection to google:" > internet_connection.txt
ping -c 3 forcesafesearch.google.com >> internet_connection.txt
