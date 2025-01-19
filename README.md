# ExpenseTracker

ExpenseTracker is a web-based application that helps users manage their personal expenses effectively. This application enables users to track their spending and stay on top of their finances.

---

## Features

### 1. **User Authentication**
- Secure user login and registration using JWT-based authentication.
- Password encryption for enhanced security.

### 2. **Expense Management**
- Add, edit, list and delete expenses.

### 3. **Responsive Design**
- Optimized for both desktop and mobile devices.

### 4. **Secure Database Interaction**
- Utilizes a robust database to store and retrieve user data securely.

---

## Technologies Used

- **Backend**: Go (Golang)
- **Modules**: `net/http`, `driver/mysql`, `jwt-go`
- **Database**: MySQL
- **Authentication**: JWT (JSON Web Tokens)
- **Version Control**: Git and GitHub

---

## Installation and Setup

Follow these steps to run the project locally:

### Prerequisites
- Go installed on your machine (v1.20+ recommended)
- PostgreSQL installed and configured
- Git installed

### Clone the Repository
```bash
git clone https://github.com/Aounjafri/ExpenseTracker.git
cd ExpenseTracker
```

### Install Dependencies
```bash
go mod tidy
```

### Set Up the Database
- Create a MySQL database named `expense_tracker`.
- Add Tables (User and Expense)

### Run the Application
```bash
go run main.go
```

The application will be available at `http://localhost:8080`.

---

## Usage

1. **Register**: Register using credentials (Username, Password, Email)
2. **Login**: Login with your credentials and get a token.
3. **Add Expenses**: Navigate to the dashboard and add your expenses.
4. **Manage Expenses**: Edit or delete expenses as needed.

---

## Contributing

Contributions are welcome! Follow these steps to contribute:

1. Fork the repository.
2. Create a new branch: `git checkout -b feature-name`.
3. Commit your changes: `git commit -m 'Add a new feature'`.
4. Push to the branch: `git push origin feature-name`.
5. Open a pull request.

---

## Contact

For any queries or feedback, feel free to contact me:
- **Email**: mailto:jafriaoun54@gmail.com
- **GitHub**: https://github.com/Aounjafri

---

Thank you for using ExpenseTracker! 🚀
