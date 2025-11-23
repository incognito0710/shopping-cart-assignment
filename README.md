🛒 Shopping Cart – Fullstack Assignment

A simple, clean and functional shopping experience built as part of the Fullstack Developer Internship assignment.

This project implements the complete flow described  — allowing a user to log in, view items, add items to their cart, place an order, and view their order history.

The solution includes:

Backend: Go (Gin) + GORM + PostgreSQL

Frontend: React + React Router

Auth: Token-based session (single token per user)

Tools: Postman Collection, README, GitHub-ready structure

Everything is lightweight, easy to run, and built exactly to match the assignment requirements.

🚀 Features Implemented

✔ User Login (with token generation)
✔ List of Items (fetched from backend)
✔ Add Items to Cart
✔ View Cart Items (via alert)
✔ Checkout (cart → order)
✔ View Order History (via alert)
✔ Responsive UI (simple & functional)
✔ Fully working backend APIs
✔ Postman collection included
✔ Follows assignment PDF step-by-step

📂 Project Structure
shoppingcart/
│
├── backend/                   # Go backend
│   ├── main.go
│   ├── database.go
│   ├── models.go
│   ├── controllers.go
│   └── routes.go
│
└── frontend/
    └── shopping-frontend/     # React app
        ├── src/
        │   ├── App.js
        │   ├── index.js
        │   └── pages/
        │       ├── Login.js
        │       ├── Items.js
        │       ├── Cart.js
        │       └── Orders.js
        └── package.json

⚙️ Backend Setup (Go + Gin + GORM + PostgreSQL)
1️⃣ Install Dependencies

Inside the backend/ folder:

go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
go get gorm.io/gorm
go get gorm.io/driver/postgres

2️⃣ Configure PostgreSQL

Create a database named:

shoppingdb


Default credentials used in the project:

host:     localhost
port:     5432
user:     postgres
password: 1234


(If your Postgres password is different, update it in database.go.)

3️⃣ Start the Backend

Inside the backend folder:

go run .


If everything is correct, you’ll see:

Listening and serving HTTP on :8080


Backend runs at:
👉 http://localhost:8080

📡 API Endpoints Summary
👤 User APIs
Method	Endpoint	Description
POST	/users	Create a new user
GET	/users	List all users
POST	/users/login	Login, returns token
🛍️ Item APIs
Method	Endpoint	Description
POST	/items	Create an item
GET	/items	List items
🛒 Cart APIs

Header required:

token: <user_token_here>

Method	Endpoint	Description
POST	/carts	Add item to cart
GET	/carts	Show all cart items
📦 Order APIs
Method	Endpoint	Description
POST	/orders	Convert cart → order
GET	/orders	List all orders
🌐 Frontend Setup (React)
1️⃣ Install Dependencies

Inside the shopping-frontend folder:

npm install

2️⃣ Start React App
npm start


React runs on:

👉 http://localhost:3000

(or http://localhost:3001
 if 3000 is busy)

🖥️ Frontend Features 
✔ Login Page

User enters username & password

On success → navigates to Items page

On failure → alert("Invalid username/password")

✔ Items Page

Fetches and displays all items

Clicking on Add to Cart adds the item to user's cart

Contains:

Checkout button

Cart button → shows cart items in alert

Order History button → shows past orders in alert

✔ Checkout

Calls backend to convert cart into an order

Displays "Order successful" alert

✔ Cart Display

A popup showing:

[{ cart_id, user_id, item_id }]

✔ Order History Display

A popup showing:

[{ order_id, user_id, cart_id }]

📬 Postman Collection Included

A ready-to-import Postman file is included:

shopping_cart_postman_collection.json


Import it into Postman → all endpoints will appear automatically.

📝 How to Test the Flow
1️⃣ Create a User

POST /users

2️⃣ Login

POST /users/login → copy token

3️⃣ Create Items

POST /items

4️⃣ Add Items to Cart

POST /carts
Header → token: <your_token>

5️⃣ Checkout

POST /orders

6️⃣ Order History

GET /orders

📤 GitHub Submission Instructions

From the root project folder:

git init
git add .
git commit -m "Shopping Cart Assignment Completed"
git branch -M main
git remote add origin <your_repo_url>
git push -u origin main


Then submit your GitHub link.

🎉 Final Notes

This project implements every requirement from the assignment PDF exactly:

Full backend in Go (Gin + GORM)

Full frontend in React

Token-based session

Cart & order flow

Postman support

Clean folder structure

Easy reproducibility

