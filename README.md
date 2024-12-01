# financial-account-management

Financial Account Management System

A web application for managing financial accounts, transactions, and generating financial reports, designed for businesses.

The system includes:
	•	Api: Built with Golang and GoFiber.
	•	Database: PostgreSQL.
	•	Frontend: ReactJS.
	•	Containerization: Docker and Docker Compose for easy deployment.

Features

	•	User Management:
	•	CRUD operations for user accounts.
	•	Financial Accounts:
	•	CRUD operations for business accounts.
	•	Transactions:
	•	Create and list transactions.
	•	Reconcile accounts.
	•	Reports:
	•	Generate financial reports (e.g., balance sheet, income statement).

Getting Started

Prerequisites

	•	Docker
	•	Docker Compose
	•	Make (optional but recommended)

Installation

	1.	Clone the repository:

git clone git@github.com:your-username/financial-system.git
cd financial-system


	2.	Set up the .env file for environment variables (optional):

cp .env.example .env

Update the values in .env as needed (e.g., database credentials).

Usage

Using Docker Compose

To manage the application, use the provided Makefile for convenience.

Start Services

make up

Stop Services

make down

View Logs

make logs

Rebuild Services

make build

Clean Containers and Volumes

make clean

API Endpoints

Users

	•	POST /api/users: Create a user.
	•	GET /api/users: Get all users.
	•	PUT /api/users/:id: Update a user.
	•	DELETE /api/users/:id: Delete a user.

Accounts

	•	POST /api/accounts: Create an account.
	•	GET /api/accounts: Get all accounts.
	•	PUT /api/accounts/:id: Update an account.
	•	DELETE /api/accounts/:id: Delete an account.

Transactions

	•	POST /api/transactions: Create a transaction.
	•	GET /api/transactions: List all transactions.
	•	POST /api/transactions/reconcile: Reconcile transactions.

Reports

	•	GET /api/reports/balance-sheet: Generate a balance sheet.
	•	GET /api/reports/income-statement: Generate an income statement.

Development

Run Backend Locally

	1.	Navigate to the backend directory:

cd backend


	2.	Start the server:

go run main.go



Run Frontend Locally

	1.	Navigate to the frontend directory:

cd frontend


	2.	Start the React development server:

npm start

Technologies Used

	•	Backend: GoFiber (Golang), GORM (ORM).
	•	Frontend: ReactJS, Axios.
	•	Database: PostgreSQL.
	•	Containerization: Docker, Docker Compose.

Contributing

Contributions are welcome! Please follow these steps:
	1.	Fork the repository.
	2.	Create a new branch:

git checkout -b feature/your-feature-name


	3.	Make your changes and commit:

git commit -m "Add your message here"


	4.	Push to your branch:

git push origin feature/your-feature-name


	5.	Open a Pull Request.

License

This project is licensed under the Apache 2.0 License.

Acknowledgements

	•	Fiber
	•	React
	•	Docker
