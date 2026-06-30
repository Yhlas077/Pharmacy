# 💊 Pharmacy Backend API

**Dermanhanalaryň işini awtomatlaşdyrmak, derman serişdeleriniň goruny dolandyrmak we sargytlary hasaba almak üçin niýetlenen durnukly, çalt we ygtybarly RESTful API.**

## ✨ Esasy Aýratynlyklary (Features)

- 🔐 **Howpsuzlyk:** JWT we Bcrypt arkaly ulanyjy awtentifikasiýasy we awtorizasiýasy (Role-based access).
- 📦 **Goruň Dolandyrylyşy (Inventory):** Dermanlary goşmak, üýtgetmek, öçürmek we galdygyny yzarlamak.
- 🛒 **Sargyt Ulgamy:** Müşderileriň sargytlaryny döretmek we sargyt ýagdaýyny (Pending, Delivered, Cancelled) dolandyrmak.
- ⚡ **Ýokary Tizlik:** Go we Gin framework arkaly iň ýokary öndürijilik we pes ýat (memory) ulanylyşy.
- 🗄️ **Ygtybarly Baza:** PostgreSQL arkaly maglumatlaryň ygtybarly we relýasion (relational) saklanmagy.

## 🛠 Tehnologiýalar (Tech Stack)

| Kategoriýa | Tehnologiýa |
| :--- | :--- |
| **Programmirleme dili** | `Go (Golang)` |
| **Web Framework** | `Gin-Gonic` |
| **Maglumatlar Bazasy** | `PostgreSQL` |
| **Awtentifikasiýa** | `JWT (JSON Web Tokens)` |
| **Gurşaw (Env) Dolandyryşy** | `godotenv` |

---

## 📂 Taslamanyň Gurluşy (Folder Structure)

Arassa arhitektura (Clean Architecture) kadalaryna laýyklykda gurlan gurluş:

```text
Pharmacy/
├── config/             # Database we Gurşaw (Env) sazlamalary
├── controllers/        # HTTP Request handler-lar (Ulgam dolandyryjylary)
├── middleware/         # Auth, CORS we Log middleware-lar
├── models/             # PostgreSQL DB gurluşlary we Structs
├── repository/         # Database queries (Maglumat bazasy bilen göni işleşýän gatlak)
├── routes/             # API-yň ähli ugurlary we gatnaw nokatlary
└── main.go             # Programmanyň başlanyş nokady

🔌 API Endpoints (Gatnaw Nokatlary)
Ulgamyň esasy API nokatlarynyň gysgaça beýany:

POST /api/v1/auth/register — Täze ulanyjy hasabyny açmak

POST /api/v1/auth/login — Ulgama girmek we JWT token almak

GET /api/v1/medications — Ähli dermanlaryň sanawy

GET /api/v1/medications/:id — Belli bir dermanyň maglumaty

POST /api/v1/medications — Täze derman goşmak (Admin)

PUT /api/v1/medications/:id — Dermanyň maglumatyny täzelemek (Admin)

DELETE /api/v1/medications/:id — Dermany sanawdan öçürmek (Admin)

POST /api/v1/orders — Täze sargyt döretmek

GET /api/v1/orders — Öz sargytlaryňy görmek

PUT /api/v1/orders/:id/status — Sargydyň ýagdaýyny üýtgetmek (Admin)

🚀 Gurnalyşy (Getting Started)
Öz lokal kompýuteriňizde işledip görmek üçin aşakdaky ädimleri yzarlaň:

1. Klounlamak (Clone the repo)
Bash
git clone [https://github.com/yhlas077/Pharmacy.git](https://github.com/yhlas077/Pharmacy.git)
cd Pharmacy
2. Gurşaw sazlamalary (.env)
Taslamanyň esasy papkasynda .env faýlyny dörediň we maglumatlaryňyzy ýazyň:

Code snippet
PORT=8080
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=siziň_parolyňyz
DB_NAME=pharmacy_db
DB_PORT=5432
JWT_SECRET=super_secret_key

3. Modullary ýükläň we Işlediň
Terminal
go mod tidy
go run main.go

Server indi http://localhost:8080 salgysynda taýýar! 🎉