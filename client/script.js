document.addEventListener("DOMContentLoaded", function () {
    const body = document.body;
    const container = document.querySelector('.container');

    // Create dark mode toggle button
    const darkModeToggle = document.createElement("button");
    darkModeToggle.innerText = "Toggle Dark Mode";
    darkModeToggle.addEventListener('click', function () {
        body.classList.toggle('dark-mode');
    });

    // Center align the button
    darkModeToggle.style.margin = '0 auto';
    darkModeToggle.style.display = 'block';

    // Append the button to the container
    container.appendChild(darkModeToggle);


    const fetchBalanceData = async (walletAddress) => {
        try {
            const response = await fetch(
                `http://localhost:8081/v1/token-balance/${walletAddress}`
            );
            const data = await response.json();
            console.log(data);
            return data;
        } catch (error) {
            console.error("Error fetching balance data:", error);
            Swal.fire({
                icon: "error",
                title: "Error",
                text: "An error occurred while fetching balance data.",
            });
            return null;
        }
    };

    const updateBalanceInfo = async () => {
        console.log("Updating balance info...");
        const walletAddress = document.getElementById("walletAddress").value;

        // Validate input
        if (!walletAddress) {
            Swal.fire({
                icon: "error",
                title: "Invalid Input",
                text: "Please enter the wallet address.",
            });
            return;
        }

        console.log(walletAddress);

        const balanceValueElement = document.getElementById("balanceValue");
        const changePercentageElement = document.getElementById("changePercentage");
        const timestampElement = document.getElementById("timestamp");
        const walletAddressDisplayElement = document.getElementById(
            "walletAddressDisplay"
        );
        const balanceSection = document.getElementById("balanceSection");

        balanceValueElement.textContent = "Loading...";
        changePercentageElement.textContent = "Loading...";

        const balanceData = await fetchBalanceData(walletAddress);

        if (balanceData) {
            balanceValueElement.textContent = balanceData.balance.toFixed(10);
            changePercentageElement.textContent =
                balanceData.changePercentage.toFixed(2) + "%";
            timestampElement.textContent = new Date(
                balanceData.timestamp
            ).toLocaleString();
            walletAddressDisplayElement.textContent = walletAddress;
            balanceSection.style.display = "block"; // Show the balance section

            // Display alert if percentage change is more than 10%
            if (balanceData.changePercentage > 10) {
                Swal.fire({
                    icon: "info",
                    title: "Percentage Change Alert",
                    text: "Percentage change is more than 10%!",
                });
            }
        } else {
            balanceValueElement.textContent = "Error";
            changePercentageElement.textContent = "Error";
            timestampElement.textContent = "";
            walletAddressDisplayElement.textContent = walletAddress;
            balanceSection.style.display = "none"; // Hide the balance section in case of an error
        }
    };

    const handleRefreshClick = async () => {
        console.log("Refresh button clicked.");
        await updateBalanceInfo();
    };

    const refreshButton = document.getElementById("refreshButton");

    refreshButton.addEventListener("click", handleRefreshClick);

    // remove this if you want to fetch data only when the user clicks refresh
    // updateBalanceInfo();
});
