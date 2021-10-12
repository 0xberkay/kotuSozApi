## Kötü Söz Api Dökümantasyon

 **Rotalar**
 

 - Post => /api/check
 - Post => /api/filter


## 127.0.0.1:3000/api/check

**İstek :**<br>
{<br>
"str":"cok kotu bir soz"<br>
}

**Cevap:**<br>
{<br>
"count": 4,
"message": "Kotu soz bulundu",<br>
"status": "success"<br>
}
## 127.0.0.1:3000/api/filter
**İstek :** <br>
{<br>
"str":"iyi bir soz kotu bir soz",<br>
"filter":"***"<br>
}

**Cevap:**<br>
{<br>
"message": "iyi bir soz *** *** ***",<br>
"status": "success"<br>
}

##
str : Gönderilen mesaj<br>
filter: Kötü sözle değiştirilecek filtre<br>
message: Geriye dönen mesaj<br>
status: Test edilme durumu<br>
cout: Kötü kelime sayısı
