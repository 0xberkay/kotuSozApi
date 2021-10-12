## Kötü Söz Api Dökümantasyon

 **Rotalar**
 

 - Post => /api/check
 - Post => /api/filter


## 127.0.0.1:3000/api/check

**İstek :**
{
"str":"cok kotu bir soz"
}

**Cevap:**
{
"count": 4,
"message": "Kotu soz bulundu",
"status": "success"
}
## 127.0.0.1:3000/api/filter
**İstek :** 
{
"str":"iyi bir soz kotu bir soz",
"filter":"***"
}

**Cevap:**
{
"message": "iyi bir soz *** *** ***",
"status": "success"
}

##
str : Gönderilen mesaj
filter: Kötü sözle değiştirilecek filtre
message: Geriye dönen mesaj
status: Test edilme durumu
cout: Kötü kelime sayısı
