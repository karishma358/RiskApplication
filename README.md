# 🛠 Risk Management API

This is a simple **Risk Management API** built using **Golang**, following best practices like structured logging, improved error handling, and a health check endpoint.

---

## 🚀 Features
- 🖥 **CRUD API** for managing risks  
- 📄 **In-memory storage** (No database required)  
- 🔄 **Proper error handling & logging**  
- 🏥 **Health check endpoint (`/health`)**  
- ✅ **Unit tests included**  

---

## 📌 **Prerequisites**
Before running the application, ensure you have:
- **Go 1.19+** installed → [Download Go](https://golang.org/dl/)
- **cURL/Postman** for API testing (optional)

---

## 🔧 **Setup Instructions**
### 1️⃣ Clone the Repository
```sh
git clone https://github.com/karishma358/RiskApplication.git
cd riskapp
```

### 2️⃣ Install Dependencies
```sh
go mod tidy
```

### 3️⃣ Run the Application
```sh
go run main.go
```
The server will start at **`http://localhost:8080`**.

---

## 🛠 **Available API Endpoints**
| Method | Endpoint | Description |
|--------|-------------|------------------------------|
| `GET` | `/health` | ✅ Check if API is running |
| `GET` | `/v1/risks` | 📄 Get all risks |
| `POST` | `/v1/risks` | ➕ Create a new risk |
| `GET` | `/v1/risks/{id}` | 🔍 Get a risk by ID |

---

## 🚀 **Testing the API**
### **1️⃣ Health Check**
✅ Open in browser:  
[http://localhost:8080/health](http://localhost:8080/health)  
Expected Response:
```json
{"status": "ok"}
```

### **2️⃣ Get All Risks (Initially Empty)**
```sh
curl -X GET http://localhost:8080/v1/risks
```
Expected Response (if no risks exist):
```json
{"error": "No risks are currently present"}
```

### **3️⃣ Create a New Risk**
```sh
curl -X POST http://localhost:8080/v1/risks \
     -H "Content-Type: application/json" \
     -d '{"state":"open","title":"Server Down","description":"The database server is down"}'
```
Expected Response:
```json
{
    "id": "c2a77f9e-1e48-432f-9f65-3a3a5efb1df7",
    "state": "open",
    "title": "Server Down",
    "description": "The database server is down"
}
```

### **4️⃣ Get a Risk by ID**
Replace `{id}` with the ID from the previous step:
```sh
curl -X GET http://localhost:8080/v1/risks/c2a77f9e-1e48-432f-9f65-3a3a5efb1df7
```

---

## 🧪 **Running Unit Tests**
To run tests, execute:
```sh
go test ./tests -v
```
Example Output:
```
=== RUN   TestCreateRisk
--- PASS: TestCreateRisk (0.02s)
=== RUN   TestGetAllRisks
--- PASS: TestGetAllRisks (0.01s)
=== RUN   TestGetRiskByID
--- PASS: TestGetRiskByID (0.02s)
PASS
```

---

## 🐳 **Run with Docker (Optional)**
1. **Build the Docker Image**
   ```sh
   docker build -t riskapp .
   ```
2. **Run the Container**
   ```sh
   docker run -p 8080:8080 riskapp
   ```
3. **Test APIs Again**
   ```sh
   curl -X GET http://localhost:8080/health
   ```

---

### 🚀 **Now You're Ready to Use the Risk Management API!** 🎯
