nomer 1 :

Program ini menghitung permutasi dan kombinasi dari dua pasang bilangan asli.

Penjelasan:

Program meminta pengguna untuk memasukkan empat bilangan asli a, b, c, dan d.
Program mengecek apakah a lebih besar atau sama dengan c dan b lebih besar atau sama dengan d. Jika ya, program melanjutkan ke langkah 3. Jika tidak, program menampilkan pesan bahwa syarat tidak terpenuhi.
Program menggunakan fungsi factorial(), permutasi(), dan kombinasi() untuk menghitung permutasi dan kombinasi dari a terhadap c dan b terhadap d.
Program menampilkan hasil perhitungan permutasi dan kombinasi.
Fungsi-fungsi yang digunakan:

factorial(n int): Menghitung faktorial dari bilangan n.
permutasi(n, r int): Menghitung permutasi dari n elemen yang diambil r elemen.
kombinasi(n, r int): Menghitung kombinasi dari n elemen yang diambil r elemen.
Syarat yang harus dipenuhi:

a harus lebih besar atau sama dengan c.
b harus lebih besar atau sama dengan d.

nomer 2 : 

Program ini mendefinisikan tiga fungsi: f(x), g(x), dan h(x). Fungsi f(x) mengembalikan kuadrat dari input x, fungsi g(x) mengembalikan input x dikurangi 2, dan fungsi h(x) mengembalikan input x ditambah 1.

Kemudian program mendefinisikan tiga fungsi lagi: fogoh(x), gohof(x), dan hofog(x).

Fungsi fogoh(x) mengembalikan hasil dari mengaplikasikan fungsi f, g, dan h secara berurutan ke input x. Artinya, fogoh(x) = f(g(h(x))).

Fungsi gohof(x) mengembalikan hasil dari mengaplikasikan fungsi g, h, dan f secara berurutan ke input x. Artinya, gohof(x) = g(h(f(x))).

Fungsi hofog(x) mengembalikan hasil dari mengaplikasikan fungsi h, f, dan g secara berurutan ke input x. Artinya, hofog(x) = h(f(g(x))).

Program kemudian meminta pengguna untuk memasukkan tiga angka integer, yang kemudian disimpan dalam variabel a, b, dan c.

Selanjutnya program memanggil fungsi fogoh, gohof, dan hofog dengan masing-masing angka a, b, dan c sebagai argumen dan menampilkan hasilnya ke layar.

Secara keseluruhan, program ini menunjukkan bagaimana fungsi dapat digabungkan untuk menciptakan fungsi baru yang kompleks.

nomer 3 :

Program ini menerima input berupa koordinat titik pusat dan jari-jari dari dua buah lingkaran, serta koordinat titik sembarang. Kemudian program menentukan posisi titik sembarang terhadap kedua lingkaran dan menampilkannya sebagai string.

Berikut adalah detail langkah program:

Input: Program menerima input berupa koordinat titik pusat (cx1, cy1) dan jari-jari (r1) dari lingkaran pertama, koordinat titik pusat (cx2, cy2) dan jari-jari (r2) dari lingkaran kedua, serta koordinat titik sembarang (x, y).

Perhitungan Jarak: Program menghitung jarak dari titik sembarang ke titik pusat lingkaran pertama (dist1) dan jarak dari titik sembarang ke titik pusat lingkaran kedua (dist2) menggunakan rumus jarak Euclidean.

Penentuan Posisi: Program membandingkan jarak yang telah dihitung dengan jari-jari masing-masing lingkaran untuk menentukan posisi titik sembarang.

Output: Program mencetak hasil berupa string yang menyatakan posisi titik sembarang, yaitu:

"Titik di dalam lingkaran 1 dan 2" jika titik sembarang berada di dalam kedua lingkaran.
"Titik di dalam lingkaran 1" jika titik sembarang berada di dalam lingkaran pertama, tetapi di luar lingkaran kedua.
"Titik di dalam lingkaran 2" jika titik sembarang berada di dalam lingkaran kedua, tetapi di luar lingkaran pertama.
"Titik di luar lingkaran 1 dan 2" jika titik sembarang berada di luar kedua lingkaran.
 
 
