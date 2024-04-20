# QR Kod Oluşturma Servisi

Bu basit Go programı, belirtilen bir URL'den QR kodu oluşturur ve istemcilere geri döndürür.

## Nasıl Çalışır

1. **Kurulum**: Projeyi klonlayın ve Go'nun yüklü olduğundan emin olun.
2. **Bağımlılıklar**: `go get` komutuyla `github.com/boombuler/barcode` ve `github.com/boombuler/barcode/qr` kütüphanelerini indirin.
3. **Sunucuyu Başlatın**: `go run main.go` komutuyla sunucuyu başlatın.
4. **Test Edin**: Tarayıcıdan veya bir API test aracından `http://localhost:8080/barcode?url=YOUR_URL` adresine bir GET isteği yaparak QR kodunu oluşturun.

## API Dökümantasyonu

### `/barcode` Endpoint'i

- **Method**: GET
- **Parametreler**: `url` parametresi, QR kodu oluşturmak için kullanılacak web sayfasının URL'sini içermelidir.
- **Cevap**: Oluşturulan QR kodunu içeren bir PNG resmi döndürür.

Örnek istek:
GET http://localhost:8080/barcode?url=https://www.example.com

![Açıklama](https://github.com/omerfdev/urlToBarcodeAPI/blob/main/Opera%20Snapshot_2024-02-22_171508_localhost.png)
