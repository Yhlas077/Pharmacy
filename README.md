# Pharmacy Management API Backend

Bu taslama, dermanhanalar we derman dolandyryş ulgamlary üçin niýetlenen, ýokary öndürijilikli we durnukly **Backend API** platformasydyr. Ulgamda ulanyjy dolandyryşy, dermanlaryň sanawy, sebet we sargyt amallary doly üpjün edilendir.

---

## Tehnologiýalar

Taslama aşakdaky professional gural we paketler arkaly gurlady:

*   **Dil:** Go (Golang 1.20+)
*   **Web Framework:** Gin Framework (`github.com/gin-gonic/gin`)
*   **Maglumatlar Binasy:** PostgreSQL
*   **Autentifikasiýa:** JWT (JSON Web Tokens)
*   **Gurşaw Sazlamalary:** `.env` (godotenv)

---

## 📂 Taslamanyň Gurluşy (Folder Structure)

```text
├── config/             # DB baglanyşygy we sazlamalar
├── controllers/        # Request-Response dolandyryjylary (Handlers)
├── repositories/         # JWT Auth we ygtyýarlylandyryş kontroly
├── models/             # PostgreSQL şablonlary (GORM Structs)
├── routes/             # API salgylarynyň paýlanyşy we toparlanmasy
├── .env                # Gizlin maglumatlar (Git-e goşulmaýar ❌)
├── .gitignore          # Git tarapyndan yzarbalmaly däl faýllar
├── go.mod              # Taslamanyň paket sanawy
└── main.go             # Serveriň giriş nokady