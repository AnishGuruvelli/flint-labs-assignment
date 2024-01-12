---

# **Flint Labs Assignment**

## **Overview**

This repository contains the source code for the Flint Labs Assignment project. It is a web application that allows users to check the token balance information for a given Ethereum wallet address. The application fetches real-time balance data and displays it along with the percentage change in the last 12 hours.

## **Features**

- **Token Balance Retrieval:** Fetches token balance information using an Ethereum wallet address.
- **Percentage Change Calculation:** Calculates and displays the percentage change in balance in the last 12 hours.
- **Dark Mode Toggle:** Allows users to toggle between light and dark modes for a personalized viewing experience.

## **Technologies Used**

- **Frontend:**
    - HTML
    - CSS
    - JavaScript
    - SweetAlert2 for displaying alerts
- **Backend:**
    - Go (Golang) - Gin framework
    - Ethereum client library: **`github.com/ethereum/go-ethereum`**
    - Logging: **`github.com/sirupsen/logrus`**

## **Project Structure**

The project is structured into frontend and backend components.

### **Frontend**

- **HTML:** **`index.html`** - Main HTML file with the web page's structure.
- **CSS:** **`style.css`** - Stylesheet for styling the web page.
- **JavaScript:** **`script.js`** - Client-side script for handling user interactions and fetching data from the backend.

### **Backend**

- **Go Files:** Located in the **`server`** directory.
    - **`main.go`** - Entry point for the backend server.
    - **`pkg/api/token_balance_controller.go`** - API controller for handling token balance requests.
    - **`pkg/client/infura/impl/infura_data_impl.go`** - Implementation of the Infura data service.
    - **`internal/usecases/impl/token_balance_usecase_impl.go`** - Implementation of the token balance use case.

## **Getting Started**

### **Prerequisites**

- Go (Golang) installed on your machine.
- Node.js and npm for frontend development.

### **Installation**

1. Clone the repository:
    
    ```bash
    git clone https://github.com/AnishGuruvelli/flint-labs-assignment.git
    cd flint-labs-assignment
    
    ```
    
2. Set up frontend dependencies:
    
    ```bash
    npm install
    
    ```
    
3. Run the backend server:
    
    ```bash
    go run server/main.go
    
    ```
    
4. Access the application in your web browser at **`http://localhost:8080`**.

## **Usage**

1. Enter an Ethereum wallet address in the input field.
2. Click the "Submit" button to fetch and display the token balance information.
3. Toggle dark mode for a different visual experience.

## **Resources**

1. https://piyopiyo.medium.com/how-to-get-ethereum-balance-with-json-rpc-api-provided-by-infura-io-6e5d22d25927

2. https://app.infura.io/key/3e581c1579624c01861c8400629bef66/active-endpoints

3. https://etherscan.io/myapikey

---